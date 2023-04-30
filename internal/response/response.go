package response

import (
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"net/http"

	"github.com/fzf-labs/fpkg/util/validutil"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// H is a shortcut for map[string]interface{}
type H map[string]interface{}

type HttpResponse struct {
	Code      int         `json:"code"`                 // HTTP Code
	Message   string      `json:"message"`              // 描述信息
	Data      interface{} `json:"data"`                 // 返回信息
	ErrMsg    string      `json:"err_msg,omitempty"`    // 错误信息
	ErrDetail string      `json:"err_detail,omitempty"` // 错误堆栈
}

func HttpSuccess(resp interface{}, lang string) *HttpResponse {
	r := &HttpResponse{
		Code:    errorx.Success.GetBusinessCode(),
		Message: errorx.Success.GetMessage(lang),
		Data:    resp,
	}
	if validutil.IsZero(r.Data) {
		r.Data = H{}
	}
	return r
}
func HttpBusinessError(err *errorx.BusinessErr, lang string) *HttpResponse {
	r := &HttpResponse{
		Code:      err.GetBusinessCode(),
		Message:   err.GetMessage(lang),
		Data:      err.GetErrData(),
		ErrMsg:    err.GetErrMsg(),
		ErrDetail: err.GetDetail(),
	}
	return r
}

func Http(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	lang := r.Header.Get(constant.HeaderLanguage)
	if err == nil {
		httpx.OkJson(w, HttpSuccess(resp, lang))
		return
	}
	Err(r, w, err)
}

func Err(r *http.Request, w http.ResponseWriter, err error) {
	//获取语言头
	lang := r.Header.Get(constant.HeaderLanguage)
	var e *errorx.BusinessErr
	causeErr := errors.Cause(err)
	//自定义错误类型检测
	if businessErr, ok := causeErr.(*errorx.BusinessErr); ok {
		e = businessErr
	} else {
		//系统级错误
		e = errorx.InternalServerError.WithDetail(err)
	}
	if e == nil {
		e = errorx.InternalServerError.WithDetail(err)
	}
	if e.ErrLevel == errorx.WarnLevel || e.ErrLevel == errorx.ErrLevel {
		//告警
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", e)
	}
	httpx.WriteJson(w, e.GetHttpCode(), HttpBusinessError(e, lang))
}
