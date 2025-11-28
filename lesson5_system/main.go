package main

import (
	"system/dao"
	"system/routers"
)

func main() {
	dao.InitDB()

	r := routers.SetRouter()

	r.Run(":80")
}
