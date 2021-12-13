package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (s *RedisSuite) TestRedis() {
	rdb := RedisConn

	err := rdb.Set(rCtx, "test", "hi-redis", 0).Err()
	s.Nil(err)

	val, err := rdb.Get(rCtx, "test").Result()
	s.Nil(err)
	s.Equal("hi-redis", val)

	_, err = rdb.Get(rCtx, "missing-key-123").Result()
	s.EqualValues(redis.Nil, err)
}

func TestRedisSuite(t *testing.T) {
	suite.Run(t, &RedisSuite{})
}

type RedisSuite struct {
	suite.Suite
}
