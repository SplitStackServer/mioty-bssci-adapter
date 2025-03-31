package common

import (
	"reflect"
	"testing"
)

func TestEui64FromInt(t *testing.T) {
	type args struct {
		in int64
	}
	tests := []struct {
		name string
		args args
		want EUI64
	}{
		{
			name: "zero",
			args: args{0},
			want: EUI64{},
		},
		{
			name: "one",
			args: args{1},
			want: EUI64{1, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "max",
			args: args{0x7FFFFFFFFFFFFFFF},
			want: EUI64{255, 255, 255, 255, 255, 255, 255, 127},
		},
		{
			name: "min",
			args: args{-1},
			want: EUI64{255, 255, 255, 255, 255, 255, 255, 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eui64FromInt(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eui64FromInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEui64FromUnsignedInt(t *testing.T) {
	type args struct {
		in uint64
	}
	tests := []struct {
		name string
		args args
		want EUI64
	}{
		{
			name: "zero",
			args: args{0},
			want: EUI64{},
		},
		{
			name: "one",
			args: args{1},
			want: EUI64{1, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "max",
			args: args{0xFFFFFFFFFFFFFFFF},
			want: EUI64{255, 255, 255, 255, 255, 255, 255, 255},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eui64FromUnsignedInt(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eui64FromUnsignedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEUI64_ToInt(t *testing.T) {
	tests := []struct {
		name string
		e    *EUI64
		want int64
	}{
		{
			name: "zero",
			e:    &EUI64{},
			want: 0,
		},
		{
			name: "one",
			e:    &EUI64{1, 0, 0, 0, 0, 0, 0, 0},
			want: 1,
		},
		{
			name: "max",
			e:    &EUI64{255, 255, 255, 255, 255, 255, 255, 127},
			want: 0x7FFFFFFFFFFFFFFF,
		},
		{
			name: "min",
			e:    &EUI64{255, 255, 255, 255, 255, 255, 255, 255},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToInt(); got != tt.want {
				t.Errorf("EUI64.ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEUI64_ToUnsignedInt(t *testing.T) {
	tests := []struct {
		name string
		e    *EUI64
		want uint64
	}{
		{
			name: "zero",
			e:    &EUI64{},
			want: 0,
		},
		{
			name: "one",
			e:    &EUI64{1, 0, 0, 0, 0, 0, 0, 0},
			want: 1,
		},
		{
			name: "max",
			e:    &EUI64{255, 255, 255, 255, 255, 255, 255, 255},
			want: 0xFFFFFFFFFFFFFFFF,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ToUnsignedInt(); got != tt.want {
				t.Errorf("EUI64.ToUnsignedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEUI64_String(t *testing.T) {
	tests := []struct {
		name string
		e    EUI64
		want string
	}{
		{
			name: "zero",
			e:    EUI64{},
			want: "0000000000000000",
		},
		{
			name: "one",
			e:    EUI64{1, 0, 0, 0, 0, 0, 0, 0},
			want: "0100000000000000",
		},
		{
			name: "max",
			e:    EUI64{255, 255, 255, 255, 255, 255, 255, 255},
			want: "ffffffffffffffff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("EUI64.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEUI64_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		e       EUI64
		want    []byte
		wantErr bool
	}{
		{
			name:    "zero",
			e:       EUI64{},
			want:    []byte("0000000000000000"),
			wantErr: false,
		},
		{
			name:    "one",
			e:       EUI64{1, 0, 0, 0, 0, 0, 0, 0},
			want:    []byte("0100000000000000"),
			wantErr: false,
		},
		{
			name:    "max",
			e:       EUI64{255, 255, 255, 255, 255, 255, 255, 255},
			want:    []byte("ffffffffffffffff"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("EUI64.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EUI64.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEUI64_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		args    args
		want    EUI64
		wantErr bool
	}{
		{
			name:    "zero",
			args:    args{[]byte("0000000000000000")},
			want:    EUI64{},
			wantErr: false,
		},
		{
			name:    "one",
			want:    EUI64{1, 0, 0, 0, 0, 0, 0, 0},
			args:    args{[]byte("0100000000000000")},
			wantErr: false,
		},
		{
			name:    "max",
			args:    args{[]byte("ffffffffffffffff")},
			want:    EUI64{255, 255, 255, 255, 255, 255, 255, 255},
			wantErr: false,
		},
		{
			name:    "prefix",
			args:    args{[]byte("0xffffffffffffffff")},
			want:    EUI64{255, 255, 255, 255, 255, 255, 255, 255},
			wantErr: false,
		},
		{
			name:    "too_short",
			args:    args{[]byte("01020304050607")},
			want:    EUI64{},
			wantErr: true,
		},
		{
			name:    "too_long",
			args:    args{[]byte("010203040506070809")},
			want:    EUI64{},
			wantErr: true,
		},
		{
			name:    "invalid",
			args:    args{[]byte("invalid")},
			want:    EUI64{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EUI64{}

			err := e.UnmarshalText(tt.args.text)

			if (err != nil) != tt.wantErr {
				t.Errorf("EUI64.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}

			if e != tt.want {
				t.Errorf("EUI64.UnmarshalText() = %v, want %v", e, tt.want)
			}

		})
	}
}

func TestEUI64_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		e       EUI64
		want    []byte
		wantErr bool
	}{
		{
			name:    "zero",
			e:       EUI64{},
			want:    []byte{0, 0, 0, 0, 0, 0, 0, 0},
			wantErr: false,
		},
		{
			name:    "one",
			e:       EUI64{1, 0, 0, 0, 0, 0, 0, 0},
			want:    []byte{0, 0, 0, 0, 0, 0, 0, 1},
			wantErr: false,
		},
		{
			name:    "max",
			e:       EUI64{255, 255, 255, 255, 255, 255, 255, 255},
			want:    []byte{255, 255, 255, 255, 255, 255, 255, 255},
			wantErr: false,
		},
		{
			name:    "little_endian",
			e:       EUI64{8, 7, 6, 5, 4, 3, 2, 1},
			want:    []byte{1, 2, 3, 4, 5, 6, 7, 8},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.MarshalBinary()
			if (err != nil) != tt.wantErr {
				t.Errorf("EUI64.MarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EUI64.MarshalBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEUI64_UnmarshalBinary(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		want    EUI64
		args    args
		wantErr bool
	}{
		{
			name:    "zero",
			args:    args{[]byte{0, 0, 0, 0, 0, 0, 0, 0}},
			want:    EUI64{},
			wantErr: false,
		},
		{
			name:    "one",
			want:    EUI64{1, 0, 0, 0, 0, 0, 0, 0},
			args:    args{[]byte{0, 0, 0, 0, 0, 0, 0, 1}},
			wantErr: false,
		},
		{
			name:    "max",
			args:    args{[]byte{255, 255, 255, 255, 255, 255, 255, 255}},
			want:    EUI64{255, 255, 255, 255, 255, 255, 255, 255},
			wantErr: false,
		},
		{
			name:    "little_endian",
			args:    args{[]byte{1, 2, 3, 4, 5, 6, 7, 8}},
			want:    EUI64{8, 7, 6, 5, 4, 3, 2, 1},
			wantErr: false,
		},
		{
			name:    "too_short",
			args:    args{[]byte{1, 2, 3, 4, 5, 6, 7}},
			want:    EUI64{},
			wantErr: true,
		},
		{
			name:    "too_long",
			args:    args{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want:    EUI64{},
			wantErr: true,
		},
		{
			name:    "invalid",
			args:    args{[]byte("invalid")},
			want:    EUI64{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			e := EUI64{}

			err := e.UnmarshalBinary(tt.args.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("EUI64.UnmarshalBinary() error = %v, wantErr %v", err, tt.wantErr)
			}

			if e != tt.want {
				t.Errorf("EUI64.UnmarshalBinary() = %v, want %v", e, tt.want)
			}
		})
	}
}

func TestEui64toInt(t *testing.T) {
	type args struct {
		e EUI64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "zero",
			args: args{EUI64{}},
			want: 0,
		},
		{
			name: "one",
			args: args{EUI64{1, 0, 0, 0, 0, 0, 0, 0}},
			want: 1,
		},
		{
			name: "max",
			args: args{EUI64{255, 255, 255, 255, 255, 255, 255, 127}},
			want: 0x7FFFFFFFFFFFFFFF,
		},
		{
			name: "min",
			args: args{EUI64{255, 255, 255, 255, 255, 255, 255, 255}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eui64toInt(tt.args.e); got != tt.want {
				t.Errorf("Eui64toInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEui64toUnsignedInt(t *testing.T) {
	type args struct {
		e EUI64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "zero",
			args: args{EUI64{}},
			want: 0,
		},
		{
			name: "one",
			args: args{EUI64{1, 0, 0, 0, 0, 0, 0, 0}},
			want: 1,
		},
		{
			name: "max",
			args: args{EUI64{255, 255, 255, 255, 255, 255, 255, 255}},
			want: 0xFFFFFFFFFFFFFFFF,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eui64toUnsignedInt(tt.args.e); got != tt.want {
				t.Errorf("Eui64toUnsignedInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
