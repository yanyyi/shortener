package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	Err404 = errors.New("404")
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// show short url: input "kiki.cn/lus7j" -> redirect real long url
	// bloom filter
	// 不存在的短链接直接返回404，不需要后续处理
	// a.基于内存版本,服务重启之后就没了，所以每次重启都要重新加载以下已有的短链接（从数据库查）

	// b.基于Redis版本,go-zero自带：https//go-zero.dev/cn/docs/blog/governance/bloom/

	exist, err := l.svcCtx.Filter.Exists([]byte(req.ShortUrl))
	if err != nil {
		logx.Errorw("Bloom Filter failed", logx.LogField{Value: err.Error(), Key: "err"})
		return
	}
	// 不存在的短链接直接返回
	if !exist {
		return nil, Err404
	}
	fmt.Println("开始查询缓存...")
	// add cache before query from SQL database
	// go-zero cache support singleflight
	u, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{Valid: true, String: req.ShortUrl})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("404")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl() failed", logx.LogField{Value: err.Error(), Key: "err"})
		return nil, err
	}
	// return the found long url
	return &types.ShowResponse{LongUrl: u.Lurl.String}, nil
}
