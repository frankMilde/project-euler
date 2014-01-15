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
var initBase string
var initExp uint64

func main() {
	ClearTerminalScreen()
	start := time.Now()

	flag.Parse()

	//base := Vectorize(initBase)
	//result2 := Pow(base, initExp)

	var f uint64 = 100

	result := []uint8{1}
	var i uint64

	for i=1;i!=f;i++{
		result= Product(result,VectorizeUint64(i))
	}


	DisplayResults(result)

	end := time.Now()
	fmt.Println("\n\nrun time:", end.Sub(start))
}

func VectorizeUint64(number uint64) []uint8 {
stringNumber := strconv.FormatUint(number,10)
	length := len(stringNumber)

	var vectorNumber []uint8

	for i := range stringNumber {
		vectorNumber = append(vectorNumber,
			uint8(stringNumber[length-1-i])-ascii_digit_offset)
	}
	return vectorNumber
}
// ===  FUNCTION  ==============================================================
//         Name:  init
//  Description:  Needed by the flag package to define the variables that are
//                parsed from the command line
// =============================================================================
func init() {
	const (
		defaultBase = "2"
		defaultExp  = 2
		usage       = "base and exponent of the number base^exponent"
	)
	flag.StringVar(&initBase, "base", defaultBase, usage)
	flag.StringVar(&initBase, "b", defaultBase, usage+" (shorthand)")
	flag.Uint64Var(&initExp, "exp", defaultExp, usage)
	flag.Uint64Var(&initExp, "e", defaultExp, usage+" (shorthand)")
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
		fmt.Println(Stringify(a),stringNumberB)
		panic(err)
	}

	var i uint64
	for i = uintNumberB - 1; i != 0; i-- {
		result = Sum(result, a)
	}
	return result
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
		DisplayProgressBar(exp-1-i, exp-1, "Calculating... ")
	}

	fmt.Print("\n\n")

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
func DisplayProgressBar(current, total uint64, action string) {
	if total > 2 {
		total--
	}
	percent := current * 100 / total
	prefix := strconv.FormatUint(percent, 10) + "%"
	//prefix := strconv.Itoa(percent) + "%"
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
func DisplayResults(summedNumbers []uint8) {
	lB := len(initBase)
	lE := len(strconv.FormatUint(initExp, 10))
	var sumOfDigits uint64
	for _, v := range summedNumbers {
		sumOfDigits += uint64(v)
	}
	fmt.Print(strings.Repeat(" ", lB), initExp)
	fmt.Print("\n", initBase, strings.Repeat(" ", lE+1), "= ", Stringify(summedNumbers))
	fmt.Print("\n\nSum of its digits is ", sumOfDigits, ".")

}
