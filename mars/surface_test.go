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
