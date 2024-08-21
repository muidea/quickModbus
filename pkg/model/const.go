package model

import (
	"fmt"
	"math"
)

const (
	aduTcpHeadLength    = 7
	aduSerialHeadLength = 1
)

const ModbusProtocol = 0

/*
| Primary tables    | Object type | Type of    | Comments                                                      |
| :---------------- | :---------- | :--------- | :------------------------------------------------------------ |
| Discrete Input    | Single bit  | Read-Only  | This type of data can be provided by an l/0 system.           |
| Coils             | Single bit  | Read-Write | This type of data can be alterable by an application program. |
| Input Registers   | 16-bit word | Read-Only  | This type of data can be provided by an l/O system            |
| Holding Registers | 16-bit word | Read-Write | This type of data can be alterable by an application program  |
*/
const (
	ReadCoils                  = byte(0x01)
	ReadDiscreteInputs         = byte(0x02)
	ReadHoldingRegisters       = byte(0x03)
	ReadInputRegisters         = byte(0x04)
	WriteSingleCoil            = byte(0x05)
	WriteSingleRegister        = byte(0x06)
	ReadExceptionStatus        = byte(0x07)
	Diagnostics                = byte(0x08)
	GetCommEventCounter        = byte(0x0B)
	GetCommEventLog            = byte(0x0C)
	WriteMultipleCoils         = byte(0x0F)
	WriteMultipleRegisters     = byte(0x10)
	ReportSlaveID              = byte(0x11)
	ReadFileRecord             = byte(0x14)
	WriteFileRecord            = byte(0x15)
	MaskWriteRegister          = byte(0x16)
	ReadWriteMultipleRegisters = byte(0x17)
	ReadFIFOQueue              = byte(0x18)
)

const (
	RequestAction  = 0
	ResponseAction = 1
)

var coilON = []byte{0xFF, 0x00}
var coilOFF = []byte{0x00, 0x00}

const (
	SuccessCode     = 0x00
	IllegalFuncCode = 0x01
	IllegalAddress  = 0x02
	IllegalCount    = 0x03
	IllegalData     = 0x04
)

func ByteArrayToBoolArrayDCBA(byteVal []byte) []bool {
	ret := []bool{}
	for _, val := range byteVal {
		ret = append(ret, ByteToBoolArrayDCBA(val)...)
	}

	return ret
}

func ByteArrayToBoolArrayABCD(byteVal []byte) []bool {
	ret := []bool{}
	for _, val := range byteVal {
		ret = append(ret, ByteToBoolArrayABCD(val)...)
	}

	return ret
}

func ByteToBoolArrayABCD(byteVal byte) []bool {
	boolArray := make([]bool, 8)

	for i := 7; i >= 0; i-- {
		bit := (byteVal >> i) & 1
		boolArray[7-i] = bit == 1
	}

	return boolArray
}

func ByteToBoolArrayDCBA(byteVal byte) []bool {
	boolArray := make([]bool, 8)

	for i := 0; i <= 7; i++ {
		bit := (byteVal >> i) & 1
		boolArray[i] = bit == 1
	}

	return boolArray
}

func BoolArrayToByteArrayABCD(boolVal []bool) []byte {
	ret := []byte{}
	idx := 0
	for {
		if idx > len(boolVal) {
			break
		}

		subArray := boolVal[idx:]
		byteVal := BoolArrayToByteABCD(subArray)
		ret = append(ret, byteVal)
		idx += 9
	}

	return ret
}

func BoolArrayToByteArrayDCBA(boolVal []bool) []byte {
	ret := []byte{}
	idx := 0
	for {
		if idx > len(boolVal) {
			break
		}

		subArray := boolVal[idx:]
		byteVal := BoolArrayToByteDCBA(subArray)
		ret = append(ret, byteVal)
		idx += 9
	}

	return ret
}

