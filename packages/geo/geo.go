package geo

import (
	"errors"
)

type Landmark struct {
	name string
	Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) error {
	if name == "" {
		return errors.New("invalid name")
	}
	l.name = name
	return nil
}

type Coordinates struct {
	longitude int
	latitude  int
}

func (c *Coordinates) SetLatitude(latitude int) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}
func (c *Coordinates) SetLongitude(longitude int) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("invalid longitude")
	}
	c.longitude = longitude
	return nil
}

func (c *Coordinates) Latitude() int {
	return c.latitude
}

func (c *Coordinates) Longitude() int {
	return c.longitude
}
