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

func TestRoverCanRotateRightOnce(t *testing.T) {
	r, err := mars.NewRover(1, 1, "N", "R")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 1 E", r.ReportLastPosition())
}

func TestRoverCanRotateRightTwice(t *testing.T) {
	r, err := mars.NewRover(1, 1, "N", "RR")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 1 S", r.ReportLastPosition())
}

func TestRoverCanRotateRightThreeTimes(t *testing.T) {
	r, err := mars.NewRover(1, 1, "N", "RRR")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 1 W", r.ReportLastPosition())
}

func TestRoverCanRotateLRightFourTimes(t *testing.T) {
	r, err := mars.NewRover(1, 1, "N", "RRRR")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 1 N", r.ReportLastPosition())
}

func TestRoverCanMoveNorth(t *testing.T) {
	r, err := mars.NewRover(1, 1, "N", "F")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 2 N", r.ReportLastPosition())
}

func TestRoverCanMoveEast(t *testing.T) {
	r, err := mars.NewRover(1, 1, "E", "F")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "2 1 E", r.ReportLastPosition())
}

func TestRoverCanMoveSouth(t *testing.T) {
	r, err := mars.NewRover(1, 1, "S", "F")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "1 0 S", r.ReportLastPosition())
}

func TestRoverCanMoveWest(t *testing.T) {
	r, err := mars.NewRover(1, 1, "W", "F")
	assert.NoError(t, err)

	executeAllRoverCommands(r)
	assert.Equal(t, "0 1 W", r.ReportLastPosition())
}

func TestIfRoverDontNotCompleteAllInstructionsItIsLost(t *testing.T) {
	r, err := mars.NewRover(1, 1, "W", "FF")
	assert.NoError(t, err)

	r.ExecuteNextCommand()
	assert.Equal(t, "0 1 W LOST", r.ReportLastPosition())
}

func TestGetNextExpectedPositionIdempotent(t *testing.T) {
	r, _ := mars.NewRover(1, 1, "W", "FF")
	assert.Equal(t, mars.Position{0, 1}, r.GetNextExpectedPosition())
	assert.Equal(t, mars.Position{0, 1}, r.GetNextExpectedPosition())
	assert.Equal(t, mars.Position{0, 1}, r.GetNextExpectedPosition())
}

func executeAllRoverCommands(r *mars.Rover) {
	for {
		ok := r.ExecuteNextCommand()
		if !ok {
			break
		}
	}
}
