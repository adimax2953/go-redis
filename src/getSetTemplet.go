package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetSet function - keys, args[] string - return string , error
func (s *MyScriptor) GetSet(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(GetSetID, keys, args)
	if err != nil {
		logtool.LogError("GetSet ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("GetSet Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// GetSet - 減少數值
const (
	GetSetID       = "GetSet"
	GetSetTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetSet
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = ARGV[2]
		local sender                                        = "GetSet.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1 = ""
			local tmp = redis.call('sismember',MAIN_KEY,v1)
			if tmp == 1 then
				r1 = v1
			end
			
			return { r1 }
		end
    `
)
