package main

import "fmt"

func main() {

	fmt.Println(fizzBuzz(3))

}

func fizzBuzz(n int) []string {
	res := make([]string, n, n)

	for i := 0; i < n; i++ {
		res[i] = "a"
	}

	return res
}
