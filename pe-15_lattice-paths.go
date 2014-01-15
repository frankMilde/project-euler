package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var ascii_digit_offset uint8 = 48
var gridSize uint64

const maxUint64 uint64 = 1<<32 - 1

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

	numerator := FactorizedFactorial(gridSize * 2)
	denominator := ProdFactor(FactorizedFactorial(gridSize), FactorizedFactorial(gridSize))
	div := DivFactor(numerator, denominator)
	nrOfPaths := ReverseFactorization(div)

	fmt.Print("Number of paths of a ", gridSize, "x", gridSize,
		" Grid:\n", gridSize*2, "!/(", gridSize, "!*", gridSize, "!)= \n")

	fmt.Print(nrOfPaths)

	end := time.Now()
	fmt.Println("\n\nrun time:", end.Sub(start))
}

// ===  FUNCTION  ==============================================================
//         Name:  init
//  Description:  Needed by the flag package to define the variables that are
//                parsed from the command line
// =============================================================================
func init() {
	const (
		defaultGridSize = 2
		usage           = "Grid Size of the lattice"
	)
	flag.Uint64Var(&gridSize, "gridSize", defaultGridSize, usage)
	flag.Uint64Var(&gridSize, "g", defaultGridSize, usage+" (shorthand)")
}

// ===  FUNCTION  ==============================================================
//         Name:  ProdFactor
//  Description:  Calculates the product of two factorizations
// =============================================================================
func ProdFactor(a, b Factorization) Factorization {
	var f Factorization
	var p PrimeFactor
	for _, av := range a {
		p.exp = av.exp
		p.prime = av.prime
		for _, bv := range b {
			if av.prime == bv.prime {
				p.exp = av.exp + bv.exp
			}
		}
		f = append(f, p)
	}

	return f
}

// ===  FUNCTION  ==============================================================
//         Name:  DivFactor
//  Description:  Calculates the division of two factorizations
// =============================================================================
func DivFactor(numerator, denominator Factorization) Factorization {
	var f Factorization
	var p PrimeFactor
	for _, nv := range numerator {
		p.exp = nv.exp
		p.prime = nv.prime
		for _, dv := range denominator {
			if nv.prime == dv.prime {
				if nv.exp < dv.exp {
					fmt.Println("ERROR")
					os.Exit(1)
				}
				p.exp = nv.exp - dv.exp
			}
		}
		if p.exp != 0 {
			f = append(f, p)
		}
	}

	return f
}

// ===  FUNCTION  ==============================================================
//         Name:  SimplifyFactorization
//  Description:  Sorts and Simplifies a factorization:
//                5^1 * 2^2 * 5^3 * 2^1 = 2^3 * 5^4
//        Input:  A Factorization
//       Output:  A sorted Factorization 
// =============================================================================
func SimplifyFactorization(f Factorization) Factorization {
	occuringPrimes := FindOccuringPrimes(f)
	sortedPrimes := Sort(occuringPrimes)
	simplifiedF := CompactFactorization(sortedPrimes, f)

	return simplifiedF
}

// ===  FUNCTION  ==============================================================
//         Name:  CompactFactorization
//  Description:  Compacts  a factorization by starting with a prime from the
//                sorted prime array and for each prime iteraters over the
//                factorization and summes all occuring exponents of the given
//                prime.
//        Input:  (1) An arry with the sorted primes of the factorization 
//                   [2 5 7]
//                (2) the factorization itsself
//                
//       Output:  A compacted factorization 
//                5^1 * 2^2 * 7^1 * 5^3 * 2^1 --> 2^3 * 5^4 *7^1
// =============================================================================
func CompactFactorization(sortedPrimes []uint64, factorization Factorization) Factorization {
	var compactedF Factorization

	// iterate over all occuring primes (in sored order)
	for _, sortedPrime := range sortedPrimes {
		var summedExponents uint64
		// for each prime iterate over factorization and sum all exponents
		for i := range factorization {
			if factorization[i].prime == sortedPrime {
				summedExponents += factorization[i].exp
			}
		}
		compactedF = append(compactedF, PrimeFactor{sortedPrime, summedExponents})
	}

	return compactedF
}

