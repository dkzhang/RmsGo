package webapi

import (
	"github.com/dkzhang/RmsGo/datebaseCommon/redisOps"
	"github.com/jmoiron/sqlx"
)

var TheContext Context

type Context struct {
	TheDb    *sqlx.DB
	TheRedis *redisOps.Redis
}
