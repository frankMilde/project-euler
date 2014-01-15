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

func main() {
	ClearTerminalScreen()
	start := time.Now()
	flag.Parse()

	target := TestForDivisibility(limit)
	fmt.Println("Result: ", target)

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
		defaultLimit = 6
		usage        = "Find all primes below this limit"
	)
	flag.Uint64Var(&limit, "limit", defaultLimit, usage)
	flag.Uint64Var(&limit, "l", defaultLimit, usage+" (shorthand)")
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
//         Name:  TestForDivisibility
//  Description:  
// =============================================================================
func TestForDivisibility(limit uint64) uint64 {

	foundTarget := make(chan bool)
	triangularNumbers := make(chan uint64)
	go FindTriangularNumbers(triangularNumbers, foundTarget)
	foundTarget <- false

	var number uint64
	for number = range triangularNumbers {
		//fmt.Println("In Test:",number)
		var counter uint64 = 0
		var i uint64 = 1
		for ; i <= number; i++ {
			if number%i == 0 {
				counter++
			}
			if counter == limit {
				fmt.Println("In Test: ", number, "is div by", counter)
				fmt.Println("In Test: Send foundTarget")
				foundTarget <- true
				break
			} 
		}
				foundTarget <- false
	}
				 close(foundTarget)
	return number
}

// ===  FUNCTION  ==============================================================
//         Name:  FindTriangularNumbers
//  Description:  
// =============================================================================
func FindTriangularNumbers(out_ch chan<- uint64, quit <-chan bool) {

	var triangularNumber uint64 = 0 // next number
	var i uint64 = 1                // next number

	//action := "Finding triangular numbers... "

	//fmt.Println(action)
	var foundTarget bool
	foundTarget = <-quit

	for ; ; i++ {
		triangularNumber += i
		out_ch <- triangularNumber
		foundTarget = <-quit
		//fmt.Println("In Find received ", foundTarget)
		if foundTarget == true{
			break
		}
	}

		fmt.Println("In Find: Aborted")
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
