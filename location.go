package geo

// Location is a location on the Earth.
type Location struct{
	Earth

	// latitude, longitude, and height are where the location resides on the Earth.
	latitude float64
	longitude float64
	height float64
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