package model

import (
	"encoding/hex"
	"testing"
)

// ReadCoils
// address: 0
// count: 10
func TestDecodeMB001(t *testing.T) {
	strVal := "00060000000601010000000A"
	byteVal, byteErr := hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	header, protocol, errCode := DecodeMBProtocol(byteVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
		return
	}

	if protocol.FuncCode() != ReadCoils {
		t.Errorf("decode ReadCoils request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBReadCoilsReq)
	if !reqOK {
		t.Errorf("decode ReadCoils request failed")
		return
	}
	if reqPtr.Address() != 0 {
		t.Errorf("decode ReadCoils request address failed")
		return
	}
	if reqPtr.Count() != 10 {
		t.Errorf("decode ReadCoils request count failed")
		return
	}

	strVal = "05CB00000005010102C103"
	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	header, protocol, errCode = DecodeMBProtocol(byteVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
		return
	}

	if protocol.FuncCode() != ReadCoils {
		t.Errorf("decode ReadCoils request failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBReadCoilsRsp)
	if !rspOK {
		t.Errorf("decode ReadCoils response failed")
		return
	}
	if rspPtr.Count() != 2 {
		t.Errorf("decode ReadCoils response count failed")
		return
	}

	// 0,6,7,8,9 = true
	// other = false
	trueSet := []int{0, 6, 7, 8, 9}
	boolArray := ByteArrayToBoolArray(rspPtr.Data())
	for idx := range boolArray {
		findFlag := false
		for _, val := range trueSet {
			if idx == val {
				findFlag = true
				break
			}
		}

		if boolArray[idx] != findFlag {
			t.Errorf("byte to bool failed")
		}
	}

	t.Logf("%v", boolArray)
}

// ReadCoils
// address: 5
// count: 13
func TestDecodeMB002(t *testing.T) {
	strVal := "07B10000000601010005000D"
	byteVal, byteErr := hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	header, protocol, errCode := DecodeMBProtocol(byteVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
		return
	}

	if protocol.FuncCode() != ReadCoils {
		t.Errorf("decode ReadCoils request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBReadCoilsReq)
	if !reqOK {
		t.Errorf("decode ReadCoils request failed")
		return
	}
	if reqPtr.Address() != 5 {
		t.Errorf("decode ReadCoils request address failed")
		return
	}
	if reqPtr.Count() != 13 {
		t.Errorf("decode ReadCoils request count failed")
		return
	}

	strVal = "07B1000000050101024718"
	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	header, protocol, errCode = DecodeMBProtocol(byteVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
		return
	}

	if protocol.FuncCode() != ReadCoils {
		t.Errorf("decode ReadCoils request failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBReadCoilsRsp)
	if !rspOK {
		t.Errorf("decode ReadCoils response failed")
		return
	}
	if rspPtr.Count() != 2 {
		t.Errorf("decode ReadCoils response count failed")
		return
	}

	// 0,1,2,6,11,12 = true
	// other = false
	trueSet := []int{0, 1, 2, 6, 11, 12}
	boolArray := ByteArrayToBoolArray(rspPtr.Data())
	for idx := range boolArray {
		findFlag := false
		for _, val := range trueSet {
			if idx == val {
				findFlag = true
				break
			}
		}

		if boolArray[idx] != findFlag {
			t.Errorf("byte to bool failed")
		}
	}

	t.Logf("%v", boolArray)
}

// ReadDiscreteInputs
// address: 0
// count: 13
func TestDecodeMB003(t *testing.T) {
	strVal := "09950000000601020000000D"
	byteVal, byteErr := hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	header, protocol, errCode := DecodeMBProtocol(byteVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
		return
	}

	if protocol.FuncCode() != ReadDiscreteInputs {
		t.Errorf("decode ReadDiscreteInputs request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBReadDiscreteInputsReq)
	if !reqOK {
		t.Errorf("decode ReadDiscreteInputs request failed")
		return
	}
	if reqPtr.Address() != 0 {
		t.Errorf("decode ReadDiscreteInputs request address failed")
		return
	}
	if reqPtr.Count() != 13 {
		t.Errorf("decode ReadDiscreteInputs request count failed")
		return
	}

	strVal = "099500000005010202E313"
	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	header, protocol, errCode = DecodeMBProtocol(byteVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
		return
	}

	if protocol.FuncCode() != ReadDiscreteInputs {
		t.Errorf("decode ReadDiscreteInputs request failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBReadDiscreteInputsRsp)
	if !rspOK {
		t.Errorf("decode ReadDiscreteInputs response failed")
		return
	}
	if rspPtr.Count() != 2 {
		t.Errorf("decode ReadDiscreteInputs response count failed")
		return
	}

	// 0, 1, 5, 6, 7, 8, 9, 12 = true
	// other = false
	trueSet := []int{0, 1, 5, 6, 7, 8, 9, 12}
	boolArray := ByteArrayToBoolArray(rspPtr.Data())
	for idx := range boolArray {
		findFlag := false
		for _, val := range trueSet {
			if idx == val {
				findFlag = true
				break
			}
		}

		if boolArray[idx] != findFlag {
			t.Errorf("byte to bool failed, idx:%v", idx)
		}
	}

	t.Logf("%v", boolArray)
}

// ReadDiscreteInputs
// address: 0
// count: 25
func TestDecodeMB004(t *testing.T) {
	strVal := "0C0F00000006010200000019"
	byteVal, byteErr := hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	header, protocol, errCode := DecodeMBProtocol(byteVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
		return
	}

	if protocol.FuncCode() != ReadDiscreteInputs {
		t.Errorf("decode ReadDiscreteInputs request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBReadDiscreteInputsReq)
	if !reqOK {
		t.Errorf("decode ReadDiscreteInputs request failed")
		return
	}
	if reqPtr.Address() != 0 {
		t.Errorf("decode ReadDiscreteInputs request address failed")
		return
	}
	if reqPtr.Count() != 25 {
		t.Errorf("decode ReadDiscreteInputs request count failed")
		return
	}

	strVal = "0C0F0000000701020487E31800"
	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	header, protocol, errCode = DecodeMBProtocol(byteVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
		return
	}

	if protocol.FuncCode() != ReadDiscreteInputs {
		t.Errorf("decode ReadDiscreteInputs request failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBReadDiscreteInputsRsp)
	if !rspOK {
		t.Errorf("decode ReadDiscreteInputs response failed")
		return
	}
	if rspPtr.Count() != 4 {
		t.Errorf("decode ReadDiscreteInputs response count failed")
		return
	}

	// 0, 1, 2, 7, 8, 9, 13, 14, 15, 19, 20 = true
	// other = false
	trueSet := []int{0, 1, 2, 7, 8, 9, 13, 14, 15, 19, 20}
	boolArray := ByteArrayToBoolArray(rspPtr.Data())
	for idx := range boolArray {
		findFlag := false
		for _, val := range trueSet {
			if idx == val {
				findFlag = true
				break
			}
		}

		if boolArray[idx] != findFlag {
			t.Errorf("byte to bool failed, idx:%v", idx)
		}
	}

	t.Logf("%v", boolArray)
}

// ReadDiscreteInputs
// address: 0
// count: 13
func TestDecodeMB005(t *testing.T) {
	strVal := "10400000000601030000000D"
	byteVal, byteErr := hex.DecodeString(strVal)
	if byteErr != nil {
		t.Errorf("hex.DecodeString, error:%v", byteErr.Error())
		return
	}

	header, protocol, errCode := DecodeMBProtocol(byteVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
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

	header, protocol, errCode = DecodeMBProtocol(byteVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
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
	if rspPtr.Count() != 13*2 {
		t.Errorf("decode ReadHoldingRegisters response count failed")
		return
	}

	u16Array, u16Err := ByteArrayToUint16ABArray(rspPtr.Data())
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

	header, protocol, errCode := DecodeMBProtocol(byteVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
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

	header, protocol, errCode = DecodeMBProtocol(byteVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
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
	if rspPtr.Count() != 13*2 {
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

	header, protocol, errCode := DecodeMBProtocol(byteVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
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

	header, protocol, errCode = DecodeMBProtocol(byteVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
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

	header, protocol, errCode := DecodeMBProtocol(byteVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
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
	if reqPtr.DataSize() != 2 {
		t.Errorf("decode WriteMultipleRegisters request dataSize failed")
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

	header, protocol, errCode = DecodeMBProtocol(byteVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
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

	header, protocol, errCode := DecodeMBProtocol(byteVal, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
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
	if reqPtr.DataSize() != 20 {
		t.Errorf("decode WriteMultipleRegisters request dataSize failed")
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

	header, protocol, errCode = DecodeMBProtocol(byteVal, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
		return
	}

	if header.Length() != aduTcpHeadLength {
		t.Errorf("decode mb header failed")
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
