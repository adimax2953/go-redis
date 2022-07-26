package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetValue function - keys, args[] string - return int64 , error
func (s *MyScriptor) GetValue(keys, args []string) (int64, error) {
	res, err := s.Scriptor.ExecSha(GetValueID, keys, args)
	if err != nil {
		logtool.LogError("GetValue ExecSha Error", err)
		return 0, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.ValueInt64, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("GetValue Value Error", err)
		return 0, err
	}

	return result.ValueInt64, nil
}

// GetValue - 取得數字
const (
	GetValueID       = "GetValue"
	GetValueTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetValue
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local sender                                        = "GetValue.lua"
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
		if DBKey and ProjectKey and TagKey and k1 and k2 then
			redis.call("select",DBKey)
			local result ={}
			result = redis.call('hget',ProjectKey..":"..TagKey..":"..k1,k2)
			return {result}
		end
    `
)
