package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// NewSet function - keys, args[] string - return string , error
func (s *MyScriptor) NewSet(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(NewSetID, keys, args)
	if err != nil {
		logtool.LogError("NewSet ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("NewSet Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// NewSet - 減少數值
const (
	NewSetID       = "NewSet"
	NewSetTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   NewSet
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = ARGV[2]
		local sender                                        = "NewSet.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1= ""
			redis.call('sadd',MAIN_KEY ,v1)
			local tmp = redis.call('sismember',MAIN_KEY ,v1)
			if tmp == 1 then
				r1 = v1
			end
			return { r1 }
		end
    `
)
