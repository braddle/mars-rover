package mars

import (
	"errors"
	"fmt"
)

type Rover struct {
	y                  int
	x                  int
	instructions       string
	direction          string
	currentInstruction int
}

const turnLeft = "L"
const turnRight = "R"

const north = "N"
const west = "W"
const south = "S"
const east = "E"

func NewRover(x int, y int, direction string, instructions string) (*Rover, error) {
	if len(instructions) > 100 {
		return nil, errors.New("Rover has more than 100 instructions")
	}

	return &Rover{x: x, y: y, instructions: instructions, direction: direction}, nil
}

func (r *Rover) GetX() int {
	return r.x
}

func (r *Rover) GetY() int {
	return r.y
}

func (r *Rover) ExecuteNextCommand() bool {
	if len(r.instructions) == 0 || r.currentInstruction >= len(r.instructions) {
		return false
	}

	if string(r.instructions[r.currentInstruction]) == turnLeft {
		if r.direction == north {
			r.direction = west
		} else if r.direction == west {
			r.direction = south
		} else if r.direction == south {
			r.direction = east
		} else if r.direction == east {
			r.direction = north
		}
	}

	if string(r.instructions[r.currentInstruction]) == turnRight {
		if r.direction == north {
			r.direction = east
		} else if r.direction == east {
			r.direction = south
		} else if r.direction == south {
			r.direction = west
		} else if r.direction == west {
			r.direction = north
		}
	}

	if string(r.instructions[r.currentInstruction]) == "F" {
		if r.direction == north {
			r.y = r.y + 1
		} else if r.direction == east {
			r.x = r.x + 1
		} else if r.direction == south {
			r.y = r.y - 1
		} else if r.direction == west {
			r.x = r.x - 1
		}
	}

	r.currentInstruction++
	return true
}

func (r *Rover) GetNextExpectedPosition() Position {
	if len(r.instructions) == 0 || r.currentInstruction >= len(r.instructions) {
		return Position{r.x, r.y}
	}

	if string(r.instructions[r.currentInstruction]) == "F" {
		if r.direction == north {
			return Position{r.x, r.y + 1}
		} else if r.direction == east {
			return Position{r.x + 1, r.y}
		} else if r.direction == south {
			return Position{r.x, r.y - 1}
		} else if r.direction == west {
			return Position{r.x - 1, r.y}
		}
	}

	return Position{r.x, r.y}
}

func (r *Rover) ReportLastPosition() string {
	status := ""
	if r.currentInstruction < len(r.instructions) {
		status = " LOST"
	}

	return fmt.Sprintf("%d %d %s%s", r.x, r.y, r.direction, status)
}
