package forcetypeassert_test

import (
	"testing"

	"github.com/gostaticanalysis/forcetypeassert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, forcetypeassert.Analyzer, "a")
}
