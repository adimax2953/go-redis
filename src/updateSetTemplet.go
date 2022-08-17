package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// UpdateSet function - keys, args[] string - return string , error
func (s *MyScriptor) UpdateSet(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(UpdateSetID, keys, args)
	if err != nil {
		logtool.LogError("UpdateSet ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("UpdateSet Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// UpdateSet - 減少數值
const (
	UpdateSetID       = "UpdateSet"
	UpdateSetTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   UpdateSet
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1} {v2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = ARGV[2]
		local v2                                            = ARGV[3]
		local sender                                        = "UpdateSet.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and v1 and v2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			redis.call('srem',MAIN_KEY ,v1)			
			local r1 = redis.call('sadd',MAIN_KEY ,v2)
			if r1 == 1 then 
				return { tostring(v2) }
			else
				r1= "-1"
				return { tostring(r1) }
			end
		end
    `
)