func BoolArrayToByteDCBA(boolArray []bool) byte {
	if len(boolArray) < 8 {
		// 在boolArray前面补充false，使其长度达到8
		padding := make([]bool, 8-len(boolArray))
		boolArray = append(padding, boolArray...)
	} else if len(boolArray) > 8 {
		boolArray = boolArray[:8] // 只取前8个元素，如果boolArray超过8个元素
	}

	var byteVal byte
	for i, bit := range boolArray {
		if bit {
			byteVal |= 1 << i
		}
	}

	return byteVal
}

func BoolArrayToByteABCD(boolArray []bool) byte {
	if len(boolArray) < 8 {
		// 在boolArray前面补充false，使其长度达到8
		padding := make([]bool, 8-len(boolArray))
		boolArray = append(padding, boolArray...)
	} else if len(boolArray) > 8 {
		boolArray = boolArray[:8] // 只取前8个元素，如果boolArray超过8个元素
	}

	var byteVal byte
	for i, bit := range boolArray {
		if bit {
			byteVal |= 1 << (7 - i)
		}
	}

	return byteVal
}

func ByteArrayToUint16ABCDArray(byteVal []byte) ([]uint16, error) {
	uint16Array := []uint16{}
	idx := 0
	totalCount := len(byteVal)
	for idx < totalCount {
		u16 := ByteToUint16ABCD(byteVal[idx : idx+2])
		uint16Array = append(uint16Array, u16)
		idx += 2
	}
	if idx > totalCount {
		return nil, fmt.Errorf("illegal byte stream")
	}

	return uint16Array, nil
}

func ByteToUint16ABCD(byteVal []byte) uint16 {
	_ = byteVal[1]
	return uint16(byteVal[1]) |
		uint16(byteVal[0])<<8
}

func AppendUint16ABCD(byteVal []byte, uVal uint16) []byte {
	return append(byteVal,
		byte(uVal>>8),
		byte(uVal))
}

func ByteToUint16BADC(byteVal []byte) uint16 {
	_ = byteVal[1]
	return uint16(byteVal[0]) |
		uint16(byteVal[1])<<8
}

func AppendUint16BADC(byteVal []byte, uVal uint16) []byte {
	return append(byteVal,
		byte(uVal),
		byte(uVal>>8),
	)
}

func ByteToUint32ABCD(byteVal []byte) uint32 {
	_ = byteVal[3]
	return uint32(byteVal[3]) |
		uint32(byteVal[2])<<8 |
		uint32(byteVal[1])<<16 |
		uint32(byteVal[0])<<24
}

func AppendUint32ABCD(byteVal []byte, uVal uint32) []byte {
	return append(byteVal,
		byte(uVal>>24),
		byte(uVal>>16),
		byte(uVal>>8),
		byte(uVal),
	)
}

func ByteToUint32BADC(byteVal []byte) uint32 {
	_ = byteVal[3]
	return uint32(byteVal[2]) |
		uint32(byteVal[3])<<8 |
		uint32(byteVal[0])<<16 |
		uint32(byteVal[1])<<24
}

func AppendUint32BADC(byteVal []byte, uVal uint32) []byte {
	return append(byteVal,
		byte(uVal>>16),
		byte(uVal>>24),
		byte(uVal),
		byte(uVal>>8),
	)
}

func ByteToUint32CDAB(byteVal []byte) uint32 {
	_ = byteVal[3]
	return uint32(byteVal[1]) |
		uint32(byteVal[0])<<8 |
		uint32(byteVal[3])<<16 |
		uint32(byteVal[2])<<24
}

func AppendUint32CDAB(byteVal []byte, uVal uint32) []byte {
	return append(byteVal,
		byte(uVal>>8),
		byte(uVal),
		byte(uVal>>24),
		byte(uVal>>16),
	)
}

func ByteToUint32DCBA(byteVal []byte) uint32 {
	_ = byteVal[3]
	return uint32(byteVal[0]) |
		uint32(byteVal[1])<<8 |
		uint32(byteVal[2])<<16 |
		uint32(byteVal[3])<<24
}

