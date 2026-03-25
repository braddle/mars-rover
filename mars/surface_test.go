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

func TestRunningWithRoverWithNoInstructionsDoesNotMove(t *testing.T) {
	mockLandableItem := new(MockLandableItem)
	mockLandableItem.On("GetX").Return(1)
	mockLandableItem.On("GetY").Return(1)
	mockLandableItem.On("GetOrientation").Return("N")
	s, _ := mars.NewSurface(2, 2)

	err := s.LandRover(mockLandableItem)
	assert.NoError(t, err)

	out := s.Run()

	assert.Equal(t, "1 1 N", out)

}

type MockLandableItem struct {
	mock.Mock
}

func (m *MockLandableItem) GetX() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockLandableItem) GetY() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockLandableItem) GetOrientation() string {
	args := m.Called()
	return args.String(0)
}
