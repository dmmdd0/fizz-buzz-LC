package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(fizzBuzz(15))

}

func fizzBuzz(n int) []string {
	var ans string
	const (
		f = "Fizz"
		b = "Bizz"
	)
	res := make([]string, n, n)

	for i := 1; i <= n; i++ {

		if i%3 == 0 {
			ans = f
		}

		if i%5 == 0 {
			ans += b
		}
		if ans == "" {
			res[i-1] = strconv.Itoa(i)
		} else {
			res[i-1] = ans
		}

		fmt.Println(i, ans)
		ans = ""

	}

	return res
}
