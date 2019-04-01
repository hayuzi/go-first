# Study Golang

This repo is just for study. It records some codes I wrote.


> refer to < The Go Programming Language >

## Sequence
- [hello.go](/hello.go)
- [syntax01](/syntax01)
    - [variable.go](/syntax01/variable.go)
    - [typeDeclare.go](/syntax01/typeDeclare.go)
    - [typeInt.go](/syntax01/typeInt.go)
    - [typeFloat.go](/syntax01/typeFloat.go)
    - [typeComplex.go](/syntax01/typeComplex.go)
    - [typeBool.go](/syntax01/typeBool.go)
    - [typeBool.go](/syntax01/typeString.go)
    - [consts.go](/syntax01/consts.go)
    - [typeArray.go](/syntax01/typeArray.go)
    - [typeSlice.go](/syntax01/typeSlice.go)
    - [typeSlice.go](/syntax01/typeMap.go)
    - [typeSlice.go](/syntax01/typeStruct.go)
    - [json.go](/syntax01/json.go)
    - [TextAndHtmlTemplate.go](/syntax01/TextAndHtmlTemplate.go)
- [syntax02](/syntax02)
    - [function.go](/syntax02/function.go)
    - [objectFunc.go](/syntax02/objectFunc.go)
    - [interface.go](/syntax02/interface.go)
    - [goroutine.go](/syntax02/goroutine.go)
    - [reflection.go](/syntax02/reflection.go)
    - [lowLevel.go](/syntax02/lowLevel.go)
- [goTool](/goTool)
    - [tool.md](/goTool/tool.md)
    - [test.md](/goTool/test.md)


## Start
```
// run it
go run hello.go

// build and run
go build hello.go
./hello

```

## gofmt
```
gofmt [flags] [path ...]
```

## package

```
# 下载 golang.org/x 下的包
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net.git

## golang官方包在github上的对应位置
golang.org/x => github.com/golang
google.golang.org => github.com/google


例如:   golang.org/x/sync => github.com/golang/sync
       google.golang.org/genproto => github.com/google/go-genproto 

```

## 终极大杀器 GOPROXY 环境变量
从 Go 1.11 版本开始，官方支持了 go module 包依赖管理工具。

其实还新增了 GOPROXY 环境变量。如果设置了该变量，下载源代码时将会通过这个环境变量设置的代理地址，而不再是以前的直接从代码库下载。

goproxy.io 这个开源项目帮我们实现好了我们想要的。该项目允许开发者一键构建自己的 GOPROXY 代理服务。
同时，也提供了公用的代理服务 https://goproxy.io ，我们只需设置该环境变量即可正常下载被墙的源码包了：
export GOPROXY=https://goproxy.io
也可以通过置空这个环境变量来关闭， export GOPROXY= 。

对于 Windows 用户，可以在 PowerShell 中设置：
$env:GOPROXY = "https://goproxy.io"
最后，我们当然推荐使用 GOPROXY 这个环境变量的解决方式，前提是 Go version >= 1.11 。



## fmt.Printf的转义字符
```
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
```




