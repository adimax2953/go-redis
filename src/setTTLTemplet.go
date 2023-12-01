package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// SetTTL function - keys, args[] string - return string , error
func (s *MyScriptor) SetTTL(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(ExpireKEYID, keys, args)
	if err != nil {
		logtool.LogError("SetTTL ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if result.Value == "" {
		logtool.LogError("SetTTL Value Is Nil")
		return "", err
	}

	return result.Value, nil
}

// SetTTL - 修改Key的TTL
const (
	SetTTLID       = "SetTTL"
	SetTTLTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   SetTTL
		EVALSHA  <script_sha1> 0 {DBKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local k1                                    		= KEYS[2]
		local v1                                        	= KEYS[3]
		local sender                                        = "SetTTL.lua"
		
		
		local MAIN_KEY = k1
	
		redis.call("select",DBKey)
		local r1= ""
		if v1=="-1" then
			redis.call('persist',MAIN_KEY)
		else
			redis.call('expire',MAIN_KEY,v1)
		end

		r1 = redis.call('ttl',MAIN_KEY)
		return { r1 }
	
    `
)
