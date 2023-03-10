package tinyurl

import (
	"context"
	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/define/cachekey"
	"fgzs-single/internal/errorx"
	"github.com/fzf-labs/fpkg/util/jsonutil"
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
	//进程内缓存查找->redis->mysql
	buildCacheKey := cachekey.TinyUrl.BuildCacheKey(req.Id)
	res, err := buildCacheKey.CollectionRocksCache(l.svcCtx.CollectionCache, l.svcCtx.RocksCache, func() (string, error) {
		tinyURLDao := dao.Use(l.svcCtx.Gorm).TinyURL
		tinyURL, err := tinyURLDao.WithContext(l.ctx).Where(tinyURLDao.TinyURL.Eq(req.Id)).First()
		if err != nil && err != gorm.ErrRecordNotFound {
			return "", errorx.DataSqlErr.WithDetail(err)
		}
		if tinyURL == nil {
			return "", nil
		}
		dbRes, err := jsonutil.EncodeToString(types.TinyUrlResp{
			OriginalUrl: tinyURL.OriginalURL,
			Expired:     tinyURL.Expired,
		})
		if err != nil {
			return "", err
		}
		return dbRes, nil
	})
	if err != nil {
		return nil, err
	}
	err = jsonutil.DecodeString(res, resp)
	if err != nil {
		return nil, errorx.ShortLinkError
	}
	return resp, nil
}
