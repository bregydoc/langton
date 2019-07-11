package langton

import (
	"errors"
	"fmt"
	"strings"
)

func isqrt(number uint) uint { // above giant shoulders
	root := uint(1)
	next := (root + number/root) >> 1
	overflow := 100

	for root != next && overflow > 0 {
		root = next
		next = (root + number/root) >> 1
		overflow--
	}

	if root*root < number {
		return root + 1
	}

	return root
}

func printState(state [][]byte) {
	w := len(state)
	h := len(state[0])
	fmt.Println(strings.Repeat("-", w*2))
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			s := state[i][j]
			fmt.Print(s, " ")
		}
		fmt.Println("")
	}
	fmt.Println(strings.Repeat("-", w*2))
}

func dataToState(info []byte) [][]byte {
	bits := uint(8 * len(info))
	d := isqrt(bits)

	s := make([][]byte, d)
	for i := range s {
		s[i] = make([]byte, d)
	}

	for i := uint(0); i < d; i++ {
		for j := uint(0); j < d; j++ {
			l := i*d + j
			if l < bits {
				// fmt.Println("l: ", l/8, " i:", l%8)
				r := info[l/8] & (1 << (l % 8))
				if r != 0 {
					s[i][j] = 1
				}
			} else {
				s[i][j] = 0
			}

		}
	}

	return s
}

func getDim(state [][]byte) (uint, error) {
	d := uint(len(state))
	if d != uint(len(state[0])) {
		return 0, errors.New("invalid state dimensions")
	}

	return d, nil
}

func stateToData(state [][]byte) ([]byte, error) {
	d, err := getDim(state)
	if err != nil {
		return nil, err
	}

	data := make([]byte, 0)
	for k := uint(0); k < d*d/8; k++ {
		data = append(data, byte(0))
	}
	for i := uint(0); i < d; i++ {
		for j := uint(0); j < d; j++ {
			l := i*d + j
			if l/8 < uint(len(data)) {
				data[l/8] += state[i][j] << (l % 8)
			}
		}
	}

	return data, nil
}
