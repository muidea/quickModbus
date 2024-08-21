package model

import (
	"bytes"
	"encoding/hex"
	"testing"
)

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

	byteBuff := bytes.NewBuffer(byteVal)
	_, protocol, errCode := DecodeMBProtocol(byteBuff, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
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

	byteBuff = bytes.NewBuffer(byteVal)
	_, protocol, errCode = DecodeMBProtocol(byteBuff, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
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

	// 0, 1, 5, 6, 7, 8, 9, 12 = true
	// other = false
	trueSet := []int{0, 1, 5, 6, 7, 8, 9, 12}
	boolArray := ByteArrayToBoolArrayDCBA(rspPtr.Data())
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

	byteBuff := bytes.NewBuffer(byteVal)
	_, protocol, errCode := DecodeMBProtocol(byteBuff, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
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

	byteBuff = bytes.NewBuffer(byteVal)
	_, protocol, errCode = DecodeMBProtocol(byteBuff, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
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

	// 0, 1, 2, 7, 8, 9, 13, 14, 15, 19, 20 = true
	// other = false
	trueSet := []int{0, 1, 2, 7, 8, 9, 13, 14, 15, 19, 20}
	boolArray := ByteArrayToBoolArrayDCBA(rspPtr.Data())
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
