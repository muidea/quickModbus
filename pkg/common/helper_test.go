package common

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestByteToBoolArray(t *testing.T) {
	// 00000001
	b01Val := byte(0x01)
	// 00010101
	b21Val := byte(0x15)

	boolSlice := []bool{true, false, true, false, true, false, false, false}
	boolArray := byteToBoolArray(b21Val)
	t.Logf("expect:%v", boolSlice)
	t.Logf("really:%v", boolArray)
	for idx := 0; idx < len(boolSlice); idx++ {
		if boolSlice[idx] != boolArray[idx] {
			t.Error("ByteToBoolArrayDCBA failed")
			return
		}
	}

	newB21Val := boolArrayToByte(boolSlice)
	if b21Val != newB21Val {
		return
	}

	// 0x01,0x15
	// [true false false false false false false false true false true false false true false false]
	boolSlice = []bool{true, false, false, false, false, false, false, false, true, false, true, false, true, false, false, false}
	boolArray = bytesToBoolArray([]byte{b01Val, b21Val})
	t.Logf("expect:%v", boolSlice)
	t.Logf("really:%v", boolArray)
	for idx := 0; idx < len(boolSlice); idx++ {
		if boolSlice[idx] != boolArray[idx] {
			t.Error("ByteToBoolArrayDCBA failed")
			return
		}
	}

	byteArray := boolArrayToByteArray(boolArray)
	if len(byteArray) != 2 {
		t.Errorf("boolArrayToByteArray failed")
		return
	}
	if byteArray[0] != b01Val || byteArray[1] != b21Val {
		t.Errorf("boolArrayToByteArray failed, missmatch byte value")
		return
	}

	boolSlice = []bool{true, false, true, false, true, false, false, false}
	boolArray = byteToBoolArray(b21Val)
	t.Logf("expect:%v", boolSlice)
	t.Logf("really:%v", boolArray)
	for idx := 0; idx < len(boolSlice); idx++ {
		if boolSlice[idx] != boolArray[idx] {
			t.Error("ByteToBoolArrayDCBA failed")
			return
		}
	}
	newB21Val = boolArrayToByte(boolSlice)
	if b21Val != newB21Val {
		return
	}

	boolSlice = []bool{true, false, true, false, true, false, false, false, true, false, false, false, false, false, false, false}
	boolArray = bytesToBoolArray([]byte{b21Val, b01Val})
	t.Logf("expect:%v", boolSlice)
	t.Logf("really:%v", boolArray)
	for idx := 0; idx < len(boolSlice); idx++ {
		if boolSlice[idx] != boolArray[idx] {
			t.Error("ByteToBoolArrayDCBA failed")
			return
		}
	}
}

func TestUint16(t *testing.T) {
	uVal1 := uint16(0x0102)
	uVal2 := uint16(0x0304)

	byteVal := []byte{}
	var byteErr error

	byteVal, byteErr = AppendUint16(byteVal, uVal1, ABCDEndian)
	if byteErr != nil {
		t.Errorf("AppendUint16 failed, error:%s", byteErr.Error())
		return
	}

	byteVal, byteErr = AppendUint16(byteVal, uVal2, ABCDEndian)
	if byteErr != nil {
		t.Errorf("AppendUint16 failed, error:%s", byteErr.Error())
		return
	}

	u16Val, u16Err := BytesToUint16Array(byteVal, ABCDEndian)
	if u16Err != nil {
		t.Errorf("BytesToUint16Array failed, error:%s", byteErr.Error())
		return
	}

	if len(u16Val) != 2 {
		t.Errorf("BytesToUint16Array failed, missmatch len")
		return
	}
	if u16Val[0] != uVal1 || u16Val[1] != uVal2 {
		t.Errorf("BytesToUint16Array failed, missmatch item value")
		return
	}

	u16Val, u16Err = BytesToUint16Array(byteVal, CDABEndian)
	if u16Err != nil {
		t.Errorf("BytesToUint16Array failed, error:%s", byteErr.Error())
		return
	}

	if len(u16Val) != 2 {
		t.Errorf("BytesToUint16Array failed, missmatch len")
		return
	}
	if u16Val[1] != uVal1 || u16Val[0] != uVal2 {
		t.Errorf("BytesToUint16Array failed, missmatch item value")
		return
	}
}

func TestUint32(t *testing.T) {
	uVal1 := uint32(0x01020304)
	uVal2 := uint32(0x03040506)

	byteVal := []byte{}
	var byteErr error

	byteVal, byteErr = AppendUint32(byteVal, uVal1, ABCDEndian)
	if byteErr != nil {
		t.Errorf("AppendUint32 failed, error:%s", byteErr.Error())
		return
	}
	byteVal, byteErr = AppendUint32(byteVal, uVal2, ABCDEndian)
	if byteErr != nil {
		t.Errorf("AppendUint32 failed, error:%s", byteErr.Error())
		return
	}

	u32Val, u32Err := BytesToUint32Array(byteVal, ABCDEndian)
	if u32Err != nil {
		t.Errorf("BytesToUint32Array failed, error:%s", byteErr.Error())
		return
	}

	if len(u32Val) != 2 {
		t.Errorf("BytesToUint32Array failed, missmatch len")
		return
	}
	if u32Val[0] != uVal1 || u32Val[1] != uVal2 {
		t.Errorf("BytesToUint32Array failed, missmatch item value")
		return
	}

	byteVal = []byte{}
	byteVal, byteErr = AppendUint32(byteVal, uVal1, BADCEndian)
	if byteErr != nil {
		t.Errorf("AppendUint32 failed, error:%s", byteErr.Error())
		return
	}
	byteVal, byteErr = AppendUint32(byteVal, uVal2, BADCEndian)
	if byteErr != nil {
		t.Errorf("AppendUint32 failed, error:%s", byteErr.Error())
		return
	}

	u32Val, u32Err = BytesToUint32Array(byteVal, BADCEndian)
	if u32Err != nil {
		t.Errorf("BytesToUint32Array failed, error:%s", byteErr.Error())
		return
	}

	if len(u32Val) != 2 {
		t.Errorf("BytesToUint32Array failed, missmatch len")
		return
	}
	if u32Val[0] != uVal1 || u32Val[1] != uVal2 {
		t.Errorf("BytesToUint32Array failed, missmatch item value")
		return
	}
}

func TestSwapArray(t *testing.T) {
	rawStr := "00000000C000405E"
	byteVal, byteErr := hex.DecodeString(rawStr)
	if byteErr != nil {
		t.Errorf("hex.DecodeString failed, error:%s", byteErr.Error())
		return
	}

	valStr := strings.ToUpper(hex.EncodeToString(byteVal))
	if rawStr != valStr {
		t.Errorf("hex.EncodeToString failed")
		return
	}

	cdabStr := "00000000405EC000"
	swapVal, _ := swapArray(byteVal, CDABEndian)
	valStr = strings.ToUpper(hex.EncodeToString(swapVal))
	if cdabStr != valStr {
		t.Errorf("hex.EncodeToString failed")
		return
	}

	badcStr := "0000000000C05E40"
	swapVal, _ = swapArray(byteVal, BADCEndian)
	valStr = strings.ToUpper(hex.EncodeToString(swapVal))
	if badcStr != valStr {
		t.Errorf("hex.EncodeToString failed")
		return
	}
}
