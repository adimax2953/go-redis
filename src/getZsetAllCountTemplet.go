package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// DelSet function - keys, args[] string
func (s *MyScriptor) GetZsetAllCount(keys, args []string) (int64, error) {

	res, err := s.Scriptor.ExecSha(GetZsetAllCountID, keys, args)
	if err != nil {
		logtool.LogError("GetZsetAllCount ExecSha Error", err)
		return 0, err
	}

	// goredis.EmptyRedisReplyValue.AsInt32()
	count, err := goredis.NewRedisReplyValue(res).AsInt64(0)
	if err != nil {
		logtool.LogError("GetZsetAllCount Value Error", err)
		return 0, err
	}
	// reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	// count := len(res.([]interface{}))
	// result := make([]RedisResult, count)

	// for i := 0; i < count; i++ {
	// 	r := &RedisResult{}
	// 	r.Value = reader.ReadString()
	// 	// r.Value2 = reader.ReadString()
	// 	result[i] = *r
	// 	if err != nil {
	// 		logtool.LogError("GetZsetAll Value Error", err)
	// 	}
	// }

	return count, nil
}

// GetZsetAll - 寫入一個字串
const (
	GetZsetAllCountID       = "GetZsetAllCount"
	GetZsetAllCountTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetZsetAllCount
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "GetZsetAllCount.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			
			local tmp = redis.call('zcard',MAIN_KEY)
		
			return tmp
		end
	`
)
