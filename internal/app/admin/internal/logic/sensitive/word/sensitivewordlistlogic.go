package word

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/repo"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/conv"
	"fgzs-single/pkg/page"
	"fgzs-single/pkg/util/timeutil"
	"github.com/jinzhu/copier"
	"strings"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SensitiveWordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSensitiveWordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SensitiveWordListLogic {
	return &SensitiveWordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SensitiveWordListLogic) SensitiveWordList(req *types.SensitiveWordListReq) (resp *types.SensitiveWordListResp, err error) {
	resp = new(types.SensitiveWordListResp)
	sensitiveWordDao := dao.Use(l.svcCtx.Gorm).SensitiveWord
	query := sensitiveWordDao.WithContext(l.ctx)
	if req.QuickSearch != "" {
		query = query.Where(sensitiveWordDao.Text.Like(req.QuickSearch))
	} else {
		for _, search := range req.Search {
			if search.Field == "id" {
				query = query.Where(sensitiveWordDao.ID.Eq(conv.Int64(search.Val)))
			}
			if search.Field == "word" {
				query = query.Where(sensitiveWordDao.Text.Eq(search.Val))
			}
			if search.Field == "categoryID" {
				query = query.Where(sensitiveWordDao.CategoryID.Eq(conv.Int64(search.Val)))
			}
			if search.Field == "createdAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sensitiveWordDao.CreatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sensitiveWordDao.CreatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
			}
			if search.Field == "updatedAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sensitiveWordDao.UpdatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sensitiveWordDao.UpdatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
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
				query = query.Order(sensitiveWordDao.ID.Desc())
			}
		case "createdAt":
			if orderParam[1] == "desc" {
				query = query.Order(sensitiveWordDao.CreatedAt.Desc())
			}
		case "updatedAt":
			if orderParam[1] == "desc" {
				query = query.Order(sensitiveWordDao.UpdatedAt.Desc())
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
	var ids []int64
	for _, v := range res {
		ids = append(ids, v.CategoryID)
	}
	idToName, err := repo.NewSensitiveCategoryRepo(l.ctx, l.svcCtx.Gorm).IdToName(ids)
	if err != nil {
		return nil, err
	}
	for _, v := range res {
		resp.List = append(resp.List, types.SensitiveWordInfo{
			Id:           v.ID,
			CategoryID:   v.CategoryID,
			CategoryName: idToName[v.CategoryID],
			Text:         v.Text,
			CreatedAt:    timeutil.ToDateTimeStringByTime(v.CreatedAt),
			UpdatedAt:    timeutil.ToDateTimeStringByTime(v.UpdatedAt),
		})
	}
	err = copier.Copy(&resp.Paginator, paginator)
	if err != nil {
		return nil, err
	}
	return
}