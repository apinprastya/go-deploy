package main

import (
	"godeploy/pkg/client"
	"os"
)

func main() {
	rootUrl := os.Getenv("ROOT_URL")

	cl := client.NewClient(os.Args[1], os.Args[2], rootUrl, os.Getenv("SECRET"))
	if err := cl.Run(); err != nil {
		panic(err)
	}

}
