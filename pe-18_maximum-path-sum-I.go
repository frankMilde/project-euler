//
//    Description:  By starting at the top of the triangle below and moving to
//                  adjacent numbers on the row below, the maximum total
//                  from top to bottom is 23.
//
//                                    3
//                                   7 4
//                                  2 4 6
//                                 8 5 9 3
//
//                  That is, 3 + 7 + 4 + 9 = 23.
//
//                  Find the maximum total from top to bottom of the
//                  triangle below:
//
//                                              75
//                                            95  64
//                                          17  47  82
//                                        18  35  87  10
//                                      20  04  82  47  65
//                                    19  01  23  75  03  34
//                                  88  02  77  73  07  63  67
//                                99  65  04  28  06  16  70  92
//                              41  41  26  56  83  40  80  70  33
//                            41  48  72  33  47  32  37  16  94  29
//                          53  71  44  65  25  43  91  52  97  51  14
//                        70  11  33  28  77  73  17  78  39  68  17  57
//                      91  71  52  38  17  14  91  43  58  50  27  29  48
//                    63  66  04  68  89  53  67  30  73  16  69  87  40  31
//                  04  62  98  27  23  09  70  98  73  93  38  53  60  04  23
//
//                  NOTE: As there are only 16384 routes, it is possible to
//                  solve this problem by trying every route. However,
//                  Problem 67, is the same challenge with a triangle
//                  containing one-hundred rows; it cannot be solved by
//                  brute force, and requires a clever method! ;o)
//
//           TODO:  Could not solve it :(.
//
//       Compiler:  go
//
//        License:  GNU General Public License
//      Copyright:  Copyright (c) 2015, Frank Milde

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// types
type Triangle [][]int
type Reference [][]bool
type Ref [][]int

type Path struct {
	path   []int
	weigth int
}

//global variables
var fileName string

func main() {
	ClearTerminalScreen()
	start := time.Now()

	flag.Parse()

	t, err := ReadNumbersFromFile(fileName)
	if err != nil {
		panic(err)
	}

	//	t := Triangle{{75}, {95, 64}, {17, 47, 82}, {18, 35, 87, 10}, {20, 4, 82, 47, 65}, {75, 1, 23,
	//19, 3, 34}}
	fmt.Print(t, "\n\n")

	r := CreateReference2(t)
	r = FindNewMax2(t, r, 1)
	iterations := 2

	for HasAtLeastOneConnectionBetweenEachLine2(r) == false {
		r = FindNewMax2(t, r, iterations)
		iterations++
	}

	//		r = FindNewMax2(t, r,iterations)
	//			iterations++

	fmt.Println("Each line has at least one connections after", iterations,
		"iterations:\n\n", r)

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
		defaultFileName = "triangle-1.dat"
		usage           = "file name from which to read-in numbers"
	)
	flag.StringVar(&fileName, "file", defaultFileName, usage)
	flag.StringVar(&fileName, "f", defaultFileName, usage+" (shorthand)")
}

// ===  FUNCTION  ==============================================================
//         Name:  CreateReference
//  Description:  Creates reference triangle
// =============================================================================
func CreateReference(t Triangle) Reference {
	ref := make(Reference, len(t))

	for i, v := range t {
		ref[i] = make([]bool, len(v))
	}
	return ref
}

func CreateReference2(t Triangle) Ref {
	ref := make(Ref, len(t))

	for i, v := range t {
		ref[i] = make([]int, len(v))
	}
	return ref
}

// ===  FUNCTION  ==============================================================
//         Name:  HasAtLeastOneConnectionBetweenEachLine
//  Description:
// =============================================================================
func HasAtLeastOneConnectionBetweenEachLine(r Reference) bool {
	var canGetToNextLine bool
	//fmt.Println("Checking\n\n", r)

	for line := 0; line != len(r)-1; line++ {
		nextLine := line + 1
		for number := 0; number != len(r[line]); number++ {
			//find marked numbers in this line
			if r[line][number] == true {
				left := number
				right := number + 1

				//are there marked numbers in next line?
				if r[nextLine][left] == false && r[nextLine][right] == false {
					canGetToNextLine = false
					//				fmt.Println("Not getting From", line+1, number+1, "to", nextLine+1)
					//				fmt.Println(r[line])
					//				fmt.Println(r[nextLine])
				} else {
					canGetToNextLine = true
				}
			}
			if canGetToNextLine == true {
				break
			}
		}
		if canGetToNextLine == false {
			break
		}
	}
	return canGetToNextLine
}

func HasAtLeastOneConnectionBetweenEachLine2(r Ref) bool {
	var canGetToNextLine bool = false
	//fmt.Println("Checking\n\n", r)

	for line := 0; line != len(r)-1; line++ {
		nextLine := line + 1
		for number := 0; number != len(r[line]); number++ {
			//find marked numbers in this line
			if r[line][number] != 0 {
				left := number
				right := number + 1

				//are there marked numbers in next line?
				if r[nextLine][left] == 0 && r[nextLine][right] == 0 {
					canGetToNextLine = false
					//				fmt.Println("Not getting From", line+1, number+1, "to", nextLine+1)
					//				fmt.Println(r[line])
					//				fmt.Println(r[nextLine])
				} else {
					canGetToNextLine = true
				}
			}
			if canGetToNextLine == true {
				break
			}
		}
		if canGetToNextLine == false {
			break
		}
	}
	return canGetToNextLine
}

