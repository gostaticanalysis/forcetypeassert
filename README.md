# fourcetypeassert

[![godoc.org][godoc-badge]][godoc]

`fourcetypeassert` finds type assertions which did forcely such as below.

```go
func f() {
	var a interface{}
	_ = a.(int) // must not do fource type assertion
}
```

<!-- links -->
[godoc]: https://godoc.org/github.com/gostaticanalysis/fourcetypeassert
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=flat-square&label=%20godoc.org

