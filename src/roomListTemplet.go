package src

import (
	logtool "github.com/adimax2953/log-tool"
	"github.com/pkg/errors"
)

// RoomList function - keys, args[] string - return interface{} , error
func (s *MyScriptor) RoomList(
	keys []string,
) (interface{}, error) {

	res, err := s.Scriptor.ExecSha(RoomListID, keys, nil)
	if err != nil {
		logtool.LogError("RoomList ExecSha Error", err)
		return "", errors.WithStack(err)
	}

	return res, nil
}

// RoomList - 減少數值
const (
	RoomListID       = "RoomList"
	RoomListTemplate = `
	local DBKey                                         = tonumber(KEYS[1])
	local ProjectKey                                    = KEYS[2]
	local TagKey                                        = KEYS[3]
	
	local function log(v)
		local s = ""
		if type(v) == 'table' then
			s = cjson.encode(v)
		else
			s = tostring(v)
		end
		redis.call("RPUSH", "log", s)
	end
	
	---@param key string
	---@return table
	local function hgetall(key)
		local entries = redis.call('HGETALL', key)
		local result = {}
		for i = 1, #entries, 2 do
			result[entries[i]] = entries[i + 1]
		end
		return result
	end
	
	---@param str string
	---@param separator string
	---@return string[]
	local function split(str, separator)
		local t = {}
		for s in string.gmatch(str, "([^" .. separator .. "]+)") do
			table.insert(t, s)
		end
		return t
	end
	
	redis.call("SELECT", DBKey)
	
	local roomIds = redis.call("SMEMBERS", ProjectKey .. "/rooms")
	
	if roomIds == false or next(roomIds) == nil then
		return "[]"
	end
	
	local rooms = {}
	for _, roomId in ipairs(roomIds) do
		local parts = split(roomId, "/")
		local platformId = parts[1]
		local gameId = parts[2]
		local countrycode = parts[3]
		local roomIdLocal = parts[4]
		local key = table.concat(
			{ProjectKey, platformId, gameId, countrycode, "room", roomIdLocal}, "/"
		)
		local room = hgetall(key)
		table.insert(rooms, room)
	end
	
	return cjson.encode(rooms)
	
    `
)
