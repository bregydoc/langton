package langton

type Ant struct {
	x int
	y int
	o int // 0, 1, 2, 3 => N, E, S, W
}

func (ant *Ant) getPosition() (int, int) {
	return ant.x, ant.y
}

func (ant *Ant) getNextPosition() (int, int) {
	if ant.o == 0 {
		// fmt.Println("ant: N")
		return ant.x, ant.y - 1
	} else if ant.o == 1 {
		// fmt.Println("ant: E")
		return ant.x + 1, ant.y
	} else if ant.o == 2 {
		// fmt.Println("ant: S")
		return ant.x, ant.y + 1
	} else if ant.o == 3 {
		// fmt.Println("ant: W")
		return ant.x - 1, ant.y
	}
	return ant.x, ant.y
}

func (ant *Ant) getPastPosition() (int, int) {
	if ant.o == 0 {
		// fmt.Println("ant: N")
		return ant.x, ant.y + 1
	} else if ant.o == 1 {
		// fmt.Println("ant: E")
		return ant.x - 1, ant.y
	} else if ant.o == 2 {
		// fmt.Println("ant: S")
		return ant.x, ant.y - 1
	} else if ant.o == 3 {
		// fmt.Println("ant: W")
		return ant.x + 1, ant.y
	}
	return ant.x, ant.y
}

func (ant *Ant) fixPosition(state [][]byte) error {
	d, err := getDim(state)
	if err != nil {
		return err
	}

	if ant.x < 0 {
		ant.x = int(d) + ant.x
	}

	if ant.y < 0 {
		ant.y = int(d) + ant.y
	}

	if ant.x >= int(d) {
		ant.x = 0
	}

	if ant.y >= int(d) {
		ant.y = 0
	}

	return nil
}