func AppendUint32DCBA(byteVal []byte, uVal uint32) []byte {
	return append(byteVal,
		byte(uVal),
		byte(uVal>>8),
		byte(uVal>>16),
		byte(uVal>>24),
	)
}

func ByteToInt32ABCD(byteVal []byte) int32 {
	_ = byteVal[3]
	return int32(byteVal[3]) |
		int32(byteVal[2])<<8 |
		int32(byteVal[1])<<16 |
		int32(byteVal[0])<<24
}

func AppendInt32ABCD(byteVal []byte, iVal int32) []byte {
	return append(byteVal,
		byte(iVal>>24),
		byte(iVal>>16),
		byte(iVal>>8),
		byte(iVal),
	)
}

func ByteToInt32BADC(byteVal []byte) int32 {
	_ = byteVal[3]
	return int32(byteVal[2]) |
		int32(byteVal[3])<<8 |
		int32(byteVal[0])<<16 |
		int32(byteVal[1])<<24
}

func AppendInt32BADC(byteVal []byte, iVal int32) []byte {
	return append(byteVal,
		byte(iVal>>16),
		byte(iVal>>24),
		byte(iVal),
		byte(iVal>>8),
	)
}

func ByteToInt32CDAB(byteVal []byte) int32 {
	_ = byteVal[3]
	return int32(byteVal[1]) |
		int32(byteVal[0])<<8 |
		int32(byteVal[3])<<16 |
		int32(byteVal[2])<<24
}

func AppendInt32CDAB(byteVal []byte, iVal int32) []byte {
	return append(byteVal,
		byte(iVal>>8),
		byte(iVal),
		byte(iVal>>24),
		byte(iVal>>16),
	)
}

func ByteToInt32DCBA(byteVal []byte) int32 {
	_ = byteVal[3]
	return int32(byteVal[0]) |
		int32(byteVal[1])<<8 |
		int32(byteVal[2])<<16 |
		int32(byteVal[3])<<24
}

func AppendInt32DCBA(byteVal []byte, iVal int32) []byte {
	return append(byteVal,
		byte(iVal),
		byte(iVal>>8),
		byte(iVal>>16),
		byte(iVal>>24),
	)
}

func ByteToFloatABCD(byteVal []byte) float32 {
	_ = byteVal[3]
	u32Val := uint32(byteVal[3]) |
		uint32(byteVal[2])<<8 |
		uint32(byteVal[1])<<16 |
		uint32(byteVal[0])<<24
	return math.Float32frombits(u32Val)
}

func AppendFloatABCD(byteVal []byte, f32Val float32) []byte {
	u32Val := math.Float32bits(f32Val)
	return append(byteVal,
		byte(u32Val>>24),
		byte(u32Val>>16),
		byte(u32Val>>8),
		byte(u32Val),
	)
}

func ByteToFloatBADC(byteVal []byte) float32 {
	_ = byteVal[3]
	u32Val := uint32(byteVal[2]) |
		uint32(byteVal[3])<<8 |
		uint32(byteVal[0])<<16 |
		uint32(byteVal[1])<<24
	return math.Float32frombits(u32Val)
}

func AppendFloatBADC(byteVal []byte, f32Val float32) []byte {
	u32Val := math.Float32bits(f32Val)
	return append(byteVal,
		byte(u32Val>>16),
		byte(u32Val>>24),
		byte(u32Val),
		byte(u32Val>>8),
	)
}

func ByteToFloatCDAB(byteVal []byte) float32 {
	_ = byteVal[3]
	u32Val := uint32(byteVal[1]) |
		uint32(byteVal[0])<<8 |
		uint32(byteVal[3])<<16 |
		uint32(byteVal[2])<<24
	return math.Float32frombits(u32Val)
}

func AppendFloatCDAB(byteVal []byte, f32Val float32) []byte {
	u32Val := math.Float32bits(f32Val)
	return append(byteVal,
		byte(u32Val>>8),
		byte(u32Val),
		byte(u32Val>>24),
		byte(u32Val>>16),
	)
}

