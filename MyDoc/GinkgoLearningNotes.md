<H1>Ginkgo学习笔记</H1>

https://blog.gmem.cc/ginkgo-study-note

<!-- TOC -->

- [1. 简介](#1-简介)
- [2. Ginkgo](#2-ginkgo)
  - [2.1. 安装](#21-安装)
  - [2.2. 起步](#22-起步)
    - [2.2.1. 创建套件](#221-创建套件)
    - [2.2.2. 添加Spec](#222-添加spec)
    - [2.2.3. 断言失败](#223-断言失败)
    - [2.2.4. 记录日志](#224-记录日志)
    - [2.2.5. 传递参数](#225-传递参数)
  - [2.3. 测试的结构](#23-测试的结构)
    - [2.3.1. It](#231-it)
    - [2.3.2. BeforeEach](#232-beforeeach)
    - [2.3.3. AfterEach](#233-aftereach)
    - [2.3.4. Describe/Context](#234-describecontext)
    - [2.3.5. JustBeforeEach](#235-justbeforeeach)
    - [2.3.6. JustAfterEach](#236-justaftereach)
    - [2.3.7. BeforeSuite/AfterSuite](#237-beforesuiteaftersuite)
    - [2.3.8. By](#238-by)
  - [2.4. 性能测试](#24-性能测试)
  - [2.5. CLI](#25-cli)
    - [2.5.1. 运行测试](#251-运行测试)
    - [2.5.2. 传递参数](#252-传递参数)
    - [2.5.3. 跳过某些包](#253-跳过某些包)
    - [2.5.4. 超时控制](#254-超时控制)
    - [2.5.5. 调试信息](#255-调试信息)
    - [2.5.6. 其他选项](#256-其他选项)
  - [2.6. Spec Runner](#26-spec-runner)
    - [2.6.1. Pending Spec](#261-pending-spec)
    - [2.6.2. Skiping Spec](#262-skiping-spec)
    - [2.6.3. Focused Specs](#263-focused-specs)
    - [2.6.4. Parallel Specs](#264-parallel-specs)
- [3. Gomega](#3-gomega)
  - [3.1. 联用](#31-联用)
    - [3.1.1. 和Ginkgo](#311-和ginkgo)
    - [3.1.2. 和Go测试框架](#312-和go测试框架)
  - [3.2. 断言](#32-断言)
    - [3.2.1. Ω/Expect](#321-ωexpect)
    - [3.2.2. 错误处理](#322-错误处理)
    - [3.2.3. 断言注解](#323-断言注解)
    - [3.2.4. 简化输出](#324-简化输出)
  - [3.3. 异步断言](#33-异步断言)
    - [3.3.1. Eventually](#331-eventually)
    - [3.3.2. Consistently](#332-consistently)
    - [3.3.3. 修改默认间隔](#333-修改默认间隔)
  - [3.4. 内置Matcher](#34-内置matcher)
    - [3.4.1. 相等性](#341-相等性)
    - [3.4.2. 接口相容](#342-接口相容)
    - [3.4.3. 空值/零值](#343-空值零值)
    - [3.4.4. 布尔值](#344-布尔值)
    - [3.4.5. 错误](#345-错误)
    - [3.4.6. 通道](#346-通道)
    - [3.4.7. 文件](#347-文件)
    - [3.4.8. 字符串](#348-字符串)
    - [3.4.9. JSON/XML/YML](#349-jsonxmlyml)
    - [3.4.10. 集合](#3410-集合)
    - [3.4.11. 数字/时间](#3411-数字时间)
    - [3.4.12. Panic](#3412-panic)
    - [3.4.13. And/Or](#3413-andor)
  - [3.5. 自定义Matcher](#35-自定义matcher)
  - [3.6. 辅助工具](#36-辅助工具)
    - [3.6.1. ghttp](#361-ghttp)
    - [3.6.2. gbytes](#362-gbytes)
    - [3.6.3. gexec](#363-gexec)
    - [3.6.4. gstruct](#364-gstruct)

<!-- /TOC -->

# 1. 简介
Ginkgo /ˈɡɪŋkoʊ / 是Go语言的一个行为驱动开发（BDD， Behavior-Driven Development）风格的测试框架，通常和库Gomega一起使用。Ginkgo在一系列的“Specs”中描述期望的程序行为。  
Ginkgo集成了Go语言的测试机制，你可以通过`go test`来运行Ginkgo测试套件。

# 2. Ginkgo

**Ginkgo官网**  
https://onsi.github.io/ginkgo/  
https://github.com/onsi/ginkgo

## 2.1. 安装
```shell script
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

## 2.2. 起步
### 2.2.1. 创建套件
假设我们想给books包编写Ginkgo测试，则首先需要使用命令创建一个Ginkgo test suite：
```shell script
cd pkg/books
ginkgo bootstrap
```
上述命令会生成文件：  
/books/books_suite_test.go
```go
package books_test
 
import (
    // 使用点号导入，把这两个包导入到当前命名空间
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)
 
func TestBooks(t *testing.T) {
    // 将Ginkgo的Fail函数传递给Gomega，Fail函数用于标记测试失败，这是Ginkgo和Gomega唯一的交互点
    // 如果Gomega断言失败，就会调用Fail进行处理
    RegisterFailHandler(Fail)
 
    // 启动测试套件
    RunSpecs(t, "Books Suite")
}
```
现在，使用命令 `ginkgo`或者`go test`即可执行测试套件。
### 2.2.2. 添加Spec
上面的空测试套件没有什么价值，我们需要在此套接下编写测试（Spec）。虽然可以在books_suite_test.go中编写测试，但是推荐分离到独立的文件中，特别是包中有多个需要被测试的源文件的情况下。

执行命令`ginkgo generate book`可以为源文件book.go生成测试：  
/books/book_test.go
```go
package books_test
 
import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    // 为了方便，被测试包被导入当前命名空间
    . "ginkgo-study/pkg/books"
)
 
// 顶级的Describe容器
 
// Describe块用于组织Specs，其中可以包含任意数量的：
//    BeforeEach：在Spec（It块）运行之前执行，嵌套Describe时最外层BeforeEach先执行
//    AfterEach：在Spec运行之后执行，嵌套Describe时最内层AfterEach先执行
//    JustBeforeEach：在It块，所有BeforeEach之后执行
//    Measurement
 
// 可以在Describe块内嵌套Describe、Context、When块
var _ = Describe("Book", func() {
 
})
```

我们可以添加一些Specs：
```go
// 使用Describe、Context容器来组织Spec
var _ = Describe("Book", func() {
    var (
        // 通过闭包在BeforeEach和It之间共享数据
        longBook  Book
        shortBook Book
    )
    // 此函数用于初始化Spec的状态，在It块之前运行。如果存在嵌套Describe，则最
    // 外面的BeforeEach最先运行
    BeforeEach(func() {
        longBook = Book{
            Title:  "Les Miserables",
            Author: "Victor Hugo",
            Pages:  1488,
        }
 
        shortBook = Book{
            Title:  "Fox In Socks",
            Author: "Dr. Seuss",
            Pages:  24,
        }
    })
 
    Describe("Categorizing book length", func() {
        Context("With more than 300 pages", func() {
            // 通过It来创建一个Spec
            It("should be a novel", func() {
                // Gomega的Expect用于断言
                Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
            })
        })
 
        Context("With fewer than 300 pages", func() {
            It("should be a short story", func() {
                Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
            })
        })
    })
})
```

### 2.2.3. 断言失败
除了调用Gomega之外，你还可以调用Fail函数直接断言失败：
```go
Fail("Failure reason")
```

Fail会记录当前进行的测试，并且触发panic，当前Spec的后续断言不会再进行。

通常情况下Ginkgo会从panic中恢复，并继续下一个测试。但是，如果你启动了一个Goroutine，
并在其中触发了断言失败，则不会自动恢复，必须手工调用GinkgoRecover：

```go
It("panics in a goroutine", func(done Done) {
    go func() {
        // 如果doSomething返回false则下面的defer会确保从panic中恢复
        defer GinkgoRecover()
        // Ω和Expect功能相同
        Ω(doSomething()).Should(BeTrue())
 
        // 在Goroutine中需要关闭done通道
        close(done)
    }()
})
```

### 2.2.4. 记录日志
全局的GinkgoWriter可以用于写日志。默认情况下GinkgoWriter仅仅在测试失败时将日志Dump到标准输出，
以冗长模式（`ginkgo -v`或`go test -ginkgo.v`）运行Ginkgo时则会立即输出。

如果通过Ctrl + C中断测试，则Ginkgo会立即输出写入到GinkgoWriter的内容。
联用`--progress`则Ginkgo会在BeforeEach/It/AfterEach之前输出通知到GinkgoWriter，
这个特性便于诊断卡住的测试。

### 2.2.5. 传递参数
直接使用flag包即可：
```go
var myFlag string
func init() {
    flag.StringVar(&myFlag, "myFlag", "defaultvalue", "myFlag is used to control my behavior")
}
```
执行测试时使用`ginkgo -- --myFlag=xxx`传递参数。

## 2.3. 测试的结构
### 2.3.1. It
你可以在Describe、Context这两种容器块内编写Spec，每个Spec写在It块中。  
为了贴合自然语言，可以使用It的别名Specify：
```go
Describe("The foobar service", func() {
  Context("when calling Foo()", func() {
    Context("when no ID is provided", func() {
      // 应该返回ErrNoID错误
      Specify("an ErrNoID error is returned", func() {
      })
    })
  })
})
```

### 2.3.2. BeforeEach
多个Spec共享的、测试准备逻辑，可以放到BeforeEach块中。  
在BeforeEach、AfterEach块中进行断言是允许的。  
存在容器嵌套时，最外层BeforeEach先运行。

### 2.3.3. AfterEach
多个Spec共享的、测试清理逻辑，可以放到AfterEach块中。存在容器嵌套时，最内层AfterEach先运行。

### 2.3.4. Describe/Context
两者的区别：
1. Describe用于描述你的代码的一个行为
2. Context用于区分上述行为的不同情况，通常为参数不同导致

下面是一个例子：
```go
// 这是关于Book服务测试
var _ = Describe("Book", func() {
    var (
        book Book
        err error
    )
 
    BeforeEach(func() {
        book, err = NewBookFromJSON(`{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`)
    })
    // 测试加载Book行为
    Describe("loading from JSON", func() {
        // 如果正常解析JSON
        Context("when the JSON parses succesfully", func() {
            It("should populate the fields correctly", func() {
                // 期望                相等
                Expect(book.Title).To(Equal("Les Miserables"))
                Expect(book.Author).To(Equal("Victor Hugo"))
                Expect(book.Pages).To(Equal(1488))
            })
 
            It("should not error", func() {
                // 期望      没有发生错误
                Expect(err).NotTo(HaveOccurred())
            })
        })
        // 如果无法解析JSON
        Context("when the JSON fails to parse", func() {
            BeforeEach(func() {
                // 这是一个BDD反模式，可以用JustBeforeEach
                book, err = NewBookFromJSON(`{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":1488oops
                }`)
            })
 
            It("should return the zero-value for the book", func() {
                // 期望          为零
                Expect(book).To(BeZero())
            })
 
            It("should error", func() {
                // 期望        发生了错误
                Expect(err).To(HaveOccurred())
            })
        })
    })
 
    Describe("Extracting the author's last name", func() {
        It("should correctly identify and return the last name", func() {
            Expect(book.AuthorLastName()).To(Equal("Hugo"))
        })
    })
})
```

### 2.3.5. JustBeforeEach
上面的例子中，内层Spec需要尝试从无效JSON创建Book，因此它调用
NewBookFromJSON对book变量进行覆盖。这种做法是推荐的，应该使用
JustBeforeEach，这种块在任何BeforeEach执行完毕后执行：
```go
var _ = Describe("Book", func() {
    var (
        book Book
        err error
        json string
    )
    // 准备默认JSON
    BeforeEach(func() {
        json = `{
            "title":"Les Miserables",
            "author":"Victor Hugo",
            "pages":1488
        }`
    })
 
    JustBeforeEach(func() {
        // 按需，根据默认数据/无效JSON创建book，避免NewBookFromJSON的重复调用（如果代价很高的话……）
        book, err = NewBookFromJSON(json)
    })
 
    Describe("loading from JSON", func() {
        Context("when the JSON parses succesfully", func() {
        })
 
        Context("when the JSON fails to parse", func() {
            BeforeEach(func() {
                // 覆盖默认JSON为无效JSON
                json = `{
                    "title":"Les Miserables",
                    "author":"Victor Hugo",
                    "pages":1488oops
                }`
            })
        })
    })
})
```
在上面的例子中，JustBeforeEach解耦了创建（Creation）和配置（Configuration）这两个阶段。

### 2.3.6. JustAfterEach
紧跟着It之后运行，在所有AfterEach执行之前。

### 2.3.7. BeforeSuite/AfterSuite
在整个测试套件执行之前/之后，进行准备/清理。和套件代码写在一起：

```go
func TestBooks(t *testing.T) {
    RegisterFailHandler(Fail)
 
    RunSpecs(t, "Books Suite")
}
 
var _ = BeforeSuite(func() {
    dbClient = db.NewClient()
    err = dbClient.Connect(dbRunner.Address())
    Expect(err).NotTo(HaveOccurred())
})
 
var _ = AfterSuite(func() {
    dbClient.Cleanup()
})
```
这两个块都支持异步执行，只需要给函数传递一个Done参数即可。 

### 2.3.8. By
此块用于给逻辑复杂的块添加文档：
```go
var _ = Describe("Browsing the library", func() {
    BeforeEach(func() {
        By("Fetching a token and logging in")
    })
 
    It("should be a pleasant experience", func() {
        By("Entering an aisle")
    })
})
```
传递给By的字符串会发送给GinkgoWriter，如果测试失败你可以看到。  
你可以传递一个可选的函数给By，此函数会立即执行。

## 2.4. 性能测试
使用Measure块可以进行性能测试，所有It能够出现的地方，都可以使用Measure。
和It一样，Measure会生成一个新的Spec。
传递给Measure的闭包函数必须具有Benchmarker入参：
```go
Measure("it should do something hard efficiently", func(b Benchmarker) {
    // 执行一段逻辑并即时
    runtime := b.Time("runtime", func() {
        output := SomethingHard()
        Expect(output).To(Equal(17))
    })
 
    // 断言 执行时间             小于 0.2 秒
    Ω(runtime.Seconds()).Should(BeNumerically("<", 0.2), "SomethingHard() shouldn't take too long.")
 
    // 录制任意数据
    b.RecordValue("disk usage (in MB)", HowMuchDiskSpaceDidYouUse())
}, 10)
```
执行时间、你录制的任意数据的最小、最大、平均值均会在测试完毕后打印出来。

## 2.5. CLI
### 2.5.1. 运行测试
```shell script
# 运行当前目录中的测试
ginkgo
# 运行其它目录中的测试
ginkgo /path/to/package /path/to/other/package ...
 
# 递归运行所有子目录中的测试
ginkgo -r ...
```

### 2.5.2. 传递参数
传递参数给测试套件：  
```shell script
ginkgo -- PASS-THROUGHS-ARGS
```

### 2.5.3. 跳过某些包
```shell script
# 跳过某些包
ginkgo -skipPackage=PACKAGES,TO,SKIP
```

### 2.5.4. 超时控制
选项`-timeout`用于控制套件的最大运行时间，如果超过此时间仍然没有完成，认为测试失败。默认24小时。

### 2.5.5. 调试信息
| 选项            | 说明 | 
| ---------------| ---- | 
| --reportPassed | 打印通过的测试的详细信息 | 
| --v            | 冗长模式 | 
| --trace        | 打印所有错误的调用栈 | 
| --progress     | 打印进度信息 | 

### 2.5.6. 其他选项
| 选项     | 说明 | 
| --------| ---- | 
| -race   | 启用竞态条件检测 | 
| -cover  | 启用覆盖率测试 | 
| -tags   | 指定编译器标记 | 

## 2.6. Spec Runner
### 2.6.1. Pending Spec
你可以标记一个Spec或容器为Pending，这样默认情况下**不会**运行它们。
定义块时使用P或X前缀：
```go
PDescribe("some behavior", func() { ... })
PContext("some scenario", func() { ... })
PIt("some assertion")
PMeasure("some measurement")
 
XDescribe("some behavior", func() { ... })
XContext("some scenario", func() { ... })
XIt("some assertion")
XMeasure("some measurement")
```
默认情况下Ginkgo会为每个Pending的Spec打印描述信息，
使用命令行选项`--noisyPendings=false`禁止该行为。 

### 2.6.2. Skiping Spec
P或X前缀会在编译期将Spec标记为Pending，你也可以在运行期跳过特定的Spec：
```go
It("should do something, if it can", func() {
    if !someCondition {
        // 跳过此Spec，不需要Return语句
        Skip("special condition wasn't met")
    }
})
```

### 2.6.3. Focused Specs
一个很常见的需求是，可以选择运行Spec的一个子集。Ginkgo提供两种机制满足此需求：

 将容器或Spec标记为Focused，这样默认情况下Ginkgo仅仅运行Focused Spec：
```go
 FDescribe("some behavior", func() { ... })
 FContext("some scenario", func() { ... })
 FIt("some assertion", func() { ... })
```
在命令行中传递正则式：`--focus=REGEXP`或/和`--skip=REGEXP`，
则Ginkgo仅仅运行/跳过匹配的Spec

### 2.6.4. Parallel Specs
Ginkgo支持并行的运行Spec，它实现方式是，创建go test子进程并在其中运行共享队列中的Spec。  
使用`ginkgo -p`可以启用并行测试，Ginkgo会自动创建适当数量的节点（进程）。
你也可以指定节点数量：`ginkgo -nodes=N`。  
如果你的测试代码需要和外部进程交互，或者创建外部进程，在并行测试上下文中需要谨慎的处理。
最简单的方式是在BeforeSuite方法中为每个节点创建外部资源。  
如果所有Spec需要共享一个外部进程，则可以利用SynchronizedBeforeSuite、SynchronizedAfterSuite：
```go
var _ = SynchronizedBeforeSuite(func() []byte {
    // 在第一个节点中执行
    port := 4000 + config.GinkgoConfig.ParallelNode
 
    dbRunner = db.NewRunner()
    err := dbRunner.Start(port)
    Expect(err).NotTo(HaveOccurred())
 
    return []byte(dbRunner.Address())
}, func(data []byte) {
    // 在所有节点中执行
    dbAddress := string(data)
 
    dbClient = db.NewClient()
    err = dbClient.Connect(dbAddress)
    Expect(err).NotTo(HaveOccurred())
})
```
上面的例子，为所有节点创建共享的数据库，然后为每个节点创建独占的客户端。 
SynchronizedAfterSuite的回调顺序则正好相反：
```go
var _ = SynchronizedAfterSuite(func() {
    // 所有节点
    dbClient.Cleanup()
}, func() {
    // 第一个节点
    dbRunner.Stop()
}) 
```



# 3. Gomega
这是Ginkgo推荐使用的断言（Matcher）库。

## 3.1. 联用
### 3.1.1. 和Ginkgo
注册Fail处理器即可：
```go
gomega.RegisterFailHandler(ginkgo.Fail)
```

### 3.1.2. 和Go测试框架
```go
func TestFarmHasCow(t *testing.T) {
    // 创建Gomega对象
    g := NewGomegaWithT(t)
 
    f := farm.New([]string{"Cow", "Horse"})
    // 进行断言
    g.Expect(f.HasCow()).To(BeTrue(), "Farm should have cow")
}
```

## 3.2. 断言
### 3.2.1. Ω/Expect
两种断言语法本质是一样的，只是命名风格有些不同：
```go
Ω(ACTUAL).Should(Equal(EXPECTED))
Expect(ACTUAL).To(Equal(EXPECTED))
 
Ω(ACTUAL).ShouldNot(Equal(EXPECTED))
Expect(ACTUAL).NotTo(Equal(EXPECTED))
Expect(ACTUAL).ToNot(Equal(EXPECTED))
```

### 3.2.2. 错误处理
对于返回多个值的函数：
```go
func DoSomethingHard() (string, error) {}
 
result, err := DoSomethingHard()
// 断言没有发生错误
Ω(err).ShouldNot(HaveOccurred())
Ω(result).Should(Equal("foo"))
```
对于仅仅返回一个error的函数： 
```go
func DoSomethingHard() (string, error) {}
 
Ω(DoSomethingSimple()).Should(Succeed())
```

### 3.2.3. 断言注解
进行断言时，可以提供格式化字符串，这样断言失败可以方便的知道原因：
```go
Ω(ACTUAL).Should(Equal(EXPECTED), "My annotation %d", foo)
 
Expect(ACTUAL).To(Equal(EXPECTED), "My annotation %d", foo)
 
Expect(ACTUAL).To(Equal(EXPECTED), func() string { return "My annotation" })
```

### 3.2.4. 简化输出
断言失败时，Gomega打印牵涉到断言的对象的递归信息，输出可能很冗长。  
format包提供了一些全局变量，调整这些变量可以简化输出。
| 变量 = 默认值                              | 说明 | 
| ------------------------------------------| ---- | 
| format.MaxDepth = 10                      | 启用竞态条件检测 | 
| format.UseStringerRepresentation = false  | 启用覆盖率测试 | 
| format.PrintContextObjects = false        | 指定编译器标记 | 
| format.TruncatedDiff = true               | 指定编译器标记 | 

## 3.3. 异步断言
Gomega提供了两个函数，用于异步断言。  
传递给Eventually、Consistently的函数，如果返回多个值，
则第一个返回值用于匹配，其它值断言为nil或零值。

### 3.3.1. Eventually
阻塞并轮询参数，直到能通过断言：
```go
// 参数是闭包，调用函数
Eventually(func() []int {
    return thing.SliceImMonitoring
}).Should(HaveLen(2))
 
// 参数是通道，读取通道
Eventually(channel).Should(BeClosed())
Eventually(channel).Should(Receive())
 
// 参数也可以是普通变量，读取变量
Eventually(myInstance.FetchNameFromNetwork).Should(Equal("archibald"))
 
// 可以和gexec包的Session配合
Eventually(session).Should(gexec.Exit(0)) // 命令最终应当以0退出
Eventually(session.Out).Should(Say("Splines reticulated")) // 检查标准输出
```
可以指定超时、轮询间隔：
```go
Eventually(func() []int {
    return thing.SliceImMonitoring
}, TIMEOUT, POLLING_INTERVAL).Should(HaveLen(2))
```

### 3.3.2. Consistently
检查断言是否在一定时间段内总是通过：
```go
Consistently(func() []int {
    return thing.MemoryUsage()
}, DURATION, POLLING_INTERVAL).Should(BeNumerically("<", 10))
```
Consistently也可以用来断言最终不会发生的事件，例如下面的例子：
```go
Consistently(channel).ShouldNot(Receive())
```

### 3.3.3. 修改默认间隔
默认情况下，Eventually每10ms轮询一次，持续1s。
Consistently每10ms轮询一次，持续100ms。调用下面的函数修改这些默认值：
```go
SetDefaultEventuallyTimeout(t time.Duration)
SetDefaultEventuallyPollingInterval(t time.Duration)
SetDefaultConsistentlyDuration(t time.Duration)
SetDefaultConsistentlyPollingInterval(t time.Duration)
```
这些调用会影响整个测试套件。

## 3.4. 内置Matcher
### 3.4.1. 相等性 
```go
// 使用reflect.DeepEqual进行比较
// 如果ACTUAL和EXPECTED都为nil，断言会失败
Ω(ACTUAL).Should(Equal(EXPECTED))
 
// 先把ACTUAL转换为EXPECTED的类型，然后使用reflect.DeepEqual进行比较
// 应当避免用来比较数字
Ω(ACTUAL).Should(BeEquivalentTo(EXPECTED))
 
// 使用 == 进行比较
BeIdenticalTo(expected interface{})
```

### 3.4.2. 接口相容
```go
Ω(ACTUAL).Should(BeAssignableToTypeOf(EXPECTED interface))
```

### 3.4.3. 空值/零值
```go
// 断言ACTUAL为Nil
Ω(ACTUAL).Should(BeNil())
 
// 断言ACTUAL为它的类型的零值，或者是Nil
Ω(ACTUAL).Should(BeZero())
```

### 3.4.4. 布尔值
```go
Ω(ACTUAL).Should(BeTrue())
Ω(ACTUAL).Should(BeFalse())
```

### 3.4.5. 错误
```go
Ω(ACTUAL).Should(HaveOccurred())
 
err := SomethingThatMightFail()
// 没有错误
Ω(err).ShouldNot(HaveOccurred())
 
 
// 如果ACTUAL为Nil则断言成功
Ω(ACTUAL).Should(Succeed())
```
可以对错误进行细粒度的匹配：
```go
Ω(ACTUAL).Should(MatchError(EXPECTED))
```
上面的EXPECTED可以是：

1. 字符串：则断言ACTUAL.Error()与之相等
2. Matcher：则断言ACTUAL.Error()与之进行匹配
3. error：则ACTUAL和error基于reflect.DeepEqual()进行比较
4. 实现了error接口的非Nil指针，调用 errors.As(ACTUAL, EXPECTED)进行检查
不符合以上条件的EXPECTED是不允许的。

### 3.4.6. 通道
```go
// 断言通道是否关闭
// Gomega会尝试读取通道进行判断，因此你需要注意：
//    如果是缓冲通道，你需要先将通道读干净
//    如果你后续需要再次读取通道，注意此断言的影响
Ω(ACTUAL).Should(BeClosed())
Ω(ACTUAL).ShouldNot(BeClosed())
 
 
// 断言能够从通道里面读取到消息
// 此断言会立即返回，如果通道已经关闭，则下面的断言失败
Ω(ACTUAL).Should(Receive(<optionalPointer>))
 
 
 
// 断言能够无阻塞的发送消息
Ω(ACTUAL).Should(BeSent(VALUE))
```

### 3.4.7. 文件
```go
// 文件或目录存在
Ω(ACTUAL).Should(BeAnExistingFile())
// 断言是普通文件
Ω(ACTUAL).Should(BeARegularFile())
// 断言是目录
BeADirectory
```

### 3.4.8. 字符串
```go
// 子串判断                        fmt.Sprintf(STRING, ARGS...)
Ω(ACTUAL).Should(ContainSubstring(STRING, ARGS...))
 
// 前缀判断
Ω(ACTUAL).Should(HavePrefix(STRING, ARGS...))
 
// 后缀判断
Ω(ACTUAL).Should(HaveSuffix(STRING, ARGS...))
 
 
// 正则式匹配
Ω(ACTUAL).Should(MatchRegexp(STRING, ARGS...))
```

### 3.4.9. JSON/XML/YML
```go
Ω(ACTUAL).Should(MatchJSON(EXPECTED))
Ω(ACTUAL).Should(MatchXML(EXPECTED))
Ω(ACTUAL).Should(MatchYAML(EXPECTED))
```
ACTUAL、EXPECTED可以是string、[]byte、Stringer。
如果两者转换为对象是reflect.DeepEqual的则匹配。

### 3.4.10. 集合
string, array, map, chan, slice都属于集合。
```go
// 断言为空
Ω(ACTUAL).Should(BeEmpty())
 
// 断言长度
Ω(ACTUAL).Should(HaveLen(INT))
 
// 断言容量
Ω(ACTUAL).Should(HaveCap(INT))
 
// 断言包含元素
Ω(ACTUAL).Should(ContainElement(ELEMENT))
 
// 断言等于                   其中之一
Ω(ACTUAL).Should(BeElementOf(ELEMENT1, ELEMENT2, ELEMENT3, ...))
 
 
// 断言元素相同，不考虑顺序
Ω(ACTUAL).Should(ConsistOf(ELEMENT1, ELEMENT2, ELEMENT3, ...))
Ω(ACTUAL).Should(ConsistOf([]SOME_TYPE{ELEMENT1, ELEMENT2, ELEMENT3, ...}))
 
// 断言存在指定的键，仅用于map
Ω(ACTUAL).Should(HaveKey(KEY))
// 断言存在指定的键值对，仅用于map
Ω(ACTUAL).Should(HaveKeyWithValue(KEY, VALUE))
```

### 3.4.11. 数字/时间
```go
// 断言数字意义（类型不感知）上的相等
Ω(ACTUAL).Should(BeNumerically("==", EXPECTED))
 
// 断言相似，无差不超过THRESHOLD（默认1e-8）
Ω(ACTUAL).Should(BeNumerically("~", EXPECTED, <THRESHOLD>))
 
 
Ω(ACTUAL).Should(BeNumerically(">", EXPECTED))
Ω(ACTUAL).Should(BeNumerically(">=", EXPECTED))
Ω(ACTUAL).Should(BeNumerically("<", EXPECTED))
Ω(ACTUAL).Should(BeNumerically("<=", EXPECTED))
 
Ω(number).Should(BeBetween(0, 10))
```
比较时间时使用BeTemporally函数，和BeNumerically类似。 

### 3.4.12. Panic
断言会发生Panic：
```go
Ω(ACTUAL).Should(Panic())
```

### 3.4.13. And/Or
```go
Expect(number).To(SatisfyAll(
            BeNumerically(">", 0),
            BeNumerically("<", 10)))
// 或者
Expect(msg).To(And(
            Equal("Success"),
            MatchRegexp(`^Error .+$`)))
 
 
 
Ω(ACTUAL).Should(SatisfyAny(MATCHER1, MATCHER2, ...))
// 或者
Ω(ACTUAL).Should(Or(MATCHER1, MATCHER2, ...))
```

## 3.5. 自定义Matcher
如果内置Matcher无法满足需要，你可以实现接口：
```go
type GomegaMatcher interface {
    Match(actual interface{}) (success bool, err error)
    FailureMessage(actual interface{}) (message string)
    NegatedFailureMessage(actual interface{}) (message string)
}
```

## 3.6. 辅助工具
### 3.6.1. ghttp
用于测试HTTP客户端，此包提供了Mock HTTP服务器的能力。

### 3.6.2. gbytes
gbytes.Buffer实现了接口io.WriteCloser，能够捕获到内存缓冲的输入。配合使用 gbytes.Say能够对流数据进行有序的断言。

### 3.6.3. gexec
简化了外部进程的测试，可以：

1. 编译Go二进制文件
2. 启动外部进程
3. 发送信号并等待外部进程退出
4. 基于退出码进行断言
5. 将输出流导入到gbytes.Buffer进行断言

### 3.6.4. gstruct
此包用于测试复杂的Go结构，提供了结构、切片、映射、指针相关的Matcher。

对所有字段进行断言：
```go
actual := struct{
    A int
    B bool
    C string
}{5, true, "foo"}
Expect(actual).To(MatchAllFields(Fields{
    "A": BeNumerically("<", 10),
    "B": BeTrue(),
    "C": Equal("foo"),
})
```
不处理某些字段： 
```go
Expect(actual).To(MatchFields(IgnoreExtras, Fields{
    "A": BeNumerically("<", 10),
    "B": BeTrue(),
    // 忽略C字段
})
 
 
Expect(actual).To(MatchFields(IgnoreMissing, Fields{
    "A": BeNumerically("<", 10),
    "B": BeTrue(),
    "C": Equal("foo"),
    "D": Equal("bar"), // 忽略多余字段
})
```
一个复杂的例子：
```go
coreID := func(element interface{}) string {
    return strconv.Itoa(element.(CoreStats).Index)
}
Expect(actual).To(MatchAllFields(Fields{
    // 忽略此字段
    "Name":      Ignore(),
    // 时间断言
    "StartTime": BeTemporally(">=", time.Now().Add(-100 * time.Hour)),
    //     解引用后再断言
    "CPU": PointTo(MatchAllFields(Fields{
        "Time":                 BeTemporally(">=", time.Now().Add(-time.Hour)),
        "UsageNanoCores":       BeNumerically("~", 1E9, 1E8),
        "UsageCoreNanoSeconds": BeNumerically(">", 1E6),
        //       包含匹配的元素， 抽取ID的函数
        "Cores": MatchElements(coreID, IgnoreExtras, Elements{
            // ID: Matcher
            "0": MatchAllFields(Fields{
                Index: Ignore(),
                "UsageNanoCores":       BeNumerically("<", 1E9),
                "UsageCoreNanoSeconds": BeNumerically(">", 1E5),
            }),
            "1": MatchAllFields(Fields{
                Index: Ignore(),
                "UsageNanoCores":       BeNumerically("<", 1E9),
                "UsageCoreNanoSeconds": BeNumerically(">", 1E5),
            }),
        }),
    }))
    "Logs":               m.Ignore(),
}))
```
