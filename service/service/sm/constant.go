package sm

const (
	SendSmsTencentCloudURL                   = "sms.tencentcloudapi.com" // 腾讯云短信url
)

// TencentCloudStatusCode 短信发送状态的错误码
type TencentCloudStatusCode string

const (
	// TencentCloudStatusCodeOK 发送成功
	TencentCloudStatusCodeOK TencentCloudStatusCode = "Ok"
	// TencentCloudStatusCodeLimitExceededPhoneNumberThirtySecondLimit 触发限频策略
	TencentCloudStatusCodeLimitExceededPhoneNumberThirtySecondLimit TencentCloudStatusCode = "LimitExceeded.PhoneNumberThirtySecondLimit"
)
