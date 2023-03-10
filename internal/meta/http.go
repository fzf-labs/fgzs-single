package meta

import (
	"context"
	"fgzs-single/internal/define/constant"
	"github.com/fzf-labs/fpkg/conv"
)

// GetUid  从Context中获取用户ID
func GetUid(ctx context.Context) string {
	return conv.String(ctx.Value(constant.ContextWithValueKey(constant.ContextUID)))
}

// SetUid 设置用户ID到Context
func SetUid(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, constant.ContextWithValueKey(constant.ContextUID), value)
}

// GetAdminId  从Context中获取用户ID
func GetAdminId(ctx context.Context) int64 {
	return conv.Int64(ctx.Value(constant.ContextWithValueKey(constant.ContextAdminId)))
}

// SetAdminId 设置用户ID到Context
func SetAdminId(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, constant.ContextWithValueKey(constant.ContextAdminId), value)
}

// GetDeviceType  从Context中获取用户ID
func GetDeviceType(ctx context.Context) string {
	return conv.String(ctx.Value(constant.ContextWithValueKey(constant.ContextDeviceType)))
}

// SetDeviceType 设置用户ID到Context
func SetDeviceType(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, constant.ContextWithValueKey(constant.ContextDeviceType), value)
}

// GetDeviceId  从Context中获取设备ID
func GetDeviceId(ctx context.Context) string {
	return conv.String(ctx.Value(constant.ContextWithValueKey(constant.ContextDeviceId)))
}

// SetDeviceId 设置设备ID
func SetDeviceId(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, constant.ContextWithValueKey(constant.ContextDeviceId), value)
}

// GetDeviceSign  从Context中获取设备签名
func GetDeviceSign(ctx context.Context) string {
	return conv.String(ctx.Value(constant.ContextWithValueKey(constant.ContextDeviceSign)))
}

// SetDeviceSign 设置用户ID到设备签名
func SetDeviceSign(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, constant.ContextWithValueKey(constant.ContextDeviceSign), value)
}

// GetLanguage 获取语言标识
func GetLanguage(ctx context.Context) string {
	return conv.String(ctx.Value(constant.ContextWithValueKey(constant.ContextLanguage)))
}

// GetHttpRequestBody 请求参数
func GetHttpRequestBody(ctx context.Context) string {
	return conv.String(ctx.Value(constant.ContextWithValueKey(constant.ContextHttpRequestBody)))
}

// GetMode 获取mode
func GetMode(ctx context.Context) string {
	return conv.String(ctx.Value(constant.ContextWithValueKey(constant.ContextMode)))
}