func ByteToFloatDCBA(byteVal []byte) float32 {
	_ = byteVal[3]
	u32Val := uint32(byteVal[0]) |
		uint32(byteVal[1])<<8 |
		uint32(byteVal[2])<<16 |
		uint32(byteVal[3])<<24
	return math.Float32frombits(u32Val)
}

func AppendFloatDCBA(byteVal []byte, f32Val float32) []byte {
	u32Val := math.Float32bits(f32Val)
	return append(byteVal,
		byte(u32Val),
		byte(u32Val>>8),
		byte(u32Val>>16),
		byte(u32Val>>24))
}

func ByteToDoubleABCD(byteVal []byte) float64 {
	_ = byteVal[7]
	u64Val := uint64(byteVal[7]) |
		uint64(byteVal[6])<<8 |
		uint64(byteVal[5])<<16 |
		uint64(byteVal[4])<<24 |
		uint64(byteVal[3])<<32 |
		uint64(byteVal[2])<<40 |
		uint64(byteVal[1])<<48 |
		uint64(byteVal[0])<<56
	return math.Float64frombits(u64Val)
}

func AppendDoubleABCD(byteVal []byte, f64Val float64) []byte {
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

func ByteToDoubleBADC(byteVal []byte) float64 {
	_ = byteVal[7]
	u64Val := uint64(byteVal[6]) |
		uint64(byteVal[7])<<8 |
		uint64(byteVal[4])<<16 |
		uint64(byteVal[5])<<24 |
		uint64(byteVal[2])<<32 |
		uint64(byteVal[3])<<40 |
		uint64(byteVal[0])<<48 |
		uint64(byteVal[1])<<56
	return math.Float64frombits(u64Val)
}

func AppendDoubleBADC(byteVal []byte, f64Val float64) []byte {
	u64Val := math.Float64bits(f64Val)
	return append(byteVal,
		byte(u64Val>>48),
		byte(u64Val>>56),
		byte(u64Val>>32),
		byte(u64Val>>40),
		byte(u64Val>>16),
		byte(u64Val>>24),
		byte(u64Val),
		byte(u64Val>>8))
}

func ByteToDoubleCDAB(byteVal []byte) float64 {
	_ = byteVal[7]
	u64Val := uint64(byteVal[5])<<16 |
		uint64(byteVal[4])<<24 |
		uint64(byteVal[7]) |
		uint64(byteVal[6])<<8 |
		uint64(byteVal[1])<<48 |
		uint64(byteVal[0])<<56 |
		uint64(byteVal[3])<<32 |
		uint64(byteVal[2])<<40
	return math.Float64frombits(u64Val)
}

func AppendDoubleCDAB(byteVal []byte, f64Val float64) []byte {
	u64Val := math.Float64bits(f64Val)
	return append(byteVal,
		byte(u64Val>>40),
		byte(u64Val>>32),
		byte(u64Val>>56),
		byte(u64Val>>48),
		byte(u64Val>>8),
		byte(u64Val),
		byte(u64Val>>24),
		byte(u64Val>>16),
	)
}

func ByteToDoubleDCBA(byteVal []byte) float64 {
	_ = byteVal[7]
	u64Val := uint64(byteVal[0]) |
		uint64(byteVal[1])<<8 |
		uint64(byteVal[2])<<16 |
		uint64(byteVal[3])<<24 |
		uint64(byteVal[4])<<32 |
		uint64(byteVal[5])<<40 |
		uint64(byteVal[6])<<48 |
		uint64(byteVal[7])<<56
	return math.Float64frombits(u64Val)
}

func AppendDoubleDCBA(byteVal []byte, f64Val float64) []byte {
	u64Val := math.Float64bits(f64Val)
	return append(byteVal,
		byte(u64Val),
		byte(u64Val>>8),
		byte(u64Val>>16),
		byte(u64Val>>24),
		byte(u64Val>>32),
		byte(u64Val>>40),
		byte(u64Val>>48),
		byte(u64Val>>56),
	)
}
