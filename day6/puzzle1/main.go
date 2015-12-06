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
	lights [][]bool
	xy     struct {
		x int
		y int
	}
)

func (l lights) on(start, end xy) {
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			l[x][y] = true
		}
	}
}

func (l lights) off(start, end xy) {
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			l[x][y] = false
		}
	}
}

func (l lights) toggle(start, end xy) {
	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {
			l[x][y] = !l[x][y]
		}
	}
}

func (l lights) count() int {
	var on int
	for x := 0; x < xsize; x++ {
		for y := 0; y < ysize; y++ {
			if l[x][y] {
				on++
			}
		}
	}
	return on
}

func parse(r io.Reader, l lights) error {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	var sx, sy, ex, ey int
	for s.Scan() {
		line := s.Text()
		switch {
		case strings.HasPrefix(line, "turn on"):
			count, err := fmt.Sscanf(line, "turn on %d,%d through %d,%d", &sx, &sy, &ex, &ey)
			if count != 4 {
				return fmt.Errorf("expected 4 items scanned, got %d", count)
			}
			if err != nil {
				return err
			}
			l.on(xy{sx, sy}, xy{ex, ey})

		case strings.HasPrefix(line, "turn off"):
			count, err := fmt.Sscanf(line, "turn off %d,%d through %d,%d", &sx, &sy, &ex, &ey)
			if count != 4 {
				return fmt.Errorf("expected 4 items scanned, got %d", count)
			}
			if err != nil {
				return err
			}
			l.off(xy{sx, sy}, xy{ex, ey})

		case strings.HasPrefix(line, "toggle"):
			count, err := fmt.Sscanf(line, "toggle %d,%d through %d,%d", &sx, &sy, &ex, &ey)
			if count != 4 {
				return fmt.Errorf("expected 4 items scanned, got %d", count)
			}
			if err != nil {
				return err
			}
			l.toggle(xy{sx, sy}, xy{ex, ey})
		}
	}
	return nil
}

func main() {
	l := make(lights, xsize)
	for i := range l {
		l[i] = make([]bool, ysize)
	}
	err := parse(os.Stdin, l)
	if err != nil {
		fmt.Println("ERROR ", err)
		os.Exit(1)
	}
	fmt.Printf("lights on: %d\n", l.count())
}
