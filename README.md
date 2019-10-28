# cyclomatic

[Go](https://golang.org/) [cyclomatic complexity](https://en.wikipedia.org/wiki/Cyclomatic_complexity) command and [`analysis.Analyzer`](https://godoc.org/golang.org/x/tools/go/analysis).

It is important to keep your functions simple and error-free. `cyclomatic` lets you know when your functions get too complex and error-prone.

## As a command

First, install `cyclomatic` command via `go get`.

```shellsession
$ go get github.com/ichiban/cyclomatic/cmd/cyclomatic
```

Then, run `cyclomatic` command at your package directory.

```shellsession
$ cyclomatic ./...
/Users/ichiban/src/example.com/package/foo.go:100:1: cyclomatic complexity of example.com/package.Foo exceeded limit 13 > 10
/Users/ichiban/src/example.com/package/foo.go:200:1: cyclomatic complexity of example.com/package.Bar exceeded limit 22 > 10
```

Also, you can relax the limit with `-limit` flag.

```shellsession
$ cyclomatic -limit 15 ./...
/Users/ichiban/src/example.com/package/foo.go:200:1: cyclomatic complexity of example.com/package.Bar exceeded limit 22 > 15
```

## As an `analysis.Analyzer`

First, install the package via `go get`.

```shellsession
$ go get github.com/ichiban/cyclomatic
```

Then, include `cyclomatic.Analyzer` in your checker.

```go
package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shift"

	"github.com/ichiban/cyclomatic"
)

func main() {
	multichecker.Main(
		// other analyzers of your choice
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,

		cyclomatic.Analyzer,
	)
}
```

## Definition of complexity

For each function definition:
- The base complexity is 1
- When we see `if` or `for`, increment the complexity
- When we see `&&` or `||`, increment the complexity
- When we see `switch` or `select`, increase the complexity by the number of values listed after `case`s

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

This package is based on these packages but slightly differs on how to count complexity:

- https://github.com/fzipp/gocyclo
  - https://github.com/sigma/gocyclo
