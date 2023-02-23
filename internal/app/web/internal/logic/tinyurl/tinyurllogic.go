package tinyurl

import (
	"context"
	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/define/cachekey"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/util/jsonutil"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type TinyUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTinyUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TinyUrlLogic {
	return &TinyUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TinyUrlLogic) TinyUrl(req *types.TinyUrlReq) (resp *types.TinyUrlResp, err error) {
	resp = new(types.TinyUrlResp)
	//进程内缓存查找->布隆过滤器查找->redis->mysql
	cacheKey := cachekey.TinyUrl.BuildCacheKey(req.Id)
	result, ok := l.svcCtx.TinyUrlCollectionCache.Get(cacheKey.Key())
	if ok {
		urlResp, ok := result.(types.TinyUrlResp)
		if ok {
			resp.OriginalUrl = urlResp.OriginalUrl
			resp.Expired = urlResp.Expired
			return resp, nil
		} else {
			return nil, errorx.ShortLinkError
		}
	}
	err = cacheKey.AutoCache(l.svcCtx.Redis, resp, func() (string, error) {
		tinyURLDao := dao.Use(l.svcCtx.Gorm).TinyURL
		tinyURL, err := tinyURLDao.WithContext(l.ctx).Where(tinyURLDao.TinyURL.Eq(req.Id)).First()
		if err != nil && err != gorm.ErrRecordNotFound {
			return "", errorx.DataSqlErr.WithDetail(err)
		}
		if tinyURL == nil {
			return "", errorx.DataRecordNotFound
		}
		res, err := jsonutil.EncodeToString(types.TinyUrlResp{
			OriginalUrl: tinyURL.OriginalURL,
			Expired:     tinyURL.Expired,
		})
		if err != nil {
			return "", err
		}
		return res, nil
	})
	if err != nil {
		return nil, err
	}
	return
}
