package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os/exec"
	"path/filepath"
)

type MysqlDsn struct {
	User string
	Pass string
	Host string
}

type Config struct {
	Source   *MysqlDsn
	Target   *MysqlDsn
	Database []string
}

var conf Config

func loadConfig(configFile string) {
	config := viper.New()
	config.AddConfigPath(filepath.Dir(configFile)) //设置读取的文件路径
	config.SetConfigName("config")                 //设置读取的文件名
	config.SetConfigType("yaml")                   //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	err := config.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
}

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()
	loadConfig(*configFile)

	if conf.Source == nil {
		fmt.Println("Source err")
		return
	}
	if conf.Target == nil {
		fmt.Println("Target err")
		return
	}

	//mysqldump src -uroot -p123456 -h127.0.0.1  --add-drop-table | mysql dst -uroot -p123456 -h127.0.0.1
	if len(conf.Database) > 0 {
		for _, v := range conf.Database {
			do(*conf.Source, *conf.Target, v)
		}
	}

}

func do(source MysqlDsn, target MysqlDsn, database string) {
	s := fmt.Sprintf("mysqldump %s -u%s -p%s -h%s  --add-drop-table | mysql %s -u%s -p%s -h%s", database, source.User, source.Pass, source.Host, database, target.User, target.Pass, target.Host)
	cmd := exec.Command("bash", "-c", s)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := stdout.String(), stderr.String()
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
