package file

import (
	"context"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"fgzs-single/pkg/oss"
	"fgzs-single/pkg/util/fileutil"
	"fgzs-single/pkg/util/uuidutil"
	"gorm.io/gorm"
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
	fileUploadDao := dao.Use(l.svcCtx.Gorm).FileUpload
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	sha1, err := fileutil.Sha1(file)
	if err != nil {
		return nil, err
	}
	fileUpload, err := fileUploadDao.WithContext(l.ctx).Where(fileUploadDao.Sha1.Eq(sha1)).First()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	if fileUpload != nil {
		resp.ID = fileUpload.ID
		return
	}
	originalFileName := strings.ToLower(fileHeader.Filename)
	fileName := uuidutil.KSUidByTime()
	ext := fileutil.Ext(originalFileName)
	path := BuildPath(req.FileCategory, fileName, ext)
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	mime, _ := fileutil.ReaderMimeTypeAndExt(file)
	switch req.Storage {
	case constant.FileStorageLocal:
		dstFile, err := fileutil.QuickOpenFile(filepath.Join(l.svcCtx.Config.Upload.Path, path))
		if err != nil {
			return nil, err
		}
		defer dstFile.Close()
		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(dstFile, file)
		if err != nil {
			return nil, err
		}
	case constant.FileStorageAliOss:
		aliOss := oss.NewAliOss(l.svcCtx.Config.AliOss)
		err := aliOss.PutObj(path, file)
		if err != nil {
			return nil, errorx.FileOSSUploadException.WithDetail(err)
		}
	case constant.FileStorageTxyOss:
	default:
		return nil, errorx.WrongFileStorageLocation

	}
	fileUpload = &model.FileUpload{
		FileCategory:     req.FileCategory,
		FileName:         fileName + ext,
		OriginalFileName: originalFileName,
		Mimetype:         mime,
		Storage:          req.Storage,
		Path:             path,
		Ext:              ext,
		Size:             fileHeader.Size,
		Sha1:             sha1,
	}
	err = fileUploadDao.WithContext(l.ctx).Create(fileUpload)
	if err != nil {
		return nil, errorx.DataSqlErr.WithDetail(err)
	}
	resp.ID = fileUpload.ID
	return
}

func BuildPath(category string, fileName, ext string) string {
	//日期
	date := time.Now().Format("20060102")
	return filepath.ToSlash(filepath.Join(category, date, fileName+ext))
}
