package model

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
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
	byteVal, resultCode = req.Encode(byteVal)
	if resultCode != SuccessCode {
		t.Error("Encode MBReadCoilsReq failed")
		return fmt.Errorf("%v", "Encode MBReadCoilsReq failed")
	}

	if len(byteVal) != int(req.Length()) {
		t.Error("Encode MBReadCoilsReq failed, missmatch length")
		return fmt.Errorf("%v", "Encode MBReadCoilsReq failed, missmatch length")
	}

	strVal := hex.EncodeToString(byteVal)
	if strings.ToUpper(strVal) != strReq {
		t.Error("Encode MBReadCoilsReq failed")
		return fmt.Errorf("%v", "Encode MBReadCoilsReq failed")
	}

	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Error("illegal byteVal")
		return byteErr
	}

	reqNew := &MBReadCoilsReq{}
	resultCode = reqNew.Decode(byteVal)
	if resultCode != SuccessCode {
		t.Error("Decode MBReadCoilsReq failed")
		return fmt.Errorf("%v", "Decode MBReadCoilsReq failed")
	}
	if reqNew.FuncCode() != ReadCoils {
		t.Error("missmatch func code")
		return fmt.Errorf("%v", "missmatch func code")
	}
	if req.count != reqNew.count || req.address != reqNew.address {
		t.Error("missmatch address or address")
		return fmt.Errorf("%v", "missmatch address or address")
	}

	return nil
}

func TestMBReadCoilsRsp(t *testing.T) {
	rsp := &MBReadCoilsRsp{
		count: byte(3),
		data:  []byte{0x15, 0x00, 0x01},
	}

	if checkCoilsRsp(t, rsp, "0103150001") != nil {
		return
	}

	rsp = &MBReadCoilsRsp{
		count: byte(7),
		data:  []byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x80, 0x03},
	}

	if checkCoilsRsp(t, rsp, "010703000000008003") != nil {
		return
	}
}

func checkCoilsRsp(t *testing.T, rsp *MBReadCoilsRsp, strRsp string) error {
	var byteVal []byte
	var byteErr error
	var resultCode byte
	byteVal, resultCode = rsp.Encode(byteVal)
	if resultCode != SuccessCode {
		t.Error("Encode MBReadCoilsRsp failed")
		return fmt.Errorf("%v", "Encode MBReadCoilsRsp failed")
	}
	if len(byteVal) != int(rsp.Length()) {
		t.Error("Encode MBReadCoilsRsp failed, missmatch length")
		return fmt.Errorf("%v", "Encode MBReadCoilsRsp failed, missmatch length")
	}

	strVal := hex.EncodeToString(byteVal)
	if strings.ToUpper(strVal) != strRsp {
		t.Error("Encode MBReadCoilsRsp failed")
		return fmt.Errorf("%v", "Encode MBReadCoilsRsp failed")
	}

	byteVal, byteErr = hex.DecodeString(strVal)
	if byteErr != nil {
		t.Error("illegal byteVal")
		return byteErr
	}

	rspNew := &MBReadCoilsRsp{}
	resultCode = rspNew.Decode(byteVal)
	if resultCode != SuccessCode {
		t.Error("Decode MBReadCoilsRsp failed")
		return fmt.Errorf("%v", "Decode MBReadCoilsRsp failed")
	}
	if rspNew.FuncCode() != ReadCoils {
		t.Error("missmatch func code")
		return fmt.Errorf("%v", "missmatch func code")
	}
	if rsp.count != rspNew.count {
		t.Error("missmatch address")
		return fmt.Errorf("%v", "missmatch address")
	}

	if bytes.Compare(rsp.data, rspNew.data) != 0 {
		t.Error("missmatch data")
		return fmt.Errorf("%v", "missmatch date")
	}

	return nil
}
