package fourcetypeassert

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "fourcetypeassert",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "fourcetypeassert is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.AssignStmt:
			if len(n.Rhs) == 0 {
				return
			}
			switch r := n.Rhs[0].(type) {
			case *ast.TypeAssertExpr:
				if r.Type != nil && len(n.Lhs) != 2 {
					pass.Reportf(n.Pos(), "must not do fource type assertion")
				}
			}
		}
	})

	return nil, nil
}
