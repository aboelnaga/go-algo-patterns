package main

import (
	"fmt"
)

var a int = 10
var x int = 20

const (
	North = iota
	South
	East
	West
)

func main() {
	// fundamentals()
	// fundamentals2()
	// fundamentals3()
	// fundamentals4()
	fundamentals5()
}

func other() {
	fmt.Println(a) // [2] ?
}

func fundamentals() {
	x = 30
	a := 40
	var z byte
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(North)
	fmt.Println(West)
	fmt.Println(South)
	fmt.Println(East)
	other()
}

func fundamentals2() {
	x := 10
	y := int64(20)
	a := float64(30.765476547654)
	b := float32(30.765476547654)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x + int(y))
	fmt.Println(y + int64(x))
	fmt.Println(a)
	fmt.Println(b)
}

func fundamentals3() {
	english := "Hello"
	arabic := "مرحبا"
	fmt.Println(len(english))
	fmt.Println(len(arabic))

	for i, v := range english {
		fmt.Println(english[i])
		fmt.Printf("en. %d %c ", i, v)
	}

	for i, v := range arabic {
		fmt.Println(arabic[i])
		fmt.Printf("ar. %d %c ", i, v)
	}
}

func fundamentals4() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func fundamentals5() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------------")

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	fmt.Println("------------------")

	y := 0
	for {
		if y == 5 {
			break
		}
		fmt.Println(y)
		y++
	}

	slices := []string{"Go", "is", "fun"}
	for _, v := range slices {
		fmt.Printf("%v ", v)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
