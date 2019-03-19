package space

type Planet string

const secondsPerEarthYear = 31557600

var EarthYearPerYear = map[Planet]float64{
	Planet("Mercury"): 0.2408467,
	Planet("Venus"):   0.61519726,
	Planet("Earth"):   1,
	Planet("Mars"):    1.8808158,
	Planet("Jupiter"): 11.862615,
	Planet("Saturn"):  29.447498,
	Planet("Uranus"):  84.016846,
	Planet("Neptune"): 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	years := seconds / secondsPerEarthYear / EarthYearPerYear[planet]
	return float64(round(years*100)) / 100
}

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}
