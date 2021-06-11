package pkg

const (
	// GenderMale 男
	GenderMale = 1
	// GenderFemale 女
	GenderFemale = 2
	// GenderUnknown 未知
	GenderUnknown = 0
)

//用户表
type User struct {
	GormModel
	Nick     string `json:"nick" form:"nick"`         // 昵称，必传
	Gender   uint   `json:"gender" form:"gender"`     // 性别，必传
	Avatar   string `json:"avatar" form:"avatar"`     // 头像，必传
	AppID    string `json:"appID" form:"appID"`       // APPID
	OpenID   string `json:"openID" form:"openID"`     // 微信openID
	Language string `json:"language" from:"language"` //语言
	City     string `json:"city" from:"city"`         //所在城市
	Province string `json:"Province" from:"Province"` //所在省份
	Country  string `json:"country" from:"country"`   //所在国家
}


// MAppUser 微信用户
type MAppUser struct {
	OpenID    string    `json:"openId"`
	NickName  string    `json:"nickName"`
	Gender    int       `json:"gender"` // 性别 0：未知、1：男、2：女
	Language  string    `json:"language"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	AvatarURL string    `json:"avatarUrl"`
	Watermark watermark `json:"watermark"`
}

// watermark 验证参数
type watermark struct {
	AppID string `json:"appid"`
}
