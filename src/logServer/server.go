package logServer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"cs255/logLayer/accessRedis"

)

var logNode_id int
var replica_id_slice []int

func Init_ids(input_logNode_id int, input_replica_id_slice []int){
	logNode_id = input_logNode_id
	replica_id_slice = input_replica_id_slice
}


func HttpServer(ip string, port string){

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	router := gin.Default()

	// router.GET("/someGet", getting)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)
	router.POST("/read", log_read)
	router.POST("/write", log_write)

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run(ip + ":" + port)
	// router.Run(":3000") hardcode 端口号

}

func log_read(c *gin.Context){
	ip := c.Param("ip")
	port := c.Param("port")
	key := c.Param("key")

	value := accessRedis.Get_v1(ip + ":" + port, key)

}

func log_write(c *gin.Context){
	
}