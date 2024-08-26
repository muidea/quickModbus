package common

import (
	"fmt"
	"github.com/muidea/magicCommon/foundation/log"
	"math"
)

func ConvertFloat64To(value float64, valueType uint16) (ret any, err error) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = fmt.Errorf("ConvertFloat64To valueType %v failed, %v", valueType, errInfo)
		}
	}()

	switch valueType {
	case Int16Value:
		ret = int16(value)
	case UInt16Value:
		ret = uint16(value)
	case Int32Value:
		ret = int32(value)
	case UInt32Value:
		ret = uint32(value)
	case Int64Value:
		ret = int64(value)
	case UInt64Value:
		ret = uint64(value)
	case Float32Value:
		ret = float32(value)
	case Float64Value:
		ret = value
	default:
		err = fmt.Errorf("illegal valueType")
	}

	return
}

func swapArrayFor64Bits[T any](valArray []T, endianType uint16) (ret []T, err error) {
	if len(valArray) < 8 {
		ret = valArray
		return
	}

	swappedVal := make([]T, len(valArray))
	copy(swappedVal, valArray)

	switch endianType {
	case ABCDEndian, BADCEndian:
	// No change needed for abcd,badc
	case CDABEndian, DCBAEndian:
		for i := 0; i+7 < len(valArray); i += 8 {
			swappedVal[i], swappedVal[i+1], swappedVal[i+2], swappedVal[i+3], swappedVal[i+4], swappedVal[i+5], swappedVal[i+6], swappedVal[i+7] =
				valArray[i+4], valArray[i+5], valArray[i+6], valArray[i+7], valArray[i], valArray[i+1], valArray[i+2], valArray[i+3]
		}
	default:
		errMsg := fmt.Sprintf("illegal endianType, endianType:%v", endianType)
		err = fmt.Errorf(errMsg)
		log.Errorf("swapArrayFor64Bits failed, error:%s", errMsg)
	}

	if err != nil {
		return
	}

	ret = swappedVal
	return
}

func swapArray[T any](valArray []T, endianType uint16) (ret []T, err error) {
	if len(valArray) < 4 {
		ret = valArray
		return
	}

	swappedVal := make([]T, len(valArray))
	copy(swappedVal, valArray)

	switch endianType {
	case ABCDEndian:
		// No change needed for abcd
	case BADCEndian:
		for i := 0; i+3 < len(valArray); i += 4 {
			swappedVal[i], swappedVal[i+1], swappedVal[i+2], swappedVal[i+3] = valArray[i+1], valArray[i], valArray[i+3], valArray[i+2]
		}
	case CDABEndian:
		for i := 0; i+3 < len(valArray); i += 4 {
			swappedVal[i], swappedVal[i+1], swappedVal[i+2], swappedVal[i+3] = valArray[i+2], valArray[i+3], valArray[i], valArray[i+1]
		}
	case DCBAEndian:
		for i := 0; i+3 < len(valArray); i += 4 {
			swappedVal[i], swappedVal[i+1], swappedVal[i+2], swappedVal[i+3] = valArray[i+3], valArray[i+2], valArray[i+1], valArray[i]
		}
	default:
		errMsg := fmt.Sprintf("illegal endianType, endianType:%v", endianType)
		err = fmt.Errorf(errMsg)
		log.Errorf("swapArray failed, error:%s", errMsg)
	}

	if err != nil {
		return
	}

	ret = swappedVal
	return
}

func checkArray[T any](arrayVal []T, lenVal int) []T {
	arrayLen := len(arrayVal)
	if arrayLen < lenVal {
		padding := make([]T, lenVal-arrayLen)
		arrayVal = append(padding, arrayVal...)
	} else if arrayLen > lenVal {
		arrayVal = arrayVal[:lenVal]
	}

	return arrayVal
}

func BytesToBoolArray(byteVal []byte, endianType uint16) (ret []bool, err error) {
	ret = bytesToBoolArray(byteVal)
	return
}

