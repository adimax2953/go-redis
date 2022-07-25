package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// ExpireKEY function - keys, args[] string - return string , error
func (s *MyScriptor) ExpireKEY(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(ExpireKEYID, keys, args)
	if err != nil {
		logtool.LogError("ExpireKEY ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if result.Value == "" {
		logtool.LogError("ExpireKEY Value Is Nil")
		return "", err
	}

	return result.Value, nil
}

// ExpireKEY - 寫入一個字串
const (
	ExpireKEYID       = "ExpireKEY"
	ExpireKEYTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   ExpireKEY
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = ARGV[2]
		local sender                                        = "ExpireKEY.lua"
		
		
		if DBKey and ProjectKey and TagKey and k1 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1= ""
			redis.call('expire',MAIN_KEY,v1)
			r1 = redis.call('ttl',MAIN_KEY)
			return { r1 }
		end
    `
)
