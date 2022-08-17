package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetSetRandom function - keys, args[subkey,request datacount] string - return *[]HashResult , error
func (s *MyScriptor) GetSetRandom(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(GetSetRandomID, keys, args)
	if err != nil {
		logtool.LogError("GetSetRandom ExecSha Error", err)
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
			logtool.LogError("GetSetRandom Value Error", err)
		}
	}
	return &result, nil
}

// GetSetRandom - 減少數值
const (
	GetSetRandomID       = "GetSetRandom"
	GetSetRandomTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetSetRandom
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local sender                                        = "GetSetRandom.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and k2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local tmp = "-1"
			tmp = redis.call('srandmember',MAIN_KEY , k2)
			
			return tmp 
		end
    `
)
