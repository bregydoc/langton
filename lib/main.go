package main

import (
	"github.com/bregydoc/langton"
)

func main() {
	l := langton.NewLangton([]byte("b"))
	l.Exec()
}
