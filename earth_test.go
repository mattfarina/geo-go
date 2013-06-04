package geo

import (
	"testing"
)

func TestBasics(t *testing.T) {

	ret := new(Earth)

	if ret.EarthRadiusSemimajor() != 6378137.0 {
		t.Error("! EarthRadiusSemimajor is reporting the wrong value.")
	}
}
