package messages

import "github.com/SplitStackServer/splitstack/api/go/v4/common"

//go:generate msgp
//msgp:tuple GeoLocation

type GeoLocation struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
	Alt float32 `json:"alt"`
}

func (m *GeoLocation) IntoProto() *common.GeoLocation {
	var message common.GeoLocation

	if m != nil {
		message = common.GeoLocation{
			Lat: float64(m.Lat),
			Lon: float64(m.Lon),
			Alt: float64(m.Alt),
		}
	}

	return &message
}
