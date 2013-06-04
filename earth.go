// Provides location information specific to Earch.
package geo

import (
	"math"
)

// Earth contains information about the shape of the Earth.
type Earth struct{}

func (e *Earth) EarthRadiusSemimajor() float64 {
	// The distance is in meters.
	return 6378137.0
}

func (e *Earth) EarthFlattening() float64 {
	return 1 / 298.257223563
}

func (e *Earth) EarthEccentricitySq() float64 {
	// The response is 2 decimal places more accurate than the PHP version.
	return 2*e.EarthFlattening() - math.Pow(e.EarthFlattening(), 2)
}

func (e *Earth) EarthRadiusSemiminor() float64 {
	return (e.EarthRadiusSemimajor() * (1 - e.EarthFlattening()))
}
