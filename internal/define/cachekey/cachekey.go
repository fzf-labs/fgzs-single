package cachekey

import (
	"fgzs-single/pkg/cache"
	"time"
)

// 缓存key前缀
var (
	UUID               = cache.NewCacheKey("uuid", time.Hour, "uuid")
	DL                 = cache.NewCacheKey("dl", time.Second*5, "分布式锁")
	Sms                = cache.NewCacheKey("sms", time.Minute*5, "短信验证")
	SmsDayNum          = cache.NewCacheKey("sms_day_num", time.Minute*5, "短信发送次数")
	SensitiveWord      = cache.NewCacheKey("sensitive_word", time.Hour*24, "敏感词")
	ResourceNewVersion = cache.NewCacheKey("resource_new_version", time.Hour*24, "本地化资源最新版本号")
	ResourceByVersion  = cache.NewCacheKey("resource_by_version", time.Hour*24, "按版本查询本地化资源")
	XiaoEAccessToken   = cache.NewCacheKey("xiaoe_access_token", time.Hour*2, "小鹅通AccessToken")
	BaiduAuth          = cache.NewCacheKey("bd", time.Hour*2, "百度授权")

	WordBookVersion    = cache.NewCacheKey("word_book_version", time.Hour*24, "单词本版本")
	WordBookNewVersion = cache.NewCacheKey("word_book_new_version", time.Hour*24, "单词本最新版本")
)
