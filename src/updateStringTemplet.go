package src

import (
	"log"

	"github.com/adimax2953/go-redis/goredis"
)

// UpdateString function - keys, args[] string - return string , eror
func (s *MyScriptor) UpdateString(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(NewStringID, keys, args)
	if err != nil {
		return "", err
	}
	result := &StringResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()

	if result.Value == "-1" {
		log.Printf("UpdateString err : %v\n", err)
		return "", err
	}

	return result.Value, nil
}

// UpdateString - 寫入一個字串
const (
	UpdateStringID       = "UpdateString"
	UpdateStringTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   UpdateString
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = ARGV[2]
		local sender                                        = "UpdateString.lua"
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
				redis.call('set',ProjectKey..":"..TagKey..":"..k1 , v1)
				r1 = redis.call('get',ProjectKey..":"..TagKey..":"..k1)
			return { r1 }
		end				
    `
)
