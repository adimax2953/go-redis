package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// HSet function - keys []string,  args ...interface{} - return *[]RedisResult , error
func (s *MyScriptor) HSet(keys []string, args ...interface{}) (*[]RedisResult, error) {

	values := make([]interface{}, len(args)-1)
	values = appendArgs(values, args)

	res, err := s.Scriptor.ExecSha(HSetID, keys, values)
	if err != nil {
		logtool.LogError("HSet ExecSha Error", err)
		return nil, err
	}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{})) / 2
	result := make([]RedisResult, count)

	for i := 0; i < count; i++ {
		r := &RedisResult{}
		r.Key = reader.ReadString()
		r.Value = reader.ReadString()
		result[i] = *r
		if err != nil {
			logtool.LogError("HSet Value Error", err)
		}
	}

	return &result, nil
}

// HSet - 批量更新Hash
const (
	HSetID       = "HSet"
	HSetTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   HSet
		EVALSHA  <script_sha1> 0 {DBKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local k1                                            = KEYS[2]
		local v1                                            = ARGV
		local sender                                        = "HSet.lua"		

		if DBKey and k1 and v1 then
		
			local MAIN_KEY = k1
			redis.call("select",DBKey)
			redis.call('hset', k1 , unpack(v1))

			local r1 = ""
			local Tmp = redis.call('hgetall',MAIN_KEY)
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end
			return r1
		end
    `
)
