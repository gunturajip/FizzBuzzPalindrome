package main

import "fmt"

func fizzBuzz(i int) string {
	out := ""
	if i%3 == 0 {
		out += "Fizz"
	}
	if i%5 == 0 {
		out += "Buzz"
	}
	if out == "" {
		out = fmt.Sprint(i)
	}
	return out
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
