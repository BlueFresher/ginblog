package model

import (
	"encoding/json"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var DBrs *redis.Client

func InitRedis() {
	redisDb := redis.NewClient(&redis.Options{
		Addr:     utils.ReHost + ":6379", // 指定
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err = redisDb.Ping().Result()
	if err != nil {
		log.Println("redis 连接失败 :", err)
	} else {
		log.Println("redis 连接成功")
	}
	DBrs = redisDb
}

func AddArticlesRedis(articles []string) {
	DBrs.RPush("articles", articles)
	// 设置2小时过期时间
	DBrs.Expire("articles", 1800*time.Second)
}

func DeleteArticlesRedis() (n int64, err error) {
	n, err = DBrs.Del("articles").Result()
	return n, err
}

func GetArticlesRedist(pageNum, pageSize int) ([]Article, int, int) {
	page := (pageNum - 1) * pageSize
	articlesRedis := DBrs.LRange("articles", int64(page), int64(page+pageSize)).Val()
	var ret = make([]Article, 0)
	var total int

	for i, v := range articlesRedis {
		if i == 0 {
			v1 := []byte(v)
			err := json.Unmarshal(v1, &total)
			if err != nil {
				return nil, 0, errmsg.ERROR_REDIS_DESERIALIZE_WRONG
			}
		} else {
			var article Article
			v1 := []byte(v)
			err := json.Unmarshal(v1, &article)
			if err != nil {
				return nil, 0, errmsg.ERROR_REDIS_DESERIALIZE_WRONG
			}
			ret = append(ret, article)
		}
	}
	return ret, total, errmsg.SUCCESS
}
