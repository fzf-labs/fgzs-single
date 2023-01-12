package file

import (
	"fgzs-single/internal/app/admin/internal/svc"
	"fgzs-single/internal/app/admin/internal/types"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// http://127.0.0.1:9800/file/static/icon/20230112/1/2kbs1db6sy3rig7qawyqlawz6tf.sql
func FileStaticHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileStaticReq
		if err := httpx.Parse(r, &req); err != nil {
			response.Err(r, w, errorx.ParamErr.WithDetail(err))
			return
		}
		handler := http.StripPrefix("/file/static/", http.FileServer(http.Dir(svcCtx.Config.Upload.Path)))
		handler.ServeHTTP(w, r)
	}
}
