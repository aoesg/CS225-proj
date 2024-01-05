package main

import (
	"cs255/logLayer/accessRedis"
	"cs255/logLayer/logServer"
	"fmt"
)

func main() {
	fmt.Println("It's my GO!!")

	version, _ := accessRedis.Incr_v1("localhost:50000", "test_count")
	fmt.Println("test_count:", version)

	logServer.Init_ids(0, []int{1})
	logServer.Init_local_redis("localhost:6379")

	logServer.HttpServer("localhost", "8080")

}
