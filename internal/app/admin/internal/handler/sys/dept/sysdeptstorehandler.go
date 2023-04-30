package dept

import (
	"fgzs-single/internal/app/admin/internal/logic/sys/dept"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/response"
	"net/http"

	"github.com/fzf-labs/fpkg/validatorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SysDeptStoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDeptStoreReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		err := validatorx.Validator.Validate(req, r.Header.Get(constant.HeaderLanguage))
		if err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		l := dept.NewSysDeptStoreLogic(r.Context(), svcCtx)
		resp, err := l.SysDeptStore(&req)
		response.Http(r, w, resp, err)
	}
}
