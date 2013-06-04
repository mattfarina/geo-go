package geo

import (
	"testing"
)

func TestBasics(t *testing.T) {

	ret := new(Earth)

	if ret.EarthRadiusSemimajor() != 6378137.0 {
		t.Error("! EarthRadiusSemimajor is reporting the wrong value.")
	}

	if ret.EarthEccentricitySq() != 0.006694379990141317 {
		t.Error("! EarthEccentricitySq is reporting the wrong value.")
	}

	if ret.EarthRadiusSemiminor() != 6.356752314245179e+06 {
		t.Error("! EarthRadiusSemiminor is reporting the wrong value.")
	}

	if ret.EarthFlattening() != 0.003352810664747481 {
		t.Error("! EarthFlattening is reporting the wrong value.")
	}
}
