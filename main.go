package main

import (
	"aries/config/app"
	"aries/config/setting"
	_ "aries/docs"
	"aries/log"
	"github.com/pkg/profile"
)

// @title Gin Swagger
// @version 1.0
// @description 基于Golang博客系统  API 接口文档
// @host localhost:8088
func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	// go tool pprof -http=:9999 mem.pprof
	// go-wrk -t=8 -c=100 -n=10000 "http://127.0.0.1:8088"
	//f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	//defer func(f *os.File) {
	//	err := f.Close()
	//	if err != nil {
	//
	//	}
	//}(f)
	//err := pprof.StartCPUProfile(f)
	//if err != nil {
	//	return
	//}
	//defer pprof.StopCPUProfile()
	engine := app.InitApp()                            // 初始化
	err := engine.Run(":" + setting.Config.Server.Port) // 运行
	if err != nil {
		log.Logger.Sugar().Panic("项目启动失败: ", err.Error())
	}

}
