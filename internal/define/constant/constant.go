package constant

type ContextWithValueKey string

const (
	HeaderLanguage   = "language"    //语言类型
	HeaderDeviceType = "device_type" //设备类型 例如: ios  android
	HeaderDeviceId   = "device_id"   //设备ID 例如：123456789
	HeaderDeviceSign = "device_sign" //设备签名
)

const (
	ContextUID             = "uid"               //用户ID
	ContextAdminId         = "admin_id"          //系统用户ID
	ContextLanguage        = "language"          //语言
	ContextMode            = "mode"              //环境
	ContextHttpRequestBody = "http_request_body" //http请求body（context元数据）

	ContextDeviceType = "device_type" //设备类型 例如: k3  m3  k3s k5
	ContextDeviceId   = "device_id"   //设备ID 例如：123456789
	ContextDeviceSign = "device_sign" //设备签名
)

// 常见状态
const (
	StatusEnable  int32 = 1
	StatusDisable int32 = 0
)

const (
	IdentityTypeWeChat = "wechat" //用户第三方登录-微信
	IdentityTypeApple  = "apple"  //用户第三方登录-苹果
	IdentityTypeQQ     = "qq"     //用户第三方登录-qq
)

// jwt业务类型
const (
	JwtTypeApp = "App"
	JwtTypeWeb = "web"
)
const (
	SmsPlatformJSms = "jsms" //短信平台类型-极光
)

// 短信发送类型
const (
	SmsTypeLogIn        = "login"         //登录
	SmsTypeRegister     = "register"      //注册
	SmsTypeAccountCheck = "account_check" //账号检查
	SmsTypeAccountBind  = "account_bind"  //账号绑定
)

// 文件上传
const (
	FileStorageLocal  = "local"
	FileStorageAliOss = "ali_oss"
	FileStorageTxyOss = "txy_oss"
)
