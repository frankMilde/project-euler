package main

import "fmt"

func main() {
	var limit uint64 = 20000000

	primes := FindPrimes(limit)
	   sum := Sum(primes)

	fmt.Println(len(primes), "th prime: ", primes[len(primes)-1])
	fmt.Println("sum: ", sum)
	//fmt.Println("   :  142913828922")
}


func Sum(vector []uint64) uint64 {
	var sum uint64 = 0
	for i := range vector {
		sum += vector[i]
	}
	return sum
}

func FindPrimes(limit uint64) []uint64 {
	var sieve_bound uint64 = (limit - 1) / 2

	sieved_numbers := make([]bool, sieve_bound)
	var i uint64 = 1 // next prime
	var j uint64 = 0 // multiples of next prime

	for ; i < sieve_bound/(2*i+1); i++ {
		if !sieved_numbers[i] {
			for j = (2 * i) * (i + 1); j < sieve_bound; j += 2*i + 1 {
				sieved_numbers[j] = true
			}
		}
	}

	primes := []uint64{}
	primes = append(primes, 2)

	for i = 1; i < sieve_bound; i++ {
		if !sieved_numbers[i] {
			primes = append(primes, (2*i + 1))
		}
	}

	return primes
}
