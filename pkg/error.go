package pkg

import "errors"

// ErrWechatAESInvalid code2session之后可能出现的错误
var ErrWechatAESInvalid = errors.New("微信返回的个人信息字符串无法识别")

// ErrWechatPhoneAESInvalid code2session之后可能出现的错误
var ErrWechatPhoneAESInvalid = errors.New("微信返回的手机号码字符串无法识别")
