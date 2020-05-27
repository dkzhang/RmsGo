# Ginkgo学习笔记

https://blog.gmem.cc/ginkgo-study-note

## 简介
Ginkgo /ˈɡɪŋkoʊ / 是Go语言的一个行为驱动开发（BDD， Behavior-Driven Development）风格的测试框架，通常和库Gomega一起使用。Ginkgo在一系列的“Specs”中描述期望的程序行为。  
Ginkgo集成了Go语言的测试机制，你可以通过`go test`来运行Ginkgo测试套件。

**Ginkgo官网**  
https://onsi.github.io/ginkgo/  
https://github.com/onsi/ginkgo

## 安装
```shell script
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

## 起步
### 创建套件
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
### 添加Spec
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

### 断言失败
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

### 记录日志
全局的GinkgoWriter可以用于写日志。默认情况下GinkgoWriter仅仅在测试失败时将日志Dump到标准输出，
以冗长模式（`ginkgo -v`或`go test -ginkgo.v`）运行Ginkgo时则会立即输出。

如果通过Ctrl + C中断测试，则Ginkgo会立即输出写入到GinkgoWriter的内容。
联用`--progress`则Ginkgo会在BeforeEach/It/AfterEach之前输出通知到GinkgoWriter，
这个特性便于诊断卡住的测试。

### 传递参数
直接使用flag包即可：
```go
var myFlag string
func init() {
    flag.StringVar(&myFlag, "myFlag", "defaultvalue", "myFlag is used to control my behavior")
}
```
执行测试时使用`ginkgo -- --myFlag=xxx`传递参数。