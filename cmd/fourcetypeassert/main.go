package main

import (
	"github.com/gostaticanalysis/fourcetypeassert"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(fourcetypeassert.Analyzer) }
