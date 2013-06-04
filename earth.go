// Provides location information specific to Earch.
package geo

type Earth struct{}

func (e *Earth) EarthRadiusSemimajor() float64 {
	return 6378137.0
}
