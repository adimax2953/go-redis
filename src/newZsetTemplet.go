package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// NewZset function - keys, args[] string - return string , error
func (s *MyScriptor) NewZset(keys, args []string) (*RedisResult, error) {
	res, err := s.Scriptor.ExecSha(NewZsetID, keys, args)
	if err != nil {
		logtool.LogError("NewZset ExecSha Error", err)
		return nil, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	result.Value2 = reader.ReadString()

	return result, nil
}

// NewZset - 減少數值
const (
	NewZsetID       = "NewZset"
	NewZsetTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   NewZset
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1} {v2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = tonumber(ARGV[2])
		local v2                                            = ARGV[3]
		local sender                                        = "NewZset.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and v1 and v2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			redis.call('zadd',MAIN_KEY ,v1,v2)
			local tmp = redis.call('zscore',MAIN_KEY ,v2)
			local r2="0"
			local r1="0"
			if tmp~=nil and tmp~="" and tmp~=false then
				r1= tmp
				r2= v2
			end
			return { r1 , tostring(r2) }
		end
    `
)
