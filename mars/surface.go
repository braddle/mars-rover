package mars

import "errors"

func NewSurface(x int, y int) (interface{}, error) {
	return nil, errors.New("X cord out of range")
}
