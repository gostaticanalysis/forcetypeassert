package a

import "fmt"

var _, _ = ((interface{})(nil)).(string)       // OK
var _ = ((interface{})(nil)).(string)          // want `type assertion must be checked`
var _, _ = ((interface{})(nil)).(string), true // want `right hand must be only type assertion`

func f() {
	var i interface{} = "hello"

	_ = i.(string) // want `type assertion must be checked`

	_, _ = i.(string), "foo" // want `right hand must be only type assertion`

	s, ok := i.(string) // ok
	s, _ = i.(string)   // ok
	fmt.Println(s, ok)

	if s, ok := i.(string); ok { // ok
		println(s)
	}

	switch n := i.(type) { // ok
	case string:
		fmt.Println(n)
	}
	switch i.(type) { // ok
	case string:
	}

	var _ = i.(string) // want `type assertion must be checked`

	var _ = *i.(*string) // want `type assertion must be checked`

	println(i.(string))   // want `type assertion must be checked`
	println(*i.(*string)) // want `type assertion must be checked`

	_ = func() int {
		println(*i.(*string)) // want `type assertion must be checked`
		return 0
	}()

	func() {
		println(*i.(*string)) // want `type assertion must be checked`
	}()
}
