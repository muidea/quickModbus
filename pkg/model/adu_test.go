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

	header, protocol, err := DecodeMBTcpProtocol(reader, ResponseAction)
	if err != SuccessCode {
		t.Errorf("DecodeMBTcpProtocol failed, error:%v", err)
		return
	}
	if header.Transaction() != 1 {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch transaction")
		return
	}
	if protocol.FuncCode() != WriteMultipleRegisters {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch funCode")
		return
	}
}

func TestDecodeReadHoldingRegisters(t *testing.T) {
	strVal := strings.ReplaceAll("00 02 00 00 00 2B 01 03 28 00 00 00 00 00 00 00 0C 00 00 00 00 00 00 00 22 00 00 00 00 00 00 00 38 00 00 00 00 00 00 00 4E 00 00 00 00 00 00 00 5A", " ", "")
	byteVal, _ := hex.DecodeString(strVal)
	reader := bytes.NewBuffer(byteVal)

	header, protocol, err := DecodeMBTcpProtocol(reader, ResponseAction)
	if err != SuccessCode {
		t.Errorf("DecodeMBTcpProtocol failed, error:%v", err)
		return
	}
	if header.Transaction() != 2 {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch transaction")
		return
	}
	if protocol.FuncCode() != ReadHoldingRegisters {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch funCode")
		return
	}
}

func TestDecodeReadWriteMultipleRegisters(t *testing.T) {
	strVal := strings.ReplaceAll("00 F9 00 00 00 0D 01 17 00 01 00 01 00 00 00 01 02 02 37", " ", "")
	byteVal, _ := hex.DecodeString(strVal)
	reader := bytes.NewBuffer(byteVal)

	header, protocol, err := DecodeMBTcpProtocol(reader, RequestAction)
	if err != SuccessCode {
		t.Errorf("DecodeMBTcpProtocol failed, error:%v", err)
		return
	}
	if header.Transaction() != 249 {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch transaction")
		return
	}
	if protocol.FuncCode() != ReadWriteMultipleRegisters {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch funCode")
		return
	}

	strVal = strings.ReplaceAll("00 07 00 00 00 0B 01 17 00 00 00 00 00 00 00 00 00", " ", "")
	byteVal, _ = hex.DecodeString(strVal)
	reader = bytes.NewBuffer(byteVal)

	header, protocol, err = DecodeMBTcpProtocol(reader, RequestAction)
	if err != SuccessCode {
		t.Errorf("DecodeMBTcpProtocol failed, error:%v", err)
		return
	}
	if header.Transaction() != 7 {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch transaction")
		return
	}
	if protocol.FuncCode() != ReadWriteMultipleRegisters {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch funCode")
		return
	}

	strVal = strings.ReplaceAll("00 03 00 00 00 0D 01 17 00 00 00 01 00 01 00 01 02 00 7B", " ", "")
	byteVal, _ = hex.DecodeString(strVal)
	reader = bytes.NewBuffer(byteVal)

	header, protocol, err = DecodeMBTcpProtocol(reader, RequestAction)
	if err != SuccessCode {
		t.Errorf("DecodeMBTcpProtocol failed, error:%v", err)
		return
	}
	if header.Transaction() != 3 {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch transaction")
		return
	}
	if protocol.FuncCode() != ReadWriteMultipleRegisters {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch funCode")
		return
	}

	strVal = strings.ReplaceAll("00 03 00 00 00 05 01 17 02 00 17", " ", "")
	byteVal, _ = hex.DecodeString(strVal)
	reader = bytes.NewBuffer(byteVal)

	header, protocol, err = DecodeMBTcpProtocol(reader, ResponseAction)
	if err != SuccessCode {
		t.Errorf("DecodeMBTcpProtocol failed, error:%v", err)
		return
	}
	if header.Transaction() != 3 {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch transaction")
		return
	}
	if protocol.FuncCode() != ReadWriteMultipleRegisters {
		t.Errorf("DecodeMBTcpProtocol failed, mismatch funCode")
		return
	}
}
