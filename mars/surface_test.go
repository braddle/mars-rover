package mars_test

import (
	"testing"
	"untitled/mars"

	"github.com/stretchr/testify/assert"
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
