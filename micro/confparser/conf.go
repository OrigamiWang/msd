package confparser

import (
	"fmt"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gopkg.in/yaml.v3"
	"os"
)

var Conf *Config // 读取配置文件的配置

type Database struct {
	Key      string                 `yaml:"key"`  // 存储在dbConn map[interface{}]int 中使用的key
	Type     string                 `yaml:"type"` // 数据库类型：MySQL、mongodb、redis
	Name     string                 `yaml:"name"`
	Host     string                 `yaml:"host"`
	Port     int                    `yaml:"port"`
	User     string                 `yaml:"user"`
	Password string                 `yaml:"password"`
	Ext      map[string]interface{} `yaml:",flow"` // 扩展
}

type Config struct {
	Dbs []Database `yaml:"databases,flow"`
}

func init() {
	Conf = LoadConf()
	if Conf == nil {
		logutil.Error("init conf failed")
	}
}

func LoadConf() *Config {
	dataBytes, err := os.ReadFile("conf.yml")
	if err != nil {
		logutil.Error("load conf failed, err: %v", err)
		return nil
	}
	config := &Config{}
	err = yaml.Unmarshal(dataBytes, config)
	if err != nil {
		logutil.Error("unmarshal config failed, err: %v", err)
		return nil
	}
	fmt.Println(config)
	return config
}
