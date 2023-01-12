package main

import (
	"database/sql"
	"fgzs-single/pkg/util/fileutil"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type Config struct {
	DSN string
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
	database := conf.DSN[strings.LastIndex(conf.DSN, "/")+1:]
	//连接数据库
	db := ConnectDB(conf.DSN)
	defer db.Close()
	var tables []string
	//查所有的table
	rows, err := db.Query(fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_schema='%s'", database))
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			return
		}
		tables = append(tables, table)
	}
	type Result struct {
		Table       string `json:"table"`
		CreateTable string `json:"create_table"`
	}
	str := ""
	for _, v := range tables {
		var res Result
		err := db.QueryRow(fmt.Sprintf("SHOW CREATE TABLE %s", v)).Scan(&res.Table, &res.CreateTable)
		if err != nil {
			fmt.Println(err)
			return
		}
		if res.CreateTable != "" {
			str += res.CreateTable + ";\n"
		}
	}
	p := "./storage/sql/" + database + ".sql"
	err = fileutil.WriteContentCover(p, str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("SQL CREATE TABLE 导出成功")
}

func ConnectDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return db
}
