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
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.AssignStmt:
			checkAssignStmt(pass, n)
		case *ast.ValueSpec:
			checkValueSpec(pass, n)
		}
	})

	return nil, nil
}

func checkAssignStmt(pass *analysis.Pass, n *ast.AssignStmt) {
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

func checkValueSpec(pass *analysis.Pass, n *ast.ValueSpec) {
	if !hasTypeAssertion(n.Values) {
		return
	}

	// if right hand has 2 or more values, assign statement can't assert boolean value which describes type assertion is succeeded
	if len(n.Values) > 1 {
		pass.Reportf(n.Pos(), "right hand must be only type assertion")
		return
	}

	if len(n.Names) == 2 {
		return
	}

	tae, ok := n.Values[0].(*ast.TypeAssertExpr)
	if !ok {
		pass.Reportf(n.Pos(), "right hand is not TypeAssertion")
		return
	}

	if tae.Type == nil {
		return
	}

	pass.Reportf(n.Pos(), "type assertion must be checked")

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
