package main

import "fmt"

var primes = []uint64{2, 3}

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

func main() {
	limit := 10
	var sum uint64
	sum = 5
	//var sum2 uint64
	
	for i := 2; i < limit; i++ {
		primes = append(primes, FindNewPrime())
		sum += primes[len(primes)-1]
	}
	//for i := range primes {
	//	sum2 += primes[i]
	//}
	
	fmt.Println("sum: ", sum)

}
