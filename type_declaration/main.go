package main

import "fmt"

type celsius float64
type fahrenheit float64

func main() {
	var t1 celsius = -273.15
	var t2 fahrenheit = 0

	// t1 = t2 // cannot use t2 (variable of type fahrenheit) as celsius value in assignment
	// fmt.Println(t1 > t2) // invalid operation: t1 > t2 (mismatched types celsius and fahrenheit)
	fmt.Println(t2 * 10)
	fmt.Println(t1 + celsius(t2))
	fmt.Printf("%T\n", t1) // main.celsius

	type myScale celsius

	t3 := myScale(500)
	//  t3 = t1 // cannot use t1 (variable of type celsius) as myScale value in assignment
	t1 = celsius(t3)
	t2 = fahrenheit(t3)
	fmt.Println(t3 > myScale(t1))
}
