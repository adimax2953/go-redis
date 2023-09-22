package src

import (
	logtool "github.com/adimax2953/log-tool"
	"github.com/pkg/errors"
)

// RoomIDPlayer function - keys, args[] string - return interface{} , error
func (s *MyScriptor) RoomIDPlayer(
	keys []string,
	platformID string,
	gameID string,
	countryCode string,
	roomID string,
) (interface{}, error) {
	args := []string{
		platformID,
		gameID,
		countryCode,
		roomID,
	}
	res, err := s.Scriptor.ExecSha(RoomIDPlayerID, keys, args)
	if err != nil {
		logtool.LogError("RoomIDPlayerID ExecSha Error", err)
		return "", errors.WithStack(err)
	}

	return res, nil
}

// RoomIDPlayer - 查詢特定房間玩家
const (
	RoomIDPlayerID       = "RoomIDPlayer"
	RoomIDPlayerTemplate = `
	local DBKey                                         = tonumber(KEYS[1])
	local ProjectKey                                    = KEYS[2]
	local TagKey                                        = KEYS[3]
	local platformId = ARGV[1]
	local gameId = ARGV[2]
	local countrycode = ARGV[3]
	local roomId = ARGV[4]

	local scope = table.concat({ProjectKey, platformId, gameId, countrycode,"room",roomId}, "/")

	local function log(v)
		local s = ""
		if type(v) == 'table' then
			s = cjson.encode(v)
		else
			s = tostring(v)
		end
		redis.call("RPUSH", "log", s)
	end
	
	---@param list table
	---@return string
	local function joinBySlash(list)
		return table.concat(list, "/")
	end
	
	---@param list table
	---@return string
	local function makeKey(list)
		return joinBySlash({scope, unpack(list)})
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
		
	redis.call("SELECT", DBKey)
	local playerlist = 	hgetall(makeKey({"playerToSeat"}))
	
	if playerlist == false or next(playerlist) == nil then
		return "[]"
	end
		
	return cjson.encode(playerlist)	
    `
)