func AppendBoolArray(byteVal []byte, boolVal []bool, endianType uint16) (ret []byte, err error) {
	bytes := boolArrayToByteArray(boolVal)
	ret = append(byteVal, bytes...)
	return
}

func BytesToUint16(byteVal []byte, endianType uint16) (ret uint16, err error) {
	ret, err = bytesToUint16(byteVal)
	return
}

func BytesToUint16Array(byteVal []byte, endianType uint16) (ret []uint16, err error) {
	for idx := 0; idx < len(byteVal); idx += 2 {
		uVal, uErr := bytesToUint16(byteVal[idx : idx+2])
		if uErr != nil {
			err = uErr
			return
		}

		ret = append(ret, uVal)
	}

	return
}

func AppendUint16(byteVal []byte, uVal, endianType uint16) (ret []byte, err error) {
	bytes := uint16ToByteArray(uVal)
	ret = append(byteVal, bytes...)
	return
}

func BytesToInt16(byteVal []byte, endianType uint16) (ret int16, err error) {
	ret, err = bytesToInt16(byteVal)
	return
}

func BytesToInt16Array(byteVal []byte, endianType uint16) (ret []int16, err error) {
	for idx := 0; idx < len(byteVal); idx += 2 {
		iVal, iErr := bytesToInt16(byteVal[idx : idx+2])
		if iErr != nil {
			err = iErr
			return
		}

		ret = append(ret, iVal)
	}

	return
}

func AppendInt16(byteVal []byte, iVal int16, endianType uint16) (ret []byte, err error) {
	bytes := int16ToByteArray(iVal)
	ret = append(byteVal, bytes...)
	return
}

func BytesToUint32(byteVal []byte, endianType uint16) (ret uint32, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToUint32 failed, error:%s", err.Error())
		return
	}

	ret, err = bytesToUint32(byteVal)
	return
}

func BytesToUint32Array(byteVal []byte, endianType uint16) (ret []uint32, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToUint32Array failed, error:%s", err.Error())
		return
	}
	for idx := 0; idx < len(byteVal); idx += 4 {
		uVal, uErr := bytesToUint32(byteVal[idx : idx+4])
		if uErr != nil {
			err = uErr
			return
		}

		ret = append(ret, uVal)
	}

	return
}

func AppendUint32(byteVal []byte, uVal uint32, endianType uint16) (ret []byte, err error) {
	bytes := uint32ToByteArray(uVal)
	bytes, err = swapArray(bytes, endianType)
	if err != nil {
		log.Errorf("AppendUint32 failed, error:%s", err.Error())
		return
	}

	ret = append(byteVal, bytes...)
	return
}

func BytesToInt32(byteVal []byte, endianType uint16) (ret int32, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToInt32 failed, error:%s", err.Error())
		return
	}

	ret, err = bytesToInt32(byteVal)
	return
}

func BytesToInt32Array(byteVal []byte, endianType uint16) (ret []int32, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToInt32Array failed, error:%s", err.Error())
		return
	}
	for idx := 0; idx < len(byteVal); idx += 4 {
		iVal, iErr := bytesToInt32(byteVal[idx : idx+4])
		if iErr != nil {
			err = iErr
			return
		}

		ret = append(ret, iVal)
	}

	return
}

func AppendInt32(byteVal []byte, iVal int32, endianType uint16) (ret []byte, err error) {
	bytes := int32ToByteArray(iVal)
	bytes, err = swapArray(bytes, endianType)
	if err != nil {
		log.Errorf("AppendInt32 failed, error:%s", err.Error())
		return
	}

	ret = append(byteVal, bytes...)
	return
}

func BytesToUint64(byteVal []byte, endianType uint16) (ret uint64, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToUint64 failed, swapArray error:%s", err.Error())
		return
	}

	byteVal, err = swapArrayFor64Bits(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToUint64 failed, swapArrayFor64Bits error:%s", err.Error())
		return
	}

	ret, err = bytesToUint64(byteVal)
	return
}

