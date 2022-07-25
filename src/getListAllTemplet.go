package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetListAll function - keys, args[] string - return *[]RedisResult] , error
func (s *MyScriptor) GetListAll(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(GetListAllID, keys, args)
	if err != nil {
		logtool.LogError("GetListAll ExecSha Error", err)
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
			logtool.LogError("GetListAll Value Error", err)
		}
	}
	return &result, nil
}

// GetListAll - 減少數值
const (
	GetListAllID       = "GetListAll"
	GetListAllTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetListAll
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "GetListAll.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1 = {""}
			local tmp = redis.call('lrange',MAIN_KEY,0,-1)
			if tmp~= nil and tmp~= false then
				r1= tmp
			else
				r1= {""}
			end
			return r1
		end
    `
)
