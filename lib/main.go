package main

import (
	"github.com/bregydoc/langton"
)

func main() {
	l := langton.New()
	data, err := l.Encode([]byte("hello world"), 10)
	if err != nil {
		panic(err)
	}
	// log.Println(string(data))

	data, err = l.Decode(data, l.Ant, 10)
	if err != nil {
		panic(err)
	}
	// log.Println(string(data))
}
