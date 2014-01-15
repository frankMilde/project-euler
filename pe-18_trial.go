package main

import (
	//	"bufio"
	//	"flag"
	"fmt"
	//	"io"
	//	"os"
	"sort"
	"strconv"
	"strings"
	//"time"
)

type Point struct {
	line, index int
}

type Connection struct {
	weight int
	start  Point
	end    Point
}

type Triangle [][]int
type Connections []*Connection

func (c Connections) Len() int      { return len(c) }
func (c Connections) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

type ByWeight struct{ Connections }

func (s ByWeight) Less(i, j int) bool {
	return s.Connections[i].weight <
		s.Connections[j].weight
}

func main() {

	t := Triangle{{3}, {7, 4}, {2, 4, 6}, {8, 5, 9, 3}}

	fmt.Println(t)
	var allConns []Connections

	for line := 0; line < len(t)-1; line++ {

		var consPerLine Connections

		for i := 0; i != len(t[line]); i++ {
			leftCon := new(Connection)
			leftCon.weight = t[line][i] + t[line+1][i]
			leftCon.start.line = line
			leftCon.start.index = i
			leftCon.end.line = line + 1
			leftCon.end.index = i

			rightCon := new(Connection)
			rightCon.weight = t[line][i] + t[line+1][i+1]
			rightCon.start.line = line
			rightCon.start.index = i
			rightCon.end.line = line + 1
			rightCon.end.index = i + 1

			consPerLine = append(consPerLine, leftCon)
			consPerLine = append(consPerLine, rightCon)
		}
		allConns = append(allConns, consPerLine)
	}

	for line := 0; line != len(t)-1; line++ {
		fmt.Println("Line", line, ":")
		nrOfConnectionsOfThisLine := 2 * len(t[line])
		sort.Sort(ByWeight{allConns[line]})
		for index := 0; index != nrOfConnectionsOfThisLine; index++ {
			fmt.Println(allConns[line][index])
		}
		fmt.Println()
	}
	fmt.Println(PrintConnections(allConns, t))
}

//func isAlreadyConnected(s []SortedConnectionsPerLine, c Connection) bool {
//for _,v:=range s{

func PointsAreEqual(p1, p2 Point) bool {
	var check bool

	if p1.line == p2.line && p1.index == p2.index {
		check = true
	} else {
		check = false
	}

	return check
}

// ===  IMPLEMENT STRING INTERFACE  ============================================
//         Name:  String
//  Description:  Defines how a Triangle type should be printed when using any of
//                the standard fmt methods. 
// =============================================================================
func PrintConnections(c []Connections, t Triangle) string {
	var s string
	//t:=c.triangle  

	for line := 0; line != len(t); line++ {
		s += strings.Repeat(" ", 2*(len(t)-1-line))
		for index := 0; index != len(t[line]); index++ {
			number := t[line][index]
			if number < 10 {
				s += " " + strconv.Itoa(number)
			} else {
				s += strconv.Itoa(number)
			}
			s += "  "
		}
		if line < len(t)-1 {
			s += "\n"
			s += strings.Repeat(" ", 2*(len(t)-1-line)-1)
			for index := 0; index != len(c[line]); index++ {
				currentPoint := Point{line, index}
				if PointsAreEqual(c[line][index].start, currentPoint) {
											fmt.Print(currentPoint)
					left := Point{line + 1, index}
					right := Point{line + 1, index + 1}
					if PointsAreEqual(c[line][index].end, left) {
						s += strconv.Itoa(c[line][index].weight) //"/ "
					}
					if PointsAreEqual(c[line][index].end, right) {
						s += "\\ "
					} else {
						s += "  "
					}
				}
			}
			s += "\n"
		}
	}
	return s
}

func PrintAllConnection(t Triangle) string {
	var s string
	//t:=c.triangle  

	for line := 0; line != len(t); line++ {
		s += strings.Repeat(" ", 2*(len(t)-1-line))
		for index := 0; index != len(t[line]); index++ {
			number := t[line][index]
			if number < 10 {
				s += " " + strconv.Itoa(number)
			} else {
				s += strconv.Itoa(number)
			}
			s += "  "

		}
		s += strings.Repeat(" ", 2*(len(t)-1-line))
		if line < len(t)-1 {
			s += "\n"
			s += strings.Repeat(" ", 2*(len(t)-1-line))
			s += strings.Repeat("/ \\ ", len(t[line]))
			s += "\n"
		}
	}

	/*
		for line := 0; line != len(t); line++ {
			offset := 2*(len(t)-1-line) + 2
			// top border
				s += strings.Repeat(" ", offset)
			for index := 0; index != len(t[line]); index++ {
				s += "+--+"
				s += strings.Repeat(" ", offset)
				s += "    "
			}
			s += strings.Repeat(" ", offset)
			s += "\n"

			// number
				s += strings.Repeat(" ", offset)
			for index := 0; index != len(t[line]); index++ {
				s += "|"
				number := t[line][index]
				if number < 10 {
					s += " " + strconv.Itoa(number)
				} else {
					s += strconv.Itoa(number)
				}
				s += "|"
				s += "  "
			}
			s += strings.Repeat(" ", offset)
			s += "\n"

			//bottom border and connections
			for index := 0; index != len(t[line]); index++ {
				s += strings.Repeat(" ", offset)
				s += "+--+"
				s += strings.Repeat(" ", offset)
				s += "    "
			}
			if line < len(t)-1 {
				s += "\n"
				s += strings.Repeat(" ", offset-1)
				s += strings.Repeat("/    \\ ", len(t[line]))
				s += "\n"
			}
		}
	*/
	return s
}
func (t Triangle) String() string {
	var s string

	for line := 0; line != len(t); line++ {
		s += strings.Repeat(" ", 2*(len(t)-1-line))
		nrOfConnectionsOfThisLine := 2 * len(t[line])
		for index := 0; index != len(t[line]); index++ {
			number := t[line][index]
			if number < 10 {
				s += " " + strconv.Itoa(number)
			} else {
				s += strconv.Itoa(number)
			}
			s += "  "

		}
		s += strings.Repeat(" ", 2*(len(t)-1-line))
		if line < len(t)-1 {
			s += strconv.Itoa(nrOfConnectionsOfThisLine)
		}
		s += "\n"
	}

	return s
}
