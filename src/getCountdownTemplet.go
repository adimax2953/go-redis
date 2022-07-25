package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetCountDown function - keys, args[] string - return int64 , int64
func (s *MyScriptor) GetCountDown(keys, args []string) (int64, int64) {
	res, err := s.Scriptor.ExecSha(GetCountDownID, keys, args)
	if err != nil {
		logtool.LogError("GetCountDown ExecSha Error", err)
		return 0, 0
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))

	result.CountDown, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("GetCountDown Error", err.Error())
	}
	result.EndTime, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("GetCountDown Error", err.Error())
	}
	return result.CountDown, result.EndTime
}

// GetCountDown - 計數器減少
const (
	GetCountDownID       = "GetCountDown"
	GetCountDownTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetCountDown
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local sender                                        = "GetCountDown.lua"
		if not DBKey or DBKey=="" then
			return  {err="invalid argument 'DBKey'", sender=sender}
		end
		if not ProjectKey or ProjectKey=="" then
			return  {err="invalid argument 'ProjectKey'", sender=sender}
		end
		if not TagKey or TagKey=="" then
			return  {err="invalid argument 'TagKey'", sender=sender}
		end
		if not k1 or k1=="" then
			return  {err="invalid argument 'k1'", sender=sender}
		end
		if not k2 or k2=="" then
			return  {err="invalid argument 'k2'", sender=sender}
		end
		if DBKey and ProjectKey and TagKey and k1 and k2 then
			redis.call("select",DBKey)
			local r1="0"
			local r2="0"
			local Tmp = redis.call('hget',ProjectKey..":"..TagKey..":"..k1 , k2)
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = string.sub(Tmp,1,string.find(Tmp,":")-1)
				r2 = string.sub(Tmp,string.find(Tmp,":")+1)
			end
			
			return { r1 , r2 }
		end
    `
)
