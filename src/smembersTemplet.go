package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// SMembers function - keys, args[] string - return *[]RedisResult , error
func (s *MyScriptor) SMembers(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(SMembersID, keys, args)
	if err != nil {
		logtool.LogError("SMembers ExecSha Error", err)
		return nil, err
	}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{}))
	result := make([]RedisResult, count)
	for i := 0; i < count; i++ {
		r := &RedisResult{}
		r.Value = reader.ReadString()
		result[i] = *r
		if err != nil {
			logtool.LogError("SMembers Value Error", err)
		}
	}

	return &result, nil
}

// SMembers - 爬key
const (
	SMembersID       = "SMembers"
	SMembersTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   SMembers
		EVALSHA  <script_sha1> 0 {DBKey} {Key} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local Key                                    		= KEYS[2]
		local sender                                        = "SMembers.lua"
		
		redis.call('select',DBKey)	
		local result = redis.call('smembers',Key )
		return result
    `
)
