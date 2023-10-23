package src

import (
	"time"

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
	time1 := time.Now()
	// 计算时间差
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
	time2 := time.Now()

	duration := time2.Sub(time1)
	logtool.LogInfo("ScanKey Count", count)

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	logtool.LogInfof("time", "小時：%d, 分鐘：%d, 秒：%d\n", hours, minutes, seconds)

	return &result, nil
}

// ScanKey - 爬key
const (
	ScanKeyID       = "ScanKey"
	ScanKeyTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   ScanKey
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local Count                                    		= KEYS[2]
		local Type                                    		= KEYS[3]
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
