package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var ascii_digit_offset uint8 = 48
var limit uint64

func main() {
	ClearTerminalScreen()
	start := time.Now()

	flag.Parse()
	var maxLength uint64 = 0
	var maxSeq uint64 = 0
	var length uint64 = 0
	var i uint64
	for i = 1; i < limit; i++ {
		length = ConstructSequence(i)
		if length > maxLength {
			maxLength = length
			maxSeq = i
		}
	}

	//DisplaySequence(limit)
	DisplayResults(maxSeq, maxLength)

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
		defaultLimit = 1000000
		usage        = "maximum starting number"
	)
	flag.Uint64Var(&limit, "limit", defaultLimit, usage)
	flag.Uint64Var(&limit, "l", defaultLimit, usage+" (shorthand)")
}

// ===  FUNCTION  ==============================================================
//         Name:  Even and Odd
//  Description:  return the corresponding Collatz condition
// =============================================================================
func Even(n uint64) uint64 { return n / 2 }
func Odd(n uint64) uint64  { return 3*n + 1 }

// ===  FUNCTION  ==============================================================
//         Name:  ConstructSequence
//  Description:  Constructs the Collatz sequence for a given starting number
//                return the length of the sequence
// =============================================================================
func ConstructSequence(number uint64) uint64 {
	var counter uint64 = 1
	for number > 1 {
		if number%2 == 0 {
			number = Even(number)
		} else {
			number = Odd(number)
		}
		counter++
	}
	return counter
}

// ===  FUNCTION  ==============================================================
//         Name:  DisplaySequence
//  Description:  
// =============================================================================
func DisplaySequence(number uint64) {
	if number < 10 {
		fmt.Print("0", number)
	} else {
		fmt.Print(number)
	}
	var counter uint64 = 1
	for number > 1 {
		if number%2 == 0 {
			number = Even(number)
		} else {
			number = Odd(number)
		}
		counter++
		if number < 10 {
			fmt.Print("-->0", number)
		} else {
			fmt.Print("-->", number)
		}
	}
	fmt.Print("\n")
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
	fmt.Println("Longest Collatz sequence under", limit, "starts at")
	fmt.Println(number)
	fmt.Println("and has a length of")
	fmt.Println(length)
}
