package units

type Distance struct {
	meters float64
}

func NewDistanceFromMeters(m float64) Distance {
	return Distance{meters: m}
}

func NewDistanceFromFeet(f float64) Distance {
	return Distance{meters: f * 0.3048}
}

func (l Distance) ToMeters() float64 {
	return l.meters
}

func (l Distance) ToFeet() float64 {
	return l.meters / 0.3048
}

func (l Distance) Add(other Distance) Distance {
	return Distance{meters: l.meters + other.meters}
}

func (l Distance) Subtract(other Distance) Distance {
	return Distance{meters: l.meters - other.meters}
}
