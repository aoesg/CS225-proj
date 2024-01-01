package logServer

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "check error: %s", err.Error())
		panic(err)
	}
}