package geo

// Location is a location on the Earth
type Location struct{
	Earth

	// latitude, longitude, and height are where the location resides on the Earth.
	latitude float64
	longitude float64
	height float64
}

func (l *Location) Latitude() float64 {
	return l.latitude
}

func (l *Location) SetLatitude(lat float64) {
	l.latitude = lat
}

func (l *Location) Longitude() float64 {
	return l.longitude
}

func (l *Location) SetLongitude(long float64) {
	l.longitude = long
}