package main

import (
	"github.com/OrigamiWang/msd/gate/biz"
	"github.com/OrigamiWang/msd/micro/framework"
)

func main() {
	root := framework.NewGinWeb()
	r := root.Group("/api")

	public := r.Group("/public")
	// auth := r.Group("/auth")

	biz.Proxy("http://localhost:8081", "/manage", public)
	root.Run("0.0.0.0:8848")
}
