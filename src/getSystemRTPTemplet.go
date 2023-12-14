package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetSystemRTP function - keys, args[] string - return map[string]interface{}, error
func (s *MyScriptor) GetSystemRTP(keys, args []string) (map[string]interface{}, error) {
	res, err := s.Scriptor.ExecSha(GetSystemRTPID, keys, args)
	if err != nil {
		logtool.LogError("GetSystemRTP ExecSha Error", err)
		return nil, err
	}

	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{})) / 2
	//result := make([]RedisResult, count)
	myMap := make(map[string]interface{}, count)

	for i := 0; i < count; i++ {
		myMap[reader.ReadString()] = reader.ReadString()
		/*r := &RedisResult{}
		r.Key = reader.ReadString()
		r.Value = reader.ReadString()
		result[i] = *r
		if err != nil {
			logtool.LogError("GetSystemRTP Value Error", err)
		}*/
	}

	return myMap, nil
}

// GetSystemRTP - 獲取系統RTP
const (
	GetSystemRTPID       = "GetSystemRTP"
	GetSystemRTPTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetSystemRTP
		EVALSHA  <script_sha1> 0 {DBKey}{ARGV} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local v1                                            = ARGV
		local sender                                        = "GetSystemRTP.lua"
		local values = {}	

		---@param mainkey string
		---@return string[]
		local function hgetall(mainkey)
			local Tmp = redis.call('hgetall', mainkey)	
			for i, key in ipairs(Tmp) do
				if i%2==0 then
					table.insert(values, key)
				else
					table.insert(values, mainkey..":"..key)
				end
			end
		end

		redis.call('select',DBKey)
		local keys = {unpack(v1)}

		for _, key in ipairs(keys) do		
			hgetall(key)

		end

		return values
    `
)
