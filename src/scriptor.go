package src

import goredis "github.com/adimax2953/go-redis"

// RedisResult -
type RedisResult struct {
	Value     string
	Value2    string
	CountDown int64
	EndTime   int64
}

type MyScriptor struct {
	Scriptor *goredis.Scriptor
}
