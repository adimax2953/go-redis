package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// IncBase62 function - keys, args[] string - return string , error
func (s *MyScriptor) IncBase62(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(IncBase62ID, keys, args)
	if err != nil {
		logtool.LogError("IncBase62 ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("IncBase62 Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// IncValue - 增加數字
const (
	IncBase62ID       = "IncBase62"
	IncBase62Template = `
	--[[
		Author      :   Adimax.Tsai
		Description :   IncBase62
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local v1                                            = ARGV[3]
		local sender                                        = "IncBase62.lua"
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
		if not v1 or v1=="" then
			return  {err="invalid argument 'v1'", sender=sender}
		end

		---@return number 
		local function string_indexOf(s,pattern,init)
			init = init or 0
			local index = string.find(s,pattern,init,true)
			return index or -1
		end
		
		---@return string 
		local function Base62Inc(s)
			local chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"		
			
			local firstChar = string.sub(s,0,1)
			if string_indexOf(chars , firstChar) == 62 then
				return Base62Inc("0"..s)
			end			

			local lastChar = string.sub(s, #s, #s)
			local fragment = string.sub(s, 1, #s-1)	

			if string_indexOf(chars,lastChar) < 62 then		
			lastChar = string.sub(chars , string_indexOf(chars,lastChar) +1 , string_indexOf(chars,lastChar) +1)
				return fragment..lastChar
			end				
				return Base62Inc(fragment).."0"				
		end


		if DBKey and ProjectKey and TagKey and k1 and k2 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
			redis.call("select",DBKey)
			local result = ""

			local s = ""
			local Tmp = redis.call('hget',MAIN_KEY , k2)
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				s = Tmp
			else
				s = v1
			end


			result = Base62Inc(s)


			redis.call('hset',MAIN_KEY, k2, result)
			return {result}
		end
    `
)
