package etcode

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"github.com/wechatapi/cruddemo/pkg"
	"github.com/wechatapi/cruddemo/util"
	"gorm.io/gorm"
)

type qrCodeSQLRepo struct {
	db *gorm.DB
}

var JdCookieRunners sync.Map

func (repo qrCodeSQLRepo) GetQrcode(c *gin.Context) (*pkg.Qrcode, error) {
	session := sessions.Default(c)
	v := session.Get("jd_token")
	if v != nil {
		token := v.(string)
		if v, ok2 := JdCookieRunners.Load(token); ok2 {
			if len(v.([]interface{})) >= 2 {
				var url = `https://plogin.m.jd.com/cgi-bin/m/tmauth?appid=300&client_type=m&token=` + token
				data, _ := qrcode.Encode(url, qrcode.Medium, 256)
				qrcode := &pkg.Qrcode{
					Url: url,
					Img: base64.StdEncoding.EncodeToString(data),
				}
				return qrcode, nil
			}
		}

	}
	var state = time.Now().Unix()
	var url = fmt.Sprintf(`https://plogin.m.jd.com/cgi-bin/mm/new_login_entrance?lang=chs&appid=300&returnurl=https://wq.jd.com/passport/LoginRedirect?state=%d&returnurl=https://home.m.jd.com/myJd/newhome.action?sceneval=2&ufc=&/myJd/home.action&source=wq_passport`,
		state)
	req := httplib.Get(url)
	req.Header("Connection", "Keep-Alive")
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Header("Accept", "application/json, text/plain, */*")
	req.Header("Accept-Language", "zh-cn")
	req.Header("Referer", url)
	req.Header("User-Agent", util.GetUserAgent())
	req.Header("Host", "plogin.m.jd.com")
	rsp, err := req.Response()
	if err != nil {
		return nil, errors.New("返回失败")
	}
	data, err := ioutil.ReadAll(rsp.Body)
	stoken := &pkg.StepOne{}
	err = json.Unmarshal(data, stoken)
	if err != nil {
		return nil, errors.New("解析失败")
	}

	cookies := strings.Join(rsp.Header.Values("Set-Cookie"), " ")
	var cookie = strings.Join([]string{
		"guid=" + util.FetchJdCookieValue("guid", cookies),
		"lang=chs",
		"lsid=" + util.FetchJdCookieValue("lsid", cookies),
		"lstoken=" + util.FetchJdCookieValue("lstoken", cookies),
	}, ";")
	state = time.Now().Unix()
	req = httplib.Post(
		fmt.Sprintf(`https://plogin.m.jd.com/cgi-bin/m/tmauthreflogurl?s_token=%s&v=%d&remember=true`,
			stoken.SToken,
			state),
	)
	req.Header("Connection", "Keep-Alive")
	req.Header("Content-Type", "application/x-www-form-urlencoded; Charset=UTF-8")
	req.Header("Accept", "application/json, text/plain, */*")
	req.Header("Cookie", cookie)
	req.Header("Referer", fmt.Sprintf(`https://plogin.m.jd.com/login/login?appid=300&returnurl=https://wqlogin2.jd.com/passport/LoginRedirect?state=%d&returnurl=//home.m.jd.com/myJd/newhome.action?sceneval=2&ufc=&/myJd/home.action&source=wq_passport`,
		state),
	)
	req.Header("User-Agent", util.GetUserAgent())
	req.Header("Host", "plogin.m.jd.com")
	req.Body(fmt.Sprintf(`{
		'lang': 'chs',
		'appid': 300,
		'returnurl': 'https://wqlogin2.jd.com/passport/LoginRedirect?state=%dreturnurl=//home.m.jd.com/myJd/newhome.action?sceneval=2&ufc=&/myJd/home.action&source=wq_passport',
	 }`, state))
	rsp, err = req.Response()
	if err != nil {
		return nil, errors.New("返回失败")
	}
	data, err = ioutil.ReadAll(rsp.Body)
	st := pkg.StepTwo{}
	err = json.Unmarshal(data, &st)
	if err != nil {
		return nil, errors.New("解析失败")
	}
	url = `https://plogin.m.jd.com/cgi-bin/m/tmauth?client_type=m&appid=300&token=` + st.Token
	cookies = strings.Join(rsp.Header.Values("Set-Cookie"), " ")
	okl_token := util.FetchJdCookieValue("okl_token", cookies)
	data, _ = qrcode.Encode(url, qrcode.Medium, 256)
	JdCookieRunners.Store(st.Token, []interface{}{cookie, okl_token})
	session.Set("jd_token", st.Token)
	session.Set("jd_cookie", cookie)
	session.Set("jd_okl_token", okl_token)
	session.Save()
	qrcode := &pkg.Qrcode{
		Url: url,
		Img: base64.StdEncoding.EncodeToString(data),
	}

	return qrcode, nil
}

