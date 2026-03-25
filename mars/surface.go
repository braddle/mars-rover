package mars

import "errors"

type LandableItem interface {
	GetX() int
	GetY() int
	ReportLastPosition() string
}

type Surface struct {
	x      int
	y      int
	rovers []LandableItem
}

const maximumSurfaceSize = 50
const minimumSurfaceSize = 0

func NewSurface(x int, y int) (*Surface, error) {
	if x > maximumSurfaceSize || x < minimumSurfaceSize {
		return nil, errors.New("X cord out of range")
	}

	if y > maximumSurfaceSize || y < minimumSurfaceSize {
		return nil, errors.New("Y cord out of range")
	}

	return &Surface{x: x, y: y, rovers: make([]LandableItem, 0)}, nil
}

func (s *Surface) LandRover(li LandableItem) error {
	if li.GetX() < 0 || li.GetX() > s.x || li.GetY() > s.y || li.GetY() < 0 {
		return errors.New("Landable item is not on the surface")
	}

	s.rovers = append(s.rovers, li)

	return nil
}

func (s *Surface) Run() (string, error) {
	if len(s.rovers) == 0 {
		return "", errors.New("No rovers on the surface")
	}

	return s.rovers[0].ReportLastPosition(), nil
}
