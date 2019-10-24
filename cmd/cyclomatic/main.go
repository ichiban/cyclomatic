package main

import (
	"github.com/ichiban/cyclomatic"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(cyclomatic.Analyzer)
}