/*
func HasAPath (r Ref,it int) bool {
	var hasAPath bool=false

	for line := len(r)-1; line != it; line-- {
		prevLine := line - 1
		for number := 0; number != len(r[line]); number++ {
		}
	}
}
*/
// ===  FUNCTION  ==============================================================
//         Name:  FindNewMax
//  Description:
// =============================================================================
func FindNewMax(t Triangle, r Reference) Reference {

	for line := 0; line != len(t); line++ {
		var max int
		for number := 0; number != len(t[line]); number++ {
			if max < t[line][number] && r[line][number] == false {
				max = t[line][number]
			}
		}
		for number := 0; number != len(t[line]); number++ {
			if t[line][number] == max && r[line][number] == false {
				r[line][number] = true
			}
		}
	}
	return r
}

func FindNewMax2(t Triangle, r Ref, it int) Ref {

	for line := 0; line != len(t); line++ {
		var max int
		for number := 0; number != len(t[line]); number++ {
			if max < t[line][number] && r[line][number] == 0 {
				max = t[line][number]
			}
		}
		for number := 0; number != len(t[line]); number++ {
			if t[line][number] == max && r[line][number] == 0 {
				r[line][number] = it
			}
		}
	}
	return r
}
func FindNewGlobalMax(t Triangle, r Ref, it int) Ref {

	var max int
	for line := 0; line != len(t); line++ {
		for number := 0; number != len(t[line]); number++ {
			if max < t[line][number] && r[line][number] == 0 {
				max = t[line][number]
			}
		}
	}
	for line := 0; line != len(t); line++ {
		for number := 0; number != len(t[line]); number++ {
			if t[line][number] == max && r[line][number] == 0 {
				r[line][number] = it
			}
		}
	}

	return r
}

// ===  IMPLEMENT STRING INTERFACE  ============================================
//         Name:  String
//  Description:  Defines how a Triangle type should be printed when using any of
//                the standard fmt methods.
// =============================================================================
func (t Triangle) String() string {
	var s string

	for i := 0; i != len(t); i++ {
		s += strings.Repeat(" ", 2*(len(t)-1-i))
		var sumOfLineElements int
		for j := 0; j != len(t[i]); j++ {
			number := t[i][j]
			sumOfLineElements += number
			if number < 10 {
				s += " " + strconv.Itoa(number)
			} else {
				s += strconv.Itoa(number)
			}
			s += "  "

		}
		s += strings.Repeat(" ", 2*(len(t)-1-i)) +
			strconv.Itoa(sumOfLineElements/len(t[i]))
		s += "\n"
	}

	return s
}

func (r Reference) String() string {
	var s string

	for i := 0; i != len(r); i++ {
		s += strings.Repeat(" ", 2*(len(r)-1-i))
		for j := 0; j != i+1; j++ {
			check := r[i][j]
			if check == true {
				s += "xx"
			} else {
				s += "__"
			}
			s += "  "
		}
		s += "\n"
	}
	return s
}

func (r Ref) String() string {
	var s string

	for i := 0; i != len(r); i++ {
		s += strings.Repeat(" ", 2*(len(r)-1-i))
		for j := 0; j != i+1; j++ {
			number := r[i][j]
			switch {
			case number == 0:
				s += "__"
			case number < 10:
				s += " " + strconv.Itoa(number)
			case number >= 10:
				s += strconv.Itoa(number)
			}
			s += "  "
		}
		s += "\n"
	}
	return s
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
//         Name:  ReadNumbersFromFile
//  Description:  Reads the numbers from file into a string slice.
// =============================================================================
func ReadNumbersFromFile(filename string) (Triangle, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Triangle{}, err
	}
	// f.Close will run when we're finished.
	// checkin on error when closing:
	// http://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file#comment20239340_9739903
	defer func() {
		if f.Close() != nil {
			panic(err)
		}
	}()

	var (
		stringNumber []string
		line         string
		errString    error
		errNr        error
		number       int
		numberDelim  string = " "
		lineDelim    byte   = '\n'
		returnSlice  [][]int
	)

	r := bufio.NewReader(f)

	for errString != io.EOF {
		line, errString = r.ReadString(lineDelim)
		// trim the trailing EOL "\n" character
		line = strings.Trim(line, string(lineDelim))
		// check if last entry was an empty line
		if line == "" {
			break
		}
		stringNumber = strings.Split(line, numberDelim)
		var temp []int
		for _, v := range stringNumber {
			number, errNr = strconv.Atoi(v)
			if errNr != nil {
				return returnSlice, errNr
			}
			temp = append(temp, number)
		}
		returnSlice = append(returnSlice, temp)
	}

	return returnSlice, nil
}
