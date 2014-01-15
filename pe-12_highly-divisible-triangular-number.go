package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var limit uint64
const ESC string = "\033["

type PrimeFactor struct {
	prime uint64
	exp   uint64
}

type Factorization []PrimeFactor

func main() {
	ClearTerminalScreen()
	start := time.Now()
	flag.Parse()

	var sum uint64 = 0
	var j uint64 = 1
	var dividends uint64 = 0

	// The loop creates triangle numbers until a number is found that has
	// more dividends than the given limit
	for ; dividends < limit; j++ {
		sum += j
		dividends = FindNumberOfDividends(Factorize(sum))
	}

	fmt.Println("The", j, "th triangular number", sum, "factors to\n")
	fmt.Print(Factorize(sum))
	fmt.Println("and has", dividends, "dividends.")

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
		defaultLimit = 500
		usage        = "Find all primes below this limit"
	)
	flag.Uint64Var(&limit, "limit", defaultLimit, usage)
	flag.Uint64Var(&limit, "l", defaultLimit, usage+" (shorthand)")
}

// ===  IMPLEMENT STRING INTERFACE  ============================================
//         Name:  String
//  Description:  Defines how a Primefactor should be printed when using any of
//                the standard fmt methods. prime^exp should be displayed as
//                     exp
//                prime
//                Therefore the cursor has to be moved down and back after
//                printing the exp. After printing the prime the cursur has to
//                move up again to print a next prime correctly. Hence, a whole
//                Factorization type will be displayed as
//                  e1   e2   e3
//                p1   p2   p3  ...
// =============================================================================
func (f PrimeFactor) String() string {
	var expTotal string
	var primeTotal string

	primeString := strconv.FormatUint(f.prime, 10)
	expString := strconv.FormatUint(f.exp, 10)
	numberOfPrimeDigits := len(primeString)
	numberOfExpDigits := len(expString)

	// construct strings
	expTotal += strings.Repeat(" ", numberOfPrimeDigits) + expString
	primeTotal += primeString + strings.Repeat(" ", numberOfExpDigits) + " "

	// cursor movment
	var cursorUp string = ESC + "1A"
	var cursorDown string = ESC + "1B"
	var cursorBack string = ESC +
		fmt.Sprintf("%dD", numberOfExpDigits+numberOfPrimeDigits)

	return expTotal + cursorDown + cursorBack + primeTotal + cursorUp
}

func (f Factorization) String() string {
	l := len(f)
	var total string
	for i := 0; i < l; i++ {
		total += fmt.Sprintf("%s", f[i])
	}
	return total + "\n\n\n"
}

// ===  FUNCTION  ==============================================================
//         Name:  FindNumberOfDividends
//  Description:  Finds the number of dividents for a given factorization of a 
//                number according to:
//                N=p1^e1 * p2^e2 * ... *  pn^en  
//                Number of Divisors is given by:  (e1+1)*(e2+1)*...*(en+1)
//        Input:  A factorization in form of a vector:
//                [p1 e1 p2 e2 ... pn en]
//       Output:  (e1+1)*(e2+1)*...*(en+1)
// =============================================================================
func FindNumberOfDividends(f Factorization) uint64 {
	l := len(f)
	var dividends uint64 = 1

	for i := 0; i < l; i++ {
		dividends *= (f[i].exp + 1)
	}
	return dividends
}

// ===  FUNCTION  ==============================================================
//         Name:  Factorize
//  Description:  Finds the factorization into exponents e of primes p for a 
//                given number according to:
//                N=p1^e1 * p2^e2 * ... *  pn^en  
//                To do so, the input number will be repeatedly divorced. 
//        Input:  A number
//       Output:  Factorization type (alias for []primeFactor) with the
//                primeFactor struct
//                [[p1 e1], [p2 e2], ... [pn en]]
// =============================================================================
func Factorize(number uint64) Factorization {
	var primeFactor = PrimeFactor{2, 0}
	var factorization Factorization

	// treat even numbers separatly
	for number%primeFactor.prime == 0 {
		primeFactor.exp++
		number = number / primeFactor.prime
	}
	if primeFactor.exp != 0 {
		factorization = append(factorization, primeFactor)
	}

	// now do the odd numbers
	for primeFactor.prime = 3; primeFactor.prime <= number+1; {
		primeFactor.exp = 0
		for number%primeFactor.prime == 0 {
			primeFactor.exp++
			number = number / primeFactor.prime
		}
		if primeFactor.exp != 0 {
			factorization = append(factorization, primeFactor)
		}
		// for odd numbers advance by 2 steps
		primeFactor.prime++
		primeFactor.prime++
	}
	return factorization
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
