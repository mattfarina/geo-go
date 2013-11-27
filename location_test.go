package geo

import (
	"testing"
)

func TestLocationBasics(t *testing.T) {
	location := new(Location)

	// Testing that Earth is available to location.
	if diff(location.EarthRadiusSemimajor(), 6378137.0) {
		t.Error("! EarthRadiusSemimajor is reporting the wrong value.")
	}
}

func TestLatitude(t *testing.T) {
	location := new(Location)

	location.SetLatitude(84.5)

	lat := location.Latitude()

	if lat != 84.5 {
		t.Error("! The latitude was not set or retrievable for the location.")
	}
}

func TestLongitude(t *testing.T) {
	location := new(Location)

	location.SetLongitude(-40.5)

	lat := location.Longitude()

	if lat != -40.5 {
		t.Error("! The longitude was not set or retrievable for the location.")
	}
}