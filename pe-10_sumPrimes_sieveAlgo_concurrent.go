//
//    Description:  The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
//       Question:  Find the sum of all the primes below two million.
//
//       Compiler:  go
//
//          Usage:  go run pe-10_sumPrimes_sieveAlgo_concurrent.go -l 2000000
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2014, Frank Milde

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const maxSliceElements uint64 = 1<<32 - 1

var limit uint64

func main() {
	ClearTerminalScreen()
	start := time.Now()
	flag.Parse()
	CheckLimitIsOkOrDie(limit)

	// using the
	// https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
	sum, last_prime, nr_of_primes := SumUpTo(limit)
	DisplayResults(sum, last_prime, limit, nr_of_primes)

	end := time.Now()

	fmt.Println("\nrun time:", end.Sub(start))
}

// ===  FUNCTION  ==============================================================
//         Name:  init
//  Description:  Needed by the flag package to define the variables that are
//                parsed from the command line
// =============================================================================
func init() {
	const (
		defaultLimit = 400000000
		usage        = "Find all primes below this limit"
	)
	flag.Uint64Var(&limit, "limit", defaultLimit, usage)
	flag.Uint64Var(&limit, "l", defaultLimit, usage+" (shorthand)")
}

// ===  FUNCTION  ==============================================================
//         Name:  CheckLimitIsOkOrDie
//  Description:  If user input of "limit" exceed the max number of slice
//                elements: Panic
// =============================================================================
func CheckLimitIsOkOrDie(limit uint64) {

	if limit > maxSliceElements {
		panicMsg := GeneratePanicMsg(limit)
		panic(panicMsg)
	}
}

// ===  FUNCTION  ==============================================================
//         Name:  GeneratePanicMsg
//  Description:
// =============================================================================
func GeneratePanicMsg(limit uint64) string {
	overMaxElements := limit - maxSliceElements
	msg :=
		"\nThe limit is too high:    " + strconv.FormatUint(limit, 10) +
			"\nThe limit cannot exceed   " + strconv.FormatUint(maxSliceElements, 10) +
			"\nYou are over the limit by " + strconv.FormatUint(overMaxElements, 10)
	return msg
}

// ===  FUNCTION  ==============================================================
//         Name:  ClearTerminalScreen
//  Description:  ClearTerminalScreen the terminal screen to have nice output
// =============================================================================
func ClearTerminalScreen() {
	os.Stdout.Write([]byte("\033[2J"))
	os.Stdout.Write([]byte("\033[H"))
	os.Stdout.Write([]byte("\n"))
}

// ===  FUNCTION  ==============================================================
//         Name:  DisplayResults
//  Description:  Pretty format of the results
// =============================================================================
func DisplayResults(sum, last_prime, limit uint64, nr_of_primes int) {
	long := len("sum of all primes below" + strconv.FormatUint(limit/1000000, 10) + "Mio:")
	short := len(strconv.Itoa(nr_of_primes) + "th prime:")
	remain := long - short
	fill := strings.Repeat(" ", remain)
	fmt.Println("\n\nResults:")
	fmt.Println(strings.Repeat("-", long+3+len(strconv.FormatUint(sum, 10))))
	fmt.Println(fill, nr_of_primes, "th prime:", last_prime)
	fmt.Println("sum of all primes below", limit/1000000, "Mio:", sum)
	fmt.Println(strings.Repeat("-", long+3+len(strconv.FormatUint(sum, 10))))
}

// ===  FUNCTION  ==============================================================
//         Name:  SumUpTo
//  Description:  Calculates the sum of all primes below a given limit.
//                Returns
//                (1) sum
//                (2) the highest prime below the limit
//                (3) total number of primes below the limit
// =============================================================================
func SumUpTo(limit uint64) (uint64, uint64, int) {
	filename := "primes.dat"

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if f.Close() != nil {
			panic(err)
		}
	}()

	primesCh := make(chan uint64)
	go FindPrimes(limit, primesCh)

	var sum uint64 = 0
	var prime uint64 = 0
	var nr_of_primes int = 0
	for prime = range primesCh {
		sum += prime
		n, err := io.WriteString(f, strconv.FormatUint(prime, 10)+"\n")
		if err != nil {
			fmt.Println(n)
			panic(err)
		}
		nr_of_primes++
	}
	fmt.Println("\n\nWriting data to " + filename)
	return sum, prime, nr_of_primes
}

// ===  FUNCTION  ==============================================================
//         Name:  FindPrimes
//  Description:  Calculates all primes below a given limit and sends them along
//                a channel.
//                Uses the "Sieve of Eratosthenes"-algo:
//                http://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
// =============================================================================
func FindPrimes(limit uint64, out_ch chan<- uint64) {
	//first prime
	out_ch <- 2

	// count only odd numbers
	var sieve_bound uint64 = (limit - 1) / 2

	sieved_numbers := make([]bool, sieve_bound)
	var i uint64 = 1 // next prime
	var j uint64 = 0 // multiples of next prime

	action := "Finding all primes below " + strconv.FormatUint(limit, 10) + "... this might take a while."

	fmt.Println(action)
	// no need to go till sieve_bound:
	// when we have found the prime 3, the highest prime we can find to sieve is bound/3
	for ; i < sieve_bound/(2*i+1); i++ {
		if !sieved_numbers[i] {
			for j = (2 * i) * (i + 1); j < sieve_bound; j += 2*i + 1 {
				sieved_numbers[j] = true
			}
		}
		DisplayProgressBar(i, sieve_bound/(2*i+1), "")
	}

	for i = 1; i < sieve_bound; i++ {
		if !sieved_numbers[i] {
			out_ch <- (2*i + 1)
		}
	}
	close(out_ch)
}

// ===  FUNCTION  ==============================================================
//         Name:  DisplayProgressBar
//  Description:
// =============================================================================
func DisplayProgressBar(current, total uint64, action string) {
	percent := current * 100 / total
	prefix := strconv.FormatUint(percent, 10) + "%"
	bar_start := " ["
	bar_end := "] "
	cols := 50

	bar_size := cols - len(prefix+bar_start+bar_end)
	amount := int(float32(current) / (float32(total) / float32(bar_size)))
	remain := bar_size - amount
	var bar string

	if current != total {
		bar = strings.Repeat("=", amount) + ">" + strings.Repeat(" ", remain)
	} else {
		bar = strings.Repeat("=", amount) + strings.Repeat(" ", remain)
	}

	os.Stdout.Write([]byte(prefix + bar_start + bar + bar_end + "\r" + action))
}
