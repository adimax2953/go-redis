package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetZset function - keys, args[] string - return int64 , error
func (s *MyScriptor) GetZset(keys, args []string) (*RedisResult, error) {
	res, err := s.Scriptor.ExecSha(GetZsetID, keys, args)
	if err != nil {
		logtool.LogError("GetZset ExecSha Error", err)
		return nil, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	result.Value2 = reader.ReadString()

	return result, nil
}

// GetZset - 寫入一個字串
const (
	GetZsetID       = "GetZset"
	GetZsetTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetZset
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = ARGV[2]
		local sender                                        = "GetZset.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)    
			local tmp = redis.call('zscore',MAIN_KEY ,v1)
			local r2="0"
			local r1="0"
			if tmp~=nil and tmp~="" and tmp~=false then
				r1= tmp
				r2= v1
			end
			return { r1 , r2 }
		end
    `
)
