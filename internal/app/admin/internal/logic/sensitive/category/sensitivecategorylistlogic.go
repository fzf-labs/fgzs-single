package category

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/errorx"
	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/page"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/jinzhu/copier"
	"strings"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSensitiveCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveCategoryListLogic {
	return &SensitiveCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SensitiveCategoryListLogic) SensitiveCategoryList(req *types.SensitiveCategoryListReq) (resp *types.SensitiveCategoryListResp, err error) {
	resp = new(types.SensitiveCategoryListResp)
	sensitiveCategoryDao := dao.Use(l.svcCtx.Gorm).SensitiveCategory
	query := sensitiveCategoryDao.WithContext(l.ctx)
	if req.QuickSearch != "" {
		query = query.Where(sensitiveCategoryDao.Name.Like(req.QuickSearch))
	} else {
		for _, search := range req.Search {
			if search.Field == "id" {
				query = query.Where(sensitiveCategoryDao.ID.Eq(conv.Int64(search.Val)))
			}
			if search.Field == "label" {
				query = query.Where(sensitiveCategoryDao.Label.Eq(search.Val))
			}
			if search.Field == "name" {
				query = query.Where(sensitiveCategoryDao.Name.Eq(search.Val))
			}
			if search.Field == "createdAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sensitiveCategoryDao.CreatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sensitiveCategoryDao.CreatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
			}
			if search.Field == "updatedAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sensitiveCategoryDao.UpdatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sensitiveCategoryDao.UpdatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
			}
		}
	}
	if req.Order != "" {
		orderParam := strings.Split(req.Order, ",")
		if len(orderParam) != 2 {
			return nil, errorx.ParamErr.WithCustomMsg("order异常")
		}
		switch orderParam[0] {
		case "id":
			if orderParam[1] == "desc" {
				query = query.Order(sensitiveCategoryDao.ID.Desc())
			}
		case "createdAt":
			if orderParam[1] == "desc" {
				query = query.Order(sensitiveCategoryDao.CreatedAt.Desc())
			}
		case "updatedAt":
			if orderParam[1] == "desc" {
				query = query.Order(sensitiveCategoryDao.UpdatedAt.Desc())
			}
		}
	}
	queryCount := query
	total, err := queryCount.Count()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	paginator := page.Paginator(req.Page, req.PageSize, int(total))
	res, err := query.Offset(paginator.Offset).Limit(paginator.Limit).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	for _, v := range res {
		resp.List = append(resp.List, types.SensitiveCategoryInfo{
			Id:        v.ID,
			Label:     v.Label,
			Name:      v.Name,
			CreatedAt: timeutil.ToDateTimeStringByTime(v.CreatedAt),
			UpdatedAt: timeutil.ToDateTimeStringByTime(v.UpdatedAt),
		})
	}
	err = copier.Copy(&resp.Paginator, paginator)
	if err != nil {
		return nil, err
	}
	return
}
