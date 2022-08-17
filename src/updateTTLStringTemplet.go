package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// UpdateTTLString function - keys, args[] string - return string , eror
func (s *MyScriptor) UpdateTTLString(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(UpdateTTLStringID, keys, args)
	if err != nil {
		logtool.LogError("UpdateTTLString ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))

	result.Value = reader.ReadString()
	if result.Value == "" {
		logtool.LogError("UpdateTTLString Value Is Nil")
		return "", err
	}

	return result.Value, nil
}

// UpdateString - 寫入一個字串
const (
	UpdateTTLStringID       = "UpdateTTLString"
	UpdateTTLStringTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   UpdateTTLString
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = ARGV[2]
		local sender                                        = "UpdateTTLString.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
		
			redis.call("select",DBKey)
			local r1= ""
				redis.call('MSET',MAIN_KEY , v1)
		
			redis.call('EXPIRE', MAIN_KEY , 10 )
				r1 = redis.call('MGET',MAIN_KEY)
			return { r1 }
		end
    `
)
