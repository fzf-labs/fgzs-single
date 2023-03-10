package tool

import (
	"fgzs-single/internal/app/web/internal/logic/tool"
	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/response"
	"github.com/fzf-labs/fpkg/validatorx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func SqlToStructHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SqlToStructReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		err := validatorx.Validator.Validate(req, r.Header.Get(constant.HeaderLanguage))
		if err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		l := tool.NewSqlToStructLogic(r.Context(), svcCtx)
		resp, err := l.SqlToStruct(&req)
		response.Http(r, w, resp, err)
	}
}
