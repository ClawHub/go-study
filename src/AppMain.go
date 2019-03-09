package main

import (
	"fmt"
	"test-log/src/log"
)

func main() {
	fmt.Println("init main")
	log.MainLogger.Debug("hello main Debug")
	log.MainLogger.Info("hello main Info")
	log.GatewayLogger.Debug("Hi Gateway Im Debug")
	log.GatewayLogger.Info("Hi Gateway  Im Info")
}
