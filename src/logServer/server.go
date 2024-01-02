package logServer

import (
	"cs255/logLayer/accessRedis"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var logNode_id int
var replica_id_slice []int

func Init_ids(input_logNode_id int, input_replica_id_slice []int) {
	logNode_id = input_logNode_id
	replica_id_slice = input_replica_id_slice
}

func HttpServer(ip string, port string) {

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
	router.GET("/read", get_wzLog)
	router.GET("/write", set_wzLog)

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run(ip + ":" + port)
	// router.Run(":3000") hardcode 端口号

}

func get_wzLog(c *gin.Context) {
	ip := c.Param("ip")
	port := c.Param("port")
	key := c.Param("key")
	// value := c.Param("value")
	timestamp := c.Param("timestamp")
	ssf_id := c.Param("ssf_id")
	step_id := c.Param("step_id")

	fmt.Printf("ssf_id=%s read (%s, t=%s) at step=%s\n", ssf_id, key, timestamp, step_id)

	value, err := accessRedis.Get_v1(ip+":"+port, key)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"value":  value,
			"status": "fail",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"value":  value,
		"status": "success",
	})
}

func set_wzLog(c *gin.Context) {
	ip := c.Param("ip")
	port := c.Param("port")
	key := c.Param("key")
	value := c.Param("value")
	timestamp := c.Param("timestamp")
	ssf_id := c.Param("ssf_id")
	step_id := c.Param("step_id")

	fmt.Printf("ssf_id=%s write (%s, t=%s):%s at step=%s\n", ssf_id, key, timestamp, value, step_id)

	err := accessRedis.Set_v1(ip+":"+port, key, value)

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "fail",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
