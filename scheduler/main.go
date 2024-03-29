package main

import (
	"fmt"
	"os/exec"

	"github.com/OrigamiWang/msd/scheduler/model"
)

func main() {
	// 调用shell启动manage
	const INSTANCE_NUM = 3
	const INSTANCE_PORT = 8081
	const MANAGE_PATH = "~/study/msd/manage"
	// init
	conf := []*model.InstanceConf{}
	for i := 0; i < INSTANCE_NUM; i++ {
		conf = append(conf, &model.InstanceConf{
			Ip:          "127.0.0.1",
			Port:        INSTANCE_PORT + i,
			InstanceId:  i,
			ProjectPath: MANAGE_PATH,
		})
	}

	// call shell to start instance
	for _, inst := range conf {
		cmdStr := fmt.Sprintf("cd %s && go run main.go -ip=%s -port=%d -instance_id=%d", inst.ProjectPath, inst.Ip, inst.Port, inst.InstanceId)
		cmd := exec.Command("sh", "-c", cmdStr)
		if err := cmd.Start(); err != nil {
			fmt.Printf("启动项目 %s 时出错: %v\n", inst.ProjectPath, err)
			continue
		}
		fmt.Printf("项目 %s 已启动，监听地址：%s，端口：%d\n", inst.ProjectPath, inst.Ip, inst.Port)
	}
}
