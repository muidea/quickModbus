package model

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/muidea/quickModbus/pkg/common"
)

func TestMBReadCoilsReq(t *testing.T) {
	req := &MBReadCoilsReq{
		address: 0,
		count:   10,
	}

	if checkCoilsReq(t, req, "010000000A") != nil {
		return
	}

	req = &MBReadCoilsReq{
		address: 2,
		count:   18,
	}

	if checkCoilsReq(t, req, "0100020012") != nil {
		return
	}
}

func checkCoilsReq(t *testing.T, req *MBReadCoilsReq, strReq string) error {
	var byteVal []byte
	var byteErr error
	var resultCode byte
	byteBuffer := bytes.NewBuffer(byteVal)
	resultCode = req.Encode(byteBuffer)
	if resultCode != SuccessCode {
		t.Error("Encode MBReadCoilsReq failed")
		return fmt.Errorf("%v", "Encode MBReadCoilsReq failed")
	}

	strVal := hex.EncodeToString(byteBuffer.Bytes())
	if strings.ToUpper(strVal) != strReq {
		t.Error("Encode MBReadCoilsReq failed")
		return fmt.Errorf("%v", "Encode MBReadCoilsReq failed")
	}

	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Error("illegal byteVal")
		return byteErr
	}

	byteBuffer = bytes.NewBuffer(byteVal)
	reqNew := &MBReadCoilsReq{}
	resultCode = reqNew.Decode(byteBuffer)
	if resultCode != SuccessCode {
		t.Error("Decode MBReadCoilsReq failed")
		return fmt.Errorf("%v", "Decode MBReadCoilsReq failed")
	}
	if reqNew.FuncCode() != ReadCoils {
		t.Error("mismatch func code")
		return fmt.Errorf("%v", "missmatch func code")
	}
	if req.count != reqNew.count || req.address != reqNew.address {
		t.Error("mismatch address or address")
		return fmt.Errorf("%v", "missmatch address or address")
	}

	return nil
}

func TestMBReadCoilsRsp(t *testing.T) {
	rsp := &MBReadCoilsRsp{
		data: []byte{0x15, 0x00, 0x01},
	}

	if checkCoilsRsp(t, rsp, "0103150001") != nil {
		return
	}

	rsp = &MBReadCoilsRsp{
		data: []byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x80, 0x03},
	}

	if checkCoilsRsp(t, rsp, "010703000000008003") != nil {
		return
	}
}