func BytesToUint64Array(byteVal []byte, endianType uint16) (ret []uint64, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToUint64Array failed, swapArray error:%s", err.Error())
		return
	}

	byteVal, err = swapArrayFor64Bits(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToUint64Array failed, swapArrayFor64Bits error:%s", err.Error())
		return
	}

	for idx := 0; idx < len(byteVal); idx += 8 {
		uVal, uErr := bytesToUint64(byteVal[idx : idx+8])
		if uErr != nil {
			err = uErr
			return
		}

		ret = append(ret, uVal)
	}

	return
}

func AppendUint64(byteVal []byte, uVal uint64, endianType uint16) (ret []byte, err error) {
	bytes := uint64ToByteArray(uVal)
	bytes, err = swapArray(bytes, endianType)
	if err != nil {
		log.Errorf("AppendUint64 failed, swapArray error:%s", err.Error())
		return
	}
	byteVal, err = swapArrayFor64Bits(byteVal, endianType)
	if err != nil {
		log.Errorf("AppendUint64 failed, swapArrayFor64Bits error:%s", err.Error())
		return
	}

	ret = append(byteVal, bytes...)
	return
}

func BytesToInt64(byteVal []byte, endianType uint16) (ret int64, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToInt64 failed, swapArray error:%s", err.Error())
		return
	}

	byteVal, err = swapArrayFor64Bits(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToInt64 failed, swapArrayFor64Bits error:%s", err.Error())
		return
	}

	ret, err = bytesToInt64(byteVal)
	return
}

func BytesToInt64Array(byteVal []byte, endianType uint16) (ret []int64, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToInt64Array failed, error:%s", err.Error())
		return
	}
	byteVal, err = swapArrayFor64Bits(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToInt64Array failed, swapArrayFor64Bits error:%s", err.Error())
		return
	}

	for idx := 0; idx < len(byteVal); idx += 8 {
		iVal, iErr := bytesToInt64(byteVal[idx : idx+8])
		if iErr != nil {
			err = iErr
			return
		}

		ret = append(ret, iVal)
	}

	return
}

func AppendInt64(byteVal []byte, iVal int64, endianType uint16) (ret []byte, err error) {
	bytes := int64ToByteArray(iVal)
	bytes, err = swapArray(bytes, endianType)
	if err != nil {
		log.Errorf("AppendInt64 failed, swapArray error:%s", err.Error())
		return
	}
	byteVal, err = swapArrayFor64Bits(byteVal, endianType)
	if err != nil {
		log.Errorf("AppendInt64 failed, swapArrayFor64Bits error:%s", err.Error())
		return
	}

	ret = append(byteVal, bytes...)
	return
}

func BytesToFloat32(byteVal []byte, endianType uint16) (ret float32, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToFloat32 failed, error:%s", err.Error())
		return
	}

	ret, err = bytesToFloat32(byteVal)
	return
}

func BytesToFloat32Array(byteVal []byte, endianType uint16) (ret []float32, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToFloat32Array failed, error:%s", err.Error())
		return
	}
	for idx := 0; idx < len(byteVal); idx += 4 {
		fVal, rErr := bytesToFloat32(byteVal[idx : idx+4])
		if rErr != nil {
			err = rErr
			return
		}

		ret = append(ret, fVal)
	}

	return
}

func AppendFloat32(byteVal []byte, fVal float32, endianType uint16) (ret []byte, err error) {
	bytes := float32ToByteArray(fVal)
	bytes, err = swapArray(bytes, endianType)
	if err != nil {
		log.Errorf("AppendFloat32 failed, error:%s", err.Error())
		return
	}

	ret = append(byteVal, bytes...)
	return
}

func BytesToFloat64(byteVal []byte, endianType uint16) (ret float64, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToFloat64 failed, swapArray error:%s", err.Error())
		return
	}
	byteVal, err = swapArrayFor64Bits(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToFloat64 failed, swapArrayFor64Bits error:%s", err.Error())
		return
	}

	ret, err = bytesToFloat64(byteVal)
	return
}

