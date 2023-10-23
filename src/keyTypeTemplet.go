package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// KeyType function - keys, args[] string - return string , error
func (s *MyScriptor) KeyType(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(ScanKeyID, keys, args)
	if err != nil {
		logtool.LogError("KeyType ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Type = reader.ReadString()
	if err != nil {
		logtool.LogError("UpdateHash Value Error", err)
		return "", err
	}
	return result.Type, nil
}

// KeyType - 爬key
const (
	KeyTypeID       = "KeyType"
	KeyTypeTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   KeyType
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} 
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local sender                                        = "KeyType.lua"

		local values = {}			
		redis.call('select',DBKey)
		local types = redis.call('type', 'key')
		table.insert(values, types)
		return values
    `
)
