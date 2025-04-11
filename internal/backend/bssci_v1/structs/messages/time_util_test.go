package messages

import (
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTimestampNsToProto(t *testing.T) {
	type args struct {
		ts int64
	}
	tests := []struct {
		name string
		args args
		want *timestamppb.Timestamp
	}{
		{
			name: "test1",
			args: args{
				ts: 1744380988827000000,
			},
			want: &timestamppb.Timestamp{
				Seconds: 1744380988,
				Nanos:   827000000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimestampNsToProto(tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimestampNsToProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDurationNsToProto(t *testing.T) {
	type args struct {
		ts int64
	}
	tests := []struct {
		name string
		args args
		want *durationpb.Duration
	}{
		{
			name: "test1",
			args: args{
				ts: 1744380988827000000,
			},
			want: &durationpb.Duration{
				Seconds: 1744380988,
				Nanos:   827000000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DurationNsToProto(tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DurationNsToProto() = %v, want %v", got, tt.want)
			}
		})
	}
}
