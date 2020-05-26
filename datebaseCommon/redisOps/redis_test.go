package redisOps

import (
	"github.com/dkzhang/RmsGo/allConfig"
	dbConfig "github.com/dkzhang/RmsGo/datebaseCommon/config"
	"os"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	os.Setenv("DbConf", "./../../Configuration/Security/database.yaml")
	allConfig.LoadAllConfig()
	opts := &RedisOpts{
		Host: dbConfig.TheDbConfig.TheRedisConfig.Host,
	}
	redis := NewRedis(opts)
	var err error
	timeoutDuration := 10 * time.Second

	if err = redis.Set("username", "silenceper", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !redis.IsExist("username") {
		t.Error("IsExist Error")
	}

	name := redis.Get("username").(string)
	if name != "silenceper" {
		t.Error("get Error")
	}

	if err = redis.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
