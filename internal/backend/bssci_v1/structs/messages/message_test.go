package messages

import (
	"encoding/json"
	"testing"

	"github.com/SplitStackServer/mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"github.com/SplitStackServer/mioty-bssci-adapter/internal/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestMessageData struct {
	name    string
	msgType MessageMsgp
	raw     []byte
	msg     MessageMsgp
	wantErr bool
	json    string
}

type TestMessageSuite struct {
	suite.Suite

	data []TestMessageData
}

func TestMessage(t *testing.T) {
	suite.Run(t, new(TestMessageSuite))
}

func (ts *TestMessageSuite) SetupSuite() {
	testVendor := "Test Vendor"
	testModel := "Test Model"
	testVersion := "1.0.0"

	testBsName := "M0007327767F3"
	testScName := "Test Name"
	testProfile := "eu"

	// equivalent to "c372c521-a778-499b-8b4e-29c783b735dd"
	testScSessionUuid := structs.SessionUuid{-61, 114, -59, 33, -89, 120, 73, -101, -117, 78, 41, -57, -125, -73, 53, -35}
	// equivalent to "f8d69e8a-a9dd-46d4-b975-11d654114a1f"
	testBsSessionUuid := structs.SessionUuid{-8, -42, -98, -118, -87, -35, 70, -44, -71, 117, 17, -42, 84, 17, 74, 31}

	testBsEui := common.EUI64{0xF3, 0x67, 0x77, 0x00, 0x00, 0x32, 0x07, 0x00}
	testScEui := common.EUI64{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
	testEpEui := common.EUI64{0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01}

	testEpSessionKey := [16]byte{3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0}

	testSubpackets := Subpackets{
		SNR:       []int32{1, 2, 3},
		RSSI:      []int32{4, 5, 6},
		Frequency: []int32{7, 8, 9},
		Phase:     &[]int32{10, 11, 12},
	}

	var testEpShAddr uint16 = 0xFFFF

	var testStatusUptime uint64 = 1000
	var testStatusTemp float64 = 45.5
	var testStatusCpu float64 = 0.5
	var testStatusMemory float64 = 0.6

	var testRxDuration uint64 = 500
	var testEqSnr float64 = 4.0

	testTrue := true
	testFalse := false
	var testByte byte = 10
	var testUint32 uint32 = 1234
	var testUint64 uint64 = 12345678
	var testFloat32 float32 = 5678.9

	ts.data = []TestMessageData{
		{
			name:    "msgAtt",
			msgType: &Att{},
			raw: []byte{0xde, 0x0, 0x12, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xa3, 0x61, 0x74, 0x74, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa5, 0x65, 0x70, 0x45, 0x75, 0x69, 0xcf, 0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1, 0xa6, 0x72, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x1, 0xaa, 0x72, 0x78, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xcd, 0x1, 0xf4, 0xa9, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x43, 0x6e, 0x74, 0x2, 0xa3, 0x73, 0x6e, 0x72, 0xcb, 0x40, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa4, 0x72, 0x73, 0x73, 0x69, 0xcb, 0xc0, 0x59, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa5, 0x65, 0x71, 0x53, 0x6e, 0x72, 0xcb, 0x40, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa7, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0xa2, 0x65, 0x75, 0xaa, 0x73, 0x75, 0x62, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x84, 0xa3, 0x73, 0x6e, 0x72, 0x93, 0x1, 0x2, 0x3, 0xa4, 0x72, 0x73, 0x73, 0x69, 0x93, 0x4, 0x5, 0x6, 0xa9, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x93, 0x7, 0x8, 0x9, 0xa5, 0x70, 0x68, 0x61, 0x73, 0x65,
				0x93, 0xa, 0xb, 0xc, 0xa5, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0xc4, 0x4, 0x4, 0x5, 0x6, 0x7, 0xa4, 0x73, 0x69, 0x67, 0x6e, 0xc4, 0x4, 0x1, 0x2, 0x3, 0x4, 0xa6, 0x73, 0x68, 0x41, 0x64, 0x64, 0x72, 0xcd, 0xff, 0xff, 0xa8, 0x64, 0x75, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0xc2, 0xaa, 0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0xc2, 0xab, 0x77, 0x69, 0x64, 0x65, 0x43, 0x61, 0x72, 0x72, 0x4f, 0x66, 0x66, 0xc2, 0xab, 0x6c, 0x6f, 0x6e, 0x67, 0x42, 0x6c, 0x6b, 0x44, 0x69, 0x73, 0x74, 0xc2},
			msg: &Att{
				Command:     structs.MsgAtt,
				OpId:        0,
				EpEui:       testEpEui,
				RxTime:      1,
				RxDuration:  &testRxDuration,
				AttachCnt:   2,
				SNR:         3.0,
				RSSI:        -100.0,
				EqSnr:       &testEqSnr,
				Profile:     &testProfile,
				Subpackets:  &testSubpackets,
				Sign:        [4]byte{1, 2, 3, 4},
				Nonce:       [4]byte{4, 5, 6, 7},
				ShAddr:      &testEpShAddr,
				DualChan:    false,
				Repetition:  false,
				WideCarrOff: false,
				LongBlkDist: false,
			},
			wantErr: false,
			json: `{
	"command": "att",
	"opId": 0,
	"epEui": "0807060504030201",
	"rxTime": 1,
	"rxDuration": 500,
	"attachCnt": 2,
	"snr": 3,
	"rssi": -100,
	"eqSnr": 4,
	"profile": "eu",
	"subpackets": {
		"snr": [
			1,
			2,
			3
		],
		"rssi": [
			4,
			5,
			6
		],
		"frequency": [
			7,
			8,
			9
		],
		"phase": [
			10,
			11,
			12
		]
	},
	"nonce": [
		4,
		5,
		6,
		7
	],
	"sign": [
		1,
		2,
		3,
		4
	],
	"shAddr": 65535,
	"dualChan": false,
	"repetition": false,
	"wideCarrOff": false,
	"longBlkDist": false
}`,
		},
		{
			name:    "msgAttRsp",
			msgType: &AttRsp{},
			raw:     []byte{132, 167, 99, 111, 109, 109, 97, 110, 100, 166, 97, 116, 116, 82, 115, 112, 164, 111, 112, 73, 100, 0, 173, 110, 119, 107, 83, 101, 115, 115, 105, 111, 110, 75, 101, 121, 196, 16, 3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0, 3, 2, 1, 0, 166, 115, 104, 65, 100, 100, 114, 205, 255, 255},
			msg: &AttRsp{
				Command:       structs.MsgAttRsp,
				OpId:          0,
				NwkSessionKey: testEpSessionKey,
				ShAddr:        &testEpShAddr,
			},
			wantErr: false,
			json: `{
	"command": "attRsp",
	"opId": 0,
	"nwkSessionKey": [
		3,
		2,
		1,
		0,
		3,
		2,
		1,
		0,
		3,
		2,
		1,
		0,
		3,
		2,
		1,
		0
	],
	"shAddr": 65535
}`,
		},
		{
			name:    "msgAttCmp",
			msgType: &AttCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 166, 97, 116, 116, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &AttCmp{
				Command: structs.MsgAttCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "attCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgAttPrp",
			msgType: &AttPrp{},
			raw:     []byte{0x8b, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xa6, 0x61, 0x74, 0x74, 0x50, 0x72, 0x70, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa5, 0x65, 0x70, 0x45, 0x75, 0x69, 0xcf, 0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1, 0xa4, 0x62, 0x69, 0x64, 0x69, 0xc2, 0xad, 0x6e, 0x77, 0x6b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0xc4, 0x10, 0x3, 0x2, 0x1, 0x0, 0x3, 0x2, 0x1, 0x0, 0x3, 0x2, 0x1, 0x0, 0x3, 0x2, 0x1, 0x0, 0xa6, 0x73, 0x68, 0x41, 0x64, 0x64, 0x72, 0xcd, 0xff, 0xff, 0xaf, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x0, 0xa8, 0x64, 0x75, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0xc2, 0xaa, 0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0xc2, 0xab, 0x77, 0x69, 0x64, 0x65, 0x43, 0x61, 0x72, 0x72, 0x4f, 0x66, 0x66, 0xc2, 0xab, 0x6c, 0x6f, 0x6e, 0x67, 0x42, 0x6c, 0x6b, 0x44, 0x69, 0x73, 0x74, 0xc2},
			msg: &AttPrp{
				Command:         structs.MsgAttPrp,
				OpId:            0,
				EpEui:           testEpEui,
				Bidi:            false,
				NwkSessionKey:   testEpSessionKey,
				ShAddr:          testEpShAddr,
				LastPacketCount: 0,
				DualChan:        false,
				Repetition:      false,
				WideCarrOff:     false,
				LongBlkDist:     false,
			},
			wantErr: false,
			json: `{
	"command": "attPrp",
	"opId": 0,
	"epEui": "0807060504030201",
	"bidi": false,
	"nwkSessionKey": [
		3,
		2,
		1,
		0,
		3,
		2,
		1,
		0,
		3,
		2,
		1,
		0,
		3,
		2,
		1,
		0
	],
	"shAddr": 65535,
	"lastPacketCount": 0,
	"dualChan": false,
	"repetition": false,
	"wideCarrOff": false,
	"longBlkDist": false
}`,
		},
		{
			name:    "msgAttPrpRsp",
			msgType: &AttPrpRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 97, 116, 116, 80, 114, 112, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &AttPrpRsp{
				Command: structs.MsgAttPrpRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "attPrpRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgAttPrpCmp",
			msgType: &AttPrpCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 97, 116, 116, 80, 114, 112, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &AttPrpCmp{
				Command: structs.MsgAttPrpCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "attPrpCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgCon",
			msgType: &Con{},
			raw:     []byte{0x89, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xa3, 0x63, 0x6f, 0x6e, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa7, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xa5, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0xa5, 0x62, 0x73, 0x45, 0x75, 0x69, 0xcf, 0xf3, 0x67, 0x77, 0x0, 0x0, 0x32, 0x7, 0x0, 0xa6, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0xab, 0x54, 0x65, 0x73, 0x74, 0x20, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0xa5, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xaa, 0x54, 0x65, 0x73, 0x74, 0x20, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xa4, 0x6e, 0x61, 0x6d, 0x65, 0xad, 0x4d, 0x30, 0x30, 0x30, 0x37, 0x33, 0x32, 0x37, 0x37, 0x36, 0x37, 0x46, 0x33, 0xa4, 0x62, 0x69, 0x64, 0x69, 0xc3, 0xa8, 0x73, 0x6e, 0x42, 0x73, 0x55, 0x75, 0x69, 0x64, 0xdc, 0x0, 0x10, 0xf8, 0xd0, 0xd6, 0xd0, 0x9e, 0xd0, 0x8a, 0xd0, 0xa9, 0xd0, 0xdd, 0x46, 0xd0, 0xd4, 0xd0, 0xb9, 0x75, 0x11, 0xd0, 0xd6, 0x54, 0x11, 0x4a, 0x1f},
			msg: &Con{
				Command:  structs.MsgCon,
				OpId:     0,
				Version:  testVersion,
				BsEui:    testBsEui,
				Vendor:   &testVendor,
				Model:    &testModel,
				Name:     &testBsName,
				SnBsUuid: testBsSessionUuid,
				Bidi:     true,
			},
			wantErr: false,
			json: `{
	"command": "con",
	"opId": 0,
	"version": "1.0.0",
	"bsEui": "f367770000320700",
	"vendor": "Test Vendor",
	"model": "Test Model",
	"name": "M0007327767F3",
	"bidi": true,
	"snBsUuid": "f8d69e8a-a9dd-46d4-b975-11d654114a1f"
}`,
		},
		{
			name:    "msgConRsp",
			msgType: &ConRsp{},
			raw:     []byte{0x89, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xa6, 0x63, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa7, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0xa5, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0xa5, 0x73, 0x63, 0x45, 0x75, 0x69, 0xcf, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0xa6, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0xab, 0x54, 0x65, 0x73, 0x74, 0x20, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0xa5, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0xaa, 0x54, 0x65, 0x73, 0x74, 0x20, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xa4, 0x6e, 0x61, 0x6d, 0x65, 0xa9, 0x54, 0x65, 0x73, 0x74, 0x20, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x73, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0xc2, 0xa8, 0x73, 0x6e, 0x53, 0x63, 0x55, 0x75, 0x69, 0x64, 0xdc, 0x0, 0x10, 0xd0, 0xc3, 0x72, 0xd0, 0xc5, 0x21, 0xd0, 0xa7, 0x78, 0x49, 0xd0, 0x9b, 0xd0, 0x8b, 0x4e, 0x29, 0xd0, 0xc7, 0xd0, 0x83, 0xd0, 0xb7, 0x35, 0xd0, 0xdd},
			msg: &ConRsp{
				Command:  structs.MsgConRsp,
				OpId:     0,
				Version:  testVersion,
				ScEui:    testScEui,
				Vendor:   &testVendor,
				Model:    &testModel,
				Name:     &testScName,
				SnResume: false,
				SnScUuid: testScSessionUuid,
			},
			wantErr: false,
			json: `{
	"command": "conRsp",
	"opId": 0,
	"version": "1.0.0",
	"scEui": "0101010101010101",
	"vendor": "Test Vendor",
	"model": "Test Model",
	"name": "Test Name",
	"snResume": false,
	"snScUuid": "c372c521-a778-499b-8b4e-29c783b735dd"
}`,
		},
		{
			name:    "msgConCmp",
			msgType: &ConCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 166, 99, 111, 110, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &ConCmp{
				Command: structs.MsgConCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "conCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgDet",
			msgType: &Det{},
			raw: []byte{0x8c, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xa3, 0x64, 0x65, 0x74, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa5, 0x65, 0x70, 0x45, 0x75, 0x69, 0xcf, 0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1, 0xa6, 0x72, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x1, 0xaa, 0x72, 0x78, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xcd, 0x1, 0xf4, 0xa9, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6e, 0x74, 0x2, 0xa3, 0x73, 0x6e, 0x72, 0xcb, 0x40, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa4, 0x72, 0x73, 0x73, 0x69, 0xcb, 0xc0, 0x59, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa5, 0x65, 0x71, 0x53, 0x6e, 0x72, 0xcb, 0x40, 0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa7, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0xa2, 0x65, 0x75, 0xaa, 0x73, 0x75, 0x62, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x84, 0xa3, 0x73, 0x6e, 0x72, 0x93, 0x1, 0x2, 0x3, 0xa4, 0x72, 0x73, 0x73, 0x69, 0x93, 0x4, 0x5, 0x6, 0xa9, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x93, 0x7, 0x8, 0x9, 0xa5, 0x70, 0x68, 0x61, 0x73, 0x65, 0x93, 0xa,
				0xb, 0xc, 0xa4, 0x73, 0x69, 0x67, 0x6e, 0xc4, 0x4, 0x1, 0x2, 0x3, 0x4},
			msg: &Det{
				Command:    structs.MsgDet,
				OpId:       0,
				EpEui:      testEpEui,
				RxTime:     1,
				RxDuration: &testRxDuration,
				PacketCnt:  2,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      &testEqSnr,
				Profile:    &testProfile,
				Subpackets: &testSubpackets,
				Sign:       [4]byte{1, 2, 3, 4},
			},
			wantErr: false,
			json: `{
	"command": "det",
	"opId": 0,
	"epEui": "0807060504030201",
	"rxTime": 1,
	"rxDuration": 500,
	"packetCnt": 2,
	"snr": 3,
	"rssi": -100,
	"eqSnr": 4,
	"profile": "eu",
	"subpackets": {
		"snr": [
			1,
			2,
			3
		],
		"rssi": [
			4,
			5,
			6
		],
		"frequency": [
			7,
			8,
			9
		],
		"phase": [
			10,
			11,
			12
		]
	},
	"sign": [
		1,
		2,
		3,
		4
	]
}`,
		},
		{
			name:    "msgDetRsp",
			msgType: &DetRsp{},
			raw:     []byte{131, 167, 99, 111, 109, 109, 97, 110, 100, 166, 100, 101, 116, 82, 115, 112, 164, 111, 112, 73, 100, 0, 164, 115, 105, 103, 110, 196, 4, 1, 2, 3, 4},
			msg: &DetRsp{
				Command: structs.MsgDetRsp,
				OpId:    0,
				Sign:    [4]byte{1, 2, 3, 4},
			},
			wantErr: false,
			json: `{
	"command": "detRsp",
	"opId": 0,
	"sign": [
		1,
		2,
		3,
		4
	]
}`,
		},
		{
			name:    "msgDetCmp",
			msgType: &DetCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 166, 100, 101, 116, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &DetCmp{
				Command: structs.MsgDetCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "detCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgDetPrp",
			msgType: &DetPrp{},
			raw:     []byte{0x83, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xa6, 0x64, 0x65, 0x74, 0x50, 0x72, 0x70, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa5, 0x65, 0x70, 0x45, 0x75, 0x69, 0xcf, 0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1},
			msg: &DetPrp{
				Command: structs.MsgDetPrp,
				OpId:    0,
				EpEui:   testEpEui,
			},
			wantErr: false,
			json: `{
	"command": "detPrp",
	"opId": 0,
	"epEui": "0807060504030201"
}`,
		},
		{
			name:    "msgDetPrpRsp",
			msgType: &DetPrpRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 100, 101, 116, 80, 114, 112, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &DetPrpRsp{
				Command: structs.MsgDetPrpRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "detPrpRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgDetPrpCmp",
			msgType: &DetPrpCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 100, 101, 116, 80, 114, 112, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &DetPrpCmp{
				Command: structs.MsgDetPrpCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "detPrpCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgDlDataQue",
			msgType: &DlDataQue{},
			raw:     []byte{0x8c, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xa9, 0x64, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa5, 0x65, 0x70, 0x45, 0x75, 0x69, 0xcf, 0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1, 0xa5, 0x71, 0x75, 0x65, 0x49, 0x64, 0xce, 0x0, 0xbc, 0x61, 0x4e, 0xa9, 0x63, 0x6e, 0x74, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0xc2, 0xa8, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x91, 0xdc, 0x0, 0x16, 0x74, 0x68, 0x69, 0x73, 0x69, 0x73, 0x61, 0x6c, 0x6f, 0x74, 0x6f, 0x66, 0x64, 0x61, 0x74, 0x61, 0x74, 0x6f, 0x73, 0x65, 0x6e, 0x64, 0xa6, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0xa, 0xa4, 0x70, 0x72, 0x69, 0x6f, 0xca, 0x45, 0xb1, 0x77, 0x33, 0xab, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x78, 0x70, 0xc2, 0xac, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x69, 0x6f, 0xc3, 0xa9, 0x64, 0x6c, 0x57, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0xc2, 0xa7, 0x65, 0x78, 0x70, 0x4f, 0x6e, 0x6c, 0x79, 0xc3},
			msg: &DlDataQue{
				Command:      structs.MsgDlDataQue,
				OpId:         0,
				EpEui:        testEpEui,
				QueId:        12345678,
				CntDepend:    false,
				PacketCnt:    nil,
				UserData:     [][]byte{[]byte("thisisalotofdatatosend")},
				Format:       &testByte,
				Prio:         &testFloat32,
				ResponseExp:  &testFalse,
				ResponsePrio: &testTrue,
				DlWindReq:    &testFalse,
				ExpOnly:      &testTrue,
			},
			wantErr: false,
			json: `{
	"command": "dlDataQue",
	"opId": 0,
	"epEui": "0807060504030201",
	"queId": 12345678,
	"cntDepend": false,
	"userData": [
		"dGhpc2lzYWxvdG9mZGF0YXRvc2VuZA=="
	],
	"format": 10,
	"prio": 5678.9,
	"responseExp": false,
	"responsePrio": true,
	"dlWindReq": false,
	"expOnly": true
}`,
		},
		{
			name:    "msgDlDataQueRsp",
			msgType: &DlDataQueRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 100, 108, 68, 97, 116, 97, 81, 117, 101, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &DlDataQueRsp{
				Command: structs.MsgDlDataQueRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "dlDataQueRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgDlDataQueCmp",
			msgType: &DlDataQueCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 100, 108, 68, 97, 116, 97, 81, 117, 101, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &DlDataQueCmp{
				Command: structs.MsgDlDataQueCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "dlDataQueCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgDlDataRes",
			msgType: &DlDataRes{},
			raw:     []byte{0x87, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xa9, 0x64, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa5, 0x65, 0x70, 0x45, 0x75, 0x69, 0xcf, 0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1, 0xa5, 0x71, 0x75, 0x65, 0x49, 0x64, 0xce, 0x0, 0xbc, 0x61, 0x4e, 0xa6, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0xa4, 0x73, 0x65, 0x6e, 0x74, 0xa6, 0x74, 0x78, 0x54, 0x69, 0x6d, 0x65, 0xce, 0x0, 0xbc, 0x61, 0x4e, 0xa9, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6e, 0x74, 0xcd, 0x4, 0xd2},
			msg: &DlDataRes{
				Command:   structs.MsgDlDataRes,
				OpId:      0,
				EpEui:     testEpEui,
				QueId:     12345678,
				Result:    dlDataResult_Sent,
				TxTime:    &testUint64,
				PacketCnt: &testUint32,
			},
			wantErr: false,
			json: `{
	"command": "dlDataRes",
	"opId": 0,
	"epEui": "0807060504030201",
	"queId": 12345678,
	"result": 1,
	"txTime": 12345678,
	"packetCnt": 1234
}`,
		},
		{
			name:    "msgDlDataResRsp",
			msgType: &DlDataResRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 100, 108, 68, 97, 116, 97, 82, 101, 115, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &DlDataResRsp{
				Command: structs.MsgDlDataResRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "dlDataResRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgDlDataResCmp",
			msgType: &DlDataResCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 100, 108, 68, 97, 116, 97, 82, 101, 115, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &DlDataResCmp{
				Command: structs.MsgDlDataResCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "dlDataResCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgDlDataRev",
			msgType: &DlDataRev{},
			raw:     []byte{0x84, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xa9, 0x64, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x76, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa5, 0x65, 0x70, 0x45, 0x75, 0x69, 0xcf, 0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1, 0xa5, 0x71, 0x75, 0x65, 0x49, 0x64, 0xce, 0x0, 0xbc, 0x61, 0x4e},
			msg: &DlDataRev{
				Command: structs.MsgDlDataRev,
				OpId:    0,
				EpEui:   testEpEui,
				QueId:   12345678,
			},
			wantErr: false,
			json: `{
	"command": "dlDataRev",
	"opId": 0,
	"epEui": "0807060504030201",
	"queId": 12345678
}`,
		},
		{
			name:    "msgDlDataRevRsp",
			msgType: &DlDataRevRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 100, 108, 68, 97, 116, 97, 82, 101, 118, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &DlDataRevRsp{
				Command: structs.MsgDlDataRevRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "dlDataRevRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgDlDataRevCmp",
			msgType: &DlDataRevCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 100, 108, 68, 97, 116, 97, 82, 101, 118, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &DlDataRevCmp{
				Command: structs.MsgDlDataRevCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "dlDataRevCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgDlRxStat",
			msgType: &DlRxStat{},
			raw:     []byte{0x87, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xa8, 0x64, 0x6c, 0x52, 0x78, 0x53, 0x74, 0x61, 0x74, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa5, 0x65, 0x70, 0x45, 0x75, 0x69, 0xcf, 0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1, 0xa6, 0x72, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x1, 0xa9, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6e, 0x74, 0xcd, 0x4, 0xd2, 0xa7, 0x64, 0x6c, 0x52, 0x78, 0x53, 0x6e, 0x72, 0xcb, 0x40, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa8, 0x64, 0x6c, 0x52, 0x78, 0x52, 0x73, 0x73, 0x69, 0xcb, 0xc0, 0x59, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			msg: &DlRxStat{
				Command:   structs.MsgDlRxStat,
				OpId:      0,
				EpEui:     testEpEui,
				RxTime:    1,
				PacketCnt: testUint32,
				DlRxSnr:   2.0,
				DlRxRssi:  -100,
			},
			wantErr: false,
			json: `{
	"command": "dlRxStat",
	"opId": 0,
	"epEui": "0807060504030201",
	"rxTime": 1,
	"packetCnt": 1234,
	"dlRxSnr": 2,
	"dlRxRssi": -100
}`,
		},
		{
			name:    "msgDlRxStatRsp",
			msgType: &DlRxStatRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 171, 100, 108, 82, 120, 83, 116, 97, 116, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &DlRxStatRsp{
				Command: structs.MsgDlRxStatRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "dlRxStatRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgDlRxStatCmp",
			msgType: &DlRxStatCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 171, 100, 108, 82, 120, 83, 116, 97, 116, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &DlRxStatCmp{
				Command: structs.MsgDlRxStatCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "dlRxStatCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgDlRxStatQry",
			msgType: &DlRxStatQry{},
			raw:     []byte{0x83, 0xa7, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xab, 0x64, 0x6c, 0x52, 0x78, 0x53, 0x74, 0x61, 0x74, 0x51, 0x72, 0x79, 0xa4, 0x6f, 0x70, 0x49, 0x64, 0x0, 0xa5, 0x65, 0x70, 0x45, 0x75, 0x69, 0xcf, 0x8, 0x7, 0x6, 0x5, 0x4, 0x3, 0x2, 0x1},
			msg: &DlRxStatQry{
				Command: structs.MsgDlRxStatQry,
				OpId:    0,
				EpEui:   testEpEui,
			},
			wantErr: false,
			json: `{
	"command": "dlRxStatQry",
	"opId": 0,
	"epEui": "0807060504030201"
}`,
		},
		{
			name:    "msgDlRxStatQryRsp",
			msgType: &DlRxStatQryRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 174, 100, 108, 82, 120, 83, 116, 97, 116, 81, 114, 121, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &DlRxStatQryRsp{
				Command: structs.MsgDlRxStatQryRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "dlRxStatQryRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgDlRxStatQryCmp",
			msgType: &DlRxStatQryCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 174, 100, 108, 82, 120, 83, 116, 97, 116, 81, 114, 121, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &DlRxStatQryCmp{
				Command: structs.MsgDlRxStatQryCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "dlRxStatQryCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgPing",
			msgType: &Ping{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			msg: &Ping{
				Command: structs.MsgPing,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "ping",
	"opId": 0
}`,
		},
		{
			name:    "msgPingRsp",
			msgType: &PingRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 167, 112, 105, 110, 103, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &PingRsp{
				Command: structs.MsgPingRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "pingRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgPingCmp",
			msgType: &PingCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 167, 112, 105, 110, 103, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &PingCmp{
				Command: structs.MsgPingCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "pingCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgStatus",
			msgType: &Status{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 166, 115, 116, 97, 116, 117, 115, 164, 111, 112, 73, 100, 0},
			msg: &Status{
				Command: structs.MsgStatus,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "status",
	"opId": 0
}`,
		},
		{
			name:    "msgStatusRsp",
			msgType: &StatusRsp{},
			raw:     []byte{138, 167, 99, 111, 109, 109, 97, 110, 100, 169, 115, 116, 97, 116, 117, 115, 82, 115, 112, 164, 111, 112, 73, 100, 0, 164, 99, 111, 100, 101, 0, 167, 109, 101, 115, 115, 97, 103, 101, 162, 111, 107, 164, 116, 105, 109, 101, 206, 59, 154, 202, 5, 169, 100, 117, 116, 121, 67, 121, 99, 108, 101, 202, 62, 204, 204, 205, 166, 117, 112, 116, 105, 109, 101, 205, 3, 232, 164, 116, 101, 109, 112, 203, 64, 70, 192, 0, 0, 0, 0, 0, 167, 99, 112, 117, 76, 111, 97, 100, 203, 63, 224, 0, 0, 0, 0, 0, 0, 167, 109, 101, 109, 76, 111, 97, 100, 203, 63, 227, 51, 51, 51, 51, 51, 51},
			msg: &StatusRsp{
				Command:     structs.MsgStatusRsp,
				OpId:        0,
				Code:        0,
				Message:     "ok",
				Time:        1000000005,
				DutyCycle:   0.4,
				GeoLocation: nil,
				Uptime:      &testStatusUptime,
				Temp:        &testStatusTemp,
				CpuLoad:     &testStatusCpu,
				MemLoad:     &testStatusMemory,
			},
			wantErr: false,
			json: `{
	"command": "statusRsp",
	"opId": 0,
	"code": 0,
	"message": "ok",
	"time": 1000000005,
	"dutyCycle": 0.4,
	"uptime": 1000,
	"temp": 45.5,
	"cpuLoad": 0.5,
	"memLoad": 0.6
}`,
		}, {
			name:    "msgStatusCmp",
			msgType: &StatusCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 115, 116, 97, 116, 117, 115, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &StatusCmp{
				Command: structs.MsgStatusCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "statusCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgUlData",
			msgType: &UlData{},
			raw:     []byte{143, 167, 99, 111, 109, 109, 97, 110, 100, 166, 117, 108, 68, 97, 116, 97, 164, 111, 112, 73, 100, 0, 165, 101, 112, 69, 117, 105, 207, 8, 7, 6, 5, 4, 3, 2, 1, 166, 114, 120, 84, 105, 109, 101, 1, 170, 114, 120, 68, 117, 114, 97, 116, 105, 111, 110, 205, 1, 244, 169, 112, 97, 99, 107, 101, 116, 67, 110, 116, 2, 163, 115, 110, 114, 203, 64, 8, 0, 0, 0, 0, 0, 0, 164, 114, 115, 115, 105, 203, 192, 89, 0, 0, 0, 0, 0, 0, 165, 101, 113, 83, 110, 114, 203, 64, 16, 0, 0, 0, 0, 0, 0, 167, 112, 114, 111, 102, 105, 108, 101, 162, 101, 117, 170, 115, 117, 98, 112, 97, 99, 107, 101, 116, 115, 132, 163, 115, 110, 114, 147, 1, 2, 3, 164, 114, 115, 115, 105, 147, 4, 5, 6, 169, 102, 114, 101, 113, 117, 101, 110, 99, 121, 147, 7, 8, 9, 165, 112, 104, 97, 115, 101, 147, 10, 11, 12, 168, 117, 115, 101, 114, 68, 97, 116, 97, 220, 0, 22, 116, 104, 105, 115, 105, 115, 97, 108, 111, 116, 111, 102, 100, 97, 116, 97, 116, 111, 115, 101, 110, 100, 166, 100, 108, 79, 112, 101, 110, 194, 171, 114, 101, 115, 112, 111, 110, 115, 101, 69, 120, 112, 194, 165, 100, 108, 65, 99, 107, 194},
			msg: &UlData{
				Command:     structs.MsgUlData,
				OpId:        0,
				EpEui:       testEpEui,
				RxTime:      1,
				RxDuration:  &testRxDuration,
				PacketCnt:   2,
				SNR:         3.0,
				RSSI:        -100.0,
				EqSnr:       &testEqSnr,
				Profile:     &testProfile,
				Mode:        nil,
				Subpackets:  &testSubpackets,
				UserData:    []byte("thisisalotofdatatosend"),
				Format:      nil,
				DlOpen:      false,
				ResponseExp: false,
				DlAck:       false,
			},
			wantErr: false,
			json: `{
	"command": "ulData",
	"opId": 0,
	"epEui": "0807060504030201",
	"rxTime": 1,
	"rxDuration": 500,
	"packetCnt": 2,
	"snr": 3,
	"rssi": -100,
	"eqSnr": 4,
	"profile": "eu",
	"subpackets": {
		"snr": [
			1,
			2,
			3
		],
		"rssi": [
			4,
			5,
			6
		],
		"frequency": [
			7,
			8,
			9
		],
		"phase": [
			10,
			11,
			12
		]
	},
	"userData": "dGhpc2lzYWxvdG9mZGF0YXRvc2VuZA==",
	"dlOpen": false,
	"responseExp": false,
	"dlAck": false
}`,
		},
		{
			name:    "msUlDataRsp",
			msgType: &UlDataRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 117, 108, 68, 97, 116, 97, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &UlDataRsp{
				Command: structs.MsgUlDataRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "ulDataRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgUlDataCmp",
			msgType: &UlDataCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 117, 108, 68, 97, 116, 97, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &UlDataCmp{
				Command: structs.MsgUlDataCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "ulDataCmp",
	"opId": 0
}`,
		},

		{
			name:    "msgError",
			msgType: &BssciError{},
			raw:     []byte{132, 167, 99, 111, 109, 109, 97, 110, 100, 165, 101, 114, 114, 111, 114, 164, 111, 112, 73, 100, 0, 164, 99, 111, 100, 101, 1, 167, 109, 101, 115, 115, 97, 103, 101, 173, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101},
			msg: &BssciError{
				Command: structs.MsgError,
				OpId:    0,
				Code:    1,
				Message: "error message",
			},
			wantErr: false,
			json: `{
	"command": "error",
	"opId": 0,
	"code": 1,
	"message": "error message"
}`,
		},
		{
			name:    "msgErrorAck",
			msgType: &BssciErrorAck{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 168, 101, 114, 114, 111, 114, 65, 99, 107, 164, 111, 112, 73, 100, 0},
			msg: &BssciErrorAck{
				Command: structs.MsgErrorAck,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "errorAck",
	"opId": 0
}`,
		},
		{
			name:    "msgVmActivate",
			msgType: &VmActivate{},
			raw:     []byte{131, 167, 99, 111, 109, 109, 97, 110, 100, 171, 118, 109, 46, 97, 99, 116, 105, 118, 97, 116, 101, 164, 111, 112, 73, 100, 0, 167, 109, 97, 99, 84, 121, 112, 101, 1},
			msg: &VmActivate{
				Command: structs.MsgVmActivate,
				OpId:    0,
				MacType: 1,
			},
			wantErr: false,
			json: `{
	"command": "vm.activate",
	"opId": 0,
	"macType": 1
}`,
		},
		{
			name:    "msgVmActivateRsp",
			msgType: &VmActivateRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 174, 118, 109, 46, 97, 99, 116, 105, 118, 97, 116, 101, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &VmActivateRsp{
				Command: structs.MsgVmActivateRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "vm.activateRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgVmActivateCmp",
			msgType: &VmActivateCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 174, 118, 109, 46, 97, 99, 116, 105, 118, 97, 116, 101, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &VmActivateCmp{
				Command: structs.MsgVmActivateCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "vm.activateCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgVmDeactivate",
			msgType: &VmDeactivate{},
			raw:     []byte{131, 167, 99, 111, 109, 109, 97, 110, 100, 173, 118, 109, 46, 100, 101, 97, 99, 116, 105, 118, 97, 116, 101, 164, 111, 112, 73, 100, 0, 167, 109, 97, 99, 84, 121, 112, 101, 1},
			msg: &VmDeactivate{
				Command: structs.MsgVmDeactivate,
				OpId:    0,
				MacType: 1,
			},
			wantErr: false,
			json: `{
	"command": "vm.deactivate",
	"opId": 0,
	"macType": 1
}`,
		},
		{
			name:    "msgVmDeactivateRsp",
			msgType: &VmDeactivateRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 176, 118, 109, 46, 100, 101, 97, 99, 116, 105, 118, 97, 116, 101, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &VmDeactivateRsp{
				Command: structs.MsgVmDeactivateRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "vm.deactivateRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgVmDeactivateCmp",
			msgType: &VmDeactivateCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 176, 118, 109, 46, 100, 101, 97, 99, 116, 105, 118, 97, 116, 101, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &VmDeactivateCmp{
				Command: structs.MsgVmDeactivateCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "vm.deactivateCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgVmStatus",
			msgType: &VmStatus{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 118, 109, 46, 115, 116, 97, 116, 117, 115, 164, 111, 112, 73, 100, 0},
			msg: &VmStatus{
				Command: structs.MsgVmStatus,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "vm.status",
	"opId": 0
}`,
		},
		{
			name:    "msgVmStatusRsp",
			msgType: &VmStatusRsp{},
			raw:     []byte{131, 167, 99, 111, 109, 109, 97, 110, 100, 172, 118, 109, 46, 115, 116, 97, 116, 117, 115, 82, 115, 112, 164, 111, 112, 73, 100, 0, 168, 109, 97, 99, 84, 121, 112, 101, 115, 147, 1, 2, 3},
			msg: &VmStatusRsp{
				Command:  structs.MsgVmStatusRsp,
				OpId:     0,
				MacTypes: []int64{1, 2, 3},
			},
			wantErr: false,
			json: `{
	"command": "vm.statusRsp",
	"opId": 0,
	"macTypes": [
		1,
		2,
		3
	]
}`,
		},
		{
			name:    "msgVmStatusCmp",
			msgType: &VmStatusCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 118, 109, 46, 115, 116, 97, 116, 117, 115, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &VmStatusCmp{
				Command: structs.MsgVmStatusCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "vm.statusCmp",
	"opId": 0
}`,
		},
		{
			name:    "msgVmUlData",
			msgType: &VmUlData{},
			raw:     []byte{143, 167, 99, 111, 109, 109, 97, 110, 100, 169, 118, 109, 46, 117, 108, 68, 97, 116, 97, 164, 111, 112, 73, 100, 10, 167, 109, 97, 99, 84, 121, 112, 101, 0, 168, 117, 115, 101, 114, 68, 97, 116, 97, 147, 1, 2, 3, 167, 116, 114, 120, 84, 105, 109, 101, 1, 167, 115, 121, 115, 84, 105, 109, 101, 1, 167, 102, 114, 101, 113, 79, 102, 102, 203, 64, 78, 0, 0, 0, 0, 0, 0, 163, 115, 110, 114, 203, 64, 8, 0, 0, 0, 0, 0, 0, 164, 114, 115, 115, 105, 203, 192, 89, 0, 0, 0, 0, 0, 0, 165, 101, 113, 83, 110, 114, 203, 64, 16, 0, 0, 0, 0, 0, 0, 170, 115, 117, 98, 112, 97, 99, 107, 101, 116, 115, 132, 163, 115, 110, 114, 147, 1, 2, 3, 164, 114, 115, 115, 105, 147, 4, 5, 6, 169, 102, 114, 101, 113, 117, 101, 110, 99, 121, 147, 7, 8, 9, 165, 112, 104, 97, 115, 101, 147, 10, 11, 12, 169, 99, 97, 114, 114, 83, 112, 97, 99, 101, 2, 167, 112, 97, 116, 116, 71, 114, 112, 3, 167, 112, 97, 116, 116, 78, 117, 109, 5, 163, 99, 114, 99, 146, 6, 77},
			msg: &VmUlData{
				Command:    structs.MsgVmUlData,
				OpId:       10,
				MacType:    0,
				UserData:   []byte{1, 2, 3},
				TrxTime:    1,
				SysTime:    1,
				FreqOff:    60,
				SNR:        3.0,
				RSSI:       -100.0,
				EqSnr:      &testEqSnr,
				Subpackets: &testSubpackets,
				CarrSpace:  2,
				PattGrp:    3,
				PattNum:    5,
				CRC:        [2]byte{6, 0x4d},
			},
			wantErr: false,
			json: `{
	"command": "vm.ulData",
	"opId": 10,
	"macType": 0,
	"userData": "AQID",
	"trxTime": 1,
	"sysTime": 1,
	"freqOff": 60,
	"snr": 3,
	"rssi": -100,
	"eqSnr": 4,
	"subpackets": {
		"snr": [
			1,
			2,
			3
		],
		"rssi": [
			4,
			5,
			6
		],
		"frequency": [
			7,
			8,
			9
		],
		"phase": [
			10,
			11,
			12
		]
	},
	"carrSpace": 2,
	"pattGrp": 3,
	"pattNum": 5,
	"crc": [
		6,
		77
	]
}`,
		},
		{
			name:    "msgVmUlDataReal",
			msgType: &VmUlData{},
			raw:     []byte{141, 167, 99, 111, 109, 109, 97, 110, 100, 169, 118, 109, 46, 117, 108, 68, 97, 116, 97, 164, 111, 112, 73, 100, 4, 167, 109, 97, 99, 84, 121, 112, 101, 1, 168, 117, 115, 101, 114, 68, 97, 116, 97, 220, 0, 86, 85, 68, 204, 165, 17, 70, 82, 121, 204, 129, 118, 7, 204, 140, 0, 204, 183, 204, 144, 15, 0, 44, 37, 62, 53, 17, 0, 204, 152, 204, 202, 204, 236, 204, 193, 18, 66, 37, 112, 122, 27, 204, 147, 50, 7, 16, 42, 68, 204, 251, 204, 169, 72, 204, 241, 204, 216, 204, 250, 24, 121, 204, 135, 204, 160, 204, 146, 11, 204, 137, 204, 141, 56, 204, 222, 3, 204, 250, 204, 159, 204, 232, 118, 204, 230, 80, 45, 204, 181, 204, 164, 103, 47, 95, 204, 132, 53, 53, 34, 93, 49, 204, 155, 104, 204, 159, 85, 87, 33, 204, 182, 204, 144, 43, 119, 70, 204, 138, 108, 167, 116, 114, 120, 84, 105, 109, 101, 207, 0, 9, 83, 63, 13, 76, 193, 199, 167, 115, 121, 115, 84, 105, 109, 101, 207, 24, 53, 85, 184, 107, 9, 219, 142, 167, 102, 114, 101, 113, 79, 102, 102, 203, 65, 201, 223, 132, 78, 0, 0, 0, 163, 115, 110, 114, 203, 64, 10, 44, 11, 192, 0, 0, 0, 164, 114, 115, 115, 105, 203, 192, 96, 28, 136, 128, 0, 0, 0, 169, 99, 97, 114, 114, 83, 112, 97, 99, 101, 1, 167, 112, 97, 116, 116, 71, 114, 112, 0, 167, 112, 97, 116, 116, 78, 117, 109, 3, 163, 99, 114, 99, 146, 87, 204, 142},
			msg: &VmUlData{
				Command:   structs.MsgVmUlData,
				OpId:      4,
				MacType:   1,
				UserData:  []byte{0x55, 0x44, 0xa5, 0x11, 0x46, 0x52, 0x79, 0x81, 0x76, 0x7, 0x8c, 0x0, 0xb7, 0x90, 0xf, 0x0, 0x2c, 0x25, 0x3e, 0x35, 0x11, 0x0, 0x98, 0xca, 0xec, 0xc1, 0x12, 0x42, 0x25, 0x70, 0x7a, 0x1b, 0x93, 0x32, 0x7, 0x10, 0x2a, 0x44, 0xfb, 0xa9, 0x48, 0xf1, 0xd8, 0xfa, 0x18, 0x79, 0x87, 0xa0, 0x92, 0xb, 0x89, 0x8d, 0x38, 0xde, 0x3, 0xfa, 0x9f, 0xe8, 0x76, 0xe6, 0x50, 0x2d, 0xb5, 0xa4, 0x67, 0x2f, 0x5f, 0x84, 0x35, 0x35, 0x22, 0x5d, 0x31, 0x9b, 0x68, 0x9f, 0x55, 0x57, 0x21, 0xb6, 0x90, 0x2b, 0x77, 0x46, 0x8a, 0x6c},
				TrxTime:   2624805061575111,
				SysTime:   1744394681234086798,
				FreqOff:   868157596,
				SNR:       3.2715067863464355,
				RSSI:      -128.89166259765625,
				CarrSpace: 1,
				PattGrp:   0,
				PattNum:   3,
				CRC:       [2]byte{0x57, 0x8e},
			},
			wantErr: false,
			json: `{
	"command": "vm.ulData",
	"opId": 4,
	"macType": 1,
	"userData": "VUSlEUZSeYF2B4wAt5APACwlPjURAJjK7MESQiVwehuTMgcQKkT7qUjx2PoYeYegkguJjTjeA/qf6HbmUC21pGcvX4Q1NSJdMZton1VXIbaQK3dGimw=",
	"trxTime": 2624805061575111,
	"sysTime": 1744394681234086798,
	"freqOff": 868157596,
	"snr": 3.2715067863464355,
	"rssi": -128.89166259765625,
	"carrSpace": 1,
	"pattGrp": 0,
	"pattNum": 3,
	"crc": [
		87,
		142
	]
}`,
		},
		{
			name:    "msgVmUlDataRsp",
			msgType: &VmUlDataRsp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 118, 109, 46, 117, 108, 68, 97, 116, 97, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			msg: &VmUlDataRsp{
				Command: structs.MsgVmUlDataRsp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "vm.ulDataRsp",
	"opId": 0
}`,
		},
		{
			name:    "msgVmUlDataCmp",
			msgType: &VmUlDataCmp{},
			raw:     []byte{130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 118, 109, 46, 117, 108, 68, 97, 116, 97, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			msg: &VmUlDataCmp{
				Command: structs.MsgVmUlDataCmp,
				OpId:    0,
			},
			wantErr: false,
			json: `{
	"command": "vm.ulDataCmp",
	"opId": 0
}`,
		},
	}
}

func (ts *TestMessageSuite) TestMessage_UnmarshalMessagePack() {
	t := ts.T()

	for _, tt := range ts.data {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			msg := tt.msgType
			_, err := msg.UnmarshalMsg(tt.raw)

			if tt.wantErr {
				assert.Error(err, "Message.UnmarshalMsg() expected error")
			} else if assert.NoError(err, "Message.UnmarshalMsg() unexpected error") {
				assert.Equal(tt.msg, msg)
			}

		})
	}
}

func (ts *TestMessageSuite) TestMessage_MarshalMessagePack() {
	t := ts.T()

	for _, tt := range ts.data {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			msg := tt.msg
			raw, err := msg.MarshalMsg(nil)

			if tt.wantErr {
				assert.Error(err, "Message.MarshalMsg() expected error")
			} else if assert.NoError(err, "Message.MarshalMsg() unexpected error") {
				assert.Equal(tt.raw, raw)
			}

		})
	}
}

func (ts *TestMessageSuite) TestMessage_UnmarshalJson() {
	t := ts.T()

	for _, tt := range ts.data {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			msg := tt.msgType
			err := json.Unmarshal([]byte(tt.json), msg)

			if tt.wantErr {
				assert.Error(err, "Message.UnmarshalJson() expected error")
			} else if assert.NoError(err, "Message.UnmarshalJson() unexpected error") {
				assert.Equal(tt.msg, msg)
			}
		})
	}
}

func (ts *TestMessageSuite) TestMessage_MarshalJson() {
	t := ts.T()

	for _, tt := range ts.data {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			msg := tt.msg

			jsonRaw, err := json.MarshalIndent(msg, "", "\t")

			value := string(jsonRaw)

			if tt.wantErr {
				assert.Error(err, "Message.MarshalMsg() expected error")
			} else if assert.NoError(err, "Message.MarshalMsg() unexpected error") {
				assert.Equal(tt.json, value)
			}
		})
	}
}

func BenchmarkMessage_UnmarshalMessagePack(b *testing.B) {
	ts := new(TestMessageSuite)
	ts.SetT(&testing.T{})
	ts.SetupSuite()
	b.ReportAllocs()
	b.ResetTimer()

	for _, tt := range ts.data {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				msg := tt.msgType
				_, _ = msg.UnmarshalMsg(tt.raw)
			}
		})
	}
}

func BenchmarkMessage_MarshalMessagePack(b *testing.B) {
	ts := new(TestMessageSuite)
	ts.SetT(&testing.T{})
	ts.SetupSuite()

	b.ReportAllocs()
	b.ResetTimer()

	for _, tt := range ts.data {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				msg := tt.msg
				raw, _ := msg.MarshalMsg(nil)
				_ = raw
			}
		})
	}
}
