package cruddemo

// import (
// 	"errors"

// 	"code.aliyun.com/mengine/wxkit/sdk/mappsdk"
// 	"code.aliyun.com/yjqu/apigate-wechatapp/util"
// 	"github.com/gin-gonic/gin"
// 	"github.com/meikeland/errkit"
// )

// // ErrWechatAESInvalid code2session之后可能出现的错误
// var ErrWechatAESInvalid = errors.New("微信返回的个人信息字符串无法识别")

// // login 登录
// func login(c *gin.Context) {
// 	param := &struct {
// 		Code          string `json:"code"`
// 		EncryptedData string `json:"encryptedData"`
// 		Iv            string `json:"iv"`
// 		Key           string `json:"key"`
// 		InviterID     uint   `json:"inviterID"`
// 	}{}
// 	if err := c.ShouldBind(param); err != nil {
// 		fail(c, errkit.Wrapf(err, "参数不正确"))
// 		return
// 	}

// 	code2Session, err := mappsdk.Code2Session(param.Code)
// 	if err != nil {
// 		fail(c, err)
// 		return
// 	}

// 	wechatUser, err := mappsdk.GetUser(code2Session.SessionKey, param.EncryptedData, param.Iv)
// 	if err != nil {
// 		if err == ErrWechatAESInvalid {
// 			fail(c, errkit.New("微信不给力了，请再试一次"))
// 			return
// 		}
// 		fail(c, err)
// 		return
// 	}

// 	var gender uint = userpkg.GenderUnknown
// 	switch wechatUser.Gender {
// 	case 1:
// 		gender = userpkg.GenderMale
// 	case 2:
// 		gender = userpkg.GenderFemale
// 	}

// 	loginParam := userpkg.LoginWechatParam{
// 		AppID:    wechatUser.Watermark.AppID,
// 		UnionID:  code2Session.UnionID,
// 		OpenID:   code2Session.OpenID,
// 		Nick:     wechatUser.NickName,
// 		Gender:   gender,
// 		Avatar:   wechatUser.AvatarURL,
// 		Language: wechatUser.Language,
// 		City:     wechatUser.City,
// 		Province: wechatUser.Province,
// 		Country:  wechatUser.Country,
// 	}

// 	userProfile, _, err := user.Manager.LoginByWechat(loginParam)
// 	if err != nil {
// 		if err.Error() == userpkg.ErrUnregister.Error() {
// 			unRegister(c)
// 		} else {
// 			fail(c, err)
// 			return
// 		}
// 	}
// 	token, err := util.Sign(userProfile.ID)
// 	if err != nil {
// 		fail(c, err)
// 		return
// 	}

// 	ok(c, resp{
// 		"userID": userProfile.ID,
// 		"token":  token.Token,
// 	})
// }
