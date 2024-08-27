package model

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestDecodeMBProtocol(t *testing.T) {
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
