package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// ScanMatchKey function - keys, args[] string - return *[]RedisResult , error
func (s *MyScriptor) ScanMatchKey(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(ScanMatchKeyID, keys, args)
	if err != nil {
		logtool.LogError("ScanMatchKey ExecSha Error", err)
		return nil, err
	}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{}))
	result := make([]RedisResult, count)
	for i := 0; i < count; i++ {
		r := &RedisResult{}
		r.Key = reader.ReadString()
		result[i] = *r
		if err != nil {
			logtool.LogError("ScanMatchKey Value Error", err)
		}
	}

	return &result, nil
}

// ScanMatchKey - 爬key
const (
	ScanMatchKeyID       = "ScanMatchKey"
	ScanMatchKeyTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   ScanMatchKey
		EVALSHA  <script_sha1> 0 {DBKey} {Key} {Count} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local Key                                           = KEYS[2]
		local Count                                    		= tonumber(KEYS[3])
		local sender                                        = "ScanMatchKey.lua"

		local values = {}	

		redis.call('select',DBKey)
		
		local result = redis.call('scan', 0, 'match' , Key.."*" ,'count', Count)
		local keys = result[2]

		for _, key in ipairs(keys) do
			table.insert(values, key)
		end

		return values
    `
)
