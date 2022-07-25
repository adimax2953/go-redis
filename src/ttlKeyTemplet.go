package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// TTLKey function - keys, args[] string - return string , error
func (s *MyScriptor) TTLKey(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(TTLKeyID, keys, args)
	if err != nil {
		logtool.LogError("TTLKey ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if result.Value == "" {
		logtool.LogError("TTLKey Value Is Nil")
		return "", err
	}

	return result.Value, nil
}

// TTLKey - 寫入一個字串
const (
	TTLKeyID       = "TTLKey"
	TTLKeyTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   TTLKey
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "TTLKey.lua"
		
		
		if DBKey and ProjectKey and TagKey and k1 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1= ""
			r1 = redis.call('ttl',MAIN_KEY)
			return { r1 }
		end
    `
)
