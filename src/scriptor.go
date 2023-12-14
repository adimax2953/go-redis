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
	Type       string
}
type RedisType int

const (
	NONE   RedisType = iota //KEY不存在
	STRING                  //字符串
	LIST                    //列表
	SET                     //集合
	ZSET                    //有序集合
	HASH                    //哈希表
)

// Redis Type 種類
func (state RedisType) String() string {
	return [...]string{
		"none",   //KEY不存在
		"string", //字符串
		"list",   //列表
		"set",    //集合
		"zset",   //有序集合
		"hash",   //哈希表
	}[state]
}

type MyScriptor struct {
	Scriptor *goredis.Scriptor
}

var LuaScripts = map[string]string{
	//Room
	RoomJoinID:       RoomJoinTemplate,
	RoomLeftID:       RoomLeftTemplate,
	RoomLeftSingleID: RoomLeftSingleTemplate,
	RoomListID:       RoomListTemplate,
	RoomPlayerID:     RoomPlayerTemplate,
	RoomIDPlayerID:   RoomIDPlayerTemplate,
	//Value
	IncValueID:              IncValueTemplate,
	GetValueID:              GetValueTemplate,
	GetValueAllID:           GetValueAllTemplate,
	UpdateValueID:           UpdateValueTemplate,
	DecValueID:              DecValueTemplate,
	DelValueID:              DelValueTemplate,
	TakeValueID:             TakeValueTemplate,
	DecNegativeValueID:      DecNegativeValueTemplate,
	IncValueBatchID:         IncValueBatchTemplate,
	DecValueBatchID:         DecValueBatchTemplate,
	IncValueBatchFixedTTLID: IncValueBatchFixedTTLTemplate,
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
	NewHashID:         NewHashTemplate,
	GetHashID:         GetHashTemplate,
	GetHashAllID:      GetHashAllTemplate,
	UpdateHashID:      UpdateHashTemplate,
	DelHashID:         DelHashTemplate,
	DelHashAllID:      DelHashAllTemplate,
	UpdateHashListID:  UpdateHashListTemplate,
	GetHashNormalID:   GetHashNormalTemplate,
	UpdateHashTTLID:   UpdateHashTTLTemplate,
	UpdateHashBatchID: UpdateHashBatchTemplate,
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
	GetZsetRangeID:    GetZsetRangeTemplate,
	UpdateZsetID:      UpdateZsetTemplate,
	DelZsetID:         DelZsetTemplate,
	DelZsetAllID:      DelZsetAllTemplate,
	//other
	GetUUIDID:       GetUUIDTemplate,
	TTLKeyID:        TTLKeyTemplate,
	ExpireKEYID:     ExpireKEYTemplate,
	IncBase62ID:     IncBase62Template,
	ExistsKEYID:     ExistsKEYTemplate,
	FlushDBID:       FlushDBTemplate,
	ScanKeyID:       ScanKeyTemplate,
	KeyTypeID:       KeyTypeTemplate,
	SetID:           SetTemplate,
	GetID:           GetTemplate,
	SAddID:          SAddTemplate,
	SMembersID:      SMembersTemplate,
	HGetAllID:       HGetAllTemplate,
	HSetID:          HSetTemplate,
	ScanMatchKeyID:  ScanMatchKeyTemplate,
	ScanMatchKeysID: ScanMatchKeysTemplate,
	SetTTLID:        SetTTLTemplate,
	GetSystemRTPID:  GetSystemRTPTemplate,
}
