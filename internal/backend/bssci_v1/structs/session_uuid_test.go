package structs

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewSessionUuid(t *testing.T) {
	type args struct {
		uuid uuid.UUID
	}
	tests := []struct {
		name  string
		args  args
		wantS SessionUuid
	}{
		{
			name:  "sessionUuid_Max",
			args:  args{uuid: uuid.Max},
			wantS: SessionUuid{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		},
		{
			name:  "sessionUuid_Zero",
			args:  args{uuid: uuid.UUID{}},
			wantS: SessionUuid{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := NewSessionUuid(tt.args.uuid); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("NewSessionUuid() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestSessionUuid_ToUuid(t *testing.T) {
	tests := []struct {
		name  string
		s     SessionUuid
		wantU uuid.UUID
	}{
		{
			name:  "sessionUuid_Max",
			s:     SessionUuid{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			wantU: uuid.Max,
		},
		{
			name:  "sessionUuid_Zero",
			s:     SessionUuid{},
			wantU: uuid.UUID{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotU := tt.s.ToUuid(); !reflect.DeepEqual(gotU, tt.wantU) {
				t.Errorf("SessionUuid.ToUuid() = %v, want %v", gotU, tt.wantU)
			}
		})
	}
}

func TestSessionUuid_String(t *testing.T) {
	tests := []struct {
		name string
		s    SessionUuid
		want string
	}{
		{
			name: "sessionUuid_Max",
			s:    SessionUuid{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			want: uuid.Max.String(),
		},
		{
			name: "sessionUuid_Zero",
			s:    SessionUuid{},
			want: uuid.UUID{}.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("SessionUuid.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionUuid_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		s       SessionUuid
		want    []byte
		wantErr bool
	}{
		{
			name:    "sessionUuid_Max",
			s:       SessionUuid{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			want:    []byte(uuid.Max.String()),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("SessionUuid.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SessionUuid.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionUuid_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *SessionUuid
		wantErr bool
	}{
		{
			name:    "sessionUuid_Max",
			args:    args{text: []byte(uuid.Max.String())},
			want:    &SessionUuid{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			wantErr: false,
		},
		{
			name:    "sessionUuid_Error",
			args:    args{text: []byte("test")},
			want:    &SessionUuid{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SessionUuid

			err := s.UnmarshalText(tt.args.text)

			if (err != nil) != tt.wantErr {
				t.Errorf("SessionUuid.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(&s, tt.want) {
				t.Errorf("SessionUuid.UnmarshalText() = %v, want %v", s, tt.want)
			}
		})
	}
}
