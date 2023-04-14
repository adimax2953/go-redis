package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// UpdateHashBatch function - keys []string,  args ...interface{} - return *[]RedisResult , error
func (s *MyScriptor) UpdateHashBatch(keys []string, args ...interface{}) (*[]RedisResult, error) {

	values := make([]interface{}, len(args)-1)
	values = appendArgs(values, args)

	res, err := s.Scriptor.ExecSha(UpdateHashBatchID, keys, values)
	if err != nil {
		logtool.LogError("UpdateHashBatch ExecSha Error", err)
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
			logtool.LogError("UpdateHashBatch Value Error", err)
		}
	}

	return &result, nil
}

// UpdateHashBatch - 批量更新Hash
const (
	UpdateHashBatchID       = "UpdateHashBatch"
	UpdateHashBatchTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   UpdateHashBatch
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = KEYS[4]
		local v1                                            = ARGV
		local sender                                        = "UpdateHashBatch.lua"		

		if DBKey and ProjectKey and TagKey and k1 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			redis.call('hset', MAIN_KEY , unpack(v1))

			local r1 = ""
			local Tmp = redis.call('hgetall',MAIN_KEY)
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end
			return r1
		end
    `
)
