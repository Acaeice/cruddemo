package pkg

type JdBound struct {
	GormModel
	JDCookie string `json:"jdCookie" gorm:"size:100"` // Cookie
	JDPin    string `json:"jdPin" gorm:"size:100"`    // JDPin
}
