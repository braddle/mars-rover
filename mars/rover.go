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

	if string(r.instructions[r.currentInstruction]) == "L" {
		if r.direction == "N" {
			r.direction = "W"
		} else if r.direction == "W" {
			r.direction = "S"
		} else if r.direction == "S" {
			r.direction = "E"
		} else if r.direction == "E" {
			r.direction = "N"
		}
	}

	r.currentInstruction++
	return true
}

func (r *Rover) ReportLastPosition() string {
	return fmt.Sprintf("1 1 %s", r.direction)
}
