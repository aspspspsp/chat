package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
)

var (
	instance *Config
	once     sync.Once
)

type Config struct {
	System struct {
		Port        int    `yaml:"port"`
		Host        string `yaml:"host"`
		UploadModel string `yaml:"uploadModel"`
	} `yaml:"system"`
	Mysql struct {
		Addr    string `yaml:"addr"`
		Port    int    `yaml:"port"`
		User    string `yaml:"user"`
		Pass    string `yaml:"pass"`
		Name    string `yaml:"name"`
		Charset string `yaml:"charset"`
	} `yaml:"mysql"`
	Consul struct {
		ServiceName string `yaml:"serviceName"`
		ServiceAddr string `yaml:"serviceAddr"`
		ServicePort int    `yaml:"servicePort"`
	} `yaml:"consul"`
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		Port     int    `yaml:"port"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
}

// GetConfig 返回单例的 Config 实例
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		err := instance.load("server.yml")
		if err != nil {
			fmt.Printf("Error loading configs: %s\n", err)
		}
	})
	return instance
}

// load 读取并解析 YAML 文件
func (c *Config) load(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}
