package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetHashAll function - keys, args[] string - return *[]HashResult , error
func (s *MyScriptor) GetHashAll(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(GetHashAllID, keys, args)
	if err != nil {
		logtool.LogError("GetHashAll ExecSha Error", err)
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
			logtool.LogError("GetHashAll Value Error", err)
		}
	}

	return &result, nil
}

// GetHashAll - 減少數值
const (
	GetHashAllID       = "GetHashAll"
	GetHashAllTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetHash
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local sender                                        = "GetHash.lua"
		if DBKey and ProjectKey and TagKey and k1 and k2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1= ""
			local Tmp = redis.call('hget',MAIN_KEY , k2)
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end
			return { r1 }
		end
    `
)
