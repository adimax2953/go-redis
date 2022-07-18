package src

import (
	"strconv"

	"github.com/pkg/errors"
)

type RoomPlayer struct {
	SeatID   string `json:"SeatID"`
	UserID   string `json:"UserID"`
	Credit   string `json:"Credit"`
	BetIndex string `json:"BetIndex"`
}

type RoomJoinResult struct {
	RoomID  string       `json:"roomId"`
	SeatID  string       `json:"seatId"`
	Players []RoomPlayer `json:"players"`
}

type RoomJoinError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// RoomJoin function - args[] string - return string , error
func (s *MyScriptor) RoomJoin(
	redisDB string,
	projectID string,
	platformID string,
	gameID string,
	countryCode string,
	playerID string,
	maxPlayerCount int,
	seatsCount int,
	date string, // YYMMDD
	isBot bool,
	roomId string,
) (string, error) {
	isBotString := ""
	if isBot {
		isBotString = "isBot"
	}
	args := []string{
		redisDB,
		projectID,
		platformID,
		gameID,
		countryCode,
		playerID,
		strconv.Itoa(maxPlayerCount),
		strconv.Itoa(seatsCount),
		date,
		isBotString,
		roomId,
	}
	res, err := s.Scriptor.ExecSha(RoomJoinID, nil, args)
	if err != nil {
		return "", errors.WithStack(err)
	}

	result := res.(string)

	return result, nil
}

// RoomJoin - 減少數值
const (
	RoomJoinID       = "RoomJoin"
	RoomJoinTemplate = `
	--[[
	--]]
	-- 配房功能
	local db = ARGV[1]
	local projectId = ARGV[2]
	local platformId = ARGV[3]
	local gameId = ARGV[4]
	local currency = ARGV[5]
	local playerId = ARGV[6]
	local maxPlayerCount = tonumber(ARGV[7])
	local seatsCount = tonumber(ARGV[8])
	local date = ARGV[9] -- YYMMDD, this is used to generate roomId
	local isBot = ARGV[10] == "isBot"
	local roomId = ARGV[11]
	local scope = table.concat({projectId, platformId, gameId, currency}, "/")
	
	if maxPlayerCount == nil or not (maxPlayerCount > 0) then
		error("maxPlayerCount should be a positive integer")
	end
	
	if seatsCount == nil or not (seatsCount > 0) then
		error("seatsCount should be a positive integer")
	end
	
	if not isBot and roomId ~= "" and roomId ~= nil then
		error("roomId must be empty when isBot is false")
	end
	
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
	---@param t table
	local function hset(key, t)
		local entries = {}
		for k, v in pairs(t) do
			table.insert(entries, k)
			table.insert(entries, v)
		end
		return redis.call('HSET', key, unpack(entries))
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
	
	---@return number 
	local function getTime()
		return redis.call("TIME")[1]
	end
	
	---@return string
	local function generateRoomId()
		local key = makeKey({"roomIdCounter"})
		local currentId = redis.call("GET", key)
	
		local lastDate
		if tonumber(currentId) == nil then
			lastDate = 0
		else
			lastDate = math.floor(currentId / 1e9)
		end
	
		local id = 0
		if (tonumber(date) > lastDate) then
			id = math.floor(date * 1e9 + 1)
		else
			id = math.floor(currentId + 1)
		end
	
		id = ("%.f"):format(id) -- tostring gives you "2.0073000000001e14"
	
		redis.call("SET", key,id)
		return id
	end
	
	---@param roomId string
	---@return table
	local function getRoom(roomId)
		return hgetall(makeKey({"room", roomId}))
	end
	
	---@return table
	local function createRoom()
		local roomId = generateRoomId()
		local room = {
			roomId = roomId,
			currentPlayerCount = 0,
			currentBotCount = 0,
			maxPlayerCount = maxPlayerCount
		}
	
		redis.call(
			"SADD", projectId .. "/rooms",
				table.concat({platformId, gameId, currency, roomId}, ":")
		)
		redis.call("ZADD", makeKey({"rooms"}), getTime(), roomId)
		hset(makeKey({"room", roomId}), room)
		redis.call("ZADD", makeKey({"roomsAvailable"}), getTime(), roomId)
	
		local seats = {}
		for i = 1, seatsCount do
			table.insert(seats, tostring(i))
		end
	
		redis.call(
			"SADD", makeKey({"room", room.roomId, "seatsAvailable"}), unpack(seats)
		)
	
		redis.call(
			"PUBLISH", "RoomOpen",
				table.concat({platformId, gameId, currency, roomId, "RoomOpen"}, "~")
		)
		return room
	end
	
	---@return table
	local function getAvailableRoom()
		local roomIds = redis.call("ZRANGE", makeKey({"roomsAvailable"}), 0, 1)
		if roomIds[1] == nil then
			return nil
		end
		local room = hgetall(makeKey({"room", roomIds[1]}))
		return room
	end
	
	---@param playerId string
	---@param room table
	---@return table, string
	local function addPlayerToRoom(playerId, room)
		room.currentPlayerCount = room.currentPlayerCount + 1
		if isBot then
			room.currentBotCount = room.currentBotCount + 1
		end
		hset(makeKey({"room", room.roomId}), room)
		if room.currentPlayerCount == tonumber(room.maxPlayerCount) then
			redis.call("ZREM", makeKey({"roomsAvailable"}), room.roomId)
		end
	
		local seatId = redis.call(
			"SPOP", joinBySlash({scope, "room", room.roomId, "seatsAvailable"})
		)
	
		redis.call(
			"HSET", makeKey({"room", room.roomId, "playerToSeat"}), playerId, seatId
		)
	
		redis.call("HSET", makeKey({"playerToRoom"}), playerId, room.roomId)
	
		return room, seatId
	end
	
	---@param roomId string
	---@param playerId string
	---@param seatId string
	---@param players table
	local function publishPlayerJoinRoom(roomId, playerId, seatId, players)
		redis.call(
			"PUBLISH", "RoomUpdate", cjson.encode(
				{
					projectId = projectId,
					platformId = platformId,
					gameId = gameId,
					currency = currency,
					roomId = roomId,
					players = players
				}
			)
		)
	end
	
	---@param playerId string
	---@return string
	local function getRoomIdOfPlayer(playerId)
		return redis.call("HGET", makeKey({"playerToRoom"}), playerId)
	end
	
	---@param roomId string
	---@return table
	local function getPlayersInRoom(roomId)
		local players = hgetall(makeKey({"room", roomId, "playerToSeat"}))
		local list = {}
		for k, v in pairs(players) do
			table.insert(list, {userid = k, seatid = v, roomid = roomId})
		end
		if #list == 0 then
			return cjson.null
		end
		return list
	end
	
	redis.call("select", db)
	
	if not isBot and getRoomIdOfPlayer(playerId) then
		return cjson.encode(
			{
				error = {
					code = "PLAYER_ALREADY_IN_ROOM",
					message = "Player is already in a room."
				}
			}
		)
	end
	
	local room
	if isBot then
		room = getRoom(roomId)
		if next(room) == nil then
			return cjson.encode(
				{error = {code = "ROOM_NOT_FOUND", message = "Room is not found."}}
			)
		end
	else
		room = getAvailableRoom() or createRoom()
		log(room)
	end
	
	local room, seatId = addPlayerToRoom(playerId, room)
	local players = getPlayersInRoom(room.roomId)
	publishPlayerJoinRoom(room.roomId, playerId, seatId, players)
	
	return cjson.encode({roomId = room.roomId, seatId = seatId, players = players})	
    `
)
