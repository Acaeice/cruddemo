package cruddemo

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"sort"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/gin-gonic/gin"
	"github.com/meikeland/errkit"
	"github.com/wechatapi/cruddemo/internal/user"
	"github.com/wechatapi/cruddemo/pkg"
	"github.com/wechatapi/cruddemo/util"
)

const (
	// 登录凭证校验地址
	urlCode2Session = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

// 创建User
// @Summary 创建User
// @Tags pc user
// @Accept json
// @Produce json
// @Param 创建User参数 body pkg.UserCreateParam true "创建User参数"
// @Success 200 {string} string "result"
// @Router /user [GET]
func getCode(c *gin.Context) {
	headerToken := c.Request.Header.Get("token")
	log.Print("heandr:", headerToken)
	if len(headerToken) > 0 {
		log.Printf("header token: %s", headerToken)
	}
	param := &struct {
		Code string `json:"code" form:"code"`
	}{}
	if err := c.ShouldBind(param); err != nil {
		fail(c, errkit.Wrapf(err, "参数不正确"))
		return
	}
	code2Session, err := user.UserRepo.GetCode(c.Request.Context(), param.Code)
	if err != nil {
		fail(c, err)
		return
	}

	ok(c, resp{
		"code2Session": code2Session,
	})
}

func wechatQuickLogin(c *gin.Context) {
	param := &struct {
		SessionKey        string `json:"session_key" form:"session_key"`
		OpenID            string `json:"openid" form:"openid"`
		UserEncryptedData string `json:"userEncryptedData" form:"userEncryptedData"`
		Iv                string `json:"iv" form:"iv"`
	}{}
	if err := c.ShouldBind(param); err != nil {
		fail(c, errkit.Wrapf(err, "参数不正确"))
		return
	}
	wechatUser, err := user.UserRepo.GetUser(param.SessionKey, param.UserEncryptedData, param.Iv)
	log.Print("sss:", wechatUser)
	if err != nil {
		if err == pkg.ErrWechatAESInvalid {
			fail(c, errkit.New("获取信息失败，请再试一次"))
			return
		}
		fail(c, err)
		return
	}
	userProfile, err := user.UserRepo.GetByOpenid(param.OpenID)
	if err != nil {
		if userProfile == nil {
			var gender uint = pkg.GenderUnknown
			switch wechatUser.Gender {
			case 1:
				gender = pkg.GenderMale
			case 2:
				gender = pkg.GenderFemale
			}
			registerParam := pkg.User{
				Nick:     wechatUser.NickName,
				Gender:   gender,
				Avatar:   wechatUser.AvatarURL,
				AppID:    wechatUser.Watermark.AppID,
				OpenID:   param.OpenID,
				Language: wechatUser.Language,
				City:     wechatUser.City,
				Province: wechatUser.Province,
				Country:  wechatUser.Country,
			}
			userProfile, err := user.UserRepo.Create(registerParam)
			if err != nil {
				fail(c, err)
				return
			}
			token, err := util.Sign(userProfile.ID)
			if err != nil {
				fail(c, err)
				return
			}
			ok(c, resp{
				"userID": userProfile.ID,
				"token":  token.Token,
			})
			return
		}

	}
	token, err := util.Sign(userProfile.ID)
	if err != nil {
		fail(c, err)
		return
	}

	ok(c, resp{
		"userID": userProfile.ID,
		"token":  token.Token,
	})
}

func getuserbyId(c *gin.Context) {
	userID, _ := util.ParseUint(c.Query("userID"))

	// 校验favoriteStoreID是否为0
	if userID <= 0 {
		fail(c, errors.New("用户Id不能为0"))
		return
	}
	User, err := user.UserRepo.GetByID(userID)

	if err != nil {
		fail(c, err)
		return
	}

	ok(c, resp{
		"User": User,
	})
}

//微信接入
func answer(c *gin.Context) {
	param := pkg.WxAccess{}
	if err := c.ShouldBind(param); err != nil {
		fail(c, errkit.Wrapf(err, "参数不正确"))
		return
	}
	var token string = "rootwang8023"
	var tempArray = []string{token, param.Timestamp, param.Nonce}
	sort.Strings(tempArray)
	sha1String := ""

	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	if sha1String == param.Signature {
		fmt.Printf("接入成功")
	} else {
		fmt.Printf("接入失败")
	}

	ok(c, resp{
		"Ok": param.Echostr,
	})
}

//阿里智能语音交互Token更新
func getToken(c *gin.Context) {
	client, err := sdk.NewClientWithAccessKey("cn-shanghai", "LTAI5tBcn1GoDktEP1VYrdbt", "UiQjLiwm5mVpRXVZSzZeRj4OxcLuzs")
	if err != nil {
		panic(err)
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Domain = "nls-meta.cn-shanghai.aliyuncs.com"
	request.ApiName = "CreateToken"
	request.Version = "2019-02-28"
	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpStatus())
	fmt.Print(response.GetHttpContentString())
}
