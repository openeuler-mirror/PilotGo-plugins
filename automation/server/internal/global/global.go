package global

import (
	"log"

	"github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
)

type AppContext struct {
	MySQL  *gorm.DB
	Redis  *redis.Client
	Logger *log.Logger
	Etcd   *clientv3.Client
}

var App = &AppContext{}
