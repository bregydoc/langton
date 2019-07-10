package main

import (
	"encoding/base64"
	"fmt"
	"math"
	"strings"
)

type ant struct {
	x int
	y int
	o int // 0, 1, 2, 3 => N, E, S, W
}

func countBits(info []byte) int {
	return 4 * len(info)
}

func printState(state [][]int) {
	w := len(state)
	h := len(state[0])

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			s := state[i][j]
			fmt.Print(s, " ")
		}
		fmt.Println("")
	}
	fmt.Println(strings.Repeat("-", w*2))
}

func performLangton(initialState [][]int, a *ant, steps int) [][]int {
	w := len(initialState)
	h := len(initialState[0])

	for i := 0; i < steps; i++ {
		if initialState[a.x][a.y] == 0 {
			a.o = a.o - 1
			if a.o < 0 {
				a.o = 3
			}
			initialState[a.x][a.y] = 1
		} else {
			a.o = a.o + 1
			if a.o > 3 {
				a.o = 0
			}
			initialState[a.x][a.y] = 0
		}

		if a.o == 0 {
			a.y--
		} else if a.o == 1 {
			a.x++
		} else if a.o == 2 {
			a.y++
		} else if a.o == 3 {
			a.x--
		}

		if a.x < 0 {
			a.x = w - 1
		}

		if a.y < 0 {
			a.y = h - 1
		}

		if a.x >= w {
			a.x = 0
		}

		if a.y >= h {
			a.y = 0
		}
	}

	return initialState
}

func diff(a [][]int, b [][]int) int {
	w := len(a)
	h := len(a[0])

	totalDiffs := 0

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if a[i][j]-b[i][j] != 0 {
				totalDiffs++
			}
		}
	}

	return totalDiffs
}

func newAntFromInfo(info []byte) *ant {
	bits := countBits(info)
	d := int(math.Ceil(math.Sqrt(float64(bits))))
	dx, dy := d, d
	return &ant{
		x: dx / 2,
		y: dy / 2,
	}
}

func infoToState(inf []byte) [][]int {
	bits := countBits(inf)
	d := int(math.Ceil(math.Sqrt(float64(bits))))
	dx, dy := d, d
	s := make([][]int, dy)
	for i := range s {
		s[i] = make([]int, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			if inf[i]&(1<<uint(j)) != 0 {
				s[i][j] = 1
			}

		}
	}

	return s
}

func stateToInfo(state [][]int) []byte {

	w := len(state)
	h := len(state[0])

	info := make([]byte, 0)
	for i := 0; i < w; i++ {
		var b byte
		for j := 0; j < h; j++ {
			if state[i][j] == 1 {
				b += 1 << uint(j)
			}
		}
		info = append(info, b)
	}

	return info
}

func main() {
	information := "bregy malpartida ramos asiduasijdndnjaksdjnaskjndjkas"
	inf := []byte(information)

	s := infoToState(inf)

	a := newAntFromInfo(inf)
	printState(s)
	for step := 0; step < 100000; step++ {
		oldS := make([][]int, len(s))
		for i := range s {
			oldS[i] = make([]int, len(s[i]))
			copy(oldS[i], s[i])
		}

		s = performLangton(s, a, 1)

		// printState(oldS)
		// printState(s)
		// fmt.Println(diff(oldS, s))
	}
	printState(s)

	final := base64.StdEncoding.EncodeToString(stateToInfo(s))
	fmt.Println(final)

	data, _ := base64.StdEncoding.DecodeString(final)
	fmt.Println(data)
}
