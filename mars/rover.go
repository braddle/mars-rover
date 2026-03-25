package mars

import (
	"errors"
	"fmt"
)

type Rover struct {
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

	return &Rover{instructions: instructions, direction: direction}, nil
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

	r.currentInstruction++
	return true
}

func (r *Rover) ReportLastPosition() string {
	return fmt.Sprintf("1 1 %s", r.direction)
}
