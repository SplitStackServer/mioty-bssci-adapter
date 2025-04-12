package messages

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// for testing / monkey patching
var getNow = time.Now

func TimestampNsToProto(ts int64) *timestamppb.Timestamp {
	pb := timestamppb.Timestamp{
		Seconds: int64(ts / 1_000_000_000),
		Nanos:   int32(ts % 1_000_000_000),
	}
	return &pb
}

func DurationNsToProto(ts int64) *durationpb.Duration {
	pb := durationpb.Duration{
		Seconds: int64(ts / 1_000_000_000),
		Nanos:   int32(ts % 1_000_000_000),
	}
	return &pb
}

