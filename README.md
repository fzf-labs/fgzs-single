# fgzs-single

## 简介

`fgzs-single`使用go-zero框架进行封装，使用简单，致力于进行快速的业务研发，同时增加了更多限制，约束项目组开发成员，规避混乱无序及自由随意的编码。

# 环境定义

- dev 开发
- test 测试
- pre 预发布
- pro 生产

# 开发环境
windows： 建议使用wsl2进行开发，否则makefile和脚本需要自行处理。
mac&linux：无特殊注意事项。

# 开发步骤

1. 定义数据库中表和字段,在MakeFile中定义数据库链接`DSN="root:123456@tcp(127.0.0.1:3306)/fgzs-single"`,执行`make gormgen`
   会自动生成数据模型,Sql等文件.
2. 定义api模板文件,例在`internal/app/web/desc/system.api`中定义接口和数据模型,注意请写好注释以便生成`swagger.json`
   文件做接口文档.然后执行`make api`生成代码.
3. 在logic中找到对应的接口,编写对应的代码.

# 服务

web:给复读机,小程序,PC官网等提供接口.
admin:管理后台提供接口

# 服务编译与启动:

makefile文件编写了相关启动命令,例:
web:

   ```shell
      make web-build
      make web-run
      make web-start
   ```

admin:

   ```shell
      make admin-build
      make admin-run
      make admin-start
   ```

# 开发流程
1.数据表建立,然后执行`make gormgen`生成数据库模型文件.
```shell
make gormgen
```
代码生成在`internal/dal/dao`,`internal/dal/model`中.

2.编写接口api模板文件,模板文件位于`internal/app/web/desc`和`internal/app/admin/desc`中,然后执行`make api`生成业务代码和接口文档.
```shell
make api
```
代码生成在`internal/app`中,
swagger.json生成在`storage/swagger`中.可将该接口文件导入到`apifox`中自动生成接口文档.

3.生成错误码
```shell
make errcode
```
错误码生成在`storage/errcode`中.

4.编写业务代码
   
自行在`internal/app/web/internal/logic`和`internal/app/admin/internal/logic`下编写业务逻辑.

5.代码提交前执行ci-lint检查
```shell
make ci-lint
```
检查代码错误并修复后提交代码