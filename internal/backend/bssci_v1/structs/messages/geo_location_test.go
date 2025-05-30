package messages

import (
	"reflect"
	"testing"

	"github.com/SplitStackServer/splitstack/api/go/v4/common"
)

func TestGeoLocation_IntoProto(t *testing.T) {
	type fields struct {
		Lat float32
		Lon float32
		Alt float32
	}
	tests := []struct {
		name   string
		fields fields
		want   *common.GeoLocation
	}{
		{
			name: "geolocation1",
			fields: fields{
				Lat: 1.0,
				Lon: 2.0,
				Alt: 3.0,
			},
			want: &common.GeoLocation{
				Lat: 1.0,
				Lon: 2.0,
				Alt: 3.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &GeoLocation{
				Lat: tt.fields.Lat,
				Lon: tt.fields.Lon,
				Alt: tt.fields.Alt,
			}
			if got := m.IntoProto(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeoLocation.IntoProto() = %v, want %v", got, tt.want)
			}
		})
	}
}
