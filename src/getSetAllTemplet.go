package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetSetAll function - keys, args[] string - return *[]HashResult , error
func (s *MyScriptor) GetSetAll(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(GetSetAllID, keys, args)
	if err != nil {
		logtool.LogError("GetSetAll ExecSha Error", err)
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
			logtool.LogError("GetSetAll Value Error", err)
		}
	}
	return &result, nil
}

// GetSetAll - 減少數值
const (
	GetSetAllID       = "GetSetAll"
	GetSetAllTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetSetAll
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		
		local sender                                        = "GetSetAll.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local tmp = ""
			tmp = redis.call('smembers',MAIN_KEY )
			
			return tmp 
		end
    `
)
