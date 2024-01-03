package logServer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"cs255/logLayer/accessRedis"
)

var logNode_id int
var replica_id_slice []int

func Init_ids(input_logNode_id int, input_replica_id_slice []int) {
	logNode_id = input_logNode_id
	replica_id_slice = input_replica_id_slice
}

var local_redis_address string

func Init_local_redis(address string) {
	local_redis_address = address
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

	router.GET("/write", set_handler)
	router.GET("/read", get_handler)

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。
	router.Run(ip + ":" + port)
	// router.Run(":3000") hardcode 端口号

}

func set_handler(c *gin.Context) {
	ip := c.Param("ip")
	port := c.Param("port")
	key := c.Param("key")
	value := c.Param("value")
	version := c.Param("version")
	ssf_id := c.Param("ssf_id")
	step_id := c.Param("step_id")

	// Start wzLog

	fmt.Printf("ssf_id=%s write (%s, ver=%s):%s at step=%s\n", ssf_id, key, version, value, step_id)

	err := accessRedis.Set_v1(ip+":"+port, key, value)

	// Start wzLog

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "fail",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func get_handler(c *gin.Context) {
	db_address := c.Param("db_address")
	key := c.Param("key")
	// value := c.Param("value")
	// version := c.Param("version")
	ssf_id := c.Param("ssf_id")
	step_id := c.Param("step_id")

	// Start wzLog

	fmt.Printf("ssf_id=%s read %s at step=%s\n", ssf_id, key, step_id)

	version, err := accessRedis.Get_v1(local_redis_address, key) //从本地获取version
	// if version == ""{
	// 	version
	// }

	log_key := fmt.Sprintf("%s:%s:%s:%s", key, version, ssf_id, step_id)
	db_key := fmt.Sprintf("%s:%s", key, version)

	value, err := accessRedis.Get_v1(local_redis_address, log_key)

	if value == "" {
		value, err = accessRedis.Get_v1(db_address, db_key) // 远程数据库
	}

	var status string
	if err == redis.Nil {
		status = "notExist"
	} else {
		status = "success"
		err = accessRedis.Set_v1(local_redis_address, log_key, value)
	}

	// End wzLog

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"value":   value,
			"version": version,
			"status":  status,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"value":  value,
		"status": status,
	})
}
