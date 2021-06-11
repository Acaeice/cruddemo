package main

import (
	"code.meikeland.com/wanghejun/cruddemo/api/cruddemo"
	_ "code.meikeland.com/wanghejun/cruddemo/docs"
)

// @title API
// @version 1.0
// @BasePath /api
func main() {
	cruddemo.Init()
}
	