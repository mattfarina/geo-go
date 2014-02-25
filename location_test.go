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

func TestDistance(t *testing.T) {
	google := NewLocation(37.422045, -122.084347) // Google HQ
	sf := NewLocation(37.77493, -122.419416)      // San Fransisco, CA
	ef := NewLocation(48.8582, 2.294407)          // Eiffel Tower
	opera := NewLocation(-33.856553, 151.214696)  // Sydney Opera House

	distance, err := google.Distance(sf)
	if err != nil {
		t.Error("! Error calculating the distance between Google and San Francisco")
	}
	if diff(distance, 49087.0658381) {
		t.Error("! An incorrect distance was calculated between Google and San Francisco", distance)
	}

	distance, err = google.Distance(ef)
	if err != nil {
		t.Error("! Error calculating the distance between Google and the Eiffel Tower")
	}
	if diff(distance, 8989724.3991584) {
		t.Error("! An incorrect distance was calculated between Google and the Eiffel Tower", distance)
	}

	distance, err = google.Distance(opera)
	if err != nil {
		t.Error("! Error calculating the distance between Google and the Sydney Opera House")
	}
	if diff(distance, 11939773.6402807) {
		t.Error("! An incorrect distance was calculated between Google and the Sydney Opera House", distance)
	}
}
