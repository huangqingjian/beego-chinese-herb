package component

import (
	"beego-chinese-herb/constant"
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	"time"
)

var (
	redisCache cache.Cache
	redis = "redis"
	redisHost = redis + constant.SMH + "host"
	redisPort = redis + constant.SMH + "port"
	redisPassword = redis + constant.SMH + "password"
	redisDb = redis + constant.SMH + "db"
	redisCollection = redis + constant.SMH + "collection"
)

// 初始化redis
func InitRedis() {
	redisHost, _ := web.AppConfig.String(redisHost)
	redisPort, _ := web.AppConfig.String(redisPort)
	redisPassword, _ := web.AppConfig.String(redisPassword)
	redisDb, _ := web.AppConfig.String(redisDb)
	redisCollection, _ := web.AppConfig.String(redisCollection)
	config := `{"key":"` + redisCollection + `","conn":"` + redisHost + `:` + redisPort + `","dbNum":"` + redisDb + `","password":"` + redisPassword + `"}`
	cache, err := cache.NewCache("redis", config)
	if err != nil {
		logs.Error("创建redis缓存失败～")
	}
	redisCache = cache
}

type Redis struct {
}

// 查找
func (r *Redis) Get(key string) (interface{}, error){
	return redisCache.Get(context.TODO(), key)
}

// 插入
func (r *Redis) Put(key string, val interface{}, timeout time.Duration) error {
	return redisCache.Put(context.TODO(), key, val, timeout)
}

// 删除
func (r *Redis) Delete(key string) (error){
	return redisCache.Delete(context.TODO(), key)
}



