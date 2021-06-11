package pkg

import (
	"errors"
)

// watermarkCreateParam watermark创建参数
type watermarkCreateParam struct {
	AppID string `json:"appID" form:"appID"` //
}

// IsValid 校验参数
func (param watermarkCreateParam) IsValid() error {

	if len(param.AppID) == 0 {
		return errors.New("watermarkCreateParam.AppID不合法")
	}

	return nil
}

// watermarkModifyParam watermark修改参数
type watermarkModifyParam struct {
	ID    uint   `json:"id" form:"id"`
	AppID string `json:"appID" form:"appID"` //
}

// IsValid 校验参数
func (param watermarkModifyParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("watermarkModifyParam.id不合法")
	}

	if len(param.AppID) == 0 {
		return errors.New("watermarkModifyParam.AppID不合法")
	}

	return nil
}

//watermarkUpdateParam watermark单独修改参数
type watermarkUpdateParam struct {
	ID    uint   `json:"id" form:"id"`
	AppID string `json:"appID" form:"appID"` //
}

// IsValid 校验参数
func (param watermarkUpdateParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("watermarkUpdateParam.id不合法")
	}
	return nil
}

// watermarkSearchParam watermark查询参数
type watermarkSearchParam struct {
	AppID    string `json:"appID" form:"appID"` //
	PageSize uint   `json:"pageSize" form:"pageSize"`
	Page     uint   `json:"page" form:"page"`
}

// watermarkListParam watermark查询参数
type watermarkListParam struct {
	AppID string `json:"appID" form:"appID"` //
}

// watermarkDeleteParam watermark删除参数
type watermarkDeleteParam struct {
	ID uint `json:"id" form:"id"`
}

// IsValid 校验参数
func (param watermarkDeleteParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("watermarkDeleteParam.id不合法")
	}
	return nil
}
