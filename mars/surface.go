package mars

import "errors"

type LandableItem interface {
	GetX() int
	GetY() int
	GetOrientation() string
}

type Surface struct {
}

func NewSurface(x int, y int) (*Surface, error) {
	if x > 50 || x < 0 {
		return nil, errors.New("X cord out of range")
	}

	if y > 50 || y < 0 {
		return nil, errors.New("Y cord out of range")
	}

	return &Surface{}, nil
}

func (s *Surface) LandRover(li LandableItem) error {
	if li.GetX() < 0 {
		return errors.New("Landable item is not on the surface")
	}

	return nil
}

func (s *Surface) Run() string {
	return "1 1 N"
}
