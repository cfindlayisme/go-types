package units

type Length struct {
	meters float64
}

func NewLengthFromMeters(m float64) Length {
	return Length{meters: m}
}

func NewLengthFromFeet(f float64) Length {
	return Length{meters: f * 0.3048}
}

func (l Length) ToMeters() float64 {
	return l.meters
}

func (l Length) ToFeet() float64 {
	return l.meters / 0.3048
}

func (l Length) Add(other Length) Length {
	return Length{meters: l.meters + other.meters}
}

func (l Length) Subtract(other Length) Length {
	return Length{meters: l.meters - other.meters}
}
