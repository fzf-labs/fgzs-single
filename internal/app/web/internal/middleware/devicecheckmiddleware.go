package middleware

import (
	"errors"
	"fgzs-single/internal/define/constant"
	"fgzs-single/internal/define/vars"
	"fgzs-single/internal/errorx"
	"fgzs-single/internal/meta"
	"fgzs-single/internal/response"
	"net/http"

	"github.com/fzf-labs/fpkg/crypt"
)

type DeviceCheckMiddleware struct {
}

func NewDeviceCheckMiddleware() *DeviceCheckMiddleware {
	return &DeviceCheckMiddleware{}
}

func (m *DeviceCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deviceId := r.Header.Get(constant.HeaderDeviceId)
		deviceType := r.Header.Get(constant.HeaderDeviceType)
		deviceSign := r.Header.Get(constant.HeaderDeviceSign)
		if deviceId == "" || deviceType == "" || deviceSign == "" {
			response.Err(r, w, errorx.ParamHeaderErr.WithDetail(errors.New("设备参数未上传")))
			return
		}
		salt := vars.DeviceSalt[deviceType]
		if salt == "" {
			response.Err(r, w, errorx.ParamHeaderErr.WithDetail(errors.New("设备类型错误")))
			return
		}
		err := crypt.Compare(deviceSign, deviceId+salt)
		if err != nil {
			response.Err(r, w, errorx.NoAccess)
			return
		}
		r = r.WithContext(meta.SetDeviceId(r.Context(), deviceId))
		r = r.WithContext(meta.SetDeviceType(r.Context(), deviceType))
		r = r.WithContext(meta.SetDeviceSign(r.Context(), deviceSign))
		next(w, r)
	}
}
