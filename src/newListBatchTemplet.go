package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// NewListBatch function - keys, args[] string - return string , error
func (s *MyScriptor) NewListBatch(keys []string, args ...interface{}) (int64, error) {

	values := make([]interface{}, len(args)-1)
	values = appendArgs(values, args)

	res, err := s.Scriptor.ExecSha(NewListBatchID, keys, values)
	if err != nil {
		logtool.LogError("NewListBatch ExecSha Error", err)
		return 0, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Count, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("NewListBatch ExecSha Count Error", err)
		return 0, err
	}
	return result.Count, nil
}

// NewListBatch - 批量新增List
const (
	NewListBatchID       = "NewListBatch"
	NewListBatchTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   NewListBatch
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {model} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                        	= KEYS[4]
		local model                                         = KEYS[5]

		local v1                                            = ARGV
		local sender                                        = "NewListBatch.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and model and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local m =""
			if model == "L" then 
				m = "lpush"
			elseif model == "R" then 
				m = "rpush"
			end
			local r1 = ""
			local tmp = redis.call(m,MAIN_KEY ,unpack(v1))
			
			return { tmp }
		end
    `
)
