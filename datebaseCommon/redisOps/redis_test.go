package redisOps

import (
	"github.com/dkzhang/RmsGo/datebaseCommon/security"
	"os"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	theDbSecurity, err := security.LoadDbSecurity(os.Getenv("DbSE"))
	if err != nil {
		t.Fatalf("dbConfig.LoadDbSecurity error, ENV DbSE = %s, error = %v", os.Getenv("DbSE"), err)
		return
	}

	opts := &RedisOpts{
		Host: theDbSecurity.TheRedisSecurity.Host,
	}
	redis := NewRedis(opts)
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
