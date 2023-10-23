package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// SAdd function - keys, args[] string - return string , error
func (s *MyScriptor) SAdd(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(SAddID, keys, args)
	if err != nil {
		logtool.LogError("SAdd ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("SAdd Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// SAdd - 爬key
const (
	SAddID       = "SAdd"
	SAddTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   SAdd
		EVALSHA  <script_sha1> 0 {DBKey} {Key} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local Key                                    		= KEYS[2]
		local Value                                    		= KEYS[3]
		local sender                                        = "SAdd.lua"
		
		redis.call('select',DBKey)	
		local result = redis.call('sadd',Key ,Value)
		return {result}
    `
)
