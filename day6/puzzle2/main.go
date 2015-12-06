package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const xsize = 1000
const ysize = 1000

type (
	lights [][]int
	xy     struct {
		x int
		y int
	}
)

func (l lights) on(start, end xy) {
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			l[x][y]++
		}
	}
}

func (l lights) off(start, end xy) {
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			if l[x][y] > 0 {
				l[x][y]--
			}
		}
	}
}

func (l lights) toggle(start, end xy) {
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			l[x][y] += 2
		}
	}
}

func (l lights) brightness() int {
	var totalBrightness int
	for x := 0; x < xsize; x++ {
		for y := 0; y < ysize; y++ {
			totalBrightness += l[x][y]
		}
	}
	return totalBrightness
}

func parseLine(line string, l lights) error {
	var start, end xy
	var cmd string
	switch {
	case strings.HasPrefix(line, "turn on"):
		line = strings.TrimPrefix(line, "turn on ")
		cmd = "on"
	case strings.HasPrefix(line, "toggle "):
		line = strings.TrimPrefix(line, "toggle ")
		cmd = "toggle"
	case strings.HasPrefix(line, "turn off "):
		line = strings.TrimPrefix(line, "turn off ")
		cmd = "off"
	default:
		return fmt.Errorf("unknown line %s", line)
	}
	count, err := fmt.Sscanf(line, "%d,%d through %d,%d", &start.x, &start.y, &end.x, &end.y)
	if count != 4 {
		return fmt.Errorf("expected 4 items scanned, got %d", count)
	}
	if err != nil {
		return err
	}
	switch cmd {
	case "on":
		l.on(start, end)
	case "off":
		l.off(start, end)
	case "toggle":
		l.off(start, end)
	}
	return nil
}

func parse(r io.Reader, l lights) error {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		err := parseLine(line, l)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	l := make(lights, xsize)
	for i := range l {
		l[i] = make([]int, ysize)
	}
	err := parse(os.Stdin, l)
	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(1)
	}
	fmt.Printf("lights brightness: %d\n", l.brightness())
}
