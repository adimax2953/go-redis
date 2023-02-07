package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetZsetRange function - keys, args[] string - return int64 , error
func (s *MyScriptor) GetZsetRange(keys, args []string) (int64, error) {
	res, err := s.Scriptor.ExecSha(GetZsetRangeID, keys, args)
	if err != nil {
		logtool.LogError("GetZsetRank ExecSha Error", err)
		return -1, err
	}

	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{}))
	result := make([]RedisResult, count)

	for i := 0; i < count; i++ {
		r := &RedisResult{}
		r.Value = reader.ReadString()
		// r.Value2 = reader.ReadString()
		result[i] = *r
		if err != nil {
			logtool.LogError("GetZsetRange Value Error", err)
		}
	}

	return rank, nil
}

// GetZsetRange - 寫入一個字串
const (
	GetZsetRangeID       = "GetZsetRange"
	GetZsetRangeTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetZsetRange
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = tonumber(ARGV[2])
		local k3                                            = tonumber(ARGV[3])
		local sender                                        = "GetZsetRange.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			
			local tmp = redis.call('zrange',MAIN_KEY , k2 ,k3 ,'withscores')
			return tmp
		end
	`
)
