<H1>Ginkgo</H1>
<H2>A Golang BDD Testing Framework</H2>

[Ginkgo](https://github.com/onsi/ginkgo) is a Go testing framework built to help 
you efficiently write expressive and comprehensive tests using 
[Behavior-Driven Development](https://en.wikipedia.org/wiki/Behavior-driven_development) (“BDD”) style. 
It is best paired with the [Gomega](https://github.com/onsi/gomega) matcher library 
but is designed to be matcher-agnostic.

These docs are written assuming you’ll be using Gomega with Ginkgo. 
They also assume you know your way around Go and have a good mental model for 
how Go organizes packages under `$GOPATH`.

<H2>Support Policy</H2>
Ginkgo provides support for versions of Go that are noted by the 
[Go release policy](https://golang.org/doc/devel/release.html#policy) i.e. N and N-1 major versions.

# Getting Ginkgo
Just go get it:
```shell script
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

This fetches `ginkgo` and installs the ginkgo executable 
under `$GOPATH/bin` – you’ll want that on your `$PATH`.

Ginkgo is tested against Go v1.6 and newer

The above commands also install the entire gomega library. 
If you want to fetch only the packages needed by your tests, 
import the packages you need and use `go get -t`.

For example, import the gomega package in your test code:
```go
import "github.com/onsi/gomega"
```

Use `go get -t` to retrieve the packages referenced in your test code:
```shell script
$ cd /path/to/my/app
$ go get -t ./...
```