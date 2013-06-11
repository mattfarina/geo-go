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

// Get the radius of the Earth at a specific latitude.
func (e *Earth) EarthRadius(latitude float64) float64 {
	ratLat := deg2rad(latitude)

	x := math.Cos(ratLat) / e.EarthRadiusSemimajor()
	y := math.Sin(ratLat) / e.EarthRadiusSemiminor()

	return 1 / (math.Sqrt((x * x) + (y * y)))
}

func (e *Earth) ConvertDecToDMS(coordinate float64) (degrees, minutes, seconds float64) {

	degrees, part := math.Modf(coordinate)

	// If the sign is negative we need to convert to a positive.
	part = math.Copysign(part, 1)

	temp := part * 3600
	minutes = math.Floor(temp / 60)
	seconds = temp - (minutes * 60)

	return
}
