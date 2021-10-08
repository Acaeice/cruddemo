package util

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

var (
	RespBody       = ""                                       //返回接收
	CheckpointsURL = "https://v2-api.jsdama.com/check-points" //检查点数URL
	UploadURL      = "https://v2-api.jsdama.com/upload"       //上传图片URL
	ReportErrorURL = "https://v2-api.jsdama.com/report-error" //结果保存URL
)

// Params 参数
type Params struct {
	SoftwareId       uint   `json:"softwareId"`       //软件ID
	SoftwareSecret   string `json:"softwareSecret"`   //联众V2接口 Secret
	Username         string `json:"username"`         //联众用户名
	Password         string `json:"password"`         //联众密码
	CaptchaData      string `json:"captchaData"`      //验证数据base64结果
	CaptchaType      uint   `json:"captchaType"`      //验证类型
	CaptchaMinLength uint   `json:"captchaMinLength"` //验证长度最小值
	CaptchaMaxLength uint   `json:"captchaMaxLength"` //验证长度最大值
	WorkerTipsId     uint   `json:"workerTipsId"`     //⼈⼯提示模板ID
	CaptchaId        string `json:"captchaId"`        //识别ID
}

type Params2 struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

//Upload 图片上传
func Upload(param Params, imgUrl string) string {
	//获取网络图片
	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
	}
	defer res.Body.Close()
	// 读取获取的[]byte数据
	data, _ := ioutil.ReadAll(res.Body)
	//转Base64
	imageBase64 := base64.StdEncoding.EncodeToString(data)

	//上传图片参数
	reqBody := Params{
		SoftwareId:       param.SoftwareId,
		SoftwareSecret:   param.SoftwareSecret,
		Username:         param.Username,
		Password:         param.Password,
		CaptchaData:      imageBase64,
		CaptchaType:      param.CaptchaType,
		CaptchaMinLength: param.CaptchaMinLength,
		CaptchaMaxLength: param.CaptchaMaxLength,
	}

	//格式化JSON
	jsons, _ := json.Marshal(reqBody)
	result := string(jsons)
	jsoninfo := strings.NewReader(result)

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //如果需要测试自签名的证书 这里需要设置跳过证书检测 否则编译报错
	req, err := http.NewRequest("POST", UploadURL, jsoninfo)
	if err != nil {
		fmt.Println("err", err)
	}
	req.Header.Set("HOST", "v2-api.jsdama.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "298")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", " text/json")
	client := &http.Client{Transport: tr, Timeout: time.Second * 60}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		defer resp.Body.Close()
		body, er := ioutil.ReadAll(resp.Body)
		if er != nil {
			fmt.Println("err:", er)
		} else {
			RespBody = string(body)
		}
	}
	return RespBody
}

//ReportError 结果报错
func ReportError(param Params) string {
	//结果报错参数
	reqBody := Params{
		SoftwareId:     param.SoftwareId,
		SoftwareSecret: param.SoftwareSecret,
		Username:       param.Username,
		Password:       param.Password,
		CaptchaId:      param.CaptchaId,
	}

	//格式化JSON
	jsons, _ := json.Marshal(reqBody)
	result := string(jsons)
	jsoninfo := strings.NewReader(result)

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //如果需要测试自签名的证书 这里需要设置跳过证书检测 否则编译报错
	req, err := http.NewRequest("POST", ReportErrorURL, jsoninfo)
	if err != nil {
		fmt.Println("err", err)
	}
	req.Header.Set("HOST", "v2-api.jsdama.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "298")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", " text/json")
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		defer resp.Body.Close()
		body, er := ioutil.ReadAll(resp.Body)
		if er != nil {
			fmt.Println("err:", er)
		} else {
			RespBody = string(body)
		}
	}
	return RespBody
}

//Checkpoints 检查点数
func Checkpoints(param Params) string {
	//检查点数所需参数
	reqBody := Params{
		SoftwareId:     param.SoftwareId,
		SoftwareSecret: param.SoftwareSecret,
		Username:       param.Username,
		Password:       param.Password,
	}

	//格式化JSON
	jsons, _ := json.Marshal(reqBody)
	result := string(jsons)
	jsoninfo := strings.NewReader(result)

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //如果需要测试自签名的证书 这里需要设置跳过证书检测 否则编译报错
	req, err := http.NewRequest("POST", CheckpointsURL, jsoninfo)
	if err != nil {
		fmt.Println("err", err)
	}
	req.Header.Set("HOST", "v2-api.jsdama.com")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "298")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Content-Type", " text/json")
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		defer resp.Body.Close()
		body, er := ioutil.ReadAll(resp.Body)
		if er != nil {
			fmt.Println("err:", er)
		} else {
			RespBody = string(body)
		}
	}

	return RespBody
}

//Test_Checkpoints 检查点数测试
func Test_Checkpoints(t *testing.T) {
	param := Params{
		SoftwareId:     27183,
		SoftwareSecret: "nABWo3cavGxZYDyd2aKQILpAYGqWczzrDPSyr7uF",
		Username:       "EThead",
		Password:       "Qianxiaoyou@",
	}
	respBody := Checkpoints(param)
	t.Log(respBody)
}

//Test_Upload 上传图片测试
func Test_Upload(t *testing.T) {
	param := Params{
		SoftwareId:       27183,
		SoftwareSecret:   "nABWo3cavGxZYDyd2aKQILpAYGqWczzrDPSyr7uF",
		Username:         "EThead",
		Password:         "Qianxiaoyou@",
		CaptchaType:      1001,
		CaptchaMinLength: 0,
		CaptchaMaxLength: 0,
	}
	respBody := Upload(param, "https://api60.maidiyun.com/page/verify.jpg?key=18888888888&r=0.48232133639654595")
	t.Log(respBody)
}

//Test_ReportError 结果报错测试
func Test_ReportError(t *testing.T) {
	param := Params{
		SoftwareId:     27183,
		SoftwareSecret: "nABWo3cavGxZYDyd2aKQILpAYGqWczzrDPSyr7uF",
		Username:       "EThead",
		Password:       "Qianxiaoyou@",
		CaptchaId:      "20210923:000000000056055265756",
	}
	respBody := ReportError(param)
	t.Log(respBody)
}
