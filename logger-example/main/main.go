package main

import (
	"logger-example/log"
	"fmt"
	"os"
)

var logger log.Logger

func main(){
	logger, err := log.NewLogger(log.Red)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.Log("Hi There!")	
}