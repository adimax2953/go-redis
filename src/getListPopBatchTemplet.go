package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetListPopBatch function - keys, args[] string - return RedisResult , error
// model :  L lpush list前 // R rpush list後
func (s *MyScriptor) GetListPopBatch(keys, args []string) ([]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(GetListPopBatchID, keys, args)
	if err != nil {
		logtool.LogError("GetListPopBatch ExecSha Error", err)
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
			logtool.LogError("GetListPopBatch Value Error", err)
		}
	}
	return result, nil
}

// GetListPop - 取出List的資料
const (
	GetListPopBatchID       = "GetListPopBatch"
	GetListPopBatchTemplate = `
	  --[[
		  Author      :   Adimax.Tsai
		  Description :   GetListPopBatch
		  EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {model}
		  --]]
		  local DBKey                                         = tonumber(KEYS[1])
		  local ProjectKey                                    = KEYS[2]
		  local TagKey                                        = KEYS[3]
		  local k1                                            = KEYS[4]
		  local model                                         = KEYS[5]
		  local count                                         = tonumber(KEYS[6])
		  local sender                                        = "GetListPopBatch.lua"
		  
		  if DBKey and ProjectKey and TagKey and model and count then
			  local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		  
			  redis.call("select",DBKey)
  
			  local length = redis.call("LLEN", MAIN_KEY)
			  if length < count then
				  count = length
			  end
				  redis.log(redis.LOG_NOTICE,length)
  
			  local start = 0
			  local send = count - 1
			  local trstart = count
			  local trend = -1
			  local rpop = false
  
			  if model == "R" then 
				  start = length - count
				  send = length - 1
				  trstart = 0 
				  trend = start - 1
				  if trend < 0 then
					  trend = 0
					  rpop = true
				  end
			  end
  
			  local result = redis.call("LRANGE", MAIN_KEY, start, send)
			  redis.call("LTRIM", MAIN_KEY, trstart , trend)
			  if rpop then
				  redis.call("RPOP", MAIN_KEY)
			  end
  
			  return  result 
		  end
	  `
)

// const (
// 	GetListPopBatchID       = "GetListPopBatch"
// 	GetListPopBatchTemplate = `
// 	--[[
// 		Author      :   Adimax.Tsai
// 		Description :   GetListPopBatch
// 		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {model}
// 		--]]
// 		local DBKey                                         = tonumber(KEYS[1])
// 		local ProjectKey                                    = KEYS[2]
// 		local TagKey                                        = KEYS[3]
// 		local k1                                            = KEYS[4]
// 		local model                                         = KEYS[5]
// 		local count                                         = tonumber(KEYS[6])
// 		local sender                                        = "GetListPopBatch.lua"

// 		if DBKey and ProjectKey and TagKey and model and count then
// 			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1

// 			redis.call("select",DBKey)

// 			local length = redis.call("LLEN", MAIN_KEY)
// 			if length < count then
// 				count = length
// 			end
// 				redis.log(redis.LOG_NOTICE,length)

// 			local start = length - count
// 			local send = length

// 			if model == "L" then
// 				start = 0
// 				send = count -1
// 			end

// 			local result = redis.call("LRANGE", MAIN_KEY, start, send)
//       		redis.call("LTRIM", MAIN_KEY, start , send)

// 			return  result
// 		end
//     `
// )
