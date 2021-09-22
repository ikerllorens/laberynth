package main

import (
	"labyrinth/labyrinth"
)

func main() {
	l := labyrinth.LabObject{
		LabyrinthMap: labyrinth.Labyrinth{},
		Roads:        0,
		Solutions:    nil,
	}
	err := l.Read()
	if err != nil {
		return
	}



	l.LabyrinthMap.Print()

}

func LookPaths(lab *labyrinth.LabObject) {
	solution := lab.LabyrinthMap

}

func Move(lab *labyrinth.LabObject) {
	

}