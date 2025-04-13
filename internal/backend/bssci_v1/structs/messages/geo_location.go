package messages

import "github.com/SplitStackServer/splitstack/api/go/v4/bs"

//go:generate msgp
//msgp:tuple GeoLocation

type GeoLocation struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	Alt float32 `json:"alt"`
}

func (m *GeoLocation) IntoProto() *bs.GeoLocation {
	var message bs.GeoLocation

	if m != nil {
		message = bs.GeoLocation{
			Lat: m.Lat,
			Lon: m.Lon,
			Alt: m.Alt,
		}
	}

	return &message
}
