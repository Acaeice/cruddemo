package user

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/meikeland/logger"
	"github.com/wechatapi/cruddemo/pkg"
	"github.com/wechatapi/cruddemo/util"
	"gorm.io/gorm"
)

type userSQLRepo struct {
	db *gorm.DB
}

const (
	// 登录凭证校验地址
	urlCode2Session = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

var (
	errWriteSubscribeCount    = errors.New("写入SubscribeCount失败，请联系技术支持")
	errReadSubscribeCount     = errors.New("查询SubscribeCount失败，请联系技术支持")
	errNotFoundSubscribeCount = errors.New("找不到SubscribeCount")
)

func (repo userSQLRepo) GetByID(id uint) (*pkg.User, error) {
	q := User{
		ID: id,
	}
	return repo.get(q)
}

func (repo userSQLRepo) GetByOpenid(openid string) (*pkg.User, error) {
	userProfile := pkg.User{}
	if err := repo.db.Where("open_id = ?", openid).First(&userProfile).Error; err != nil {
		return nil, errors.New("查询错误")
	}
	return &userProfile, nil
}

func (repo userSQLRepo) Create(param pkg.User) (*pkg.User, error) {
	contextLogger := logger.WithFields(logger.Fields{
		"param": param,
	})
	obj := &pkg.User{
		Nick:     param.Nick,
		Gender:   param.Gender,
		Avatar:   param.Avatar,
		AppID:    param.AppID,
		OpenID:   param.OpenID,
		Language: param.Language,
		City:     param.City,
		Province: param.Province,
		Country:  param.Country,
		JdFk:     param.JdFk,
		Unionid:  param.Unionid,
	}
	if err := repo.db.Create(obj).Error; err != nil {
		contextLogger.Errorf("Create User err: %v", err)
		return nil, errWriteSubscribeCount
	}
	return repo.GetByID(obj.ID)
}

func (repo userSQLRepo) GetCode(ctx context.Context, code string) (*pkg.Code2Session, error) {
	if len(code) == 0 {
		return nil, errors.New("code不能为空")
	}
	url := fmt.Sprintf(urlCode2Session, "wx6e415ee37673960b", "8b7b17b84d90dea852ad9e9bc117ff2a", code)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	code2Session := pkg.Code2Session{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&code2Session); err != nil {
		return nil, err
	}

	log.Print(&code2Session)

	// 判断微信接口返回的是否是一个异常情况
	if code2Session.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", code2Session.ErrCode, code2Session.ErrMsg))
	}

	return &code2Session, nil
}

func (repo userSQLRepo) GetUser(sessionKey, encryptedData, iv string) (*pkg.MAppUser, error) {
	var (
		aesKey, aesIV, aesCipherText []byte
		aesBlock                     cipher.Block
		err                          error
	)

	if aesKey, err = base64.StdEncoding.DecodeString(sessionKey); err != nil {
		fmt.Printf("++++++++++++++++++++++++++++++++++++++：%s", sessionKey)
		return nil, errors.New("解析微信用户信息失败")
	}
	if aesIV, err = base64.StdEncoding.DecodeString(iv); err != nil {
		fmt.Printf("++++++++++++++++++++++++++++++++++++++：%s", iv)
		return nil, errors.New("解析微信用户信息失败")
	}
	if aesCipherText, err = base64.StdEncoding.DecodeString(encryptedData); err != nil {
		fmt.Printf("++++++++++++++++++++++++++++++++++++++：%s", encryptedData)
		return nil, errors.New("解析微信用户信息失败")
	}
	aesPlantText := make([]byte, len(aesCipherText))
	if aesBlock, err = aes.NewCipher(aesKey); err != nil {
		return nil, errors.New("解析微信用户信息失败")
	}
	mode := cipher.NewCBCDecrypter(aesBlock, aesIV)
	mode.CryptBlocks(aesPlantText, aesCipherText)
	aesPlantText = util.PKCS7UnPadding(aesPlantText)
	wechatUser := &pkg.MAppUser{}
	log.Printf("从微信返回字符串解析微信用户, aesPlantText: %s", aesPlantText)
	if err := json.Unmarshal(aesPlantText, wechatUser); err != nil {
		log.Println("解析微信用户信息失败")
		log.Print(aesPlantText)
		return nil, pkg.ErrWechatAESInvalid
	}
	if wechatUser.Watermark.AppID != "wx6e415ee37673960b" {
		return nil, errors.New("解析微信用户信息失败")
	}
	return wechatUser, nil

}

func (repo userSQLRepo) get(query User) (*pkg.User, error) {
	contextLogger := logger.WithFields(logger.Fields{
		"query": query,
	})

	obj := &pkg.User{}
	if err := repo.db.Scopes(query.where(), query.preload()).First(obj).Error; err != nil {
		contextLogger.Errorf("get SubscribeCount err: %v", err)
		if err == gorm.ErrRecordNotFound {
			return nil, errNotFoundSubscribeCount
		}
		return nil, errReadSubscribeCount
	}
	return obj, nil
}
