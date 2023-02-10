package system

import (
	"context"
	"fgzs-single/internal/core"
	"strings"

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
	redis := core.NewGoRedis(l.svcCtx.Config.Redis, 0)
	commandstats := core.RedisInfo(redis, "commandstats")
	for k, v := range commandstats {
		resp.CommandStats = append(resp.CommandStats, map[string]string{
			"name":  strings.Split(k, "_")[1],
			"value": v[strings.Index(v, "=")+1 : strings.Index(v, ",")],
		})
	}
	resp.Info = core.RedisInfo(redis)
	resp.DbSize = core.DBSize(redis)
	return resp, nil
}
