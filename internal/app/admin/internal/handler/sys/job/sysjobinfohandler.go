package job

import (
	"fgzs-single/internal/app/admin/internal/logic/sys/job"
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/response"
	"net/http"

	"github.com/fzf-labs/fpkg/validatorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SysJobInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysJobInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		err := validatorx.Validator.Validate(req, r.Header.Get(constant.HeaderLanguage))
		if err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		l := job.NewSysJobInfoLogic(r.Context(), svcCtx)
		resp, err := l.SysJobInfo(&req)
		response.Http(r, w, resp, err)
	}
}
