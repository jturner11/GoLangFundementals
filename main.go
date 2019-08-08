package main

import (
	// packages do not need comma seperation
	"errors"
	"fmt"
	"math"
)

func main() {
	// Line 9 & 10 are the same - 9 is a short hand - declare varibles
	x := 23
	var y int = 12
	var sum int = x + y

	fmt.Println(sum)

	if x > y {
		fmt.Println(`x is greater than y`)
	}

	//These two are the same however top one is shorthand - set index in array
	b := [5]int{0, 0, 5, 0, 0}

	var a [5]int
	a[2] = 7
	fmt.Println(a, b)

	// adding to an array
	c := []int{1, 2, 3, 4, 5}
	c = append(c, 7)
	fmt.Println(c)

	// mapping over key values
	hola := make(map[string]int)
	hola["hello"] = 1
	hola["world"] = 2
	hola["worldyyy"] = 3

	// delete(hola, "world")   <== this deletes the key value

	fmt.Println(hola)
	fmt.Println(hola["world"])

	// loops - go only has for loop
	for d := 0; d < 5; d++ {
		fmt.Println(d)
	}

	// Iterate over elements in array using range
	ranger := []string{"a", "bn", "c"}
	for index, value := range ranger {
		fmt.Println("index", index, "value", value)
	}

	// calling the function sum which is outside of main
	sumResult := suma(2, 5)
	fmt.Println(sumResult)

	// Sqrt function
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Calling the struct defined below
	p := person{name: "your name", age: 23}
	fmt.Println(p)
	fmt.Println(p.name)

	// memoery pointeer
	i := 7
	inc(&i)
	fmt.Println(i)

}

// you can also make functions outside of main
func suma(e int, f int) int {
	return e + f
}

// sqrt function
func sqrt(g float64) (float64, error) {
	if g < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return math.Sqrt(g), nil
}

// Create a Struct type
type person struct {
	name string
	age  int
}

// Pointer
func inc(x *int) {
	*x++
}
