package pkg

import (
	"errors"
)

// UserCreateParam User创建参数
type UserCreateParam struct {
	Nick     string `json:"nick" form:"nick"`         //昵称，必传
	Gender   uint   `json:"gender" form:"gender"`     //性别，必传
	Avatar   string `json:"avatar" form:"avatar"`     //头像，必传
	AppID    string `json:"appID" form:"appID"`       //APPID
	OpenID   string `json:"openID" form:"openID"`     //微信openID
	Language string `json:"language" form:"language"` //语言
	City     string `json:"city" form:"city"`         //所在城市
	Province string `json:"province" form:"province"` //所在省份
	Country  string `json:"country" form:"country"`   //所在国家

}

// IsValid 校验参数
func (param UserCreateParam) IsValid() error {

	if len(param.Nick) == 0 {
		return errors.New("UserCreateParam.Nick不合法")
	}

	if param.Gender == 0 {
		return errors.New("UserCreateParam.Gender不合法")
	}

	if len(param.Avatar) == 0 {
		return errors.New("UserCreateParam.Avatar不合法")
	}

	if len(param.AppID) == 0 {
		return errors.New("UserCreateParam.AppID不合法")
	}

	if len(param.OpenID) == 0 {
		return errors.New("UserCreateParam.OpenID不合法")
	}

	if len(param.Language) == 0 {
		return errors.New("UserCreateParam.Language不合法")
	}

	if len(param.City) == 0 {
		return errors.New("UserCreateParam.City不合法")
	}

	if len(param.Province) == 0 {
		return errors.New("UserCreateParam.Province不合法")
	}

	if len(param.Country) == 0 {
		return errors.New("UserCreateParam.Country不合法")
	}

	return nil
}

// UserModifyParam User修改参数
type UserModifyParam struct {
	ID       uint   `json:"id" form:"id"`
	Nick     string `json:"nick" form:"nick"`         //昵称，必传
	Gender   uint   `json:"gender" form:"gender"`     //性别，必传
	Avatar   string `json:"avatar" form:"avatar"`     //头像，必传
	AppID    string `json:"appID" form:"appID"`       //APPID
	OpenID   string `json:"openID" form:"openID"`     //微信openID
	Language string `json:"language" form:"language"` //语言
	City     string `json:"city" form:"city"`         //所在城市
	Province string `json:"province" form:"province"` //所在省份
	Country  string `json:"country" form:"country"`   //所在国家

}

// IsValid 校验参数
func (param UserModifyParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("UserModifyParam.id不合法")
	}

	if len(param.Nick) == 0 {
		return errors.New("UserModifyParam.Nick不合法")
	}

	if param.Gender == 0 {
		return errors.New("UserModifyParam.Gender不合法")
	}

	if len(param.Avatar) == 0 {
		return errors.New("UserModifyParam.Avatar不合法")
	}

	if len(param.AppID) == 0 {
		return errors.New("UserModifyParam.AppID不合法")
	}

	if len(param.OpenID) == 0 {
		return errors.New("UserModifyParam.OpenID不合法")
	}

	if len(param.Language) == 0 {
		return errors.New("UserModifyParam.Language不合法")
	}

	if len(param.City) == 0 {
		return errors.New("UserModifyParam.City不合法")
	}

	if len(param.Province) == 0 {
		return errors.New("UserModifyParam.Province不合法")
	}

	if len(param.Country) == 0 {
		return errors.New("UserModifyParam.Country不合法")
	}

	return nil
}

//UserUpdateParam User单独修改参数
type UserUpdateParam struct {
	ID       uint   `json:"id" form:"id"`
	Nick     string `json:"nick" form:"nick"`         //昵称，必传
	Gender   uint   `json:"gender" form:"gender"`     //性别，必传
	Avatar   string `json:"avatar" form:"avatar"`     //头像，必传
	AppID    string `json:"appID" form:"appID"`       //APPID
	OpenID   string `json:"openID" form:"openID"`     //微信openID
	Language string `json:"language" form:"language"` //语言
	City     string `json:"city" form:"city"`         //所在城市
	Province string `json:"province" form:"province"` //所在省份
	Country  string `json:"country" form:"country"`   //所在国家

}

// IsValid 校验参数
func (param UserUpdateParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("UserUpdateParam.id不合法")
	}
	return nil
}

// UserSearchParam User查询参数
type UserSearchParam struct {
	Nick     string `json:"nick" form:"nick"`         //昵称，必传
	Gender   uint   `json:"gender" form:"gender"`     //性别，必传
	Avatar   string `json:"avatar" form:"avatar"`     //头像，必传
	AppID    string `json:"appID" form:"appID"`       //APPID
	OpenID   string `json:"openID" form:"openID"`     //微信openID
	Language string `json:"language" form:"language"` //语言
	City     string `json:"city" form:"city"`         //所在城市
	Province string `json:"province" form:"province"` //所在省份
	Country  string `json:"country" form:"country"`   //所在国家

	PageSize uint `json:"pageSize" form:"pageSize"`
	Page     uint `json:"page" form:"page"`
}

// UserListParam User查询参数
type UserListParam struct {
	Nick     string `json:"nick" form:"nick"`         //昵称，必传
	Gender   uint   `json:"gender" form:"gender"`     //性别，必传
	Avatar   string `json:"avatar" form:"avatar"`     //头像，必传
	AppID    string `json:"appID" form:"appID"`       //APPID
	OpenID   string `json:"openID" form:"openID"`     //微信openID
	Language string `json:"language" form:"language"` //语言
	City     string `json:"city" form:"city"`         //所在城市
	Province string `json:"province" form:"province"` //所在省份
	Country  string `json:"country" form:"country"`   //所在国家

}

// UserDeleteParam User删除参数
type UserDeleteParam struct {
	ID uint `json:"id" form:"id"`
}

// IsValid 校验参数
func (param UserDeleteParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("UserDeleteParam.id不合法")
	}
	return nil
}
