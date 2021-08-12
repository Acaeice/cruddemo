package cruddemo

import (
	"github.com/gin-gonic/gin"
	"github.com/wechatapi/cruddemo/internal/etcode"
)

func getQrcode(c *gin.Context) {
	qrcode, err := etcode.QrcodeRepo.GetQrcode(c)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"qrcode": qrcode,
	})
	// session := sessions.Default(c)
	// v := session.Get("jd_token")
	// if v != nil {
	// 	token := v.(string)
	// 	if v, ok2 := JdCookieRunners.Load(token); ok2 {
	// 		if len(v.([]interface{})) >= 2 {
	// 			var url = `https://plogin.m.jd.com/cgi-bin/m/tmauth?appid=300&client_type=m&token=` + token
	// 			data, _ := qrcode.Encode(url, qrcode.Medium, 256)
	// 			qrcode := &pkg.Qrcode{
	// 				Url: url,
	// 				Img: base64.StdEncoding.EncodeToString(data),
	// 			}
	// 			ok(c, resp{
	// 				"qrcode": qrcode,
	// 			})
	// 			return
	// 		}
	// 	}

	// }
	// var state = time.Now().Unix()
	// var url = fmt.Sprintf(`https://plogin.m.jd.com/cgi-bin/mm/new_login_entrance?lang=chs&appid=300&returnurl=https://wq.jd.com/passport/LoginRedirect?state=%d&returnurl=https://home.m.jd.com/myJd/newhome.action?sceneval=2&ufc=&/myJd/home.action&source=wq_passport`,
	// 	state)
	// req := httplib.Get(url)
	// req.Header("Connection", "Keep-Alive")
	// req.Header("Content-Type", "application/x-www-form-urlencoded")
	// req.Header("Accept", "application/json, text/plain, */*")
	// req.Header("Accept-Language", "zh-cn")
	// req.Header("Referer", url)
	// req.Header("User-Agent", util.GetUserAgent())
	// req.Header("Host", "plogin.m.jd.com")
	// rsp, err := req.Response()
	// if err != nil {
	// 	fail(c, err)
	// 	return
	// }
	// data, err := ioutil.ReadAll(rsp.Body)
	// stoken := &pkg.StepOne{}
	// err = json.Unmarshal(data, stoken)
	// log.Printf("%vstoken:", stoken)
	// if err != nil {
	// 	fail(c, err)
	// 	return
	// }

	// cookies := strings.Join(rsp.Header.Values("Set-Cookie"), " ")
	// var cookie = strings.Join([]string{
	// 	"guid=" + util.FetchJdCookieValue("guid", cookies),
	// 	"lang=chs",
	// 	"lsid=" + util.FetchJdCookieValue("lsid", cookies),
	// 	"lstoken=" + util.FetchJdCookieValue("lstoken", cookies),
	// }, ";")
	// state = time.Now().Unix()
	// req = httplib.Post(
	// 	fmt.Sprintf(`https://plogin.m.jd.com/cgi-bin/m/tmauthreflogurl?s_token=%s&v=%d&remember=true`,
	// 		stoken.SToken,
	// 		state),
	// )
	// req.Header("Connection", "Keep-Alive")
	// req.Header("Content-Type", "application/x-www-form-urlencoded; Charset=UTF-8")
	// req.Header("Accept", "application/json, text/plain, */*")
	// req.Header("Cookie", cookie)
	// req.Header("Referer", fmt.Sprintf(`https://plogin.m.jd.com/login/login?appid=300&returnurl=https://wqlogin2.jd.com/passport/LoginRedirect?state=%d&returnurl=//home.m.jd.com/myJd/newhome.action?sceneval=2&ufc=&/myJd/home.action&source=wq_passport`,
	// 	state),
	// )
	// req.Header("User-Agent", util.GetUserAgent())
	// req.Header("Host", "plogin.m.jd.com")
	// req.Body(fmt.Sprintf(`{
	// 	'lang': 'chs',
	// 	'appid': 300,
	// 	'returnurl': 'https://wqlogin2.jd.com/passport/LoginRedirect?state=%dreturnurl=//home.m.jd.com/myJd/newhome.action?sceneval=2&ufc=&/myJd/home.action&source=wq_passport',
	//  }`, state))
	// rsp, err = req.Response()
	// if err != nil {
	// 	fail(c, err)
	// 	return
	// }
	// data, err = ioutil.ReadAll(rsp.Body)
	// st := pkg.StepTwo{}
	// err = json.Unmarshal(data, &st)
	// if err != nil {
	// 	fail(c, err)
	// 	return
	// }
	// url = `https://plogin.m.jd.com/cgi-bin/m/tmauth?client_type=m&appid=300&token=` + st.Token
	// cookies = strings.Join(rsp.Header.Values("Set-Cookie"), " ")
	// okl_token := util.FetchJdCookieValue("okl_token", cookies)
	// data, _ = qrcode.Encode(url, qrcode.Medium, 256)
	// JdCookieRunners.Store(st.Token, []interface{}{cookie, okl_token})
	// log.Print("%v+++++++++++++++++++++++++", st.Token)
	// session.Set("jd_token", st.Token)
	// session.Set("jd_cookie", cookie)
	// session.Set("jd_okl_token", okl_token)
	// session.Save()
	// qrcode := &pkg.Qrcode{
	// 	Url: url,
	// 	Img: base64.StdEncoding.EncodeToString(data),
	// }
	// ok(c, resp{
	// 	"qrcode": qrcode,
	// })
}

func getjQuery(c *gin.Context) {
	jquery, err := etcode.QrcodeRepo.GetjQuery(c)
	if err != nil {
		fail(c, err)
		return
	}
	ok(c, resp{
		"jquery": jquery,
	})
	// session := sessions.Default(c)
	// if v := session.Get("jd_token"); v == nil {
	// 	ok(c, resp{
	// 		"msg": "未获取二维码，请获取二维码",
	// 	})
	// 	return
	// } else {
	// 	token := v.(string)
	// 	if v, ok1 := JdCookieRunners.Load(token); !ok1 {
	// 		log.Print("+++++++++++++++++++++++++++++++++++")
	// 		ok(c, resp{
	// 			"msg": "重新获取二维码",
	// 		})
	// 		return
	// 	} else {
	// 		if len(v.([]interface{})) >= 2 {
	// 			log.Print("+++++++++++++%v++++++++++++++++++++++", v)
	// 			ok(c, resp{
	// 				"msg": "二维码未扫描，请扫描二维码",
	// 			})
	// 			return
	// 		} else {
	// 			log.Print("+++++++++++++%v++++++++++++++++++++++", v)
	// 			pin := v.([]interface{})[0].(string)
	// 			log.Print("+++++++++++++%v++++++++++++++++++++++", pin)
	// 			session.Set("pin", pin)
	// 			session.Save()
	// 			// if note := c.GetString("note"); note != "" {
	// 			// 	if ck := models.GetJdCookie(pin); ck != nil {
	// 			// 		ck.Updates(models.Note, note)
	// 			// 	}
	// 			// }
	// 			ok(c, resp{
	// 				"msg": "已扫描二维码",
	// 			})
	// 			return
	// 		}
	// 	}
	// }
}
