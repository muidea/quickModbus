package model

import "testing"

func TestByteToBoolArray(t *testing.T) {
	// 00000001
	b01Val := byte(0x01)
	// 00010101
	b21Val := byte(0x15)

	boolSlice := []bool{true, false, true, false, true, false, false, false}
	boolArray := ByteToBoolArrayDCBA(b21Val)
	t.Logf("expect:%v", boolSlice)
	t.Logf("really:%v", boolArray)
	for idx := 0; idx < len(boolSlice); idx++ {
		if boolSlice[idx] != boolArray[idx] {
			t.Error("ByteToBoolArrayDCBA failed")
			return
		}
	}

	newB21Val := BoolArrayToByteABCD(boolSlice)
	if b21Val != newB21Val {
		return
	}

	boolSlice = []bool{false, false, false, false, false, false, false, true, false, false, false, true, false, true, false, true}
	boolArray = ByteArrayToBoolArrayDCBA([]byte{b01Val, b21Val})
	t.Logf("expect:%v", boolSlice)
	t.Logf("really:%v", boolArray)
	for idx := 0; idx < len(boolSlice); idx++ {
		if boolSlice[idx] != boolArray[idx] {
			t.Error("ByteToBoolArrayDCBA failed")
			return
		}
	}

	byteArray := BoolArrayToByteArrayABCD(boolArray)
	if len(byteArray) != 2 {
		t.Errorf("BoolArrayToByteArrayABCD failed")
		return
	}
	if byteArray[0] != b01Val || byteArray[1] != b21Val {
		t.Errorf("BoolArrayToByteArrayABCD failed, missmatch byte value")
		return
	}

	boolSlice = []bool{true, false, true, false, true, false, false, false}
	boolArray = ByteToBoolArrayABCD(b21Val)
	t.Logf("expect:%v", boolSlice)
	t.Logf("really:%v", boolArray)
	for idx := 0; idx < len(boolSlice); idx++ {
		if boolSlice[idx] != boolArray[idx] {
			t.Error("ByteToBoolArrayDCBA failed")
			return
		}
	}
	newB21Val = BoolArrayToByteDCBA(boolSlice)
	if b21Val != newB21Val {
		return
	}

	boolSlice = []bool{true, false, true, false, true, false, false, false, true, false, false, false, false, false, false, false}
	boolArray = ByteArrayToBoolArrayABCD([]byte{b21Val, b01Val})
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

	byteVal = AppendUint16ABCD(byteVal, uVal1)
	nVal1 := ByteToUint16ABCD(byteVal[:])
	if nVal1 != uVal1 {
		t.Error("Encode byte failed")
		return
	}
	byteVal = AppendUint16ABCD(byteVal, uVal2)
	nVal1 = ByteToUint16ABCD(byteVal[:2])
	if nVal1 != uVal1 {
		t.Error("Encode byte failed")
		return
	}
	nVal2 := ByteToUint16ABCD(byteVal[2:])
	if nVal2 != uVal2 {
		t.Error("Encode byte failed")
		return
	}
}

func TestUint32(t *testing.T) {
	uVal1 := uint32(0x01020304)
	uVal2 := uint32(0x03040506)

	byteVal := []byte{}

	byteVal = AppendUint32ABCD(byteVal, uVal1)
	nVal1 := ByteToUint32ABCD(byteVal[:])
	if nVal1 != uVal1 {
		t.Error("Encode byte failed")
		return
	}
	byteVal = AppendUint32ABCD(byteVal, uVal2)
	nVal1 = ByteToUint32ABCD(byteVal[:4])
	if nVal1 != uVal1 {
		t.Error("Encode byte failed")
		return
	}
	nVal2 := ByteToUint32ABCD(byteVal[4:])
	if nVal2 != uVal2 {
		t.Error("Encode byte failed")
		return
	}
}
