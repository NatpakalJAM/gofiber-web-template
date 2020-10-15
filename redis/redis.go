package redis

import (
	"context"
	"fmt"
	"gofiber-web-template/cfg"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

//Init init redis package
func Init() {
	rc := cfg.C.Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", rc.Host, rc.Port),
		Password: rc.Password,
		DB:       rc.Db,
	})
}
