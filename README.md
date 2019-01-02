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

```



