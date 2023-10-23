package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// Set function - keys, args[] string - return string , error
func (s *MyScriptor) Set(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(SetID, keys, args)
	if err != nil {
		logtool.LogError("Set ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("Set Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// Set - 爬key
const (
	SetID       = "Set"
	SetTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   Set
		EVALSHA  <script_sha1> 0 {DBKey} {Key} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local Key                                    		= KEYS[2]
		local Value                                    		= KEYS[3]
		local sender                                        = "Set.lua"
		
		redis.call('select',DBKey)	
		local result = redis.call('set',Key ,Value)
		return result
    `
)
