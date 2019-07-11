package langton

import "fmt"

type Langton struct {
	Ant   *AntDescriptor
	State [][]byte
}

func NewLangton(info []byte) *Langton {
	s := dataToState(info)
	return &Langton{
		State: s,
		Ant: &AntDescriptor{
			x: len(s) / 2,
			y: len(s) / 2,
			o: 0,
		},
	}
}

func (l *Langton) encode(steps int) error {
	n := 0
	for n < steps {
		if err := l.Ant.fixPosition(l.State); err != nil {
			return err
		}

		x, y := l.Ant.getPosition()

		if l.State[x][y] == 0 {
			l.Ant.o = l.Ant.o - 1
			if l.Ant.o < 0 {
				l.Ant.o = 3
			}
			l.State[x][y] = 1
		} else {
			l.Ant.o = l.Ant.o + 1
			if l.Ant.o > 3 {
				l.Ant.o = 0
			}
			l.State[x][y] = 0
		}

		l.Ant.x, l.Ant.y = l.Ant.getNextPosition()
		n++
	}

	return nil
}

func (l *Langton) decode(steps int) error {
	n := 0
	for n < steps {
		if err := l.Ant.fixPosition(l.State); err != nil {
			return err
		}

		x, y := l.Ant.getPosition()

		if l.State[x][y] == 0 {
			l.Ant.o = l.Ant.o - 1
			if l.Ant.o < 0 {
				l.Ant.o = 3
			}
			l.State[x][y] = 1
		} else {
			l.Ant.o = l.Ant.o + 1
			if l.Ant.o > 3 {
				l.Ant.o = 0
			}
			l.State[x][y] = 0
		}

		l.Ant.x, l.Ant.y = l.Ant.getPastPosition()
		n++
	}

	return nil
}

func (l *Langton) Exec() {
	data, _ := stateToData(l.State)

	fmt.Printf("ant: x:%d, y:%d, o:%d\n", l.Ant.x, l.Ant.y, l.Ant.o)
	fmt.Println(string(data))
	printState(l.State)

	if err := l.encode(1); err != nil {
		panic(err)
	}

	data, _ = stateToData(l.State)

	fmt.Printf("ant: x:%d, y:%d, o:%d\n", l.Ant.x, l.Ant.y, l.Ant.o)
	fmt.Println(string(data))
	printState(l.State)

	if err := l.decode(1); err != nil {
		panic(err)
	}

	data, _ = stateToData(l.State)
	fmt.Printf("ant: x:%d, y:%d, o:%d\n", l.Ant.x, l.Ant.y, l.Ant.o)
	fmt.Println(string(data))
	printState(l.State)
}
