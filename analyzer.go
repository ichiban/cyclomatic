package cyclomatic

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer checks cyclomatic complexity of defined functions and reports if it exceeded the given limit.
var Analyzer = &analysis.Analyzer{
	Name:      "cyclomatic",
	Doc:       `check cyclomatic complexity of functions.`,
	Requires:  []*analysis.Analyzer{inspect.Analyzer},
	Run:       run,
	FactTypes: []analysis.Fact{new(complexity)},
}

var limit int // -limit flag

func init() {
	Analyzer.Flags.IntVar(&limit, "limit", 10, "limit of cyclomatic complexity")
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	var c complexity
	inspect.Nodes([]ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.IfStmt)(nil),
		(*ast.ForStmt)(nil),
		(*ast.CaseClause)(nil),
		(*ast.CommClause)(nil),
		(*ast.BinaryExpr)(nil),
	}, func(n ast.Node, push bool) bool {
		if push {
			c.add(n)

			return true
		}

		if fd, ok := n.(*ast.FuncDecl); ok {
			o := pass.TypesInfo.Defs[fd.Name]
			f := o.(*types.Func)

			c := c
			pass.ExportObjectFact(f, &c)

			if c > complexity(limit) {
				pass.Reportf(n.Pos(), "cyclomatic complexity of %s exceeded limit %d > %d", f.FullName(), c, limit)
			}
		}

		return true
	})

	return nil, nil
}

type complexity int

func (c *complexity) AFact()         {}
func (c *complexity) String() string { return fmt.Sprintf("complexity(%d)", *c) }

func (c *complexity) add(n ast.Node) {
	switch n := n.(type) {
	case *ast.FuncDecl:
		*c = 1
	case *ast.IfStmt, *ast.ForStmt, *ast.CommClause:
		*c++
	case *ast.CaseClause:
		*c += complexity(len(n.List))
	case *ast.BinaryExpr:
		switch n.Op {
		case token.LAND, token.LOR:
			*c++
		}
	}
}
