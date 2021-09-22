package labyrinth

import "fmt"

func (l *Labyrinth) Print() {
	for j := 0; j < l.Size.Y; j++ {
		for i := 0; i < l.Size.X; i++ {
			fmt.Printf("%c", l.Labyrinth[i+(j*l.Size.X)])
			//fmt.Println(i+(j*l.Size.Y))
		}
		fmt.Printf("\n")
	}

	fmt.Println("Start:", l.Start.X, l.Start.Y)
	fmt.Println("End:", l.End.X, l.End.Y)
}
