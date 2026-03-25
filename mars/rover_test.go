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
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 1 N", r.ReportLastPosition())
}

func TestRoverCanRotateLeftOnce(t *testing.T) {
	r, err := mars.NewRover(1, 1, "N", "L")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 1 W", r.ReportLastPosition())
}

func TestRoverCanRotateLeftTwice(t *testing.T) {
	r, err := mars.NewRover(1, 1, "N", "LL")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 1 S", r.ReportLastPosition())
}

func TestRoverCanRotateLeftThreeTimes(t *testing.T) {
	r, err := mars.NewRover(1, 1, "N", "LLL")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 1 E", r.ReportLastPosition())
}

func TestRoverCanRotateLeftFourTimes(t *testing.T) {
	r, err := mars.NewRover(1, 1, "N", "LLLL")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 1 N", r.ReportLastPosition())
}

func executeAllRoverCommands(r *mars.Rover) {
	for {
		ok := r.ExecuteNextCommand()
		if !ok {
			break
		}
	}
}
