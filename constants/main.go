package main

import "fmt"

const b = 1

func main() {
	a := 2
	c := a / b

	fmt.Println(c)

	type person struct {
		name string
	}
	//const Charles = person{name: "Charles"} // person{…} (value of type person) is not constant

	const r = 5 + b
	fmt.Println(r)

	//s := "lalal"
	//const name = "x" + s // "x" + s (value of type string) is not constant

	type Kg int
	const weight Kg = 83
	fmt.Printf("%T\n", weight) // main.Kg

	const (
		x = 1
		y
		z = 3
	)
	fmt.Println(y) // 1

	const (
		w Kg = 10
		v
	)
	fmt.Printf("%T\n", v) // main.Kg

	iotaExample()
}

/* глава 3.6.1 */
func iotaExample() {
	const base = 10

	const (
		_   = 1 << (base * iota)
		kib // 1024
		mib // 1048576
	)

	const (
		a = " asdf asd"
		b = 2
		c = iota
		d
		f = 10
		g = iota
	)
	const (
		x = string(rune(iota)) + " kg"
		y
		z
	)
	fmt.Println(c, d, g)
	fmt.Printf("%T %T %T\n", x, y, z)
	fmt.Println(x, y, z)

	const (
		t = iota + iota*iota // можно несколько раз но каждая  = 0
		r
		k
	)
	fmt.Println(t, r, k) // 0 2 6

}
