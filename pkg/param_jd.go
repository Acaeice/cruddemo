package pkg

import (
	"errors"
)

// UserCreateParam User创建参数
type JdBoundCreateParam struct {
	JDCookie string `json:"jdCookie" form:"jdCookie"` // Cookie
	JDPin    string `json:"jdPin" form:"jdPin"`       // JDPin

}

// IsValid 校验参数
func (param JdBoundCreateParam) IsValid() error {

	if len(param.JDCookie) == 0 {
		return errors.New("JdBoundCreateParam.JDCookie不合法")
	}

	if len(param.JDPin) == 0 {
		return errors.New("UserCreateParam.JDPin不合法")
	}

	return nil
}

// UserModifyParam User修改参数
type JdBoundModifyParam struct {
	ID       uint   `json:"id" form:"id"`
	JDCookie string `json:"jdCookie" form:"jdCookie"` // Cookie
	JDPin    string `json:"jdPin" form:"jdPin"`       // JDPin

}

// IsValid 校验参数
func (param JdBoundModifyParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("UserModifyParam.id不合法")
	}

	if len(param.JDCookie) == 0 {
		return errors.New("JdBoundCreateParam.JDCookie不合法")
	}

	if len(param.JDPin) == 0 {
		return errors.New("UserCreateParam.JDPin不合法")
	}

	return nil
}

//JdBoundUpdateParam User单独修改参数
type JdBoundUpdateParam struct {
	ID       uint   `json:"id" form:"id"`
	JDCookie string `json:"jdCookie" form:"jdCookie"` // Cookie
	JDPin    string `json:"jdPin" form:"jdPin"`       // JDPin
}

// IsValid 校验参数
func (param JdBoundUpdateParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("UserUpdateParam.id不合法")
	}
	return nil
}

// UserSearchParam User查询参数
type JdBoundSearchParam struct {
	JDCookie string `json:"jdCookie" form:"jdCookie"` // Cookie
	JDPin    string `json:"jdPin" form:"jdPin"`       // JDPin

	PageSize uint `json:"pageSize" form:"pageSize"`
	Page     uint `json:"page" form:"page"`
}

// UserListParam User查询参数
type JdBoundListParam struct {
	JDCookie string `json:"jdCookie" form:"jdCookie"` // Cookie
	JDPin    string `json:"jdPin" form:"jdPin"`       // JDPin

}

// UserDeleteParam User删除参数
type JdBoundDeleteParam struct {
	ID uint `json:"id" form:"id"`
}

// IsValid 校验参数
func (param JdBoundDeleteParam) IsValid() error {
	if param.ID == 0 {
		return errors.New("UserDeleteParam.id不合法")
	}
	return nil
}
