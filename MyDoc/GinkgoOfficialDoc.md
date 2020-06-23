<H1>Ginkgo</H1>
<H2>A Golang BDD Testing Framework</H2>

[Ginkgo](https://github.com/onsi/ginkgo) is a Go testing framework built to help you efficiently write expressive and comprehensive tests using [Behavior-Driven Development](https://en.wikipedia.org/wiki/Behavior-driven_development) (“BDD”) style.  It is best paired with the [Gomega](https://github.com/onsi/gomega) matcher library but is designed to be matcher-agnostic. These docs are written assuming you’ll be using Gomega with Ginkgo. They also assume you know your way around Go and have a good mental model for how Go organizes packages under `$GOPATH`.

<H2>Support Policy</H2>
Ginkgo provides support for versions of Go that are noted by the [Go release policy](https://golang.org/doc/devel/release.html#policy) i.e. N and N-1 major versions.



# Getting Ginkgo

Just go get it:
```shell script
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

This fetches ginkgo and installs the `ginkgo` executable under `$GOPATH/bin` – you’ll want that on your `$PATH`.

**Ginkgo is tested against Go v1.6 and newer** To install Go, follow the [installation instructions](https://golang.org/doc/install)

The above commands also install the entire gomega library. If you want to fetch only the packages needed by your tests, import the packages you need and use `go get -t`.

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
Ginkgo hooks into Go’s existing `testing` infrastructure. This allows you to run a Ginkgo suite using `go test`.

> This also means that Ginkgo tests can live alongside traditional Go testing tests. Both go test and ginkgo will run all the tests in your suite.

## Bootstrapping a Suite
To write Ginkgo tests for a package you must first bootstrap a Ginkgo test suite.  Say you have a package named books:
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

- Go allows us to specify the `books_test` package alongside the `books` package.  Using `books_test` instead of `books` allows us to respect the encapsulation of the `books` package:  your tests will need to import `books` and access it from the outside, like any other package. 
This is preferred to reaching into the package and testing its internals and leads to more behavioral tests.  You can, of course, opt out of this – just change package `books_test` to package `books`
- We import the `ginkgo` and `gomega` packages into the test’s top-level namespace by performing a dot-import.  If you’d rather not do this, check out the [Avoiding Dot Imports](#Avoiding Dot Imports) section below.
- `TestBooks` is a `testing` test. The Go test runner will run this function when you run `go test` or `ginkgo`.
- `RegisterFailHandler(Fail)`: A Ginkgo test signals failure by calling Ginkgo’s `Fail(description string)` function. We pass this function to Gomega using `RegisterFailHandler`. This is the sole connection point between Ginkgo and Gomega.
- `RunSpecs(t *testing.T, suiteDescription string)` tells Ginkgo to start the test suite. Ginkgo will automatically fail the `testing.T` if any of your specs fail.

At this point you can run your suite:

```shell
$ ginkgo #or go test

=== RUN TestBootstrap

Running Suite: Books Suite
==========================
Random Seed: 1378936983

Will run 0 of 0 specs


Ran 0 of 0 Specs in 0.000 seconds
SUCCESS! -- 0 Passed | 0 Failed | 0 Pending | 0 Skipped

--- PASS: TestBootstrap (0.00 seconds)
PASS
ok      books   0.019s
```



## Adding Specs to a Suite

An empty test suite is not very interesting. While you can start to add tests directly into `books_suite_test.go` you’ll probably prefer to separate your tests into separate files (especially for packages with multiple files). Let’s add a test file for our `book.go` model:

```shell
$ ginkgo generate book
```

This will generate a file named `book_test.go` containing:

```go
package books_test

import (
    "/path/to/books"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {

})
```

Let’s break this down:



# The Ginkgo CLI

## Avoiding Dot Imports

Ginkgo and Gomega provide a DSL and, by default, the `ginkgo bootstrap` and `ginkgo generate` commands import both packages into the top-level namespace using dot imports.

There are certain, rare, cases where you need to avoid this. For example, your code may define methods with names that conflict with the methods defined in Ginkgo and/or Gomega. In such cases you can either import your code into its own namespace (i.e. drop the `.` in front of your package import). Or, you can drop the `.` in front of Ginkgo and/or Gomega. The latter comes at the cost of constantly having to preface your `Describe`s and `It`s with `ginkgo.` and your `Expect`s and `ContainSubstring`s with `gomega.`.

There is a *third* option that the ginkgo CLI provides, however. If you need to (or simply want to!) avoid dot imports you can:

```shell
$ ginkgo bootstrap --nodot
```

and

```shell
$ ginkgo generate --nodot <filename>
```

This will create a bootstrap file that *explicitly* imports all the exported identifiers in Ginkgo and Gomega into the top level namespace. This happens at the bottom of your bootstrap file and generates code that looks something like:

```go
import (
    github.com/onsi/ginkgo
    ...
)

...

// Declarations for Ginkgo DSL
var Describe = ginkgo.Describe
var Context = ginkgo.Context
var It = ginkgo.It
// etc...
```

This allows you to write tests using `Describe`, `Context`, and `It` without dot imports and without the `ginkgo.` prefix. Crucially, it also allows you to redefine any conflicting identifiers (or even cook up your own semantics!). For example:

```go
var _ = ginkgo.Describe
var When = ginkgo.Context
var Then = ginkgo.It
```

This will avoid importing `Describe` and will rename `Context` to `When` and `It` to `Then`.

As new matchers are added to Gomega you may need to update the set of imports identifiers. You can do this by entering the directory containing your bootstrap file and running:

```shell
$ ginkgo nodot
```

this will update the imports, preserving any renames that you’ve provided.