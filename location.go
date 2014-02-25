package geo

import (
	"errors"
	"github.com/Masterminds/convert/radial"
	"math"
)

// Location is a location on the Earth.
type Location struct {
	Earth

	// latitude, longitude, and height are where the location resides on the Earth.
	latitude  float64
	longitude float64
	height    float64
}

// NewLocation creates a new Location based on a latitude and longitude.
func NewLocation(lat, long float64) Location {
	l := Location{}
	l.SetLatitude(lat)
	l.SetLongitude(long)
	return l
}

// Latitude returns the latitude for this location on the Earth.
func (l *Location) Latitude() float64 {
	return l.latitude
}

// SetLatitude sets the latitude for this location on the Earth.
func (l *Location) SetLatitude(lat float64) {
	l.latitude = lat
}

// Longitude returns the longitude for this location on the Earth.
func (l *Location) Longitude() float64 {
	return l.longitude
}

// SetLongitude sets the longitude for this location on the Earth.
func (l *Location) SetLongitude(long float64) {
	l.longitude = long
}

// func (l *Location) LatitudeRange(distance float64) (north, south Location) {

// }

// func (l *Location) LongitudeRange(distance float64) (east, west Location) {

// }

// Distance returns the distance, in meters, between two locations. One location
// is this one and the second location is passed in. Vincenty's Forumlae is
// used to calculate the distance. For more information see
// https://en.wikipedia.org/wiki/Vincenty%27s_formulae
func (l *Location) Distance(location Location) (float64, error) {
	lat1 := radial.DegToRad(l.Latitude())
	long1 := radial.DegToRad(l.Longitude())
	lat2 := radial.DegToRad(location.Latitude())
	long2 := radial.DegToRad(location.Longitude())

	a := l.EarthRadiusSemimajor()
	b := l.EarthRadiusSemiminor()
	f := l.EarthFlattening()

	la := long2 - long1
	lambda := la
	lambdaP := 2 * math.Pi

	u1 := math.Atan((1 - f) * math.Tan(lat1))
	u2 := math.Atan((1 - f) * math.Tan(lat2))

	sinU1 := math.Sin(u1)
	cosU1 := math.Cos(u1)
	sinU2 := math.Sin(u2)
	cosU2 := math.Cos(u2)

	var sinLambda, cosLambda, sinSigma, cosSigma, sigma, alpha, cosSqAlpha float64
	var cos2SigmaM, c float64

	iterLimit := 20

	for ; (math.Abs(lambda-lambdaP) > 1e-12) && iterLimit > 0; iterLimit-- {
		sinLambda = math.Sin(lambda)
		cosLambda = math.Cos(lambda)
		sinSigma = math.Sqrt((cosU2*sinLambda)*(cosU2*sinLambda) + (cosU1*sinU2-sinU1*cosU2*cosLambda)*(cosU1*sinU2-sinU1*cosU2*cosLambda))

		if sinSigma == 0 {
			return 0, nil
		}

		cosSigma = sinU1*sinU2 + cosU1*cosU2*cosLambda
		sigma = math.Atan2(sinSigma, cosSigma)
		alpha = math.Asin(cosU1 * cosU2 * sinLambda / sinSigma)
		cosSqAlpha = math.Cos(alpha) * math.Cos(alpha)
		cos2SigmaM = cosSigma - 2*sinU1*sinU2/cosSqAlpha
		c = f / 16 * cosSqAlpha * (4 + f*(4-3*cosSqAlpha))
		lambdaP = lambda
		lambda = la + (1-c)*f*math.Sin(alpha)*(sigma+c*sinSigma*(cos2SigmaM+c*cosSigma*(-1+2*cos2SigmaM*cos2SigmaM)))
	}

	// This would mean we have an error
	if iterLimit == 0 {
		return 0, errors.New("Geo: Error calculating distance.")
	}

	uSq := cosSqAlpha * (a*a - b*b) / (b * b)
	A := 1 + uSq/16384*(4096+uSq*(-768+uSq*(320-175*uSq)))
	B := uSq / 1024 * (256 + uSq*(-128+uSq*(74-47*uSq)))

	deltaSigma := B * sinSigma * (cos2SigmaM + B/4*(cosSigma*(-1+2*cos2SigmaM*cos2SigmaM)-B/6*cos2SigmaM*(-3+4*sinSigma*sinSigma)*(-3+4*cos2SigmaM*cos2SigmaM)))

	s := b * A * (sigma - deltaSigma)

	return s, nil
}
