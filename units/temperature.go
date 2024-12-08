package units

type Temperature struct {
	celsius float64
}

func NewTemperatureFromCelsius(c float64) Temperature {
	return Temperature{celsius: c}
}

func NewTemperatureFromFahrenheit(f float64) Temperature {
	return Temperature{celsius: (f - 32) * 5 / 9}
}

func NewTemperatureFromKelvin(k float64) Temperature {
	return Temperature{celsius: k - 273.15}
}

func (t Temperature) ToCelsius() float64 {
	return t.celsius
}

func (t Temperature) ToFahrenheit() float64 {
	return t.celsius*9/5 + 32
}

func (t Temperature) ToKelvin() float64 {
	return t.celsius + 273.15
}
