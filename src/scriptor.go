package src

import goredis "github.com/adimax2953/go-redis"

// RedisResult -
type RedisResult struct {
	Value      string
	Value2     string
	CountDown  int64
	EndTime    int64
	ValueInt64 int64
	Key        string
}

type MyScriptor struct {
	Scriptor *goredis.Scriptor
}

var LuaScripts = map[string]string{
	//Room
	RoomJoinID:   RoomJoinTemplate,
	RoomLeftID:   RoomLeftTemplate,
	RoomListID:   RoomListTemplate,
	RoomPlayerID: RoomPlayerTemplate,
	//Value
	IncValueID:         IncValueTemplate,
	GetValueID:         GetValueTemplate,
	GetValueAllID:      GetValueAllTemplate,
	UpdateValueID:      UpdateValueTemplate,
	DecValueID:         DecValueTemplate,
	DelValueID:         DelValueTemplate,
	TakeValueID:        TakeValueTemplate,
	DecNegativeValueID: DecNegativeValueTemplate,
	//CountDown
	IncCountDownID: IncCountDownTemplate,
	GetCountDownID: GetCountDownTemplate,
	DecCountDownID: DecCountDownTemplate,
	DelCountDownID: DelCountDownTemplate,
	//String
	NewStringID:       NewStringTemplate,
	UpdateStringID:    UpdateStringTemplate,
	GetStringID:       GetStringTemplate,
	DelStringID:       DelStringTemplate,
	UpdateTTLStringID: UpdateTTLStringTemplate,
	//Hash
	NewHashID:        NewHashTemplate,
	GetHashID:        GetHashTemplate,
	GetHashAllID:     GetHashAllTemplate,
	UpdateHashID:     UpdateHashTemplate,
	DelHashID:        DelHashTemplate,
	DelHashAllID:     DelHashAllTemplate,
	UpdateHashListID: UpdateHashListTemplate,
	GetHashNormalID:  GetHashNormalTemplate,
	UpdateHashTTLID:  UpdateHashTTLTemplate,
	//List
	NewListID:       NewListTemplate,
	GetListID:       GetListTemplate,
	GetListPopID:    GetListPopTemplate,
	GetListAllID:    GetListAllTemplate,
	GetListLengthID: GetListLengthTemplate,
	UpdateListID:    UpdateListTemplate,
	DelListID:       DelListTemplate,
	DelListAllID:    DelListAllTemplate,
	//Set
	NewSetID:       NewSetTemplate,
	GetSetID:       GetSetTemplate,
	GetSetRandomID: GetSetRandomTemplate,
	GetSetPopID:    GetSetPopTemplate,
	GetSetAllID:    GetSetAllTemplate,
	UpdateSetID:    UpdateSetTemplate,
	DelSetID:       DelSetTemplate,
	DelSetAllID:    DelSetAllTemplate,
	//Zset
	NewZsetID:         NewZsetTemplate,
	GetZsetID:         GetZsetTemplate,
	GetZsetAllID:      GetZsetAllTemplate,
	GetZsetAllRevID:   GetZsetAllRevTemplate,
	GetZsetRankID:     GetZsetRankTemplate,
	GetZsetCountID:    GetZsetCountTemplate,
	GetZsetAllCountID: GetZsetAllCountTemplate,
	GetZsetRangeID:GetZsetRangeTemplate,
	UpdateZsetID:      UpdateZsetTemplate,
	DelZsetID:         DelZsetTemplate,
	DelZsetAllID:      DelZsetAllTemplate,
	//other
	GetUUIDID:   GetUUIDTemplate,
	TTLKeyID:    TTLKeyTemplate,
	ExpireKEYID: ExpireKEYTemplate,
	IncBase62ID: IncBase62Template,
}
