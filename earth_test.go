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

func TestBasics(t *testing.T) {

	ret := new(Earth)

	if diff(ret.EarthRadiusSemimajor(), 6378137.0) {
		t.Error("! EarthRadiusSemimajor is reporting the wrong value.")
	}

	if diff(ret.EarthEccentricitySq(), 0.006694379990141317) {
		t.Error("! EarthEccentricitySq is reporting the wrong value.")
	}

	if diff(ret.EarthRadiusSemiminor(), 6.356752314245179e+06) {
		t.Error("! EarthRadiusSemiminor is reporting the wrong value.")
	}

	if diff(ret.EarthFlattening(), 0.003352810664747481) {
		t.Error("! EarthFlattening is reporting the wrong value.")
	}
}
