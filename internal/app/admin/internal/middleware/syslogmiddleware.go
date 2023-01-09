package middleware

import (
	"bytes"
	"context"
	"fgzs-single/internal/dal/dao"
	"fgzs-single/internal/dal/model"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/meta"
	"fgzs-single/internal/response"
	"fgzs-single/pkg/util/jsonutil"
	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"io"
	"net"
	"net/http"
	"strings"
)

type SysLogMiddleware struct {
	DB *gorm.DB
}

func NewSysLogMiddleware(db *gorm.DB) *SysLogMiddleware {
	return &SysLogMiddleware{
		DB: db,
	}
}

func (m *SysLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyLogWriter := &CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: w}
		w = bodyLogWriter
		adminId := meta.GetAdminId(r.Context())
		//获取url
		uri := r.RequestURI
		realIP := GetRealIP(r)
		userAgent := r.Header.Get("User-Agent")
		//获取header (精简参数)
		header := map[string]string{
			"Content-Type":  r.Header.Get("Content-Type"),
			"Authorization": r.Header.Get("Authorization"),
			"language":      r.Header.Get(constant.HeaderLanguage),
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
		logx.WithContext(r.Context()).Infow("syslog", logx.LogField{
			Key:   "adminId",
			Value: adminId,
		}, logx.LogField{
			Key:   "uri",
			Value: uri,
		}, logx.LogField{
			Key:   "ip",
			Value: realIP,
		}, logx.LogField{
			Key:   "user_agent",
			Value: userAgent,
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
		sysLogDao := dao.Use(m.DB).SysLog
		_ = sysLogDao.WithContext(context.Background()).Create(&model.SysLog{
			AdminID:   adminId,
			IP:        realIP,
			URI:       uri,
			Useragent: userAgent,
			Header:    datatypes.JSON(headerString),
			Req:       datatypes.JSON(req),
			Resp:      datatypes.JSON(resp),
		})
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

// GetRealIP 获取真实IP
func GetRealIP(ctx *http.Request) (ip string) {
	var index int
	if ip = ctx.Header.Get("X-Forwarded-For"); ip != "" {
		index = strings.IndexByte(ip, ',')
		if index < 0 {
			return ip
		}
		if ip = ip[:index]; ip != "" {
			return ip
		}
	}
	if ip = ctx.Header.Get("X-Real-Ip"); ip != "" {
		index = strings.IndexByte(ip, ',')
		if index < 0 {
			return ip
		}
		if ip = ip[:index]; ip != "" {
			return ip
		}
	}
	if ip = ctx.Header.Get("Proxy-Forwarded-For"); ip != "" {
		index = strings.IndexByte(ip, ',')
		if index < 0 {
			return ip
		}
		if ip = ip[:index]; ip != "" {
			return ip
		}
	}
	ip, _, _ = net.SplitHostPort(ctx.RemoteAddr)
	return ip
}
