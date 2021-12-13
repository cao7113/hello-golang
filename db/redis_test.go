package db

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
	"testing"
)

func (s *RedisSuite) TestRedis() {
	k := "test"
	v := "hi-redis"
	err := RedisConn.Set(rCtx, k, v, 0).Err()
	s.Nil(err)

	val, err := RedisConn.Get(rCtx, k).Result()
	s.Nil(err)
	s.Equal(v, val)
	fmt.Printf("%s redis value: %s\n", k, v)

	_, err = RedisConn.Get(rCtx, k+"-missing-key-123").Result()
	s.EqualValues(redis.Nil, err)
}

func TestRedisSuite(t *testing.T) {
	suite.Run(t, &RedisSuite{})
}

type RedisSuite struct {
	suite.Suite
}
