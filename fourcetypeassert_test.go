package fourcetypeassert_test

import (
	"testing"

	"github.com/gostaticanalysis/fourcetypeassert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, fourcetypeassert.Analyzer, "a")
}
