package geo

import (
	"testing"
)

const MAX_DIFF float64 = 0.0000001

func diff(a, b float64) bool {
	if a > b {
		return (a - b) > MAX_DIFF
	}
	return (b - a) > MAX_DIFF
}

func TestDeg2rad(t *testing.T) {
	if diff(deg2rad(1), 0.0174533) {
		t.Error("! deg2rad is not accurately converting between degrees and radians.")
	}

	if diff(deg2rad(5), 0.0872665) {
		t.Error("! deg2rad is not accurately converting between degrees and radians.")
	}
}
