package logServer

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HttpServer(ip string, port string, logNode_id int, replica_ids []int){

	router := gin.Default()


}