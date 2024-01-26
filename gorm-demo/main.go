package main

import (
	"fmt"

	"github.com/OrigamiWang/msd/micro/framework"
)

func main() {
	fmt.Println("Hello, world.")
	framework.NewGinWeb()
}
