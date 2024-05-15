package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// IncValueBefore function - keys, args[] string - return int64 , error
func (s *MyScriptor) IncValueBefore(keys, args []string) (int64, int64, error) {
	res, err := s.Scriptor.ExecSha(IncValueBeforeID, keys, args)
	if err != nil {
		logtool.LogError("IncValueBefore ExecSha Error", err)
		return 0, 0, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.ValueInt64, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("IncValueBefore ValueInt64 Error", err)
		return 0, 0, err
	}
	result.Value2Int64, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("IncValueBefore Value2Int64 Error", err)
		return 0, 0, err
	}
	return result.ValueInt64, result.Value2Int64, nil
}

// IncValueBefore - 增加數字且查詢原始數字
const (
	IncValueBeforeID       = "IncValueBefore"
	IncValueBeforeTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   IncValueBefore
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local v1                                            = tonumber(ARGV[3])
		local sender                                        = "IncValueBefore.lua"
		if not DBKey or DBKey=="" then
			return  {err="invalid argument 'DBKey'", sender=sender}
		end
		if not ProjectKey or ProjectKey=="" then
			return  {err="invalid argument 'ProjectKey'", sender=sender}
		end
		if not TagKey or TagKey=="" then
			return  {err="invalid argument 'TagKey'", sender=sender}
		end
		if not k1 or k1=="" then
			return  {err="invalid argument 'k1'", sender=sender}
		end
		if not k2 or k2=="" then
			return  {err="invalid argument 'k2'", sender=sender}
		end
		if not v1 or v1=="" then
			return  {err="invalid argument 'v1'", sender=sender}
		end
		
		---@return number 
		local function getTime()
			return redis.call("TIME")[1]
		end

		if DBKey and ProjectKey and TagKey and k1 and k2 and v1 then

			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
			redis.call("select",DBKey)
			local afther ={}
			local before ={}
				before = redis.call('hget',MAIN_KEY,k2)
				afther = redis.call('hincrby',MAIN_KEY,k2,v1)
				redis.call('hset',MAIN_KEY, 'lastUpdateTime', getTime())

			return {before,afther}
		end
    `
)
