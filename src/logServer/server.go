package logServer

import (
	"fmt"
	"net/http"
	"strings"

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

func resp_wzLog(c *gin.Context, value string, status int, message string) {
	c.JSON(http.StatusServiceUnavailable, gin.H{
		"data":    value,
		"status":  int(status),
		"message": message,
	})
}

func set_handler(c *gin.Context) {
	db_address := c.Param("db_address")
	key := c.Param("key")
	value := c.Param("value")
	ssf_id := c.Param("ssf_id")
	step_id := c.Param("step_id")

	fmt.Printf("ssf_id=%s write %s:%s in %s at step=%s\n", ssf_id, key, value, db_address, step_id)

	// Start wzLog

	// get本地(db_address:key:ssf_id:step_id)
	log_key := fmt.Sprintf("%s:%s:%s:%s", db_address, key, ssf_id, step_id)
	value_version_string, err := accessRedis.Get_v1(local_redis_address, log_key) // get 本地
	value_version_arr := strings.Split(value_version_string, ":")

	// if 本地(db_address:key:ssf_id:step_id)存在
	if err != redis.Nil {
		value := value_version_arr[0]
		version := value_version_arr[1]

		if value != "" {
			// if value 存在
			resp_wzLog(c, "", 1, "set success")
		} else {
			// if 只有 version
			db_key := fmt.Sprintf("%s:%s", key, version)
			err = accessRedis.Set_v1(db_address, db_key, value) // set 远程db (key:version)
			log_record := fmt.Sprintf("%s:%s", value, version)
			err = accessRedis.Set_v1(local_redis_address, log_key, log_record) // set 本地 "value:version"

			resp_wzLog(c, "", 1, "set success")
		}
	} else {
		// if value&version 都不存在，本地无记录
		db_version_key := key
		version, _ := accessRedis.Incr_v1(db_address, db_version_key) // incr 远程db (key)
		log_record_onlyVersion := fmt.Sprintf(":%s", version)
		err = accessRedis.Set_v1(local_redis_address, log_key, log_record_onlyVersion) // set 本地 ":version"

		db_key := fmt.Sprintf("%s:%s", key, version)
		err = accessRedis.Set_v1(db_address, db_key, value) // set 远程db (key:version)
		log_record := fmt.Sprintf("%s:%s", value, version)
		err = accessRedis.Set_v1(local_redis_address, log_key, log_record) // set 本地 "value:version"

		resp_wzLog(c, "", 1, "set success")

	}

	version, _ := accessRedis.Incr_v1(db_address, key)

	if value != "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"value":  value,
			"status": "success",
		})
	}

	db_key := fmt.Sprintf("%s:%s", db_address, key, ssf_id, step_id)

	err := accessRedis.Set_v1(db_address, key, value)

	// Start wzLog
}

func get_handler(c *gin.Context) {
	db_address := c.Param("db_address")
	key := c.Param("key")
	// value := c.Param("value")
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
