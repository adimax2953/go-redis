package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// ScanKey function - keys, args[] string - return *[]RedisResult , error
func (s *MyScriptor) ScanKey(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(ScanKeyID, keys, args)
	if err != nil {
		logtool.LogError("ScanKey ExecSha Error", err)
		return nil, err
	}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{}))
	result := make([]RedisResult, count)
	for i := 0; i < count; i++ {
		r := &RedisResult{}
		r.Key = reader.ReadString()
		result[i] = *r
		if err != nil {
			logtool.LogError("ScanKey Value Error", err)
		}
	}

	return &result, nil
}

// ScanKey - 爬key
const (
	ScanKeyID       = "ScanKey"
	ScanKeyTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   ScanKey
		EVALSHA  <script_sha1> 0 {DBKey} {Count} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local Count                                    		= KEYS[2]
		local sender                                        = "ScanKey.lua"

		local values = {}	


		---@param cursor string
		---@return string[]
		local function scan_keys(cursor)
			local result = redis.call('scan', cursor)
			cursor = tonumber(result[1])
			local keys = result[2]

			for _, key in ipairs(keys) do
				table.insert(values, key)
			end

			if cursor == 0 then
				return keys
			else
				return scan_keys(cursor)
			end
		end
		
		redis.call('select',DBKey)	
		scan_keys(Count)
		return values
    `
)
