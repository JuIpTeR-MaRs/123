package main

import (
	"shop_server/config"
	"shop_server/pkg/logs"
	"shop_server/pkg/mysqldb"
	"shop_server/pkg/redisdb"
	"shop_server/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化配置文件 viper
	config.Init()
	//初始化log
	logs.InitLogger(config.CONFIG.Logger.LogTypes, config.CONFIG.Logger.Dir, logs.LogEnvType(config.CONFIG.System.Mode), config.CONFIG.Logger.LogMaxAge)
	//初始化数据库链接
	mysqldb.InitMysql()
	//初始化redis
	redisdb.Init()
	gin.SetMode(config.CONFIG.System.Mode)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	routers.InitRouter(r)
	err := r.Run(":" + config.CONFIG.System.Port)
	if err != nil {
		panic(err)
	}
}
