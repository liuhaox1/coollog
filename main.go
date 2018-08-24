package main

import (
	"github.com/cool/coollog/logger"
)
func main(){
	var log *logger.Log=logger.GetLog("ssss")
	log.Info("sasa")
}