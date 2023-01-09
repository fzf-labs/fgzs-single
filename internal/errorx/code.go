package errorx

import "net/http"

// @Description: 默认为中文错误码,在NewError()时请传入中文
// @return
var (
	Success = NewError(1, "成功")
	Fail    = NewError(0, "失败")

	//服务级错误
	InternalServerError   = NewError(10001, "服务器发生异常", WithLevel(ErrLevel))
	ErrServiceUnavailable = NewError(10002, "服务不可用", WithLevel(ErrLevel))
	UnknownError          = NewError(10003, "未知错误,请联系管理员", WithLevel(ErrLevel))
	ErrDataException      = NewError(10004, "数据异常", WithLevel(ErrLevel))
)

// 请求相关
var (
	TooManyRequests           = NewError(10101, "请求过多")
	RequestFrequencyIsTooFast = NewError(10105, "请求频率太快了")
	NoAccess                  = NewError(10106, "无访问权限")
)

// 路由权限相关
var (
	RouteNoAccess                     = NewError(10201, "无权限访问")
	RoutePermissionVerificationFailed = NewError(10203, "路由权限校验失败")
	RouteMethodNoAccess               = NewError(10204, "无访问该路由权限")
)

// 参数相关
var (
	ParamBindErr          = NewError(20001, "参数绑定到结构时发生错误", WithLevel(WarnLevel))
	ParamErr              = NewError(20002, "参数有误", WithLevel(WarnLevel))
	ParamValidationErr    = NewError(20003, "参数验证失败")
	ParamNotJsonRequest   = NewError(20004, "请使用JSON请求", WithLevel(WarnLevel))
	ParamHeaderErr        = NewError(20005, "Header参数有误")
	VerificationCodeError = NewError(20006, "验证码错误")
)

// 数据查询相关
var (
	DataSqlErr           = NewError(20100, "数据异常(S)", WithLevel(ErrLevel))
	DataRedisErr         = NewError(20101, "数据异常(R)", WithLevel(ErrLevel))
	DataRecordNotFound   = NewError(20102, "数据不存在")
	DataDuplicateRecords = NewError(20103, "记录重复")
	DataFormattingError  = NewError(20104, "数据格式化错误", WithLevel(ErrLevel))
)

// token,签名 ,校验相关
var (
	TokenNotRequest          = NewError(20200, "请求中未携带令牌")
	TokenFormatErr           = NewError(20201, "令牌格式化错误")
	TokenErr                 = NewError(20202, "错误的token", WithHttpCode(http.StatusUnauthorized))
	TokenInvalidErr          = NewError(20203, "令牌无效", WithHttpCode(http.StatusUnauthorized))
	TokenExpired             = NewError(20205, "令牌过期", WithHttpCode(http.StatusUnauthorized))
	TokenVerificationFailed  = NewError(20206, "您的登录状态已失效,或在其他设备登录,请您重新登录", WithHttpCode(http.StatusUnauthorized))
	TokenRefreshErr          = NewError(20207, "令牌刷新失败")
	TokenStorageFailed       = NewError(20208, "令牌储存失败")
	TokenErrSignatureParam   = NewError(20209, "签名参数缺失")
	TokenWrongTypeOfBusiness = NewError(20210, "错误的业务类型")
	TokenGenerationFailed    = NewError(20211, "Token生成失败")
)

// 文件上传
var (
	FileParsingError            = NewError(20301, "文件解析错误")
	FileNotExist                = NewError(20302, "上传文件不存在")
	FileError                   = NewError(20303, "文件错误")
	FileClassificationException = NewError(20304, "文件分类异常")
	FileOSSUploadException      = NewError(20305, "OSS上传异常")
	FileSizeExceedsLimit        = NewError(20306, "文件大小超过限制范围")
	WrongFileStorageLocation    = NewError(20307, "文件存储位置错误")
)

// 短信
var (
	SmsSendOverClock    = NewError(20400, "短信发送超频")
	SmsCodeInvalid      = NewError(20401, "短信验证码无效")
	SmsCodeExpired      = NewError(20402, "短信验证码未发送或已失效,请重新发送")
	SmsCodeVerified     = NewError(20403, "短信验证码已验证")
	SmsRepeatSend       = NewError(20404, "短信重复发送")
	SmsRequestOverClock = NewError(20405, "短信请求超频")
	SmsSendFailed       = NewError(20406, "短信发送失败")
	SmsTimesLimit       = NewError(20407, "同一手机号,一天只能发%s次")
	SmsCodeBeenSent     = NewError(20408, "短信发送频繁，请%s秒后重试")
	SmsTypeErr          = NewError(20409, "短信类型错误")
)

// 第三方请求
var (
	ThirdPartyRequestFail = NewError(20500, "第三方请求失败")
)

// 用户账号
var (
	AccountNotExist                    = NewError(21000, "账号不存在")
	FailedToObtainAccountInformation   = NewError(21001, "账号信息获取失败")
	UserIsLocked                       = NewError(21002, "用户已锁定.请联系客服")
	UserIsLoggedOut                    = NewError(21003, "用户已注销")
	AccountError                       = NewError(21004, "账号错误")
	WrongPassword                      = NewError(21005, "密码错误")
	AccountIsBanned                    = NewError(21006, "账号封禁中")
	UserUpdateFailed                   = NewError(21007, "用户更新失败")
	DuplicateUsername                  = NewError(21008, "用户名重复")
	AbnormalAccountStatus              = NewError(21009, "账号状态异常,请重试")
	UserNicknameIsSuspectedOfViolation = NewError(21010, "用户昵称涉嫌违规")
	UserNotBoundRole                   = NewError(21011, "用户未绑定角色")
)

// 百度
var (
	BaiduPermissionNotObtained                         = NewError(21200, "未获取到百度云盘授权")
	BaiduCloudDiskNoTransferredFiles                   = NewError(21201, "百度云盘无可转移文件")
	BaiduCloudDiskAlreadyHasExclusiveSpaceForListeners = NewError(21202, "百度云盘已有倾听者专属空间")
	BaiduFailedToCreateTheListenerOnlySpaceFolder      = NewError(21203, "倾听者专属空间文件夹创建失败")
	BaiduCloudDiskFileTransferFailed                   = NewError(21204, "百度云盘文件转移失败")
)
