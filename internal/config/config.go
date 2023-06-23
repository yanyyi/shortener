package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	ShortUrlDB struct {
		DSN string
	}

	Sequence struct {
		DSN       string
		RedisAddr string
	}

	BaseString string

	ShortUrlBlackList []string

	ShortDomain string

	CacheRedis cache.CacheConf
}
