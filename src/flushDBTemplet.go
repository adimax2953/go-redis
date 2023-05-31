package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// FlushDB function - keys, args[] string - return string , error
func (s *MyScriptor) FlushDB(keys []string) (string, error) {
	res, err := s.Scriptor.ExecSha(FlushDBID, keys)
	if err != nil {
		logtool.LogError("FlushDB ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if result.Value == "" {
		logtool.LogError("FlushDB Value Is Nil")
		return "", err
	}

	return result.Value, nil
}

// FlushDB - 清空指定DB
const (
	FlushDBID       = "FlushDB"
	FlushDBTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   FlushDB
		EVALSHA  <script_sha1> 0 {DBKey} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local sender                                        = "FlushDB.lua"
		
		if DBKey then
		
			redis.call("select",DBKey)   
			local r1= ""
			r1 = redis.call('flushdb')
			return { r1 }
		end
    `
)
