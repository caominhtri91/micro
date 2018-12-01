package main

import (
	"fmt"

	"github.com/caominhtri91/micro/accountservice/dbclient"

	"github.com/caominhtri91/micro/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDB()
	service.DBClient.Seed()
}
