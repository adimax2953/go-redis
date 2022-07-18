package src

import (
	"strconv"

	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetUUID function - StringCount int64 - return string , error
func (s *MyScriptor) GetUUID(stringCount int64) (string, error) {
	args := []string{strconv.FormatInt(stringCount, 10)}

	res, err := s.Scriptor.ExecSha(GetUUIDID, nil, args)
	if err != nil {
		logtool.LogError("GetUUID ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if result.Value == "" {
		logtool.LogError("GetUUID Value Is Nil")
		return "", err
	}

	return result.Value, nil
}

// GetUUID - 寫入一個字串
const (
	GetUUIDID       = "GetUUID"
	GetUUIDTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetUUID
		EVALSHA  <script_sha1> 0 {k1} 
		--]]
		local n  = ARGV[1]

		local time = redis.call('TIME')
		math.randomseed(time)
		
		local t = {
			"0","1","2","3","4","5","6","7","8","9",
			"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z",
			"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z",    
		}
		local s = ""
		for i =1, n do
			if i==9 or i==14 or i==19 or i==24 then
				s = s.."-" 
				else 
				s = s .. t[math.random(#t)]     
				end  
		end
		
		return {s}				
    `
)
