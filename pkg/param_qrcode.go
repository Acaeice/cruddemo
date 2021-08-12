package pkg

import "sync"

type Qrcode struct {
	Url string `json:"url"`
	Img string `json:"img"`
}

type Query struct {
	Msg  string `json:"msg"`
	Code uint   `json:"code"`
}

type QrcodeParam struct {
	JdCookie   string `json:"jdCookie"`
	JdOklToken string `json:"jdOklToken"`
}

type QrcodeJdToken struct {
	JdToken string `json:"jdToken"`
}

type StepOne struct {
	SToken string `json:"s_token"`
}

type StepTwo struct {
	Token string `json:"token"`
}

type JdCookieRunners struct {
	CookieRunners *sync.Map
}

type StepThree struct {
	CheckIP int    `json:"check_ip"`
	Errcode int    `json:"errcode"`
	Message string `json:"message"`
}
