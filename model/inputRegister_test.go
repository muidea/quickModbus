package model

import (
	"bytes"
	"encoding/hex"
	"testing"
)

// ReadInputRegisters
// address: 0
// count: 13
func TestDecodeMB006(t *testing.T) {
	strVal := "1D7A0000000601040000000D"
	byteVal, byteErr := hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	buffVal := bytes.NewBuffer(byteVal)
	_, protocol, errCode := DecodeMBProtocol(buffVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if protocol.FuncCode() != ReadInputRegisters {
		t.Errorf("decode ReadInputRegisters request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBReadInputRegistersReq)
	if !reqOK {
		t.Errorf("decode ReadInputRegisters request failed")
		return
	}
	if reqPtr.Address() != 0 {
		t.Errorf("decode ReadInputRegisters request address failed")
		return
	}
	if reqPtr.Count() != 13 {
		t.Errorf("decode ReadInputRegisters request count failed")
		return
	}

	strVal = "1D7A0000001D01041A0200010000000000000000000000000000000000000004001000"
	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	buffVal = bytes.NewBuffer(byteVal)
	_, protocol, errCode = DecodeMBProtocol(buffVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if protocol.FuncCode() != ReadInputRegisters {
		t.Errorf("decode ReadInputRegisters request failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBReadInputRegistersRsp)
	if !rspOK {
		t.Errorf("decode ReadInputRegisters response failed")
		return
	}
	if len(rspPtr.Data()) != 13*2 {
		t.Errorf("decode ReadInputRegisters response count failed")
		return
	}

	u16Array, u16Err := ByteArrayToUint16ABArray(rspPtr.Data())
	if u16Err != nil {
		t.Errorf("decode ReadInputRegisters response, error:%s", u16Err.Error())
		return
	}

	u16Set := []uint16{512, 256, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1024, 4096}
	for idx := range u16Array {
		if u16Set[idx] != u16Array[idx] {
			t.Errorf("byte to u16 failed, idx:%v", idx)
		}
	}
	t.Logf("%v", u16Array)
}

// WriteSingleRegister
// address: 0
// value: 6789
func TestDecodeMB007(t *testing.T) {
	strVal := "2B1B00000006010600001A85"
	byteVal, byteErr := hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	buffVal := bytes.NewBuffer(byteVal)
	_, protocol, errCode := DecodeMBProtocol(buffVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if protocol.FuncCode() != WriteSingleRegister {
		t.Errorf("decode WriteSingleRegister request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBWriteSingleRegisterReq)
	if !reqOK {
		t.Errorf("decode WriteSingleRegister request failed")
		return
	}
	if reqPtr.Address() != 0 {
		t.Errorf("decode WriteSingleRegister request address failed")
		return
	}
	if len(reqPtr.Data()) != 2 {
		t.Errorf("decode WriteSingleRegister request data count failed")
		return
	}
	u16 := ByteToUint16AB(reqPtr.Data())
	if u16 != 6789 {
		t.Errorf("decode WriteSingleRegister request data failed")
		return
	}

	strVal = "2B1B00000006010600001A85"
	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	buffVal = bytes.NewBuffer(byteVal)
	_, protocol, errCode = DecodeMBProtocol(buffVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if protocol.FuncCode() != WriteSingleRegister {
		t.Errorf("decode WriteSingleRegister response failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBWriteSingleRegisterRsp)
	if !rspOK {
		t.Errorf("decode WriteSingleRegister response failed")
		return
	}
	if len(rspPtr.Data()) != 2 {
		t.Errorf("decode WriteSingleRegister response count failed")
		return
	}

	u16Val := ByteToUint16AB(rspPtr.Data())
	if u16Val != 6789 {
		t.Errorf("byte to u16 failed")
	}
}

// WriteMultipleRegisters
// address: 0
// count: 1
// value: 6789
func TestDecodeMB008(t *testing.T) {
	strVal := "310300000009011000000001021A85"
	byteVal, byteErr := hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	buffVal := bytes.NewBuffer(byteVal)
	_, protocol, errCode := DecodeMBProtocol(buffVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if protocol.FuncCode() != WriteMultipleRegisters {
		t.Errorf("decode WriteMultipleRegisters request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBWriteMultipleRegistersReq)
	if !reqOK {
		t.Errorf("decode WriteMultipleRegisters request failed")
		return
	}
	if reqPtr.Address() != 0 {
		t.Errorf("decode WriteMultipleRegisters request address failed")
		return
	}
	if reqPtr.Count() != 1 {
		t.Errorf("decode WriteMultipleRegisters request count failed")
		return
	}
	if len(reqPtr.Data()) != 2 {
		t.Errorf("decode WriteMultipleRegisters request data len failed")
		return
	}

	u16Array, u16Err := ByteArrayToUint16ABArray(reqPtr.Data())
	if u16Err != nil || len(u16Array) != 1 {
		t.Errorf("decode WriteMultipleRegisters request data value failed")
		return
	}
	if u16Array[0] != 6789 {
		t.Errorf("decode WriteMultipleRegisters request data value failed")
		return
	}

	strVal = "310300000006011000000001"
	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	buffVal = bytes.NewBuffer(byteVal)
	_, protocol, errCode = DecodeMBProtocol(buffVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if protocol.FuncCode() != WriteMultipleRegisters {
		t.Errorf("decode WriteMultipleRegisters response failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBWriteMultipleRegistersRsp)
	if !rspOK {
		t.Errorf("decode WriteMultipleRegisters response failed")
		return
	}
	if rspPtr.Address() != 0 {
		t.Errorf("decode WriteMultipleRegisters response address failed")
		return
	}
	if rspPtr.Count() != 1 {
		t.Errorf("decode WriteMultipleRegisters response count failed")
		return
	}
}

// WriteMultipleRegisters
// address: 0
// count: 10
// value: 12,23,34,45,56,67,78,90,100
func TestDecodeMB009(t *testing.T) {
	strVal := "39770000001B01100000000A14000C00170022002D00380043004E0059005A0064"
	byteVal, byteErr := hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	buffVal := bytes.NewBuffer(byteVal)
	_, protocol, errCode := DecodeMBProtocol(buffVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if protocol.FuncCode() != WriteMultipleRegisters {
		t.Errorf("decode WriteMultipleRegisters request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBWriteMultipleRegistersReq)
	if !reqOK {
		t.Errorf("decode WriteMultipleRegisters request failed")
		return
	}
	if reqPtr.Address() != 0 {
		t.Errorf("decode WriteMultipleRegisters request address failed")
		return
	}
	if reqPtr.Count() != 10 {
		t.Errorf("decode WriteMultipleRegisters request count failed")
		return
	}
	if len(reqPtr.Data()) != 20 {
		t.Errorf("decode WriteMultipleRegisters request data len failed")
		return
	}

	u16Array, u16Err := ByteArrayToUint16ABArray(reqPtr.Data())
	if u16Err != nil || len(u16Array) != 10 {
		t.Errorf("decode WriteMultipleRegisters request data value failed")
		return
	}
	if u16Array[0] != 12 {
		t.Errorf("decode WriteMultipleRegisters request data value failed")
		return
	}

	strVal = "39770000000601100000000A"
	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	buffVal = bytes.NewBuffer(byteVal)
	_, protocol, errCode = DecodeMBProtocol(buffVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if protocol.FuncCode() != WriteMultipleRegisters {
		t.Errorf("decode WriteMultipleRegisters response failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBWriteMultipleRegistersRsp)
	if !rspOK {
		t.Errorf("decode WriteMultipleRegisters response failed")
		return
	}
	if rspPtr.Address() != 0 {
		t.Errorf("decode WriteMultipleRegisters response address failed")
		return
	}
	if rspPtr.Count() != 10 {
		t.Errorf("decode WriteMultipleRegisters response count failed")
		return
	}
}
