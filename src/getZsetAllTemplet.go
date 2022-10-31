package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetZsetAll function - keys, args[] string - return int64 , error
func (s *MyScriptor) GetZsetAll(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(GetZsetAllID, keys, args)
	if err != nil {
		logtool.LogError("GetZsetAll ExecSha Error", err)
		return nil, err
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
			logtool.LogError("GetZsetAll Value Error", err)
		}
	}

	return &result, nil
}

// GetZsetAll - 寫入一個字串
const (
	GetZsetAllID       = "GetZsetAll"
	GetZsetAllTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetZsetAll
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = tonumber(ARGV[2])
		local k3                                            = tonumber(ARGV[3])
		local sender                                        = "GetZsetAll.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			
			local tmp = redis.call('zrange',MAIN_KEY , k2, k3)
		
			return tmp
		end
	`
)
