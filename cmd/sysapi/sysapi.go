package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
)

var fileName = "internal/app/admin/desc/admin.api"

func main() {
	fmt.Println("sysapi 生成开始")
	do(fileName)
	fmt.Println("sysapi 生成结束")
}

func do(file string) {
	apis := ParseContent(file)
	if len(apis) > 0 {
		sql := "INSERT INTO sys_api (`group`,`method`,`path`,`desc`) VALUES"
		for _, api := range apis {
			sql += fmt.Sprintf("('%s','%s','%s','%s'),", api.Group, api.Method, api.Path, api.Desc)
		}
		err := WriteWithIo("./storage/sysapi/sysapi.sql", strings.Trim(sql, ","))
		if err != nil {
			return
		}
	}
}

type Api struct {
	Group  string `json:"group"`  //分组
	Method string `json:"method"` //方法
	Path   string `json:"path"`   //路径
	Desc   string `json:"desc"`   //描述
}

func ParseContent(api string) []Api {
	apis := make([]Api, 0)
	sp, err := parser.Parse(api)
	if err != nil {
		return nil
	}

	if len(sp.Service.Groups) > 0 {
		for _, group := range sp.Service.Groups {
			groupName := group.Annotation.Properties["group"]
			if len(group.Routes) > 0 {
				for _, route := range group.Routes {
					method := route.Method
					path := route.Path
					desc := strings.Trim(route.AtDoc.Properties["summary"], "\"")
					a := Api{
						Group:  groupName,
						Method: strings.ToUpper(method),
						Path:   "/" + groupName + path,
						Desc:   strings.TrimSpace(strings.TrimLeft(desc, "/")),
					}
					apis = append(apis, a)
				}
			}
		}
	}
	return apis
}

// WriteWithIo 使用io.WriteString()函数进行数据的写入，不存在则创建
func WriteWithIo(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0665)
	if err != nil {
		return err
	}
	defer file.Close()

	if content != "" {
		_, err := io.WriteString(file, content)
		if err != nil {
			return err
		}
	}
	return nil
}
