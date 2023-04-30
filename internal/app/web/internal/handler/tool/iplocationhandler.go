package tool

import (
	"fgzs-single/internal/app/web/internal/logic/tool"
	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/response"
	"net/http"

	"github.com/fzf-labs/fpkg/util/iputil"
	"github.com/fzf-labs/fpkg/validatorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func IpLocationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IpLocationReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		err := validatorx.Validator.Validate(req, r.Header.Get(constant.HeaderLanguage))
		if err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		l := tool.NewIpLocationLogic(r.Context(), svcCtx)
		if req.Ip == "" {
			req.Ip = iputil.GetRealIP(r)
		}
		resp, err := l.IpLocation(&req)
		response.Http(r, w, resp, err)
	}
}
