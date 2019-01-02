go tool
===

### 1. go 工具
```sh
go help 查看内置的文档


```


### 2. 项目目录组织

```
GOPATH/
    |__ src (包导入路径)
    |__ bin
    |__ pkg (编译后的包的位置)

```


### 3. 包下载

```
go get

go get github.com/gplang/lint/golint

go get -u 获取最新版本

```


### 4. 包的构建
```
go build

```


### 5. 包的文档化
> 文档注释
包还可以有 doc.go 的单独文件来记录特别长的注释

```
go doc time

# 浏览自己的包文档，运行一个 godoc 实例
godoc -http :8000

# 加上 -analysis=type 和  -analysis=pointer 标记使文档内容丰富，同时提供源代码的高级静态分析结果

```


### 6. 包的查询

```
go list

```


### 7. go test 工具

参考 [test.md](./test.md)
