package accessRedis

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func checkError(err error) {
	if err != nil && err != redis.Nil {
		fmt.Fprintf(os.Stderr, "check error: %s", err.Error())
		panic(err)
	}
}

func Example() {
	fmt.Println("Print an Example")
}
