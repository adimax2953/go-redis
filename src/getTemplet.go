package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// Get function - keys, args[] string - return string , error
func (s *MyScriptor) Get(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(GetID, keys, args)
	if err != nil {
		logtool.LogError("Get ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("Get Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// Get - 爬key
const (
	GetID       = "Get"
	GetTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   Get
		EVALSHA  <script_sha1> 0 {DBKey} {Key} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local Key                                    		= KEYS[2]
		local sender                                        = "Get.lua"
		
		redis.call('select',DBKey)	
		local result = redis.call('get',Key)
		return {result}
    `
)
