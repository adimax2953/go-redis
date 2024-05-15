package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// ZRange function - keys, args[] string - return string , error
func (s *MyScriptor) ZRange(keys, args []string) ([]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(ZRangeID, keys, args)
	if err != nil {
		logtool.LogError("ZRange ExecSha Error", err)
		return nil, err
	}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{})) / 2
	result := make([]RedisResult, count)

	for i := 0; i < count; i++ {
		r := &RedisResult{}
		r.Value = reader.ReadString()
		r.ValueInt64, err = reader.ReadInt64(0)
		if err != nil {
			logtool.LogError("ZRange Error", err.Error())
		}
		result[i] = *r
		if err != nil {
			logtool.LogError("ZRange Value Error", err)
		}
	}

	return result, nil
}

// ZRange - 爬key
const (
	ZRangeID       = "ZRange"
	ZRangeTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   ZRange
		EVALSHA  <script_sha1> 0 {DBKey} {Key} {Value}  {Score} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local Key                                    		= KEYS[2]
		local sender                                        = "ZRange.lua"
		
		redis.call('select',DBKey)	
		local result = redis.call('zrange',Key ,0 ,-1, 'withscores')

		return result
    `
)
