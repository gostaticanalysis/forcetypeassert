package b

import "fmt"

var _, _ = ((interface{})(nil)).(string)       // OK
var _ = ((interface{})(nil)).(string)          // want `panicable`
var _, _ = ((interface{})(nil)).(string), true // OK

func f() {
	var i interface{} = "hello"

	_ = i.(string) // want `panicable`

	_, _ = i.(string), "foo" // OK

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

	var _ = i.(string) // want `panicable`

	var _ = *i.(*string) // want `panicable`

	println(i.(string))   // want `panicable`
	println(*i.(*string)) // want `panicable`

	_ = func() int {
		println(*i.(*string)) // want `panicable`
		return 0
	}()

	func() {
		println(*i.(*string)) // want `panicable`
	}()
}
