package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var dataMap = map[string]func(detailType string) (dataType string){
	"int":     func(detailType string) (dataType string) { return "int64" },
	"tinyint": func(detailType string) (dataType string) { return "int32" },
	"json":    func(string) string { return "datatypes.JSON" },
}

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

	db := ConnectDB(conf.DSN)
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/dal/dao",
		ModelPkgPath: "./internal/dal/model",
	})
	g.UseDB(db)
	g.WithDataTypeMap(dataMap)
	g.WithJSONTagNameStrategy(func(c string) string {
		return LowerCamelCase(c)
	})
	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)
	// apply diy interfaces on structs or table models
	//g.ApplyInterface(func(model.Method) {}, g.GenerateModel("user"))
	g.Execute()
}

func ConnectDB(dsn string) *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %s", err))
	}
	return db
}

// UpperCamelCase 下划线单词转为大写驼峰单词
func UpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	causer := cases.Title(language.English)
	s = causer.String(s)
	return strings.Replace(s, " ", "", -1)
}

// LowerCamelCase 下划线单词转为小写驼峰单词
func LowerCamelCase(s string) string {
	s = UpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
