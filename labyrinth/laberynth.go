package labyrinth

type Labyrinth struct {
	Labyrinth []rune
	Size      Coordinates
	Start     Coordinates
	End       Coordinates
}

type Coordinates struct {
	X int
	Y int
}

type LabObject struct {
	LabyrinthMap Labyrinth
	Roads        int
	Solutions    []Solution
}

type Solution struct {
	SolutionLabyrinthMap Labyrinth
	Steps                int
}
