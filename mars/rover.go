package mars

import "errors"

type Rover struct {
}

func NewRover(x int, y int, direction string, instructions string) (*Rover, error) {
	if len(instructions) > 100 {
		return nil, errors.New("Rover has more than 100 instructions")
	}

	return &Rover{}, nil
}

func (r *Rover) ExecuteNextCommand() {
}

func (r *Rover) ReportLastPosition() string {
	return "1 1 N"
}
