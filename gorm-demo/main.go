package main

import (
	"fmt"

	"github.com/OrigamiWang/msd/micro/server"
)

func main() {
	fmt.Println("Hello, world.")
	server.NewGinWeb()
}
