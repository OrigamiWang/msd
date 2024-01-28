package confparser

import (
	"fmt"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strings"
	"time"
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
	EXT      map[string]interface{} `yaml:",flow"` // 扩展
}

type Config struct {
	Dbs []Database             `yaml:"databases,flow"`
	EXT map[string]interface{} `yaml:"ext,flow"`
}

func init() {
	Conf = LoadConf()
}

func LoadConf() *Config {
	fileName := "conf.yml"
	if _, err := os.Stat(fileName); err != nil {
		logutil.Warn("config file not exist, fileName: %v", fileName)
		return nil
	}
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
	return config
}

func (c *Config) Ext(keys string, defaultVal ...interface{}) interface{} {
	res, err := c.ExtSep(keys, ".")
	if err != nil || res == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		} else {
			panic(err)
		}
	}
	return res
}

func (c *Config) ExtSep(keys, sep string) (interface{}, error) {
	ks := strings.Split(keys, sep)
	var res interface{}
	var success, isFinal bool
	res = c.EXT
	for _, k := range ks {
		res, success, isFinal = find(res, k)
		if !success {
			return "", fmt.Errorf("no such key: %v", k)
		} else if isFinal {
			break
		}
	}
	if success {
		return res, nil
	} else {
		return "", fmt.Errorf("not found the key: %v", keys)
	}
}

func find(v interface{}, key interface{}) (res interface{}, success, isFinal bool) {
	switch m := v.(type) {
	case map[string]interface{}:
		res, success = m[key.(string)]
		isFinal = reflect.TypeOf(res) != nil && reflect.TypeOf(res).Kind() != reflect.Map
	case map[interface{}]interface{}:
		res, success = m[key]
		isFinal = reflect.TypeOf(res) != nil && reflect.TypeOf(res).Kind() != reflect.Map
	}
	return
}

// shortcut of Conf.ExtString
func ExtString(keys string, defaultVal ...interface{}) string {
	return Conf.ExtString(keys, defaultVal...)
}

func ExtInt(key string, defaultVal ...interface{}) int {
	return Conf.ExtInt(key, defaultVal...)
}
func ExtFloat32(key string, defaultVal ...interface{}) float32 {
	return Conf.ExtFloat32(key, defaultVal...)
}
func ExtFloat64(key string, defaultVal ...interface{}) float64 {
	return Conf.ExtFloat64(key, defaultVal...)
}
func ExtBool(key string, defaultVal ...interface{}) bool {
	return Conf.ExtBool(key, defaultVal...)
}

func ExtDuration(key string, defaultVal ...interface{}) time.Duration {
	return Conf.ExtDuration(key, defaultVal...)
}

func (c *Config) ExtString(keys string, defaultVal ...interface{}) string {
	return c.Ext(keys, defaultVal...).(string)
}

func (c *Config) ExtInt(key string, defaultVal ...interface{}) int {
	return c.Ext(key, defaultVal...).(int)
}
func (c *Config) ExtFloat32(key string, defaultVal ...interface{}) float32 {
	return c.Ext(key, defaultVal...).(float32)
}
func (c *Config) ExtFloat64(key string, defaultVal ...interface{}) float64 {
	return c.Ext(key, defaultVal...).(float64)
}
func (c *Config) ExtBool(key string, defaultVal ...interface{}) bool {
	return c.Ext(key, defaultVal...).(bool)
}

func (c *Config) ExtDuration(key string, defaultVal ...interface{}) time.Duration {
	str := fmt.Sprintf("%v", c.Ext(key, defaultVal...))
	t, _ := time.ParseDuration(str)
	return t
}

func (d *Database) Ext(key string, defaultVal ...interface{}) interface{} {
	if res, exist := d.EXT[key]; exist {
		return res
	}
	if len(defaultVal) > 0 {
		return defaultVal[0]
	}
	logutil.Error("the key is not exist: %v", key)
	return ""
}

func (d *Database) ExtString(key string, defaultVal ...interface{}) string {
	return d.Ext(key, defaultVal...).(string)
}
func (d *Database) ExtInt(key string, defaultVal ...interface{}) int {
	return d.Ext(key, defaultVal...).(int)
}
func (d *Database) ExtFloat32(key string, defaultVal ...interface{}) float32 {
	return d.Ext(key, defaultVal...).(float32)
}
func (d *Database) ExtFloat64(key string, defaultVal ...interface{}) float64 {
	return d.Ext(key, defaultVal...).(float64)
}
func (d *Database) ExtBool(key string, defaultVal ...interface{}) bool {
	return d.Ext(key, defaultVal...).(bool)
}
func (d *Database) ExtDuration(key string, defaultVal ...interface{}) time.Duration {
	str := fmt.Sprintf("%v", d.Ext(key, defaultVal...))
	t, _ := time.ParseDuration(str)
	return t
}