func BytesToFloat64Array(byteVal []byte, endianType uint16) (ret []float64, err error) {
	byteVal, err = swapArray(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToFloat64Array failed, swapArray error:%s", err.Error())
		return
	}
	byteVal, err = swapArrayFor64Bits(byteVal, endianType)
	if err != nil {
		log.Errorf("BytesToFloat64Array failed, swapArrayFor64Bits error:%s", err.Error())
		return
	}

	for idx := 0; idx < len(byteVal); idx += 8 {
		fVal, rErr := bytesToFloat64(byteVal[idx : idx+8])
		if rErr != nil {
			err = rErr
			return
		}

		ret = append(ret, fVal)
	}

	return
}

func AppendFloat64(byteVal []byte, fVal float64, endianType uint16) (ret []byte, err error) {
	bytes := float64ToByteArray(fVal)
	bytes, err = swapArray(bytes, endianType)
	if err != nil {
		log.Errorf("AppendFloat64 failed, swapArray error:%s", err.Error())
		return
	}
	byteVal, err = swapArrayFor64Bits(byteVal, endianType)
	if err != nil {
		log.Errorf("AppendFloat64 failed, swapArrayFor64Bits error:%s", err.Error())
		return
	}

	ret = append(byteVal, bytes...)
	return
}

func bytesToBoolArray(byteVal []byte) []bool {
	ret := []bool{}
	for _, val := range byteVal {
		ret = append(ret, byteToBoolArray(val)...)
	}

	return ret
}

func byteToBoolArray(byteVal byte) []bool {
	boolArray := make([]bool, 8)

	for i := 7; i >= 0; i-- {
		boolArray[i] = (byteVal & (1 << i)) != 0
	}

	return boolArray
}

func boolArrayToByteArray(boolVal []bool) []byte {
	ret := []byte{}
	idx := 0
	for {
		if idx >= len(boolVal) {
			break
		}

		subArray := boolVal[idx:]
		byteVal := boolArrayToByte(subArray)
		ret = append(ret, byteVal)
		idx += 8
	}

	return ret
}

func boolArrayToByte(boolArray []bool) byte {
	const sizeVal = 8
	boolArray = checkArray(boolArray, sizeVal)

	_ = boolArray[7]
	var byteVal byte
	for i, bit := range boolArray {
		if bit {
			byteVal |= 1 << i
		}
	}

	return byteVal
}

func bytesToUint16(byteVal []byte) (uint16, error) {
	const sizeVal = 2
	byteVal = checkArray(byteVal, sizeVal)

	_ = byteVal[1]
	return uint16(byteVal[1]) |
		uint16(byteVal[0])<<8, nil
}

func uint16ToByteArray(uVal uint16) []byte {
	byteVal := []byte{}
	return append(byteVal,
		byte(uVal>>8),
		byte(uVal))
}

func bytesToInt16(byteVal []byte) (int16, error) {
	const sizeVal = 2
	byteVal = checkArray(byteVal, sizeVal)

	_ = byteVal[1]
	return int16(byteVal[1]) |
		int16(byteVal[0])<<8, nil
}

func int16ToByteArray(iVal int16) []byte {
	byteVal := []byte{}
	return append(byteVal,
		byte(iVal>>8),
		byte(iVal))
}

func bytesToUint32(byteVal []byte) (uint32, error) {
	const sizeVal = 4
	byteVal = checkArray(byteVal, sizeVal)

	_ = byteVal[3]
	return uint32(byteVal[3]) |
		uint32(byteVal[2])<<8 |
		uint32(byteVal[1])<<16 |
		uint32(byteVal[0])<<24, nil
}

func uint32ToByteArray(uVal uint32) []byte {
	byteVal := []byte{}
	return append(byteVal,
		byte(uVal>>24),
		byte(uVal>>16),
		byte(uVal>>8),
		byte(uVal),
	)
}

