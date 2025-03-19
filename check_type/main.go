package main

import (
	"fmt"
	"reflect"
)

func main() {
	type myInt int
	type myInt2 int

	a := myInt2(5)

	fmt.Println(a)

	fmt.Println(reflect.TypeOf(a).Kind())

}
