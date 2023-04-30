package middleware

import (
	"bytes"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/response"
	"io"
	"net/http"

	"github.com/fzf-labs/fpkg/util/jsonutil"
	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
)

type RequestLogMiddleware struct {
}

func NewRequestLogMiddleware() *RequestLogMiddleware {
	return &RequestLogMiddleware{}
}

func (r *RequestLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//忽略 websocket
		if r.Header.Get("Connection") == "Upgrade" || r.Header.Get("Upgrade") == "websocket" {
			next(w, r)
		} else {
			bodyLogWriter := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: w}
			w = bodyLogWriter
			//获取url
			url := r.RequestURI
			//获取header (精简参数)
			header := map[string]string{
				"Content-Type":  r.Header.Get("Content-Type"),
				"User-Agent":    r.Header.Get("User-Agent"),
				"Authorization": r.Header.Get("Authorization"),
				"language":      r.Header.Get(constant.HeaderLanguage),
				"device_type":   r.Header.Get(constant.HeaderDeviceType),
				"device_id":     r.Header.Get(constant.HeaderDeviceId),
				"device_sign":   r.Header.Get(constant.HeaderDeviceSign),
			}
			var body []byte
			var err error
			//不记录文件
			if r.FormValue("file") != "" {
				body = []byte("文件")
			} else {
				//获取Request
				body, err = io.ReadAll(r.Body)
				if err != nil {
					response.Err(r, w, errorx.InternalServerError.WithDetail(err))
					return
				}
				r.Body = io.NopCloser(bytes.NewBuffer(body))
			}
			next(w, r)
			var req string
			if gjson.ValidBytes(body) {
				req = gjson.ParseBytes(body).String()
			} else {
				req = string(body)
			}
			headerString, _ := jsonutil.EncodeToString(header)
			//获取response
			resp := bodyLogWriter.body.String()
			logx.WithContext(r.Context()).Infow("request log", logx.LogField{
				Key:   "url",
				Value: url,
			}, logx.LogField{
				Key:   "header",
				Value: headerString,
			}, logx.LogField{
				Key:   "req",
				Value: req,
			}, logx.LogField{
				Key:   "resp",
				Value: resp,
			})
		}
	}
}

type CustomResponseWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