func checkCoilsRsp(t *testing.T, rsp *MBReadCoilsRsp, strRsp string) error {
	var byteVal []byte
	var byteErr error
	var resultCode byte
	byteBuff := bytes.NewBuffer(byteVal)
	resultCode = rsp.Encode(byteBuff)
	if resultCode != SuccessCode {
		t.Error("Encode MBReadCoilsRsp failed")
		return fmt.Errorf("%v", "Encode MBReadCoilsRsp failed")
	}

	strVal := hex.EncodeToString(byteBuff.Bytes())
	if strings.ToUpper(strVal) != strRsp {
		t.Error("Encode MBReadCoilsRsp failed")
		return fmt.Errorf("%v", "Encode MBReadCoilsRsp failed")
	}

	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Error("illegal byteVal")
		return byteErr
	}

	byteBuff = bytes.NewBuffer(byteVal)
	rspNew := &MBReadCoilsRsp{}
	resultCode = rspNew.Decode(byteBuff)
	if resultCode != SuccessCode {
		t.Error("Decode MBReadCoilsRsp failed")
		return fmt.Errorf("%v", "Decode MBReadCoilsRsp failed")
	}
	if rspNew.FuncCode() != ReadCoils {
		t.Error("mismatch func code")
		return fmt.Errorf("%v", "missmatch func code")
	}

	if bytes.Compare(rsp.data, rspNew.data) != 0 {
		t.Error("mismatch data")
		return fmt.Errorf("%v", "missmatch date")
	}

	return nil
}

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

	byteBuff := bytes.NewBuffer(byteVal)
	_, protocol, errCode := DecodeMBProtocol(byteBuff, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
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

	byteBuff = bytes.NewBuffer(byteVal)
	_, protocol, errCode = DecodeMBProtocol(byteBuff, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
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

	// 0,6,7,8,9 = true
	// other = false
	trueSet := []int{0, 6, 7, 8, 9}
	boolArray, err := common.BytesToBoolArray(rspPtr.Data())
	if err != nil {
		t.Errorf("common.BytesToBoolArray failed, err:%s", err.Error())
		return
	}

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

	byteBuff := bytes.NewBuffer(byteVal)
	_, protocol, errCode := DecodeMBProtocol(byteBuff, RequestAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
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

	byteBuff = bytes.NewBuffer(byteVal)
	_, protocol, errCode = DecodeMBProtocol(byteBuff, ResponseAction)
	if errCode != SuccessCode {
		t.Errorf("DecodeMBProtocol failed, error code :%v", errCode)
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

	// 0,1,2,6,11,12 = true
	// other = false
	trueSet := []int{0, 1, 2, 6, 11, 12}
	boolArray, err := common.BytesToBoolArray(rspPtr.Data())
	if err != nil {
		t.Errorf("common.BytesToBoolArray failed, err:%s", err.Error())
		return
	}
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

// WriteSingleCoil
// address: 2
// value: true
func TestDecodeMB010(t *testing.T) {
	strVal := "3CA40000000601050002FF00"
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

	if protocol.FuncCode() != WriteSingleCoil {
		t.Errorf("decode WriteSingleCoil request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBWriteSingleCoilReq)
	if !reqOK {
		t.Errorf("decode WriteSingleCoil request failed")
		return
	}
	if reqPtr.Address() != 2 {
		t.Errorf("decode WriteSingleCoil request address failed")
		return
	}
	if len(reqPtr.Data()) != 2 {
		t.Errorf("decode WriteSingleCoil request data count failed")
		return
	}

	if bytes.Compare(reqPtr.Data(), CoilON) != 0 {
		t.Errorf("decode WriteSingleCoil request data failed")
		return
	}

	strVal = "3CA40000000601050002FF00"
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

	if protocol.FuncCode() != WriteSingleCoil {
		t.Errorf("decode WriteSingleCoil response failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBWriteSingleCoilRsp)
	if !rspOK {
		t.Errorf("decode WriteSingleCoil response failed")
		return
	}
	if len(rspPtr.Data()) != 2 {
		t.Errorf("decode WriteSingleRegister response data len failed")
		return
	}

	if bytes.Compare(rspPtr.Data(), CoilON) != 0 {
		t.Errorf("decode WriteSingleCoil response data failed")
		return
	}
}

// WriteMultipleCoils
// address: 0
// value: (0,2,4,5,6)true
func TestDecodeMB011(t *testing.T) {
	strVal := "076600000009010F0000000A027500"
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

	if protocol.FuncCode() != WriteMultipleCoils {
		t.Errorf("decode WriteMultipleCoils request failed")
		return
	}

	reqPtr, reqOK := protocol.(*MBWriteMultipleCoilsReq)
	if !reqOK {
		t.Errorf("decode WriteMultipleCoils request failed")
		return
	}
	if reqPtr.Address() != 0 {
		t.Errorf("decode WriteMultipleCoils request address failed")
		return
	}
	if reqPtr.Count() != 10 {
		t.Errorf("decode WriteMultipleCoils request count failed")
		return
	}
	if len(reqPtr.Data()) != 2 {
		t.Errorf("decode WriteMultipleCoils request data count failed")
		return
	}

	valSet := []bool{true, false, true, false, true, true, true, false, false, false}
	boolArray, err := common.BytesToBoolArray(reqPtr.Data())
	if err != nil {
		t.Errorf("common.BytesToBoolArray failed, err:%s", err.Error())
		return
	}

	idx := uint16(0)
	for ; idx < reqPtr.Count(); idx++ {
		if valSet[idx] != boolArray[idx] {
			t.Errorf("decode WriteMultipleCoils request data count failed")
		}
	}
	for idx < uint16(len(boolArray)) {
		if boolArray[idx] {
			t.Errorf("decode WriteMultipleCoils request data count failed")
		}

		idx++
	}

	strVal = "076600000006010F0000000A"
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

	if protocol.FuncCode() != WriteMultipleCoils {
		t.Errorf("decode WriteMultipleCoils response failed")
		return
	}

	rspPtr, rspOK := protocol.(*MBWriteMultipleCoilsRsp)
	if !rspOK {
		t.Errorf("decode WriteMultipleCoils response failed")
		return
	}
	if rspPtr.Address() != 0 {
		t.Errorf("decode WriteMultipleCoils response address failed")
		return
	}

	if rspPtr.Count() != 10 {
		t.Errorf("decode WriteMultipleCoils response data count failed")
		return
	}
}
