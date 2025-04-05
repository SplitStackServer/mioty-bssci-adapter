package bssci_v1

import (
	"context"

	"errors"

	"mioty-bssci-adapter/internal/api/cmd"
	"mioty-bssci-adapter/internal/api/msg"
	"mioty-bssci-adapter/internal/api/rsp"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs"
	"mioty-bssci-adapter/internal/backend/bssci_v1/structs/messages"
	"mioty-bssci-adapter/internal/backend/events"
	"mioty-bssci-adapter/internal/common"
	"mioty-bssci-adapter/internal/config"
	"net"
	"os"

	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type TestBackendSuite struct {
	suite.Suite

	backend *Backend
	bs_eui  common.EUI64

	handleBasestationMessagesTestCases []testCaseHandleBasestationMessages
	handleServerCommandTestCases       []testCaseHandleServerCommand
	handleServerResponseTestCases      []testCaseHandleServerResponse
}

type testCaseHandleBasestationMessages struct {
	name                    string
	payload                 []byte
	expectResponse          bool
	expectedResponseCommand structs.Command
}

type testCaseHandleServerCommand struct {
	name    string
	cmd     *cmd.ProtoCommand
	wantErr bool
}

type testCaseHandleServerResponse struct {
	name    string
	rsp     *rsp.ProtoResponse
	wantErr bool
}

func TestBackend(t *testing.T) {
	suite.Run(t, new(TestBackendSuite))
}

func (ts *TestBackendSuite) SetupSuite() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	ts.bs_eui = common.EUI64{0, 1, 2, 3, 4, 5, 6, 7}

	ts.handleBasestationMessagesTestCases = []testCaseHandleBasestationMessages{
		{
			name:                    "unsupported",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 27, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 171, 117, 110, 115, 117, 112, 112, 111, 114, 116, 101, 100, 164, 111, 112, 73, 100, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgError,
		},
		{
			name:                    "con",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 160, 0, 0, 0, 137, 167, 99, 111, 109, 109, 97, 110, 100, 163, 99, 111, 110, 164, 111, 112, 73, 100, 0, 167, 118, 101, 114, 115, 105, 111, 110, 165, 49, 46, 48, 46, 48, 165, 98, 115, 69, 117, 105, 207, 0, 7, 50, 0, 0, 119, 103, 243, 166, 118, 101, 110, 100, 111, 114, 174, 68, 105, 101, 104, 108, 32, 77, 101, 116, 101, 114, 105, 110, 103, 165, 109, 111, 100, 101, 108, 181, 77, 73, 79, 84, 89, 32, 80, 114, 101, 109, 105, 117, 109, 32, 71, 97, 116, 101, 119, 97, 121, 164, 110, 97, 109, 101, 173, 77, 48, 48, 48, 55, 51, 50, 55, 55, 54, 55, 70, 51, 164, 98, 105, 100, 105, 195, 168, 115, 110, 66, 115, 85, 117, 105, 100, 220, 0, 16, 208, 195, 114, 208, 197, 33, 208, 167, 120, 73, 208, 155, 208, 139, 78, 41, 208, 199, 208, 131, 208, 183, 53, 208, 221},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgConRsp,
		},
		{
			name:                    "conCmp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 22, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 166, 99, 111, 110, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgConCmp,
		},
		{
			name:                    "detPrpRsp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 25, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 100, 101, 116, 80, 114, 112, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgDetPrpCmp,
		},
		{
			name:                    "attPrpRsp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 25, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 97, 116, 116, 80, 114, 112, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgAttPrpCmp,
		},
		{
			name:                    "error",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 49, 0, 0, 0, 132, 167, 99, 111, 109, 109, 97, 110, 100, 165, 101, 114, 114, 111, 114, 164, 111, 112, 73, 100, 0, 164, 99, 111, 100, 101, 1, 167, 109, 101, 115, 115, 97, 103, 101, 173, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgErrorAck,
		},
		{
			name:                    "errorAck",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 24, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 168, 101, 114, 114, 111, 114, 65, 99, 107, 164, 111, 112, 73, 100, 0},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgErrorAck,
		},
		{
			name:                    "ping",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 20, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 164, 112, 105, 110, 103, 164, 111, 112, 73, 100, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgPingRsp,
		},
		{
			name:                    "pingRsp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 23, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 167, 112, 105, 110, 103, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgPingCmp,
		},
		{
			name:                    "pingCmp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 23, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 167, 112, 105, 110, 103, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgPingCmp,
		},
		{
			name:                    "statusRsp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 125, 0, 0, 0, 138, 167, 99, 111, 109, 109, 97, 110, 100, 169, 115, 116, 97, 116, 117, 115, 82, 115, 112, 164, 111, 112, 73, 100, 0, 164, 99, 111, 100, 101, 0, 167, 109, 101, 115, 115, 97, 103, 101, 162, 111, 107, 164, 116, 105, 109, 101, 206, 59, 154, 202, 5, 169, 100, 117, 116, 121, 67, 121, 99, 108, 101, 202, 62, 204, 204, 205, 166, 117, 112, 116, 105, 109, 101, 205, 3, 232, 164, 116, 101, 109, 112, 203, 64, 70, 192, 0, 0, 0, 0, 0, 167, 99, 112, 117, 76, 111, 97, 100, 203, 63, 224, 0, 0, 0, 0, 0, 0, 167, 109, 101, 109, 76, 111, 97, 100, 203, 63, 227, 51, 51, 51, 51, 51, 51},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgStatusCmp,
		},
		{
			name:                    "att",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 0, 1, 0, 0, 222, 0, 18, 167, 99, 111, 109, 109, 97, 110, 100, 163, 97, 116, 116, 164, 111, 112, 73, 100, 0, 165, 101, 112, 69, 117, 105, 211, 8, 7, 6, 5, 4, 3, 2, 1, 166, 114, 120, 84, 105, 109, 101, 1, 170, 114, 120, 68, 117, 114, 97, 116, 105, 111, 110, 205, 1, 244, 169, 97, 116, 116, 97, 99, 104, 67, 110, 116, 2, 163, 115, 110, 114, 203, 64, 8, 0, 0, 0, 0, 0, 0, 164, 114, 115, 115, 105, 203, 192, 89, 0, 0, 0, 0, 0, 0, 165, 101, 113, 83, 110, 114, 203, 64, 16, 0, 0, 0, 0, 0, 0, 167, 112, 114, 111, 102, 105, 108, 101, 162, 101, 117, 170, 115, 117, 98, 112, 97, 99, 107, 101, 116, 115, 132, 163, 115, 110, 114, 147, 1, 2, 3, 164, 114, 115, 115, 105, 147, 4, 5, 6, 169, 102, 114, 101, 113, 117, 101, 110, 99, 121, 147, 7, 8, 9, 165, 112, 104, 97, 115, 101, 147, 10, 11, 12, 165, 110, 111, 110, 99, 101, 196, 4, 4, 5, 6, 7, 164, 115, 105, 103, 110, 196, 4, 1, 2, 3, 4, 166, 115, 104, 65, 100, 100, 114, 205, 255, 255, 168, 100, 117, 97, 108, 67, 104, 97, 110, 194, 170, 114, 101, 112, 101, 116, 105, 116, 105, 111, 110, 194, 171, 119, 105, 100, 101, 67, 97, 114, 114, 79, 102, 102, 194, 171, 108, 111, 110, 103, 66, 108, 107, 68, 105, 115, 116, 194},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgAtt,
		},
		{
			name:                    "det",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 184, 0, 0, 0, 140, 167, 99, 111, 109, 109, 97, 110, 100, 163, 100, 101, 116, 164, 111, 112, 73, 100, 0, 165, 101, 112, 69, 117, 105, 211, 8, 7, 6, 5, 4, 3, 2, 1, 166, 114, 120, 84, 105, 109, 101, 1, 170, 114, 120, 68, 117, 114, 97, 116, 105, 111, 110, 205, 1, 244, 169, 112, 97, 99, 107, 101, 116, 67, 110, 116, 2, 163, 115, 110, 114, 203, 64, 8, 0, 0, 0, 0, 0, 0, 164, 114, 115, 115, 105, 203, 192, 89, 0, 0, 0, 0, 0, 0, 165, 101, 113, 83, 110, 114, 203, 64, 16, 0, 0, 0, 0, 0, 0, 167, 112, 114, 111, 102, 105, 108, 101, 162, 101, 117, 170, 115, 117, 98, 112, 97, 99, 107, 101, 116, 115, 132, 163, 115, 110, 114, 147, 1, 2, 3, 164, 114, 115, 115, 105, 147, 4, 5, 6, 169, 102, 114, 101, 113, 117, 101, 110, 99, 121, 147, 7, 8, 9, 165, 112, 104, 97, 115, 101, 147, 10, 11, 12, 164, 115, 105, 103, 110, 196, 4, 1, 2, 3, 4},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgDet,
		},
		{
			name:                    "ulData",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 237, 0, 0, 0, 143, 167, 99, 111, 109, 109, 97, 110, 100, 166, 117, 108, 68, 97, 116, 97, 164, 111, 112, 73, 100, 0, 165, 101, 112, 69, 117, 105, 211, 8, 7, 6, 5, 4, 3, 2, 1, 166, 114, 120, 84, 105, 109, 101, 1, 170, 114, 120, 68, 117, 114, 97, 116, 105, 111, 110, 205, 1, 244, 169, 112, 97, 99, 107, 101, 116, 67, 110, 116, 2, 163, 115, 110, 114, 203, 64, 8, 0, 0, 0, 0, 0, 0, 164, 114, 115, 115, 105, 203, 192, 89, 0, 0, 0, 0, 0, 0, 165, 101, 113, 83, 110, 114, 203, 64, 16, 0, 0, 0, 0, 0, 0, 167, 112, 114, 111, 102, 105, 108, 101, 162, 101, 117, 170, 115, 117, 98, 112, 97, 99, 107, 101, 116, 115, 132, 163, 115, 110, 114, 147, 1, 2, 3, 164, 114, 115, 115, 105, 147, 4, 5, 6, 169, 102, 114, 101, 113, 117, 101, 110, 99, 121, 147, 7, 8, 9, 165, 112, 104, 97, 115, 101, 147, 10, 11, 12, 168, 117, 115, 101, 114, 68, 97, 116, 97, 196, 22, 116, 104, 105, 115, 105, 115, 97, 108, 111, 116, 111, 102, 100, 97, 116, 97, 116, 111, 115, 101, 110, 100, 166, 100, 108, 79, 112, 101, 110, 194, 171, 114, 101, 115, 112, 111, 110, 115, 101, 69, 120, 112, 194, 165, 100, 108, 65, 99, 107, 194},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgUlDataRsp,
		},
		{
			name:                    "dlDataRes",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 91, 0, 0, 0, 135, 167, 99, 111, 109, 109, 97, 110, 100, 169, 100, 108, 68, 97, 116, 97, 82, 101, 115, 164, 111, 112, 73, 100, 0, 165, 101, 112, 69, 117, 105, 211, 8, 7, 6, 5, 4, 3, 2, 1, 165, 113, 117, 101, 73, 100, 206, 0, 188, 97, 78, 166, 114, 101, 115, 117, 108, 116, 167, 115, 117, 99, 99, 101, 115, 115, 166, 116, 120, 84, 105, 109, 101, 206, 0, 188, 97, 78, 169, 112, 97, 99, 107, 101, 116, 67, 110, 116, 205, 4, 210},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgDlDataResRsp,
		},
		{
			name:                    "dlRxStat",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 95, 0, 0, 0, 135, 167, 99, 111, 109, 109, 97, 110, 100, 168, 100, 108, 82, 120, 83, 116, 97, 116, 164, 111, 112, 73, 100, 0, 165, 101, 112, 69, 117, 105, 211, 8, 7, 6, 5, 4, 3, 2, 1, 166, 114, 120, 84, 105, 109, 101, 1, 169, 112, 97, 99, 107, 101, 116, 67, 110, 116, 205, 4, 210, 167, 100, 108, 82, 120, 83, 110, 114, 203, 64, 0, 0, 0, 0, 0, 0, 0, 168, 100, 108, 82, 120, 82, 115, 115, 105, 203, 192, 89, 0, 0, 0, 0, 0, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgDlRxStatRsp,
		},
		{
			name:                    "dlDataRevRsp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 28, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 100, 108, 68, 97, 116, 97, 82, 101, 118, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgDlDataRevCmp,
		},
		{
			name:                    "dlDlDataQueRsp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 28, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 100, 108, 68, 97, 116, 97, 81, 117, 101, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgDlDataQueCmp,
		},
		{
			name:                    "dlRxStatQryRsp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 30, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 174, 100, 108, 82, 120, 83, 116, 97, 116, 81, 114, 121, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgDlRxStatQryCmp,
		},
		{
			name:                    "ulDataCmp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 25, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 169, 117, 108, 68, 97, 116, 97, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgUlDataCmp,
		},
		{
			name:                    "attCmp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 22, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 166, 97, 116, 116, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgAttCmp,
		},
		{
			name:                    "detCmp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 22, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 166, 100, 101, 116, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgDetCmp,
		},
		{
			name:                    "dlDataResCmp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 28, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 100, 108, 68, 97, 116, 97, 82, 101, 115, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgDlDataResCmp,
		},
		{
			name:                    "dlRxStatCmp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 27, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 171, 100, 108, 82, 120, 83, 116, 97, 116, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgDlRxStatCmp,
		},
		{
			name:                    "vmUlData",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 207, 0, 0, 0, 143, 167, 99, 111, 109, 109, 97, 110, 100, 169, 118, 109, 46, 117, 108, 68, 97, 116, 97, 164, 111, 112, 73, 100, 10, 167, 109, 97, 99, 84, 121, 112, 101, 0, 168, 117, 115, 101, 114, 68, 97, 116, 97, 196, 3, 1, 2, 3, 167, 116, 114, 120, 84, 105, 109, 101, 1, 167, 115, 121, 115, 84, 105, 109, 101, 1, 167, 102, 114, 101, 113, 79, 102, 102, 60, 163, 115, 110, 114, 203, 64, 8, 0, 0, 0, 0, 0, 0, 164, 114, 115, 115, 105, 203, 192, 89, 0, 0, 0, 0, 0, 0, 165, 101, 113, 83, 110, 114, 203, 64, 16, 0, 0, 0, 0, 0, 0, 170, 115, 117, 98, 112, 97, 99, 107, 101, 116, 115, 132, 163, 115, 110, 114, 147, 1, 2, 3, 164, 114, 115, 115, 105, 147, 4, 5, 6, 169, 102, 114, 101, 113, 117, 101, 110, 99, 121, 147, 7, 8, 9, 165, 112, 104, 97, 115, 101, 147, 10, 11, 12, 169, 99, 97, 114, 114, 83, 112, 97, 99, 101, 2, 167, 112, 97, 116, 116, 71, 114, 112, 3, 167, 112, 97, 116, 116, 78, 117, 109, 5, 163, 99, 114, 99, 196, 2, 6, 7},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgVmUlDataRsp,
		},
		{
			name:                    "vmUlDataCmp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 28, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 172, 118, 109, 46, 117, 108, 68, 97, 116, 97, 67, 109, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          false,
			expectedResponseCommand: structs.MsgVmUlDataCmp,
		},
		{
			name:                    "vmStatusRsp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 41, 0, 0, 0, 131, 167, 99, 111, 109, 109, 97, 110, 100, 172, 118, 109, 46, 115, 116, 97, 116, 117, 115, 82, 115, 112, 164, 111, 112, 73, 100, 0, 168, 109, 97, 99, 84, 121, 112, 101, 115, 147, 1, 2, 3},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgVmStatusCmp,
		},
		{
			name:                    "vmActivateRsp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 30, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 174, 118, 109, 46, 97, 99, 116, 105, 118, 97, 116, 101, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgVmActivateCmp,
		},
		{
			name:                    "vmDeactivateRsp",
			payload:                 []byte{77, 73, 79, 84, 89, 66, 48, 49, 32, 0, 0, 0, 130, 167, 99, 111, 109, 109, 97, 110, 100, 176, 118, 109, 46, 100, 101, 97, 99, 116, 105, 118, 97, 116, 101, 82, 115, 112, 164, 111, 112, 73, 100, 0},
			expectResponse:          true,
			expectedResponseCommand: structs.MsgVmDeactivateCmp,
		},
	}

	ts.handleServerCommandTestCases = []testCaseHandleServerCommand{
		{
			name: "ProtoCommand_DlDataQue",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_DlDataQue{
					DlDataQue: &cmd.EnqueDownlink{
						EndnodeEui:     123,
						DlQueId:        456,
						Priority:       new(float32),
						Format:         new(uint32),
						Payload:        &cmd.EnqueDownlink_Ack{Ack: &cmd.Acknowledgement{}},
						ResponseExp:    new(bool),
						ResponsePrio:   new(bool),
						ReqDlWindow:    new(bool),
						OnlyIfExpected: new(bool),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoCommand_DlDataQue_Err",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_DlDataQue{
					DlDataQue: &cmd.EnqueDownlink{
						EndnodeEui:     123,
						DlQueId:        456,
						Priority:       new(float32),
						Format:         new(uint32),
						Payload:        nil,
						ResponseExp:    new(bool),
						ResponsePrio:   new(bool),
						ReqDlWindow:    new(bool),
						OnlyIfExpected: new(bool),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "ProtoCommand_DlDataRev",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_DlDataRev{
					DlDataRev: &cmd.RevokeDownlink{
						EndnodeEui: 123,
						DlQueId:    456,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoCommand_DlDataRev_Err",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_DlDataRev{
					DlDataRev: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "ProtoCommand_DlRxStatQry",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_DlRxStatQry{
					DlRxStatQry: &cmd.DownlinkRxStatusQuery{
						EndnodeEui: 123,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoCommand_DlRxStatQry_Err",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_DlRxStatQry{
					DlRxStatQry: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "ProtoCommand_AttPrp",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_AttPrp{
					AttPrp: &cmd.AttachPropagate{
						EndnodeEui:    123,
						NwkSessionKey: []byte{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoCommand_AttPrp_Err",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_AttPrp{
					AttPrp: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "ProtoCommand_DetPrp",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_DetPrp{
					DetPrp: &cmd.DetachPropagate{
						EndnodeEui: 123,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoCommand_DetPrp_Err",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_DetPrp{
					DetPrp: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "ProtoCommand_ReqStatus",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_ReqStatus{
					ReqStatus: &cmd.RequestStatus{},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoCommand_ReqStatus_Err",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_ReqStatus{
					ReqStatus: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "ProtoCommand_VmActivate",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_VmActivate{
					VmActivate: &cmd.EnableVariableMac{},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoCommand_VmActivate_Err",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_VmActivate{
					VmActivate: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "ProtoCommand_VmDeactivate",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_VmDeactivate{
					VmDeactivate: &cmd.DisableVariableMac{},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoCommand_VmDeactivate_Err",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_VmDeactivate{
					VmDeactivate: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "ProtoCommand_VmStatus",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_VmStatus{
					VmStatus: &cmd.RequestVariableMacStatus{},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoCommand_VmStatus_Err",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &cmd.ProtoCommand_VmStatus{
					VmStatus: nil,
				},
			},
			wantErr: true,
		},
		{
			name:    "Nil",
			cmd:     nil,
			wantErr: true,
		},
		{
			name:    "Empty",
			cmd:     &cmd.ProtoCommand{},
			wantErr: true,
		},
		{
			name: "Basestation Error",
			cmd: &cmd.ProtoCommand{
				BsEui: ts.bs_eui.ToUnsignedInt() + 1,
				V1: &cmd.ProtoCommand_ReqStatus{
					ReqStatus: &cmd.RequestStatus{},
				},
			},
			wantErr: true,
		},
	}

	ts.handleServerResponseTestCases = []testCaseHandleServerResponse{
		{
			name: "ProtoResponse_AttRsp",
			rsp: &rsp.ProtoResponse{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &rsp.ProtoResponse_AttRsp{
					AttRsp: &rsp.EndnodeAttachResponse{
						EndnodeEui:    123,
						NwkSessionKey: []byte{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoResponse_AttRsp_Err",
			rsp: &rsp.ProtoResponse{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &rsp.ProtoResponse_AttRsp{
					AttRsp: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "ProtoResponse_DetRsp",
			rsp: &rsp.ProtoResponse{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &rsp.ProtoResponse_DetRsp{
					DetRsp: &rsp.EndnodeDetachResponse{
						EndnodeEui: 123,
						Sign:       456,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "ProtoResponse_DetRsp_Err",
			rsp: &rsp.ProtoResponse{
				BsEui: ts.bs_eui.ToUnsignedInt(),
				V1: &rsp.ProtoResponse_DetRsp{
					DetRsp: nil,
				},
			},
			wantErr: true,
		},
		{
			name:    "Nil",
			rsp:     nil,
			wantErr: true,
		},
		{
			name:    "Empty",
			rsp:     &rsp.ProtoResponse{},
			wantErr: true,
		},
		{
			name: "Basestation Error",
			rsp: &rsp.ProtoResponse{
				BsEui: ts.bs_eui.ToUnsignedInt() + 1,
				V1: &rsp.ProtoResponse_DetRsp{
					DetRsp: &rsp.EndnodeDetachResponse{
						EndnodeEui: 123,
						Sign:       456,
					},
				},
			},
			wantErr: true,
		},
	}

}

func (ts *TestBackendSuite) SetupTest() {
	var err error
	assert := require.New(ts.T())

	// setup config
	var conf config.Config
	conf.Backend.Type = "bssci_v1"
	conf.Backend.BssciV1.Bind = "127.0.0.1:0"
	conf.Backend.BssciV1.StatsInterval = time.Minute
	conf.Backend.BssciV1.PingInterval = 30 * time.Second
	conf.Backend.BssciV1.ReadTimeout = 2 * time.Minute
	conf.Backend.BssciV1.WriteTimeout = time.Second

	backend, err := NewBackend(conf)
	assert.NoError(err)

	subscribeChan := make(chan events.Subscribe, 1)

	// setup backend listener
	backend.SetSubscribeEventHandler(func(pl events.Subscribe) {
		subscribeChan <- pl
	})
	backend.SetBasestationMessageHandler(func(*msg.ProtoBasestationMessage) {})
	backend.SetEndnodeMessageHandler(func(*msg.ProtoEndnodeMessage) {})

	ts.backend = backend
}

func (ts *TestBackendSuite) TearDownTest() {

	assert := require.New(ts.T())

	assert.NoError(ts.backend.Stop())
}

func (ts *TestBackendSuite) TestBackend_GetBssciVersion() {
	assert := require.New(ts.T())

	version := ts.backend.GetBssciVersion()

	assert.Equal("1.0.0", version)
}

func (ts *TestBackendSuite) TestBackend_HandleBasestationMessages() {
	t := ts.T()

	for _, tt := range ts.handleBasestationMessagesTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			serverConn, clientConn := net.Pipe()

			bsConnection := connection{
				conn: clientConn,
				// stats:      stats.NewCollector(),
				lastActive: time.Now(),
				opId:       -1,
				SnBsUuid:   uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7},
				SnScUuid:   uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7},
			}

			go func() {
				defer clientConn.Close()
				clientConn.SetReadDeadline(time.Now().Add(time.Second))
				ctx := context.Background()
				assert.NoError(ts.backend.handleBasestationMessages(ctx, ts.bs_eui, &bsConnection))
			}()

			serverConn.SetDeadline(time.Now().Add(time.Second))
			_, err := serverConn.Write(tt.payload)

			assert.NoError(err)
			if tt.expectResponse {
				cmd, _, err := ReadBssciMessage(serverConn)
				assert.NoError(err)
				assert.Equal(tt.expectedResponseCommand, cmd.Command)
			}
		})
	}
}

func BenchmarkBackend_HandleBasestationMessages(b *testing.B) {
	ts := new(TestBackendSuite)
	ts.SetT(&testing.T{})
	ts.SetupSuite()
	log.Logger = zerolog.Nop()

	b.ReportAllocs()

	for _, tt := range ts.handleBasestationMessagesTestCases {
		b.Run(tt.name, func(b *testing.B) {
			ts.SetupTest()
			serverConn, clientConn := net.Pipe()

			bsConnection := connection{
				conn: clientConn,
				// stats:      stats.NewCollector(),
				lastActive: time.Now(),
				opId:       -1,
				SnBsUuid:   uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7},
				SnScUuid:   uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7},
			}

			go func() {
				defer clientConn.Close()
				clientConn.SetReadDeadline(time.Now().Add(time.Second))
				ctx := context.Background()
				ts.backend.handleBasestationMessages(ctx, ts.bs_eui, &bsConnection)
			}()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				serverConn.SetDeadline(time.Now().Add(time.Second))
				serverConn.Write(tt.payload)
				if tt.expectResponse {
					ReadBssciMessage(serverConn)
				}
			}
		})
	}
}

func (ts *TestBackendSuite) TestBackend_HandleServerCommand() {
	t := ts.T()

	serverConn, clientConn := net.Pipe()

	bsConnection := connection{
		conn: clientConn,
		// stats:      stats.NewCollector(),
		lastActive: time.Now(),
		opId:       -1,
		SnBsUuid:   uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7},
		SnScUuid:   uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7},
	}
	ts.backend.basestations.set(ts.bs_eui, &bsConnection)

	go func() {
		defer serverConn.Close()

		for {
			serverConn.SetReadDeadline(time.Now().Add(time.Second))
			buf := make([]byte, 12)
			serverConn.Read(buf)
		}
	}()

	for _, tt := range ts.handleServerCommandTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			err := ts.backend.HandleServerCommand(tt.cmd)

			if tt.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}

func (ts *TestBackendSuite) TestBackend_HandleServerResponse() {
	t := ts.T()

	serverConn, clientConn := net.Pipe()

	bsConnection := connection{
		conn: clientConn,
		// stats:      stats.NewCollector(),
		lastActive: time.Now(),
		opId:       -1,
		SnBsUuid:   uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7},
		SnScUuid:   uuid.UUID{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7},
	}
	ts.backend.basestations.set(ts.bs_eui, &bsConnection)

	go func() {
		defer serverConn.Close()

		for {
			serverConn.SetReadDeadline(time.Now().Add(time.Second))
			buf := make([]byte, 12)
			serverConn.Read(buf)
		}
	}()

	for _, tt := range ts.handleServerResponseTestCases {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			err := ts.backend.HandleServerResponse(tt.rsp)

			if tt.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}

func (ts *TestBackendSuite) TestBackend_initBasestation() {
	t := ts.T()

	type args struct {
		con     messages.Con
		handler func(ctx context.Context, eui common.EUI64, conn *connection) error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "new_con",
			args: args{
				con: messages.Con{
					Command:  structs.MsgCon,
					OpId:     0,
					Version:  "1.0.0",
					BsEui:    common.EUI64{1},
					Bidi:     false,
					SnBsUuid: structs.SessionUuid{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
				},
				handler: func(ctx context.Context, eui common.EUI64, conn *connection) error {
					return errors.New("exit test")
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			ctx := context.Background()
			logger := log.Logger
			ctx = logger.WithContext(ctx)
			server, client := net.Pipe()

			go func() {
				defer client.Close()
				cmd, _, err := ReadBssciMessage(client)
				assert.NoError(err)
				assert.Equal(structs.MsgConRsp, cmd.Command)
			}()

			err := ts.backend.initBasestation(ctx, tt.args.con, server, tt.args.handler)
			assert.NoError(err)
		})
	}
}

func (ts *TestBackendSuite) TestBackend_Start() {
	t := ts.T()

	tests := []struct {
		name string
		msg  messages.Message
	}{
		{
			name: "valid",
			msg: &messages.Con{
				Command:  structs.MsgCon,
				OpId:     0,
				Version:  "1.0.0",
				BsEui:    common.EUI64{1},
				Bidi:     false,
				SnBsUuid: structs.SessionUuid{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			},
		},
		{
			name: "invalid",
			msg: &messages.Ping{
				Command: structs.MsgPing,
				OpId:    0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			err := ts.backend.Start()
			assert.NoError(err)

			addr := ts.backend.listener.Addr().String()

			conn, err := net.Dial("tcp", addr)
			assert.NoError(err)

			go func() {
				err := WriteBssciMessage(conn, tt.msg)
				assert.NoError(err)

			}()

		})
	}
}
