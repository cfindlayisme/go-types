package units

type Coordinates struct {
	latitude  float64
	longitude float64
}

func NewCoordinateFromLatLong(latitude, longitude float64) Coordinates {
	return Coordinates{latitude: latitude, longitude: longitude}
}
