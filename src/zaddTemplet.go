package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// ZAdd function - keys, args[] string - return string , error
func (s *MyScriptor) ZAdd(keys, args []string) (string, int64, error) {
	res, err := s.Scriptor.ExecSha(ZAddID, keys, args)
	if err != nil {
		logtool.LogError("ZAdd ExecSha Error", err)
		return "", 0, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	result.ValueInt64, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("ZAdd ReadInt64 Error", err)
		return result.Value, 0, err
	}
	return result.Value, result.ValueInt64, nil
}

// ZAdd - 爬key
const (
	ZAddID       = "ZAdd"
	ZAddTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   ZAdd
		EVALSHA  <script_sha1> 0 {DBKey} {Key} {Value}  {Score} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local Key                                    		= KEYS[2]
		local Value                                    		= ARGV[1]
		local Score                                    		= tonumber(ARGV[2])
		local sender                                        = "ZAdd.lua"
		
		redis.call('select',DBKey)	
		local result = redis.call('zadd',Key ,Score ,Value)
		return {result}
    `
)
