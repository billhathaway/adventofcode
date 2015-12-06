package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type xy struct {
	x int
	y int
}

var houses = make(map[xy]int)
var presents int

func deliver(loc xy) {
	fmt.Printf("Delivering to %v\n", loc)
	houses[loc]++
	presents++
}
func parse(santas int, r io.Reader) (int, error) {
	santaXY := make([]xy, santas)
	var santaIdx int
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanBytes)
	houses[santaXY[0]] = 2
	presents += 2

	for s.Scan() {
		val := s.Text()
		switch val {
		case "^":
			santaXY[santaIdx].y++
			deliver(santaXY[santaIdx])
		case "v":
			santaXY[santaIdx].y--
			deliver(santaXY[santaIdx])
		case "<":
			santaXY[santaIdx].x--
			deliver(santaXY[santaIdx])
		case ">":
			santaXY[santaIdx].x++
			deliver(santaXY[santaIdx])
		case "\n":
			fmt.Printf("presents = %d\n", presents)
			return len(houses), nil
		default:
			return 0, fmt.Errorf("invalid token: [%s]\n", val)
		}
		fmt.Printf("santaIdx=%d\n", santaIdx)
		santaIdx++
		santaIdx %= santas
	}
	return len(houses), nil
}

func main() {
	houses, err := parse(2, os.Stdin)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	fmt.Printf("Houses %d\n", houses)
}
