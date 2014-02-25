package geo

import (
	"github.com/Masterminds/convert/radial"
	"math"
)

// Earth contains information about the shape of the Earth.
type Earth struct{}

// EarthRadiusSemimajor returns the Semi-major radius of the Earth in meters.
func (e *Earth) EarthRadiusSemimajor() float64 {
	// The distance is in meters.
	return 6378137.0
}

// EarthFlattening performs a flattening operation on the Earth. This is an
// ellipsoid operation. For more details see http://en.wikipedia.org/wiki/Flattening
func (e *Earth) EarthFlattening() float64 {
	return 1 / 298.257223563
}

// EarthEccentricitySq returns the result of eccentricity operations for the
// Earth. For more information see http://en.wikipedia.org/wiki/Eccentricity_%28mathematics%29
func (e *Earth) EarthEccentricitySq() float64 {
	// The response is 2 decimal places more accurate than the PHP version.
	return 2 * e.EarthFlattening() - math.Pow(e.EarthFlattening(), 2)
}

// EarthRadiusSemiminor returns the semi-minor value for the Earth, an ellipsoid.
// For more information see http://en.wikipedia.org/wiki/Reference_ellipsoid
func (e *Earth) EarthRadiusSemiminor() float64 {
	return (e.EarthRadiusSemimajor() * (1 - e.EarthFlattening()))
}

// EarthRadius returns the radius of the Earth in meters at a specific latitude.
func (e *Earth) EarthRadius(latitude float64) float64 {
	ratLat := radial.DegToRad(latitude)

	x := math.Cos(ratLat) / e.EarthRadiusSemimajor()
	y := math.Sin(ratLat) / e.EarthRadiusSemiminor()

	return 1 / (math.Sqrt((x * x) + (y * y)))
}

// ConvertDecToDMS turns a coordinate into a value represented by degrees,
// minutes, and seconds.
func (e *Earth) ConvertDecToDMS(coordinate float64) (degrees, minutes, seconds float64) {

	degrees, part := math.Modf(coordinate)

	// If the sign is negative we need to convert to a positive.
	part = math.Copysign(part, 1)

	temp := part * 3600
	minutes = math.Floor(temp / 60)
	seconds = temp - (minutes * 60)

	return
}

// ConvertDMStoDec turns a value represented in degrees, minutes, and seconds
// into a coordinate value.
func (e *Earth) ConvertDMStoDec(degrees, minutes, seconds float64) float64 {
	if degrees < 0 {
		return degrees - (((minutes * 60) + seconds) / 3600)
	}

	return degrees + (((minutes * 60) + seconds) / 3600)
}
