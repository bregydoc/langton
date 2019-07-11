package langton

type Langton struct {
	Ant   *Ant
	State [][]byte
}

func NewWithState(info []byte) *Langton {
	s := dataToState(info)
	return &Langton{
		State: s,
		Ant: &Ant{
			x: len(s) / 2,
			y: len(s) / 2,
			o: 0,
		},
	}
}

func New() *Langton {
	return &Langton{
		Ant: &Ant{},
	}
}

func (l *Langton) setState(info []byte) {
	s := dataToState(info)
	l.State = s
}

func (l *Langton) setAnt(x, y, o int) {
	l.Ant.x = x
	l.Ant.y = y
	l.Ant.o = o
}

func (l *Langton) encode(steps int64) error {
	n := int64(0)
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

func (l *Langton) decode(steps int64) error {
	n := int64(0)
	for n < steps {

		l.Ant.x, l.Ant.y = l.Ant.getPastPosition()

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

		n++
	}

	return nil
}