func bytesToInt32(byteVal []byte) (int32, error) {
	const sizeVal = 4
	byteVal = checkArray(byteVal, sizeVal)

	_ = byteVal[3]
	return int32(byteVal[3]) |
		int32(byteVal[2])<<8 |
		int32(byteVal[1])<<16 |
		int32(byteVal[0])<<24, nil
}

func int32ToByteArray(iVal int32) []byte {
	byteVal := []byte{}
	return append(byteVal,
		byte(iVal>>24),
		byte(iVal>>16),
		byte(iVal>>8),
		byte(iVal),
	)
}

func bytesToUint64(byteVal []byte) (uint64, error) {
	const sizeVal = 8
	byteVal = checkArray(byteVal, sizeVal)

	_ = byteVal[7]
	return uint64(byteVal[3]) |
		uint64(byteVal[2])<<8 |
		uint64(byteVal[1])<<16 |
		uint64(byteVal[0])<<24 |
		uint64(byteVal[7])<<32 |
		uint64(byteVal[6])<<40 |
		uint64(byteVal[5])<<48 |
		uint64(byteVal[4])<<56, nil
}

func uint64ToByteArray(uVal uint64) []byte {
	byteVal := []byte{}
	return append(byteVal,
		byte(uVal>>56),
		byte(uVal>>48),
		byte(uVal>>40),
		byte(uVal>>32),
		byte(uVal>>24),
		byte(uVal>>16),
		byte(uVal>>8),
		byte(uVal))
}

func bytesToInt64(byteVal []byte) (int64, error) {
	const sizeVal = 8
	byteVal = checkArray(byteVal, sizeVal)

	_ = byteVal[7]
	return int64(byteVal[3]) |
		int64(byteVal[2])<<8 |
		int64(byteVal[1])<<16 |
		int64(byteVal[0])<<24 |
		int64(byteVal[7])<<32 |
		int64(byteVal[6])<<40 |
		int64(byteVal[5])<<48 |
		int64(byteVal[4])<<56, nil
}

func int64ToByteArray(iVal int64) []byte {
	byteVal := []byte{}
	return append(byteVal,
		byte(iVal>>56),
		byte(iVal>>48),
		byte(iVal>>40),
		byte(iVal>>32),
		byte(iVal>>24),
		byte(iVal>>16),
		byte(iVal>>8),
		byte(iVal))
}

func bytesToFloat32(byteVal []byte) (float32, error) {
	const sizeVal = 4
	byteVal = checkArray(byteVal, sizeVal)

	_ = byteVal[3]
	u32Val := uint32(byteVal[3]) |
		uint32(byteVal[2])<<8 |
		uint32(byteVal[1])<<16 |
		uint32(byteVal[0])<<24
	return math.Float32frombits(u32Val), nil
}

func float32ToByteArray(f32Val float32) []byte {
	byteVal := []byte{}
	u32Val := math.Float32bits(f32Val)
	return append(byteVal,
		byte(u32Val>>24),
		byte(u32Val>>16),
		byte(u32Val>>8),
		byte(u32Val),
	)
}

func bytesToFloat64(byteVal []byte) (float64, error) {
	const sizeVal = 8
	byteVal = checkArray(byteVal, sizeVal)

	_ = byteVal[7]
	u64Val := uint64(byteVal[7]) |
		uint64(byteVal[6])<<8 |
		uint64(byteVal[5])<<16 |
		uint64(byteVal[4])<<24 |
		uint64(byteVal[3])<<32 |
		uint64(byteVal[2])<<40 |
		uint64(byteVal[1])<<48 |
		uint64(byteVal[0])<<56
	return math.Float64frombits(u64Val), nil
}

func float64ToByteArray(f64Val float64) []byte {
	byteVal := []byte{}
	u64Val := math.Float64bits(f64Val)
	return append(byteVal,
		byte(u64Val>>56),
		byte(u64Val>>48),
		byte(u64Val>>40),
		byte(u64Val>>32),
		byte(u64Val>>24),
		byte(u64Val>>16),
		byte(u64Val>>8),
		byte(u64Val))
}
