package labyrinth

import (
	"io/ioutil"
	"sync"
)

func (l *LabObject) Read() error {
	bytes, err := ioutil.ReadFile("intro.lab")
	if err != nil {
		return err
	}

	str := string(bytes)
	l.LabyrinthMap = Labyrinth{
		Labyrinth: make([]rune, 0),
	}

	sx, sy := sizeLab(str)
	l.LabyrinthMap.Size.X = sx
	l.LabyrinthMap.Size.Y = sy

	for i := range str {
		if str[i] != '\n' && (str[i] == '@' || str[i] == ' ' || str[i] == 'E' || str[i] == 'S') {
			l.LabyrinthMap.Labyrinth = append(l.LabyrinthMap.Labyrinth, rune(str[i]))
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		l.LabyrinthMap.findStart()
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		l.LabyrinthMap.findEnd()
		wg.Done()
	}(&wg)


	wg.Wait()

	return nil
}

func sizeLab(str string) (int, int) {
	y := 0
	x := 0

	for i := range str {
		if str[i] == '\n' {
			y++
		} else {
			if y == 0 {
				x++
			}
		}
	}

	return x, y
}

func (l *Labyrinth) findStart() {
	for i := range l.Labyrinth {
		if l.Labyrinth[i] == 'E' {
			l.Start.X = i % l.Size.X
			l.Start.Y = (i % l.Size.Y) + (i/l.Size.X)
		}
	}

	//TODO: Validate
}

func (l *Labyrinth)findEnd() {
	for i := range l.Labyrinth {
		if l.Labyrinth[i] == 'S' {
			l.End.X = i % l.Size.X
			l.End.Y = (i % l.Size.Y) + (i/l.Size.X)
		}
	}

	//TODO: Validate
}