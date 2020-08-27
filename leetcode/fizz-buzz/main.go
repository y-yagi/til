package main

import "fmt"

const fizz = "Fizz"
const buzz = "Buzz"

func fizzBuzz(n int) []string {
	a := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			a[i-1] = fizz + buzz
		} else if i%5 == 0 {
			a[i-1] = buzz
		} else if i%3 == 0 {
			a[i-1] = fizz
		} else {
			a[i-1] = fmt.Sprintf("%d", i)
		}
	}

	return a
}
