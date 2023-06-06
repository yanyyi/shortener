package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortener/internal/config"
	"shortener/model"
	"shortener/sequence"
)

type ServiceContext struct {
	Config        config.Config
	ShortUrlModel model.ShortUrlMapModel //short_url_map
	Sequence      sequence.Sequence      //自己定义的结构体
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	return &ServiceContext{
		Config:        c,
		ShortUrlModel: model.NewShortUrlMapModel(conn),
		//Sequence:      sequence.NewMySQL(c.Sequence.DSN),
		Sequence: sequence.NewRedis(c.Sequence.RedisAddr),
	}
}
