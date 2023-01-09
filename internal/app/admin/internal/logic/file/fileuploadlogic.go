package file

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/oss"
	"fgzs-single/pkg/util/fileutil"
	"fgzs-single/pkg/util/uuidutil"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadReq, fileHeader *multipart.FileHeader) (resp *types.FileUploadResp, err error) {
	resp = new(types.FileUploadResp)
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	originalFileName := strings.ToLower(fileHeader.Filename)
	fileName := uuidutil.GenUUID()
	ext := fileutil.Ext(originalFileName)
	path := BuildPath(req.FileCategory, fileName, ext)
	switch req.Storage {
	case "local":
		dstFile, err := fileutil.QuickOpenFile(filepath.Join(l.svcCtx.Config.LocalPath, path))
		if err != nil {
			return nil, err
		}
		defer dstFile.Close()
		_, err = io.Copy(dstFile, file)
		if err != nil {
			return nil, err
		}
	case "ali_oss":
		aliOss := oss.NewAliOss(l.svcCtx.Config.AliOss)
		err := aliOss.PutObj(path, file)
		if err != nil {
			return nil, errorx.FileOSSUploadException.WithDetail(err)
		}
	case "txy_oss":
	default:
		return nil, errorx.WrongFileStorageLocation

	}
	fileUploadDao := dao.Use(l.svcCtx.Gorm).FileUpload
	fileUpload := &model.FileUpload{
		FileCategory:     req.FileCategory,
		FileName:         fileName,
		OriginalFileName: originalFileName,
		Storage:          req.Storage,
		Path:             path,
		Ext:              ext,
		Size:             fileHeader.Size,
	}
	err = fileUploadDao.WithContext(l.ctx).Create(fileUpload)
	if err != nil {
		return nil, err
	}
	resp.ID = fileUpload.ID
	return
}

func BuildPath(category string, fileName, ext string) string {
	//日期
	date := time.Now().Format("20060102")
	return filepath.ToSlash(filepath.Join(category, date, fileName+ext))
}
