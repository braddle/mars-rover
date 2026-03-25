package mars

import "errors"

func NewSurface(x int, y int) (interface{}, error) {
	if x > 50 || x < 0 {
		return nil, errors.New("X cord out of range")
	}

	return nil, errors.New("Y cord out of range")
}
