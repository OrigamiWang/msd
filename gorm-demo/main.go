package main

import (
	"fmt"

	"github.com/OrigamiWang/msd/micro/framework"
	"github.com/OrigamiWang/msd/micro/midware"
)

func main() {
	fmt.Println("Hello, world.")
	router := framework.NewGinWeb()
	v1 := router.Group("/v1")
	{
		v1.GET("/test", midware.PostHandler())
	}
}
