package geo

import (
	"testing"
)

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

func TestEarthRadius(t *testing.T) {

	ret := new(Earth)

	m := map[float64]float64{
		1:  6378130.4536189,
		10: 6377489.0140512,
		45: 6367417.7249667,
		90: 6356752.3142452,
	}

	for key, val := range m {
		if diff(ret.EarthRadius(key), val) {
			t.Error("! EarthRadius is reporting the wrong value.")
		}
	}
}

func TestConvertDecToDMS(t *testing.T) {

	earth := new(Earth)

	// The location we are testing is for Michigan State University.
	degrees, minutes, seconds := earth.ConvertDecToDMS(42.7186)
	if degrees != 42 || minutes != 43 || diff(seconds, 6.96) {
		t.Error("! ConvertDecToDMS incorrectly converting.")
	}

	degrees, minutes, seconds = earth.ConvertDecToDMS(-84.468466)
	if degrees != -84 || minutes != 28 || diff(seconds, 6.4776) {
		t.Error("! ConvertDecToDMS incorrectly converting.")
	}
}
