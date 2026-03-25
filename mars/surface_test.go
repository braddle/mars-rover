package mars_test

import (
	"mars_rover/mars"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSurfaceXCordMustBeInRange(t *testing.T) {
	_, err := mars.NewSurface(51, 10)

	assert.ErrorContains(t, err, "X cord out of range")
}

func TestSurfaceYCordMustBeInRange(t *testing.T) {
	_, err := mars.NewSurface(10, 51)

	assert.ErrorContains(t, err, "Y cord out of range")
}

func TestXCordMustBePositiveNumber(t *testing.T) {
	_, err := mars.NewSurface(-1, 10)

	assert.ErrorContains(t, err, "X cord out of range")
}

func TestYCordMustBePositiveNumber(t *testing.T) {
	_, err := mars.NewSurface(10, -1)

	assert.ErrorContains(t, err, "Y cord out of range")
}

func TestRoverCannotBePlaceOffTheBottomOfTheSurface(t *testing.T) {
	mockLandableItem := new(MockLandableItem)
	mockLandableItem.On("GetX").Return(-1)
	s, _ := mars.NewSurface(3, 3)
	err := s.LandRover(mockLandableItem)

	assert.ErrorContains(t, err, "Landable item is not on the surface")
}

func TestRoverCannotBePlaceOffTheTopOfTheSurface(t *testing.T) {
	mockLandableItem := new(MockLandableItem)
	mockLandableItem.On("GetX").Return(4)
	s, _ := mars.NewSurface(3, 3)
	err := s.LandRover(mockLandableItem)

	assert.ErrorContains(t, err, "Landable item is not on the surface")
}

func TestRoverCannotBePlaceOffTheRightOfTheSurface(t *testing.T) {
	mockLandableItem := new(MockLandableItem)
	mockLandableItem.On("GetX").Return(1)
	mockLandableItem.On("GetY").Return(4)

	s, _ := mars.NewSurface(3, 3)
	err := s.LandRover(mockLandableItem)

	assert.ErrorContains(t, err, "Landable item is not on the surface")
}

func TestRoverCannotBePlaceOffTheLeftOfTheSurface(t *testing.T) {
	mockLandableItem := new(MockLandableItem)
	mockLandableItem.On("GetX").Return(1)
	mockLandableItem.On("GetY").Return(-1)

	s, _ := mars.NewSurface(3, 3)
	err := s.LandRover(mockLandableItem)

	assert.ErrorContains(t, err, "Landable item is not on the surface")
}

func TestRunningWithRoverWithNoInstructionsDoesNotMove(t *testing.T) {
	mockLandableItem := new(MockLandableItem)
	mockLandableItem.On("GetX").Return(1)
	mockLandableItem.On("GetY").Return(1)
	mockLandableItem.On("ReportLastPosition").Return("1 1 N")

	s, _ := mars.NewSurface(2, 2)

	err := s.LandRover(mockLandableItem)
	assert.NoError(t, err)

	out, err := s.Run()

	assert.NoError(t, err)
	assert.Equal(t, "1 1 N", out)
}

func TestRunningFailsIfNoRoversOnSurface(t *testing.T) {
	s, _ := mars.NewSurface(2, 2)

	_, err := s.Run()

	assert.ErrorContains(t, err, "No rovers on the surface")
}

func TestRealRoverCanRun(t *testing.T) {
	s, _ := mars.NewSurface(5, 5)
	r, _ := mars.NewRover(2, 2, "N", "FLF")

	s.LandRover(r)
	out, _ := s.Run()
	assert.Equal(t, "1 3 W", out)
}

func TestMultipleRealRoversCanRun(t *testing.T) {
	s, _ := mars.NewSurface(5, 5)

	r, _ := mars.NewRover(2, 2, "N", "FF")
	s.LandRover(r)
	r, _ = mars.NewRover(2, 2, "E", "FF")
	s.LandRover(r)
	r, _ = mars.NewRover(2, 2, "S", "FF")
	s.LandRover(r)
	r, _ = mars.NewRover(2, 2, "W", "FF")
	s.LandRover(r)

	out, _ := s.Run()
	exp := "2 4 N\n4 2 E\n2 0 S\n0 2 W"
	assert.Equal(t, exp, out)
}

func TestRoverRunningNorthOffTheSurface(t *testing.T) {
	s, _ := mars.NewSurface(5, 5)
	r, _ := mars.NewRover(2, 2, "N", "FFFFF")

	s.LandRover(r)
	out, _ := s.Run()
	assert.Equal(t, "2 5 N LOST", out)
}

func TestRoverRunningEastOffTheSurface(t *testing.T) {
	s, _ := mars.NewSurface(5, 5)
	r, _ := mars.NewRover(2, 2, "E", "FFFFF")

	s.LandRover(r)
	out, _ := s.Run()
	assert.Equal(t, "5 2 E LOST", out)
}

func TestRoverRunningSouthOffTheSurface(t *testing.T) {
	s, _ := mars.NewSurface(5, 5)
	r, _ := mars.NewRover(2, 2, "S", "FFFFF")

	s.LandRover(r)
	out, _ := s.Run()
	assert.Equal(t, "2 0 S LOST", out)
}

func TestRoverRunningWestOffTheSurface(t *testing.T) {
	s, _ := mars.NewSurface(5, 5)
	r, _ := mars.NewRover(2, 2, "W", "FFFFF")

	s.LandRover(r)
	out, _ := s.Run()
	assert.Equal(t, "0 2 W LOST", out)
}

type MockLandableItem struct {
	mock.Mock
}

func (m *MockLandableItem) ExecuteNextCommand() bool {
	return false
}

func (m *MockLandableItem) GetNextExpectedPosition() mars.Position {
	return mars.Position{0, 0}
}

func (m *MockLandableItem) GetX() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockLandableItem) GetY() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockLandableItem) ReportLastPosition() string {
	args := m.Called()
	return args.String(0)
}
