package main

import (
	"cs255/logLayer/accessRedis"
	"cs255/logLayer/logServer"
	"fmt"
)

func main() {
	fmt.Println("It's my GO!!")

	redis_address := "localhost:6379"
	accessRedis.Set_v1(redis_address, "x", "1")
	accessRedis.Get_v1(redis_address, "x")

	logServer.Init_ids(0, []int{1})
	logServer.HttpServer("localhost", "8080")

}
