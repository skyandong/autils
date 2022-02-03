package redis

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var conf Conf

const (
	testRedis1 = "testredis1"
	testRedis2 = "testredis2"
)

func init() {
	conf = make(Conf)
	redis1 := &entry{
		Address:        "localhost:6379",
		Password:       "testredisxixihaha123",
		Db:             1,
		MaxIdle:        32,
		PoolSize:       128,
		ConnectTimeout: time.Second,
		IdleTimeout:    time.Minute,
		ReadTimeout:    100 * time.Microsecond,
		WriteTimeout:   100 * time.Microsecond,
	}
	conf[testRedis1] = redis1

	redis2 := &entry{
		Address:        "localhost:6379",
		Password:       "testredisxixihaha123",
		Db:             2,
		MaxIdle:        32,
		PoolSize:       128,
		ConnectTimeout: time.Second,
		IdleTimeout:    time.Minute,
		ReadTimeout:    100 * time.Microsecond,
		WriteTimeout:   100 * time.Microsecond,
	}
	conf[testRedis2] = redis2
}

func TestGet(t *testing.T) {
	r := conf.Get(testRedis1)
	assert.NotNil(t, r)
}

func TestEnsure(t *testing.T) {
	err := conf.Ensure([]string{testRedis1, testRedis2})
	assert.NoError(t, err)
}
