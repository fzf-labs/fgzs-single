package file

import (
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/page"
	"github.com/fzf-labs/fpkg/util/timeutil"
	"github.com/jinzhu/copier"
	url2 "net/url"
	"strings"

	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListLogic {
	return &FileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileListLogic) FileList(req *types.FileListReq) (resp *types.FileListResp, err error) {
	resp = new(types.FileListResp)
	fileUploadDao := dao.Use(l.svcCtx.Gorm).FileUpload
	query := fileUploadDao.WithContext(l.ctx)
	if req.QuickSearch != "" {
		query = query.Where(fileUploadDao.FileName.Like(req.QuickSearch))
	} else {
		for _, search := range req.Search {
			if search.Field == "id" {
				query = query.Where(fileUploadDao.ID.Eq(conv.Int64(search.Val)))
			}
			if search.Field == "originalFileName" {
				query = query.Where(fileUploadDao.OriginalFileName.Eq(search.Val))
			}
			if search.Field == "filename" {
				query = query.Where(fileUploadDao.FileName.Eq(search.Val))
			}
			if search.Field == "storage" {
				query = query.Where(fileUploadDao.Storage.Eq(search.Val))
			}
			if search.Field == "status" {
				query = query.Where(fileUploadDao.Status.Eq(conv.Int32(search.Val)))
			}
			if search.Field == "createdAt" {
				ss := strings.Split(search.Val, ",")
				query = query.Where(fileUploadDao.CreatedAt.Gte(timeutil.Carbon().Parse(ss[0]).Carbon2Time()), fileUploadDao.CreatedAt.Lte(timeutil.Carbon().Parse(ss[1]).Carbon2Time()))
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
				query = query.Order(fileUploadDao.ID.Desc())
			}
		case "createdAt":
			if orderParam[1] == "desc" {
				query = query.Order(fileUploadDao.CreatedAt.Desc())
			}
		}
	}
	queryCount := query
	total, err := queryCount.Count()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	paginator := page.Paginator(req.Page, req.PageSize, int(total))
	result, err := query.Offset(paginator.Offset).Limit(paginator.Limit).Find()
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}

	for _, v := range result {
		var url string
		switch v.Storage {
		case constant.FileStorageLocal:
			url, err = url2.JoinPath(l.svcCtx.Config.Upload.Host, v.Path)
			if err != nil {
				return nil, err
			}
		case constant.FileStorageAliOss:
			url, err = url2.JoinPath(l.svcCtx.Config.AliOss.Host, v.Path)
			if err != nil {
				return nil, err
			}
		case constant.FileStorageTxyOss:

		default:

		}
		resp.List = append(resp.List, types.FileInfo{
			Id:               v.ID,
			FileCategory:     v.FileCategory,
			FileName:         v.FileName,
			OriginalFileName: v.OriginalFileName,
			Storage:          v.Storage,
			Path:             v.Path,
			Url:              url,
			Ext:              v.Ext,
			Size:             v.Size,
			Sha1:             v.Sha1,
			Status:           v.Status,
			CreatedAt:        timeutil.ToDateTimeStringByTime(v.CreatedAt),
			UpdatedAt:        timeutil.ToDateTimeStringByTime(v.UpdatedAt),
		})
	}
	err = copier.Copy(&resp.Paginator, paginator)
	if err != nil {
		return nil, err
	}
	return
}
