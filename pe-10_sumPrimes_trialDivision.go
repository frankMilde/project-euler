//
//    Description:  The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
//       Question:  Find the sum of all the primes below two million.
//
//       Compiler:  go
//
//          Usage:  go run pe-10_sumPrimes_trialDivision.go
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2014, Frank Milde

package main

import "fmt"

var primes = []uint64{2, 3}

func main() {
	limit := 10
	var sum uint64
	sum = 5

	// Simple trial and error method to find primes. Not very effective.
	// Better use the sieve versions.
	for i := 2; i < limit; i++ {
		primes = append(primes, FindNewPrime())
		sum += primes[len(primes)-1]
	}

	fmt.Println("sum: ", sum)

}

func CheckIfIsNewPrime(number uint64) bool {
	check_if_number_is_new_prime := false
	var remainder uint64 = 0

	for i := range primes {
		remainder = number % primes[i]
		if remainder == 0 {
			check_if_number_is_new_prime = false
			break
		} else {
			check_if_number_is_new_prime = true
		}
	}
	return check_if_number_is_new_prime
}

func FindNewPrime() uint64 {
	for counter := primes[len(primes)-1]; ; counter += 2 {
		if CheckIfIsNewPrime(counter) == true {
			return counter
		}
	}
	return 0
}
