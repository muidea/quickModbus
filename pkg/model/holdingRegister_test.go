package model

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/muidea/quickModbus/pkg/common"
)

// ReadHoldingRegisters
// address: 0
// count: 13
func TestDecodeMB005(t *testing.T) {
	strVal := "10400000000601030000000D"
	byteVal, byteErr := hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	byteBuff := bytes.NewBuffer(byteVal)
	_, protocol, errCode := DecodeMBProtocol(byteBuff, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if protocol.FuncCode() != ReadHoldingRegisters {
		t.Errorf("decode ReadHoldingRegisters request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBReadHoldingRegistersReq)
	if !reqOK {
		t.Errorf("decode ReadHoldingRegisters request failed")
		return
	}
	if reqPtr.Address() != 0 {
		t.Errorf("decode ReadHoldingRegisters request address failed")
		return
	}
	if reqPtr.Count() != 13 {
		t.Errorf("decode ReadHoldingRegisters request count failed")
		return
	}

	strVal = "10400000001D01031A00010002000300000000000000070008000900000000000C000D"
	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	byteBuff = bytes.NewBuffer(byteVal)
	_, protocol, errCode = DecodeMBProtocol(byteBuff, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if protocol.FuncCode() != ReadHoldingRegisters {
		t.Errorf("decode ReadHoldingRegisters request failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBReadHoldingRegistersRsp)
	if !rspOK {
		t.Errorf("decode ReadHoldingRegisters response failed")
		return
	}
	if len(rspPtr.Data()) != 13*2 {
		t.Errorf("decode ReadHoldingRegisters response count failed")
		return
	}

	u16Array, u16Err := common.BytesToUint16Array(rspPtr.Data(), common.ABCDEndian)
	if u16Err != nil {
		t.Errorf("decode ReadHoldingRegisters response, error:%s", u16Err.Error())
		return
	}

	u16Set := []uint16{1, 2, 3, 0, 0, 0, 7, 8, 9, 0, 0, 12, 13}
	for idx := range u16Array {
		if u16Set[idx] != u16Array[idx] {
			t.Errorf("byte to u16 failed, idx:%v", idx)
		}
	}
	t.Logf("%v", u16Array)
}
