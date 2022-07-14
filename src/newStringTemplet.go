package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// StringResult - 字串類型回傳
type StringResult struct {
	Value string
}

// NewString function - keys, args[] string - return string , error
func (s *MyScriptor) NewString(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(NewStringID, keys, args)
	if err != nil {
		return "", err
	}
	result := &StringResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if result.Value == "" {
		logtool.LogError("NewString err : %v\n", err)
		return "", err
	}

	return result.Value, nil
}

// NewString - 寫入一個字串
const (
	NewStringID       = "NewString"
	NewStringTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   NewString
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = ARGV[2]
		local sender                                        = "NewString.lua"
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
		if not v1 or v1=="" then
			return  {err="invalid argument 'v1'", sender=sender}
		end
		if DBKey and ProjectKey and TagKey and k1 and v1 then
			redis.call("select",DBKey)
			local r1= ""
			local Tmp = redis.call('get',ProjectKey..":"..TagKey..":"..k1)
			if Tmp==nil or Tmp=="" or Tmp==false then
				redis.call('set',ProjectKey..":"..TagKey..":"..k1 , v1)
				r1 = redis.call('get',ProjectKey..":"..TagKey..":"..k1)
			elseif Tmp~=nil and Tmp~=false then 
				r1 = ""
			end
			
			return { r1 }
		end		
    `
)
