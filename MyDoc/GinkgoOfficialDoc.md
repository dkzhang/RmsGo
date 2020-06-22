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

# Getting Started: Writing Your First Test
Ginkgo hooks into Go’s existing `testing` infrastructure. 
This allows you to run a Ginkgo suite using `go test`.

> This also means that Ginkgo tests can live alongside traditional Go testing tests. Both go test and ginkgo will run all the tests in your suite.

## Bootstrapping a Suite
To write Ginkgo tests for a package you must first bootstrap a Ginkgo test suite. 
Say you have a package named books:
```shell script
$ cd path/to/books
$ ginkgo bootstrap
```
This will generate a file named `books_suite_test.go` containing:
```go
package books_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestBooks(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Books Suite")
}
```
Let’s break this down:

- Go allows us to specify the `books_test` package alongside the `books` package. 
Using `books_test` instead of `books` allows us to respect the encapsulation of the `books` package: 
your tests will need to import `books` and access it from the outside, like any other package. 
This is preferred to reaching into the package and testing its internals and leads to more behavioral tests. 
You can, of course, opt out of this – just change package `books_test` to package `books`
- We import the `ginkgo` and `gomega` packages into the test’s top-level namespace by performing a dot-import. 
If you’d rather not do this, check out the [Avoiding Dot Imports](#Avoiding Dot Imports) section below.

#The Ginkgo CLI
##Avoiding Dot Imports