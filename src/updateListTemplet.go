package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// UpdateList function - keys, args[] string - return string , error
func (s *MyScriptor) UpdateList(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(UpdateListID, keys, args)
	if err != nil {
		logtool.LogError("UpdateList ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("UpdateList Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// UpdateList - 減少數值
const (
	UpdateListID       = "UpdateList"
	UpdateListTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   UpdateList
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1} {v2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = tonumber(ARGV[2])    
		local v2                                            = ARGV[3]
		local sender                                        = "UpdateList.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and v1 and v2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local chkexists = 0
			local chkllen = 0
			local r1 = ""
			chkexists = redis.call('exists',MAIN_KEY)
			chkllen = redis.call('LLEN',MAIN_KEY)
			if chkexists == 1 then
				if tonumber(v1) < chkllen then
					local tmp = redis.call('lset',MAIN_KEY,v1,v2)
					if tmp~=nil and tmp~=false and tmp~="" then
						r1 = v2
						return { r1 }
					end
				elseif tonumber(v1) >= chkllen then
						return { "err , index out of range" }
				end
			end
		
		end
    `
)
