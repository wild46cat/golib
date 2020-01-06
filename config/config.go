package config

import (
	"github.com/jinzhu/configor"
	"log"
	"os"
)

func init() {
	configFilePath := "config.toml"
	err := configor.Load(&Configuration, configFilePath)
	if err != nil {
		log.Fatalf("load config fail from path: %v \n", configFilePath)
		os.Exit(1)
	}
}

// 配置结构
var Configuration struct {
	// 日志配置
	Logger struct {
		LogFilePath             string `default:""`                        //日志文件路径
		LogFileName             string `default:"game-server.log"`         //日志文件名称
		LogFileNameSuffixFormat string `default:".%Y%m%d"`                 //日志文件名称后缀格式化
		LogLevel                string `default:"info"`                    //日志级别
		LogFormat               string `default:"text"`                    //日志文件格式
		LogDatetimeFormat       string `default:"2006-01-02 15:04:05.000"` //时间戳格式化
		LogMaxAgeHour           int    `default:"720"`                     //日志最大保存时间（小时）
		LogRotationTimeHour     int    `default:"24"`                      //日志滚动时间（小时）
	}
	// Redis配置
	Redis struct {
		Addrs        []string `default:["127.0.0.1:6379"]` //Redis地址
		Password     string   `default:""`                 //Redis鉴权密码
		PoolSize     int      `default:"100"`
		DialTimeout  int      `default:"2000"`
		ReadTimeout  int      `default:"200"`
		WriteTimeout int      `default:"200"`
		MinIdleConns int      `default:"10"`
		PoolTimeOut  int      `default:"2000"`
		IdleTimeOut  int      `default:"2000"`
	}

	Rabbitmq struct {
		User     string `default:"guest"`
		Password string `default:"guest"`
		Host     string `default:"127.0.0.1"`
		Port     int    `default:"5672"`
	}

	Mysql struct {
		UserName string `default:"root"`
		Password string `default:"123456"`
		Addr     string `default:"127.0.0.1:3306"`
		Db       string `default:""`
	}

	Kafka struct {
		BrokerList     string `default:"localhost:9092"`
		DealyTopicName string `default:"test"`
		GroupId        string `default:"consumerGroupId-1"`
	}
}
