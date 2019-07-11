package langton

import (
	"fmt"
	"time"
)

func (l *Langton) Encode(data []byte, steps int64) ([]byte, error) {
	l.setState(data)
	l.setAnt(len(data)/2, len(data)/2, 0) // ant in the exactly middle
	if err := l.encode(steps); err != nil {
		return nil, err
	}

	return stateToData(l.State)
}

func (l *Langton) Decode(data []byte, ant *Ant, steps int64) ([]byte, error) {
	l.setState(data)
	l.setAnt(ant.x, ant.y, ant.o) // ant in the exactly middle
	if err := l.decode(steps); err != nil {
		return nil, err
	}

	return stateToData(l.State)
}

func (l *Langton) Exec() {
	data, _ := stateToData(l.State)
	steps := time.Now().Unix() % 1000000
	fmt.Println(steps)
	// fmt.Printf("ant: x:%d, y:%d, o:%d\n", l.Ant.x, l.Ant.y, l.Ant.o)
	fmt.Println(string(data))
	printState(l.State)

	if err := l.encode(steps); err != nil {
		panic(err)
	}

	data, _ = stateToData(l.State)

	// fmt.Printf("ant: x:%d, y:%d, o:%d\n", l.Ant.x, l.Ant.y, l.Ant.o)
	fmt.Println(string(data))
	printState(l.State)

	if err := l.decode(steps); err != nil {
		panic(err)
	}

	data, _ = stateToData(l.State)
	// fmt.Printf("ant: x:%d, y:%d, o:%d\n", l.Ant.x, l.Ant.y, l.Ant.o)
	fmt.Println(string(data))
	printState(l.State)
}
