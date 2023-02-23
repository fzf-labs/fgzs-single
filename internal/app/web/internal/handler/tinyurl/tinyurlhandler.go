package tinyurl

import (
	"fgzs-single/internal/app/web/internal/logic/tinyurl"
	"fgzs-single/internal/app/web/internal/svc"
	"fgzs-single/internal/app/web/internal/types"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/response"
	"fgzs-single/pkg/util/timeutil"
	"fgzs-single/pkg/validatorx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func TinyUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TinyUrlReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		err := validatorx.Validator.Validate(req, r.Header.Get(constant.HeaderLanguage))
		if err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		l := tinyurl.NewTinyUrlLogic(r.Context(), svcCtx)
		resp, err := l.TinyUrl(&req)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		if resp.Expired < timeutil.NowUnix() {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, resp.OriginalUrl, http.StatusFound)
	}
}
