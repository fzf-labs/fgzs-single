package job

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

type SysJobListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysJobListLogic {
	return &SysJobListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysJobListLogic) SysJobList(req *types.SysJobListReq) (resp *types.SysJobListResp, err error) {
	resp = new(types.SysJobListResp)
	sysJobDao := dao.Use(l.svcCtx.Gorm).SysJob
	query := sysJobDao.WithContext(l.ctx)
	if req.QuickSearch != "" {
		query = query.Where(sysJobDao.Name.Like(req.QuickSearch))
	} else {
		for _, search := range req.Search {
			if search.Field == "id" {
				query = query.Where(sysJobDao.ID.Eq(conv.Int64(search.Val)))
			}
			if search.Field == "name" {
				query = query.Where(sysJobDao.Name.Eq(search.Val))
			}
			if search.Field == "status" {
				query = query.Where(sysJobDao.Status.Eq(conv.Int32(search.Val)))
			}
			if search.Field == "createdAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sysJobDao.CreatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sysJobDao.CreatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
			}
			if search.Field == "updatedAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(sysJobDao.UpdatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), sysJobDao.UpdatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
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
				query = query.Order(sysJobDao.ID.Desc())
			}
		case "createdAt":
			if orderParam[1] == "desc" {
				query = query.Order(sysJobDao.CreatedAt.Desc())
			}
		case "updatedAt":
			if orderParam[1] == "desc" {
				query = query.Order(sysJobDao.UpdatedAt.Desc())
			}
		}
	}
	queryCount := query
	total, err := queryCount.Count()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	paginator := page.Paginator(req.Page, req.PageSize, int(total))
	sysJobs, err := query.Offset(paginator.Offset).Limit(paginator.Limit).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}

	for _, v := range sysJobs {
		resp.List = append(resp.List, types.SysJobInfo{
			ID:        v.ID,
			Name:      v.Name,
			Code:      v.Code,
			Remark:    v.Remark,
			Status:    v.Status,
			Sort:      v.Sort,
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
