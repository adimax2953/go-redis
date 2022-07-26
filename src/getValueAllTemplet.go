package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetValueAll function - keys, args[] string - return *[]HashResult , error
func (s *MyScriptor) GetValueAll(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(GetValueAllID, keys, args)
	if err != nil {
		logtool.LogError("GetValueAll ExecSha Error", err)
		return nil, err
	}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{})) / 2
	result := make([]RedisResult, count)

	for i := 0; i < count; i++ {
		r := &RedisResult{}
		r.Key = reader.ReadString()
		r.ValueInt64, err = reader.ReadInt64(0)
		result[i] = *r
		if err != nil {
			logtool.LogError("GetValueAll Value Error", err)
		}
	}

	return &result, nil
}

// GetValueAll - 取得數值陣列
const (
	GetValueAllID       = "GetValueAll"
	GetValueAllTemplate = `
	--[[
		Author      :   Adima.Tsai
		Description :   GetValueAll
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "GetValueAll.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1= "0"
			local Tmp = redis.call('hgetall',MAIN_KEY )
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end
			return r1
		end
    `
)
