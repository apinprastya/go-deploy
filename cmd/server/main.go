package main

import (
	"fmt"
	"godeploy/pkg/api"
	"godeploy/pkg/service"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Go-Deploy")

	rootPath := os.Getenv("ROOT_FOLDER")
	address := os.Getenv("SERVER_ADDRESS")
	port, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	secret := os.Getenv("SECRET")
	srv := service.NewService(rootPath)
	api := api.NewApi(secret, srv)

	if err := api.Run(address, port); err != nil {
		fmt.Errorf("can not run server: %s", err.Error())
	}
}
