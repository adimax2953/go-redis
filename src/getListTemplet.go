package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetList function - keys, args[] string - return string , error
func (s *MyScriptor) GetList(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(GetListID, keys, args)
	if err != nil {
		logtool.LogError("GetList ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("GetList Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// GetList - 減少數值
const (
	GetListID       = "GetList"
	GetListTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetList
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = tonumber(ARGV[2]) --idx
		local sender                                        = "GetList.lua"
		
		if DBKey and ProjectKey and TagKey and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1 = ""
			local tmp = redis.call('lindex',MAIN_KEY ,v1)
			if tmp~= nil and tmp~= false then
				r1= tmp
			else
				r1= ""
			end
			return { r1 }
		end
    `
)
