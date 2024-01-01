package main

import (
	"fmt"
	"cs255/logLayer/accessRedis"
)

func main() {
	fmt.Println("It's my GO!!")

	redis_address := "localhost:6379"
	
	accessRedis.Set_v1(redis_address, "x", "1")

	accessRedis.Get_v1(redis_address, "x")

}
