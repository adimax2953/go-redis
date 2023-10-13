package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// IncValueBatchFixedTTL function - keys []string,  args ...interface{} - return *[]RedisResult , error
func (s *MyScriptor) IncValueBatchFixedTTL(keys []string, args ...interface{}) (*[]RedisResult, error) {
	values := make([]interface{}, len(args)-1)
	values = appendArgs(values, args)

	res, err := s.Scriptor.ExecSha(IncValueBatchFixedTTLID, keys, values)
	if err != nil {
		logtool.LogError("IncValueBatchFixedTTL ExecSha Error", err)
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
			logtool.LogError("IncValueBatchFixedTTL Value Error", err)
		}
	}

	return &result, nil
}

// IncValueBatchFixedTTL - 批量增加數字(設定一個固定的TTL)
const (
	IncValueBatchFixedTTLID       = "IncValueBatchFixedTTL"
	IncValueBatchFixedTTLTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   IncValueBatchFixedTTL
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = KEYS[4]
		local k2                                            = tonumber(KEYS[5])
		local v1                                            = ARGV
		local sender                                        = "IncValueBatch.lua"
		
		---@return number 
		local function getTime()
			return redis.call("TIME")[1]
		end

		if DBKey and ProjectKey and TagKey and k1 and v1 then

			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
			redis.call("select",DBKey)

			local count = #v1
			for i = 1,count,2 do
				redis.call('hincrby',MAIN_KEY , v1[i] ,v1[i+1])
			end

			redis.call('hset',MAIN_KEY, 'lastUpdateTime', getTime())
			
			local r1 = ""
			local Tmp = redis.call('hgetall',MAIN_KEY)
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end

			local r = redis.call('ttl',MAIN_KEY)
			if r == -1 and k2 then
				redis.call('expire',MAIN_KEY,k2)
			end

			return r1
		end
    `
)
