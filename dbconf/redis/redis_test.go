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
		Address:        "ip:port",
		Password:       "",
		Db:             0,
		MaxIdle:        32,
		PoolSize:       128,
		ConnectTimeout: time.Second,
		IdleTimeout:    time.Minute,
		ReadTimeout:    100 * time.Millisecond,
		WriteTimeout:   100 * time.Millisecond,
	}
	conf[testRedis1] = redis1

	redis2 := &entry{
		Address:        "ip:port",
		Password:       "",
		Db:             1,
		MaxIdle:        32,
		PoolSize:       128,
		ConnectTimeout: time.Second,
		IdleTimeout:    time.Minute,
		ReadTimeout:    100 * time.Millisecond,
		WriteTimeout:   100 * time.Millisecond,
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
