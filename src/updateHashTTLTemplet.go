package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// UpdateHashTTL function - keys, args[] string - return string , error
func (s *MyScriptor) UpdateHashTTL(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(UpdateHashTTLID, keys, args)
	if err != nil {
		logtool.LogError("UpdateHashTTL ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("UpdateHashTTL Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// UpdateHashTTL - 減少數值
const (
	UpdateHashTTLID       = "UpdateHashTTL"
	UpdateHashTTLTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   UpdateHashTTL
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local v1                                            = ARGV[3]
		local v2                                            = tonumber(ARGV[4])
		local sender                                        = "UpdateHashTTL.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and k2 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1= ""
			redis.call('hset',MAIN_KEY , k2 ,v1)
			r1 = redis.call('hget',MAIN_KEY , k2)
			redis.call('EXPIRE', MAIN_KEY , v2 )
			return { r1 }
		end
    `
)
