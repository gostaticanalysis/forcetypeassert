package forcetypeassert

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "forcetypeassert",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "forcetypeassert is finds type assertions which did forcely such as below."

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
		(*ast.ValueSpec)(nil),
		(*ast.TypeAssertExpr)(nil),
	}

	inspect.Nodes(nodeFilter, func(n ast.Node, push bool) bool {
		if !push {
			return false
		}
		switch n := n.(type) {
		case *ast.AssignStmt:
			return checkAssignStmt(pass, n)
		case *ast.ValueSpec:
			return checkValueSpec(pass, n)
		case *ast.TypeAssertExpr:
			if n.Type != nil {
				pass.Reportf(n.Pos(), "type assertion must be checked")
			}
			return false
		}

		return true
	})

	return nil, nil
}

func checkAssignStmt(pass *analysis.Pass, n *ast.AssignStmt) bool {
	tae := findTypeAssertion(n.Rhs)
	if tae == nil {
		return true
	}

	switch {
	// if right hand has 2 or more values, assign statement can't assert boolean value which describes type assertion is succeeded
	case len(n.Rhs) > 1 :
		pass.Reportf(n.Pos(), "right hand must be only type assertion")
		return false
	case len(n.Lhs) != 2 && tae.Type != nil:
		pass.Reportf(n.Pos(), "type assertion must be checked")
		return false
	case len(n.Lhs) == 2:
		return false
	}

	return true
}

func checkValueSpec(pass *analysis.Pass, n *ast.ValueSpec) bool {
	tae := findTypeAssertion(n.Values)
	if tae == nil {
		return true
	}

	switch {
	// if right hand has 2 or more values, assign statement can't assert boolean value which describes type assertion is succeeded
	case len(n.Values) > 1 :
		pass.Reportf(n.Pos(), "right hand must be only type assertion")
		return false
	case len(n.Names) != 2 && tae.Type != nil:
		pass.Reportf(n.Pos(), "type assertion must be checked")
		return false
	case len(n.Names) == 2:
		return false
	}

	return true
}

func findTypeAssertion(exprs []ast.Expr) *ast.TypeAssertExpr {
	for _, expr := range exprs {
		var typeAssertExpr *ast.TypeAssertExpr
		ast.Inspect(expr, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.FuncLit:
				return false
			case *ast.TypeAssertExpr:
				typeAssertExpr = n
				return false
			}
			return true
		})
		if typeAssertExpr != nil {
			return typeAssertExpr
		}
	}
	return nil
}
