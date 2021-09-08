package forcetypeassert_test

import (
	"testing"

	"github.com/gostaticanalysis/forcetypeassert"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, forcetypeassert.Analyzer, "a")
}

func TestResult(t *testing.T) {
	testdata := analysistest.TestData()
	a := &analysis.Analyzer{
		Name: "test",
		Doc:  "test",
		Requires: []*analysis.Analyzer{
			forcetypeassert.Analyzer,
		},
		Run: func(pass *analysis.Pass) (interface{}, error) {
			panicable, _ := pass.ResultOf[forcetypeassert.Analyzer].(*forcetypeassert.Panicable)
			for i := 0; i < panicable.Len(); i++ {
				n := panicable.At(i)
				pass.Reportf(n.Pos(), "panicable")
			}
			return nil, nil
		},
	}

	analysistest.Run(t, testdata, a, "b")
}
