package mars_test

import (
	"mars_rover/mars"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoverCannotHaveMoreThat100Instructions(t *testing.T) {
	_, err := mars.NewRover(1, 1, "N", strings.Repeat("L", 101))

	assert.ErrorContains(t, err, "Rover has more than 100 instructions")
}

func TestRoverReportInitalPositionIfItHasNoInstructions(t *testing.T) {
	r, err := mars.NewRover(1, 1, "N", "")
	r.ExecuteNextCommand()

	assert.NoError(t, err)
	assert.Equal(t, "1 1 N", r.ReportLastPosition())
}
