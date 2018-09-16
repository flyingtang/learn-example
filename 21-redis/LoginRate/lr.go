package LoginRate

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"fmt"
)

func NewRedis(addr string) (client *redis.Client) {

	client = redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	return
}

func IsBeyondMax(id string, client *redis.Client) (bool, error) {

	sc := client.HGetAll(id)
	r, err := sc.Result()
	var count, lastLoginUnix int64

	fmt.Println(r)

	count, err = strconv.ParseInt(r["count"], 10, 32)

	if err != nil {
		return false, err
	}
	lastLoginUnix, err = strconv.ParseInt(r["lastLoginUnix"], 10, 64)

	if err != nil {
		return false, err
	}

	fmt.Println(count,"\n", time.Now().Unix()-lastLoginUnix,"\n", lastLoginUnix)

	if count > 6 && time.Now().Unix()-lastLoginUnix < 3600 {
		return true, nil
	}
	return false, nil
}
