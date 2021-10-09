[iris-go 文档](https://www.iris-go.com/docs/#/?id=installation)
[iris-go 文档](https://docs.iris-go.com/iris/getting-started/installation)
[iris12 文档](https://www.kancloud.cn/terry/iris/1683304)
[go语言介绍](http://www.topgoer.com/)

[GORM V2 中文文档](https://www.kancloud.cn/sliver_horn/gorm/1861157)
[GORM V2 英文文档](https://gorm.io/docs/gorm_config.html)


### [本地https测试](https://github.com/FiloSottile/mkcert)
```
mkcert -install
mkcert local.cn "*.local.cn" localhost 127.0.0.1 ::1
```

### go 指令
1. `go clean --modcache` 清除模块缓存`$GOPATH/pkg/mod`
### 安装打包器(已设置环境变量)
```
go get -u github.com/go-bindata/go-bindata/v3/go-bindata

go-bindata -h
```

### 安装[gormt](https://github.com/xxjwxc/gormt/blob/master/README_zh_cn.md) mysql数据库转 struct 工具
需要修改根目录下的`config.yml`
```
go get -u -v github.com/xxjwxc/gormt@master

gormt -h
gormt
```

1. 修改配置文件后，需要运行打包命令
```sh
go-bindata -pkg config_data -o app/bindata/config/config_data.go config/...
```

### [go mod使用](https://www.jianshu.com/p/760c97ff644c)
## 问题
1.  go: updates to go.mod needed, disabled by -mod=readonly : packages.Load error
```
go mod tidy
```
2.
