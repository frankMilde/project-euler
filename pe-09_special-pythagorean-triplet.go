package main

import "fmt"

func condition(x int, y int,bound int) int {
	return bound*(x+y) - x*y
}
func FindC(a int, b int, bound int) int {
	return bound - a - b
}

func CheckResult(a int, b int, c int, bound int) bool {
	if a*a+b*b == c*c && a+b+c == bound {
		return true
	} else {
		return false
	}
	return false
}

func main() {
	var bound int = 1000
	for b := 2; b < bound; b++ {
		for a := 1; a < b; a++ {
			if condition(a, b, bound) == bound*bound/2 {
				fmt.Println("b: ", b, " a: ", a, "c: ", FindC(a, b, bound))
				fmt.Println("a,b and c are correct: ", CheckResult(a, b, FindC(a, b, bound), bound))
				fmt.Println("a*b*c: ", a*b*FindC(a, b,bound))
			}
		}
	}
}