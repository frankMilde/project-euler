//    Description:  A Pythagorean triplet is a set of three natural numbers,
//                  a < b < c, for which, a^2 + b^2 = c^2
//                  For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//                  There exists exactly one Pythagorean triplet for which
//                  a + b + c = 1000.
//
//       Question:  Find the product abc.
//
//       Compiler:  go
//
//          Usage:  go run pe-09_special-pythagorean-triplet.go
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2014, Frank Milde

package main

import "fmt"

func main() {
	var bound int = 1000
	for b := 2; b < bound; b++ {
		for a := 1; a < b; a++ {
			if condition(a, b, bound) == bound*bound/2 {
				fmt.Println("b: ", b, " a: ", a, "c: ", FindC(a, b, bound))
				fmt.Println("a,b and c are correct: ", CheckResult(a, b, FindC(a, b, bound), bound))
				fmt.Println("a*b*c: ", a*b*FindC(a, b, bound))
			}
		}
	}
}

func condition(x int, y int, bound int) int {
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