// ===  FUNCTION  ==============================================================
//         Name:  Sort
//  Description:  Uses the Insertion-Sort-Algorithm to sort an array
//                [5 2 7] --> [2 5 7]
// =============================================================================
func Sort(a []uint64) []uint64 {

	for i := 1; i != len(a); i++ {
		j := i
		for (j > 0) && (a[j] < a[j-1]) {
			a[j], a[j-1] = a[j-1], a[j]
			j--
		}
	}
	return a
}

// ===  FUNCTION  ==============================================================
//         Name:  FindOccuringPrimes
//  Description:  Finds all primes of a factorization
//                5^1 * 2^2 * 5^3 * 2^1 ->  [5 2 ]
//        Input:  A Factorization
//       Output:  A sorted slice of primes 
// =============================================================================
func FindOccuringPrimes(f Factorization) []uint64 {
	l := len(f)
	var occuringPrimes []uint64
	var isAlreadyRegistered = false

	occuringPrimes = append(occuringPrimes, f[0].prime)
	for i := 1; i < l; i++ {
		isAlreadyRegistered = false
		for j := 0; j < len(occuringPrimes); j++ {
			if f[i].prime == occuringPrimes[j] {
				isAlreadyRegistered = true
			}
		}
		if isAlreadyRegistered != true {
			occuringPrimes = append(occuringPrimes, f[i].prime)
		}
	}
	return occuringPrimes
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
//         Name:  ReverseFactorization
//  Description:  Calculates the product of a factorization
//                 2  2
//                2  3  = 36
// =============================================================================
func ReverseFactorization(f Factorization) string {
	var result []uint8 = []uint8{1}

fmt.Print("\n\n")
	for i, v := range f {
		prime := Vectorize(strconv.FormatUint(v.prime, 10))
		result = Product(result, Pow(prime, v.exp))
		DisplayProgressBar(i, len(f)-1, "Calculating... ")
	}

fmt.Print("\n\n")
	return Stringify(result)
}

// ===  FUNCTION  ==============================================================
//         Name:  Pow
//  Description:  Calculates base^exp as a series of products
//                b^e=b*b*...*b*b
//                       e-times
// =============================================================================
func Pow(base []uint8, exp uint64) []uint8 {
	var result []uint8 = base

	for i := exp - 1; i != 0; i-- {
		result = Product(result, base)
	}
	return result
}

// ===  FUNCTION  ==============================================================
//         Name:  Product
//  Description:  Calculates the product of two arbitrary long numbers as a
//                number of sum: a*b = a+a+...+a+a
//                                        b-times
//                The numbers are represented as arrays of uint8s in reverse
//                order.
// =============================================================================
func Product(a, b []uint8) []uint8 {
	var result []uint8 = a
	stringNumberB := Stringify(b)
	uintNumberB, err := strconv.ParseUint(stringNumberB, 10, 64)
	if err != nil {
		panic(err)
	}

	var i uint64
	for i = uintNumberB - 1; i != 0; i-- {
		result = Sum(result, a)
	}
	return result
}

// ===  FUNCTION  ==============================================================
//         Name:  Sum
//  Description:  Sums two integers by:
//								(1) Checking which number (represented by an array of single
//								    digits) has more digits 
//  							(2) The number with fewer digits is filled with zeros until
//  							    both arrays have an equal length
//								(3) Performs the addition of the two numbers
// =============================================================================
func Sum(a, b []uint8) []uint8 {
	switch {
	case len(a) < len(b):
		a, b = FillSmallerWithZeros(a, b)
	case len(b) < len(a):
		b, a = FillSmallerWithZeros(b, a)
	}
	return Addition(a, b)
}

// ===  FUNCTION  ==============================================================
//         Name:  Addition
//  Description:  Adds two numbers (of equal length) by a pen-and-paper addition
//  							algorithm. Numbers are given in a form where the first element
//  							of the vector contains the unit position, second element
//  							contains the decade, ect.
//                See also function Vectorize()
//  =============================================================================
func Addition(a, b []uint8) []uint8 {
	var sum uint8 = 0
	var carry uint8 = 0
	var result []uint8

	if len(a) != len(b) {
		fmt.Println("ERROR: Addants have no equal length!")
		os.Exit(1)
	}

	for i := 0; i != len(a); i++ {
		sum = a[i] + b[i] + carry
		if sum >= 10 {
			carry = 1
			result = append(result, sum%10)
		} else {
			carry = 0
			result = append(result, sum)
		}
		//	DisplayProgressBar(i,len(a),"Adding numbers... ")
		//fmt.Println(sum,carry,result)
	}
	if carry > 0 {
		result = append(result, carry)
	}
	return result
}

// ===  FUNCTION  ==============================================================
//         Name:  FillSmallerWithZeros
//  Description:  Fill the smaller integer slice with zeros until the length of
//  							both slices is equal
// =============================================================================
func FillSmallerWithZeros(smaller, bigger []uint8) ([]uint8, []uint8) {
	diff := len(bigger) - len(smaller)

	for i := 0; i != diff; i++ {
		smaller = append(smaller, 0)
	}
	return smaller, bigger
}

// ===  FUNCTION  ==============================================================
//         Name:  FactorizedFactorial
//  Description:  returns the Factorial in prime factorized form:
//                10! = 
// =============================================================================
func FactorizedFactorial(n uint64) Factorization {
	var f Factorization
	var i uint64

	for i = 2; i <= n; i++ {
		for j := 0; j != len(Factorize(i)); j++ {
			f = append(f, Factorize(i)[j])
		}
	}
	return SimplifyFactorization(f)
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
//         Name:  Stringify
//  Description:  Transforms a vector of digits into a string, like
//  							[1 2 3] -> "321"
//  =============================================================================
func Stringify(number []uint8) string {
	length := len(number)
	var string_number []byte
	var buffer bytes.Buffer

	for i := range number {
		string_number = append(string_number,
			byte(number[length-1-i]+ascii_digit_offset))
	}
	buffer.Write(string_number)

	return buffer.String()
}

// ===  FUNCTION  ==============================================================
//         Name:  Vectorize
//  Description:  Transforms a string of digits into an vector of digits, and
//                reversing the order, so that the first element of the vector
//                is the unit position, the second the decade position, etc. like
//  							"321"  -> [1 2 3]
//  =============================================================================
func Vectorize(number string) []uint8 {
	length := len(number)

	var vector_number []uint8

	for i := range number {
		vector_number = append(vector_number,
			uint8(number[length-1-i])-ascii_digit_offset)
	}
	return vector_number
}
func VectorizeAll(stringNumbers []string) [][]uint8 {
	var vectorNumbers [][]uint8

	for i := range stringNumbers {
		vectorNumbers = append(vectorNumbers, Vectorize(stringNumbers[i]))
	}
	return vectorNumbers
}

// ===  FUNCTION  ==============================================================
//         Name:  GeneratePanicMsg
//  Description:  
// =============================================================================
func GeneratePanicMsg(limit uint64) string {
	overMaxElements := limit - maxUint64
	msg :=
		"\nThe numerator is too high:    " + strconv.FormatUint(limit, 10) +
			"\nThe numerator cannot exceed   " + strconv.FormatUint(maxUint64, 10) +
			"\nYou are over the limit by " + strconv.FormatUint(overMaxElements, 10)
	return msg
}

// ===  FUNCTION  ==============================================================
//         Name:  ClearTerminalScreen
//  Description:  Clears the terminal screen to have nice output
// =============================================================================
func ClearTerminalScreen() {
	os.Stdout.Write([]byte("\033[2J"))
	os.Stdout.Write([]byte("\033[H"))
	os.Stdout.Write([]byte("\n"))
}

// ===  FUNCTION  ==============================================================
//         Name:  DisplayProgressBar
//  Description:  
// =============================================================================
func DisplayProgressBar(current, total int, action string) {
	percent := current * 100 / total
	//prefix := strconv.FormatUint(percent, 10) + "%"
	prefix := strconv.Itoa(percent) + "%"
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

// ===  FUNCTION  ==============================================================
//         Name:  DisplayResults
//  Description:  Pretty format of the results
// =============================================================================
func DisplayResults(number uint64, length uint64) {
	fmt.Println("Longest Collatz sequence under", gridSize, "starts at")
	fmt.Println(number)
	fmt.Println("and has a length of")
	fmt.Println(length)
}

// ===  FUNCTION  ==============================================================
//         Name:  CheckLimitIsOkOrDie
//  Description:  If user input of "limit" exceed the max number of slice 
//                elements: Panic
// =============================================================================
func CheckNumberIsOkOrDie(limit uint64) {

	if limit > maxUint64 {
		panicMsg := GeneratePanicMsg(limit)
		panic(panicMsg)
	}
}
