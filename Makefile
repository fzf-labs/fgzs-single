SHELL := /bin/bash
BASEDIR = $(shell pwd)

export GOPATH := $(shell go env GOPATH)
export GOPROXY := https://goproxy.cn/,direct

DSN="root:123456@tcp(0.0.0.0:3306)/fgzs_single"
versionDir = "pkg/version"
gitBranch = $(shell git rev-parse --abbrev-ref HEAD)
gitCommit = $(shell git rev-parse HEAD)
buildTime = $(shell TZ=Asia/Shanghai date +%Y%m%d%H%M%S)
ldflags="-w -X ${versionDir}.buildTime=${buildTime} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitBranch=${gitBranch}"

.PHONY: init
# make init  初始化安装必要扩展库
init:
	@go install github.com/zeromicro/go-zero/tools/goctl@latest
	@go install github.com/zeromicro/goctl-swagger@latest
	@goctl env check -i -f --verbose
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

.PHONY: mod
# make mod  golang库更新
mod:
	#环境更新-开始
	@go mod download
	@go mod tidy
	#环境更新-结束

.PHONY: fmt
# make fmt  格式化代码
fmt:
	@gofmt -s -w .

.PHONY: vet
# make vet golang官方命令,用于检查代码中的问题.
vet:
	@go vet ./...

.PHONY: ci-lint
# make ci-lint  golang使用最多的第三方静态程序分析工具
ci-lint:
	@golangci-lint run ./...

.PHONY: git-clean
# make git-clean  git clean
git-clean:
	#清除开始
	@git checkout --orphan latest_branch
	@git add -A
	@git commit -am "clean"
	@git branch -D ${gitBranch}
	@git branch -m ${gitBranch}
	@git push -f origin ${gitBranch}
	#清除结束

.PHONY: api
# make api  go-zero的生成代码模板工具,编写api后一键生成代码,提供开发效率
api: swagger
	@goctl api go -api ./internal/app/web/desc/web.api -dir ./internal/app/web/ -home ./deploy/goctl
	@goctl api go -api ./internal/app/admin/desc/admin.api -dir ./internal/app/admin/ -home ./deploy/goctl

.PHONY: swagger
# make swagger   go-zero中的swagger插件,生成api接口文档
swagger:
	@goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api ./internal/app/web/desc/web.api -dir ./storage/swagger/web
	@goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api ./internal/app/admin/desc/admin.api -dir ./storage/swagger/admin

.PHONY: gormgen
# make gormgen 生成gorm模型文件和Sql DDL
gormgen:
	@go run ./cmd/gormgen/main.go -dsn ${DSN}
	@go run ./cmd/sqldump/main.go -dsn ${DSN}

.PHONY: errcode
# make errcode 生成错误码
errcode:
	@go run ./cmd/errcode/main.go

.PHONY: sysapi
# make sysapi 生成管理后台接口权限表数据
sysapi:
	@go run ./cmd/sysapi/sysapi.go

.PHONY: code-update
# make code-update 代码更新
code-update:
	#代码更新-开始
	@git checkout ${gitBranch}
	@git pull origin ${gitBranch}
	@git rev-parse --abbrev-ref HEAD #打印分支
	@git rev-parse --short HEAD #打印短commit id
	#代码更新-结束

.PHONY: web-build
# make web-build  web服务编译
web-build:mod
	#编译-开始
	#配置
	@mkdir -p bin/web/etc/ && cp internal/app/web/etc/web.${gitBranch}.yaml bin/web/etc/web.yaml
	#编译代码
	@go build -v -ldflags ${ldflags} -o bin/web/fgzs_single_web internal/app/web/web.go
	#编译-结束


.PHONY: web-run
# make web-run web服务执行
web-run:
	#代码执行-开始
	@bash ./scripts/manage.sh bin/web/fgzs_single_web restart
	#代码执行-结束

.PHONY: web-start
# make web-start web服务代码更新 编译 执行
web-start: code-update mod web-build web-run


.PHONY: admin-build
# make admin-build admin服务编译
admin-build:mod
	##配置
	@mkdir -p bin/admin/etc/ && cp internal/app/admin/etc/admin.${gitBranch}.yaml bin/admin/etc/admin.yaml
	##编译
	@go build -v -ldflags ${ldflags} -o bin/admin/fgzs_single_admin internal/app/admin/admin.go

.PHONY: admin-run
# make admin-run admin服务执行
admin-run:
	#代码执行-开始
	@bash scripts/manage.sh  bin/admin/fgzs_single_admin restart
	#代码执行-结束

.PHONY: admin-start
# make admin-start admin服务 代码更新 编译 执行
admin-start: code-update mod admin-build admin-run



# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

