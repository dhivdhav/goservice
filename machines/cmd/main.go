package main

import (
	"fmt"
	"machines/configserver/handlers"
	machineSpec "machines/configserver/spec/machine"

	"github.com/gin-gonic/gin"
)

func initService() {
	router := gin.Default()

	// this handler will return all details of a particular machine
	router.GET(machineSpec.ListPath, handlers.ListMachineFeeds)

	// this handler will create new entry in machine config svc db
	router.POST(machineSpec.CreatePath, handlers.CreateMachineFeed)

	router.Run(":8080")
}

func main() {
	fmt.Println("Starting the machine config service")
	initService()
	fmt.Println("Machine config service stopped")
}
