package cachekey

import (
	"fgzs-single/pkg/cache"
	"time"
)

// 缓存key前缀

// web
var (
	UUID           = cache.NewCacheKey("uuid", time.Hour, "uuid")
	DL             = cache.NewCacheKey("dl", time.Second*5, "分布式锁")
	Sms            = cache.NewCacheKey("sms", time.Minute*5, "短信验证")
	SmsDayNum      = cache.NewCacheKey("sms_day_num", time.Minute*5, "短信发送次数")
	SensitiveWord  = cache.NewCacheKey("sensitive_word", time.Hour*24, "敏感词")
	TinyUrl        = cache.NewCacheKey("tiny_url", time.Hour*24, "短连接")
	ChatGPTMessage = cache.NewCacheKey("chat_gpt_message", time.Minute*30, "chatGPT聊天信息")
)

// admin
var (
	SysAdminInfo     = cache.NewCacheKey("sys_admin_info", time.Minute*5, "管理员信息")
	SysAdminPermmenu = cache.NewCacheKey("sys_admin_permmenu", time.Minute*5, "管理员权限")
)
