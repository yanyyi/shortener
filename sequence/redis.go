package sequence

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"time"
)

const AutoIncrId = "shortener:id"

// 基于Redis实现一个发号器
type Redis struct {
	//redis连接
	redis.RedisConf
}

func NewRedis(redisAddr string) Sequence {
	return &Redis{
		redis.RedisConf{
			Host:        redisAddr,
			Type:        "node",
			Pass:        "",
			Tls:         false,
			NonBlock:    false,
			PingTimeout: time.Second,
		},
	}
}

// Next 取下一个号
func (r *Redis) Next() (seq uint64, err error) {
	rds := redis.MustNewRedis(r.RedisConf)
	ctx := context.Background()
	_, err = rds.IncrCtx(ctx, AutoIncrId)
	if err != nil {
		fmt.Printf("rds.IncrCtx() failed,  error:%v\n", err)
		return
	}
	id, err := rds.Get(AutoIncrId)
	if err != nil {
		fmt.Printf("rds.Get() failed,  error:%v\n", err)
		return
	}
	idInt, _ := strconv.Atoi(id)
	return uint64(idInt), nil
}
