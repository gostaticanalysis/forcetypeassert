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
			if !hasTypeAssertion(n.Rhs) {
				return
			}
			// if right hand has 2 or more values, assign statement can't assert boolean value which describes type assertion is succeeded
			if len(n.Rhs) > 1 {
				pass.Reportf(n.Pos(), "right hand must be only type assertion")
				return
			}
			if len(n.Lhs) == 2 {
				return
			}

			tae, ok := n.Rhs[0].(*ast.TypeAssertExpr)
			if !ok {
				pass.Reportf(n.Pos(), "right hand is not TypeAssertion")
				return
			}
			if tae.Type == nil {
				return
			}
			pass.Reportf(n.Pos(), "type assertion must be checked")
		}
	})

	return nil, nil
}

func hasTypeAssertion(exprs []ast.Expr) bool {
	for _, node := range exprs {
		_, ok := node.(*ast.TypeAssertExpr)
		if ok {
			return true
		}
	}
	return false
}
