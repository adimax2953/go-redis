package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetZsetAll function - keys, args[] string - return int64 , error
func (s *MyScriptor) GetZsetRank(keys, args []string) (int64, error) {
	res, err := s.Scriptor.ExecSha(GetZsetRankID, keys, args)
	if err != nil {
		logtool.LogError("GetZsetRank ExecSha Error", err)
		return -1, err
	}

	rank, err := goredis.NewRedisReplyValue(res).AsInt64(0)
	if err != nil {
		logtool.LogError("GetZsetRank Value Error", err)
		return -1, err
	}

	return rank, nil
}

// GetZsetAll - 寫入一個字串
const (
	GetZsetRankID       = "GetZsetRank"
	GetZsetRankTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetZsetRank
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local sender                                        = "GetZsetRank.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			
			local tmp = redis.call('zrank',MAIN_KEY , k2)
		
			return tmp
		end
	`
)
