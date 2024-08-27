package model

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"
)

func TestDecodeWriteMultipleRegisters(t *testing.T) {
	strVal := "000100000006011000000002"
	byteVal, _ := hex.DecodeString(strVal)
	reader := bytes.NewBuffer(byteVal)

	header, protocol, err := DecodeMBProtocol(reader, ResponseAction)
	if err != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error:%v", err)
		return
	}
	if header.Transaction() != 1 {
		t.Errorf("DecodeMBProtocol failed, mismatch transaction")
		return
	}
	if protocol.FuncCode() != WriteMultipleRegisters {
		t.Errorf("DecodeMBProtocol failed, mismatch funCode")
		return
	}
}

func TestDecodeReadHoldingRegisters(t *testing.T) {
	strVal := strings.ReplaceAll("00 02 00 00 00 2B 01 03 28 00 00 00 00 00 00 00 0C 00 00 00 00 00 00 00 22 00 00 00 00 00 00 00 38 00 00 00 00 00 00 00 4E 00 00 00 00 00 00 00 5A", " ", "")
	byteVal, _ := hex.DecodeString(strVal)
	reader := bytes.NewBuffer(byteVal)

	header, protocol, err := DecodeMBProtocol(reader, ResponseAction)
	if err != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error:%v", err)
		return
	}
	if header.Transaction() != 2 {
		t.Errorf("DecodeMBProtocol failed, mismatch transaction")
		return
	}
	if protocol.FuncCode() != ReadHoldingRegisters {
		t.Errorf("DecodeMBProtocol failed, mismatch funCode")
		return
	}
}
