package units

type WindSpeed struct {
	metersPerSecond float64
}

func NewWindSpeedFromMetersPerSecond(mps float64) WindSpeed {
	return WindSpeed{metersPerSecond: mps}
}

func NewWindSpeedFromKilometersPerHour(kph float64) WindSpeed {
	return WindSpeed{metersPerSecond: kph / 3.6}
}

func NewWindSpeedFromMilesPerHour(mph float64) WindSpeed {
	return WindSpeed{metersPerSecond: mph * 0.44704}
}

func (w WindSpeed) ToMetersPerSecond() float64 {
	return w.metersPerSecond
}

func (w WindSpeed) ToKilometersPerHour() float64 {
	return w.metersPerSecond * 3.6
}

func (w WindSpeed) ToMilesPerHour() float64 {
	return w.metersPerSecond / 0.44704
}
