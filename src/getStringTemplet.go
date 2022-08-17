package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetString function - keys, args[] string - return string , error
func (s *MyScriptor) GetString(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(GetStringID, keys, args)
	if err != nil {
		logtool.LogError("GetString ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if result.Value == "" {
		logtool.LogError("GetString Value Is Nil")
		return "", err
	}

	return result.Value, nil
}

// GetString - 寫入一個字串
const (
	GetStringID       = "GetString"
	GetStringTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetString
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "GetString.lua"
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
		if DBKey and ProjectKey and TagKey and k1 then
			redis.call("select",DBKey)
			local r1= ""
			local Tmp = redis.call('get',ProjectKey..":"..TagKey..":"..k1)
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end
			return { r1 }
		end				
    `
)
