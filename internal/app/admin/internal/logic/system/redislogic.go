package system

import (
	"context"
	"strings"

	"github.com/fzf-labs/fpkg/cache"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisLogic {
	return &RedisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *RedisLogic) Redis(req *types.RedisReq) (*types.RedisResp, error) {
	resp := new(types.RedisResp)
	commandstats := cache.RedisInfo(l.svcCtx.GoRedis, "commandstats")
	for k, v := range commandstats {
		resp.CommandStats = append(resp.CommandStats, map[string]string{
			"name":  strings.Split(k, "_")[1],
			"value": v[strings.Index(v, "=")+1 : strings.Index(v, ",")],
		})
	}
	resp.Info = cache.RedisInfo(l.svcCtx.GoRedis)
	resp.DbSize = cache.DBSize(l.svcCtx.GoRedis)
	return resp, nil
}
