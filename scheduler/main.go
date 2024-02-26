package main

import (
	"fmt"

	"github.com/OrigamiWang/msd/scheduler/model"
)

func main() {
	// 调用shell启动manage
	const INSTANCE_NUM = 3
	const INSTANCE_PORT = 8081
	const MANAGE_PATH = "~/study/msd/manage"
	// init
	conf := []*model.InstanceConf{}
	for i := 1; i <= INSTANCE_NUM; i++ {
		conf = append(conf, &model.InstanceConf{
			Ip:          "127.0.0.1",
			Port:        INSTANCE_PORT + i,
			InstanceId:  i,
			ProjectPath: MANAGE_PATH,
		})
	}

	// call shell to start instance
	for _, inst := range conf {
		cmsStr := fmt.Sprintf("cd %s && go run main.go -ip=%s -port=%d", inst.ProjectPath, inst.Ip, inst.Port)
		fmt.Println(cmsStr)
	}
}
