package file

import (
	"fgzs-single/internal/app/admin/internal/logic/file"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/response"
	"fgzs-single/pkg/validatorx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func FileInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		err := validatorx.Validator.Validate(req, r.Header.Get(constant.HeaderLanguage))
		if err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		l := file.NewFileInfoLogic(r.Context(), svcCtx)
		resp, err := l.FileInfo(&req)
		response.Http(r, w, resp, err)
	}
}
