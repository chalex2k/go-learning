package main

import "fmt"

/*
Функция new создаёт и возвращает указатель на переменную инициализированную нулевым значением.
*/

func main() {
	i := new(int)
	fmt.Printf("%T, %v\n", i, *i) //*int, 0

	type person struct {
		name string
	}
	p := new(person)
	p2 := &person{}
	fmt.Println(*p == *p2) // true

	j := 0
	pi1 := &j
	pi2 := new(int)
	fmt.Println(*pi1 == *pi2) // true

	k := new(*int)
	fmt.Printf("%T, %v\n", k, *k) // **int, <nil>

}