func (repo qrCodeSQLRepo) GetjQuery(c *gin.Context) (*pkg.Query, error) {
	session := sessions.Default(c)
	if v := session.Get("jd_token"); v == nil {
		log.Print("+++++++++++++++++++++++++++++++++++=", v)
		jqueryMsg := &pkg.Query{
			Code: 202,
			Msg:  "重新获取二维码",
		}
		return jqueryMsg, nil
	} else {
		token := v.(string)
		if v, ok1 := JdCookieRunners.Load(token); !ok1 {
			jqueryMsg := &pkg.Query{
				Msg: "重新获取二维码",
			}
			return jqueryMsg, nil
		} else {
			if len(v.([]interface{})) >= 2 {
				jqueryMsg := &pkg.Query{
					Code: 201,
					Msg:  "二维码未扫描,请扫描二维码",
				}
				return jqueryMsg, nil
			} else {
				pin := v.([]interface{})[0].(string)
				session.Set("pin", pin)
				session.Save()
				// if note := c.GetString("note"); note != "" {
				// 	if ck := models.GetJdCookie(pin); ck != nil {
				// 		ck.Updates(models.Note, note)
				// 	}
				// }
				jqueryMsg := &pkg.Query{
					Code: 200,
					Msg:  "成功",
				}
				return jqueryMsg, nil
			}
		}
	}
}

func init() {
	go func() {
		for {
			// log.Print("!!!!!!!!!!!!!!!!!!!!!")
			time.Sleep(time.Second)
			JdCookieRunners.Range(func(k, v interface{}) bool {
				jd_token := k.(string)
				vv := v.([]interface{})
				if len(vv) >= 2 {
					cookie := vv[0].(string)
					okl_token := vv[1].(string)
					// fmt.Println(jd_token, cookie, okl_token)
					result := QrcodeRepo.CheckLogin(jd_token, cookie, okl_token)
					log.Print("+++++++result+++++", result)
					// fmt.Println(result)
					// switch result {
					// case "成功":
					// 	models.SendTgMsg(tgid, "扫码成功")
					// case "授权登录未确认":
					// case "":
					// default: //失效
					// 	models.SendTgMsg(tgid, "扫码失败")
					// }
				}
				return true
			})
		}
	}()
}

func (repo qrCodeSQLRepo) CheckLogin(token, cookie, okl_token string) string {
	state := time.Now().Unix()
	req := httplib.Post(
		fmt.Sprintf(`https://plogin.m.jd.com/cgi-bin/m/tmauthchecktoken?&token=%s&ou_state=0&okl_token=%s`,
			token,
			okl_token,
		),
	)
	req.Header("Referer", fmt.Sprintf(`https://plogin.m.jd.com/login/login?appid=300&returnurl=https://wqlogin2.jd.com/passport/LoginRedirect?state=%d&returnurl=//home.m.jd.com/myJd/newhome.action?sceneval=2&ufc=&/myJd/home.action&source=wq_passport`,
		state),
	)
	req.Header("Cookie", cookie)
	req.Header("Connection", "Keep-Alive")
	req.Header("Content-Type", "application/x-www-form-urlencoded; Charset=UTF-8")
	req.Header("Accept", "application/json, text/plain, */*")
	req.Header("User-Agent", util.GetUserAgent())
	req.Header("Host", "plogin.m.jd.com")

	req.Param("lang", "chs")
	req.Param("appid", "300")
	req.Param("returnurl", fmt.Sprintf("https://wqlogin2.jd.com/passport/LoginRedirect?state=%d&returnurl=//home.m.jd.com/myJd/newhome.action?sceneval=2&ufc=&/myJd/home.action", state))
	req.Param("source", "wq_passport")

	rsp, err := req.Response()
	if err != nil {
		return "" //err.Error()
	}
	data, err := ioutil.ReadAll(rsp.Body)
	sth := &pkg.StepThree{}
	err = json.Unmarshal(data, &sth)
	if err != nil {
		return "" //err.Error()
	}
	log.Print("+++++sth++++++", sth)
	switch sth.Errcode {
	case 0:
		cookies := strings.Join(rsp.Header.Values("Set-Cookie"), " ")
		pt_key := util.FetchJdCookieValue("pt_key", cookies)
		pt_pin := util.FetchJdCookieValue("pt_pin", cookies)
		log.Print("++++++++pt_key", pt_key)
		log.Print("++++++++pt_pin", pt_pin)
		if pt_pin == "" {
			JdCookieRunners.Delete(token)
			return sth.Message
		}
		// go func() {
		// 	ck := models.JdCookie{
		// 		PtKey: pt_key,
		// 		PtPin: pt_pin,
		// 	}
		// 	if nck := models.GetJdCookie(ck.PtPin); nck != nil {
		// 		ck.ToPool(ck.PtKey)
		// 		msg := fmt.Sprintf("更新账号，%s", ck.PtPin)
		// 		(&models.JdCookie{}).Push(msg)
		// 		logs.Info(msg)
		// 	} else {
		// 		models.NewJdCookie(ck)
		// 		msg := fmt.Sprintf("添加账号，%s", ck.PtPin)
		// 		(&models.JdCookie{}).Push(msg)
		// 		logs.Info(msg)
		// 	}
		// 	go func() {
		// 		models.Save <- &ck
		// 	}()
		// }()
		JdCookieRunners.Store(token, []interface{}{pt_pin})
		return "成功"
	case 19: //Token无效，请退出重试
		JdCookieRunners.Delete(token)
		return sth.Message
	case 21: //Token不存在，请退出重试
		JdCookieRunners.Delete(token)
		return sth.Message
	case 176: //授权登录未确认
		return sth.Message
	case 258: //务异常，请稍后重试
		return ""
	case 264: //出错了，请退出重试
		JdCookieRunners.Delete(token)
		return sth.Message
	default:
		JdCookieRunners.Delete(token)
		fmt.Println(sth)
	}
	return ""
}
