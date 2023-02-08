package tool

import (
	"bytes"
	"context"

	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"
	"github.com/cascax/sql2gorm/parser"

	"github.com/zeromicro/go-zero/core/logx"
)

type SqlToStructLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSqlToStructLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SqlToStructLogic {
	return &SqlToStructLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SqlToStructLogic) SqlToStruct(req *types.SqlToStructReq) (resp *types.SqlToStructResp, err error) {
	resp = new(types.SqlToStructResp)
	code := new(bytes.Buffer)
	err = parser.ParseSqlToWrite(req.Sql, code, parser.WithTablePrefix("t_"), parser.WithJsonTag(), parser.WithGormType())
	if err != nil {
		return
	}
	resp.Code = code.String()
	return
}
