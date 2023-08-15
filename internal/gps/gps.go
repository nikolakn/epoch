package gps

import (
	"fmt"
	"math"
)

const (
	mileInKilometres        = 1.60934         // kilometres -> miles conversion.
	degToRad                = math.Pi / 180.0 // degrees -> radians conversion.
	earthRadiusInKilometres = 6371
)

type Degrees float64

func (self Degrees) Radians() Radians {
	return Radians(self * degToRad)
}

type Radians float64

func (self Radians) Degrees() Degrees {
	return Degrees(self / degToRad)
}

type GPS struct {
	Latitude  Degrees `json:"latitude"`
	Longitude Degrees `json:"longitude"`
}

func NewGPS(latitude, longitude Degrees) GPS {
	return GPS{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func HaversineDistance(one, two GPS) float64 {
	lat1 := one.Latitude.Radians()
	lng1 := one.Longitude.Radians()
	lat2 := two.Latitude.Radians()
	lng2 := two.Longitude.Radians()

	deltaLng := float64(lng2 - lng1)
	deltaLat := float64(lat2 - lat1)
	a := math.Pow((math.Sin(deltaLat/2)), 2.0) + math.Cos(float64(lat1))*
		math.Cos(float64(lat2))*math.Pow(math.Sin(deltaLng/2), 2.0)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0-a))

	return earthRadiusInKilometres * c
}

func (g GPS) String() string {
	t := fmt.Sprintf("\tgps: %f, %f", g.Latitude, g.Longitude)
	return t
}
