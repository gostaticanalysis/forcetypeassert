package a

import "fmt"

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
}
