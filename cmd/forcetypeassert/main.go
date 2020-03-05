package main

import (
	"github.com/gostaticanalysis/forcetypeassert"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(forcetypeassert.Analyzer) }
