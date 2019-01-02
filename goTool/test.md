Go语言测试
===


> go test子命令是 Go语言包的测试驱动程序，这些包根据某些约定组织在一起,
在一个包目录中，以 _test.go为结尾的文件不是 go build命令编译的目标，
而是 go test编译的目标

在 *_test.go文件中，三种函数需要特殊对待

- 功能测试函数
- 基准测试函数
- 示例函数

### 1. 功能测试函数 (Test函数)

每一个测试文件必须导入 testing包，这些函数的函数签名如下：
```
func TestName(t *testing.T) {
    // ...
}
```

```
# -v 可以输出包中每个测试用例的时间 
go test -v 

# -run 的参数是一个正则表达式，它可以使 go test只运行那些测试函数名称匹配给定模式的函数
go test -run

```

> 一些常用测试方式
- 测试功能函数必须以Test开头，可选的后缀名称必须以大些字母开头
- 可以编写基于表的测试用例 （举例子，参考书中示例）
- 可以在表之外做随机测试
- 白盒测试
    - 有些函数需要调用一些比价敏感操作，我们在测试时候想要避免，可以使用如下方式
        - 可以保留原来的方法并做defer恢复，可以参看下面代码块
- 外部测试包
    - 有些包中的库会有依赖，高级的包依赖于低级包，但是测试低级包有时候需要导入高级包功能
    - 这样会造成 Go规范中禁止的 循环引用
    - 所以需要引入外部测试包, 将测试独立在一个包中
        - 包申明为 xxx_test
    - 可以将包内的方法在 _test.go中添加一些声明，将内部功能暴露给外部测试
        - 如果一个源文件的功能仅在于此，他们一般称为 export_test.go
- 编写有效测试
    - 规范：
        - 结果暂时尽量统一
        - 例如： func(...params) , want foo, got bar
- 避免脆弱测试
- 覆盖率
    - 语句覆盖率最常用
    - 但是覆盖率很少做到 100% (有些语句几乎都走不到)
    
```
// 白盒测试
func TestXxx(t *testing.T) {
    // 保留等待恢复的方法
    saved := waitRecoveredFunc
    defer func() { waitRecoveredFunc = saved }()
    
    // 设置测试的伪方法
    var waitRecoveredFunc = func(){
        // 需要替代的实现
    }

    // ...其余测试部分...
}
```

```
# -coverprofile=filename 启用覆盖数据收集
go test -run=xxx -coverprofile=c.out

# 只需要汇总信息
go test -cover

# 生成HTML报告
go tool cover -html=c.out


```



> 有一些要点要记得, 在测试代码中不要调用 log.Fatal 或者 os.Exit
因为这两个函数会阻止跟踪中的过程。


### 2. 基准测试函数 (Benchmark函数)

固定格式
```
func BenchmarkXxx(b *testing B){
    for i := 0; i < b.N, i++ {
        // Xxx方法调用代码
    }
}

```

### 3. 性能剖析
可以获取新能分析报告
```
# 输出到文件
go test -cpuprofile=cpu.out
go test -blockprofile=block.out
go test -memprofile=mem.out

```


### 4. 示例函数 (Example函数)
```
func ExampleXxx(){
    Xxx(param)
    // 输出:
    // Aaa
}
```

如果示例函数代码中包含 // 输出: 这样的一个注释，测试驱动程序将执行只鹅个函数并检查输出到终端的内容匹配这个注释中的文本
