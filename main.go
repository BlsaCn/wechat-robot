package main

import (
	_ "wechat-robot/db"
	"wechat-robot/router"
)

func main() {
	r := router.Router()
	_ = r.Run(":8001")
}
