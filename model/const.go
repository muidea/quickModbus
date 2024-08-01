package model

import "fmt"

const (
	minReqDataLength    = 5
	minRspDataLength    = 2
	pduReqHeadLength    = 5
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
	ReportServerID             = byte(0x11)
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

func ByteArrayToBoolArray(byteVal []byte) []bool {
	ret := []bool{}
	for _, val := range byteVal {
		ret = append(ret, ByteToBoolArrayForBigEndian(val)...)
	}

	return ret
}

func ByteArrayToBoolArrayForLittleEndian(byteVal []byte) []bool {
	ret := []bool{}
	for _, val := range byteVal {
		ret = append(ret, ByteToBoolArrayForLittleEndian(val)...)
	}

	return ret
}

func ByteToBoolArrayForLittleEndian(byteVal byte) []bool {
	boolArray := make([]bool, 8)

	for i := 7; i >= 0; i-- {
		bit := (byteVal >> i) & 1
		boolArray[7-i] = bit == 1
	}

	return boolArray
}

func ByteToBoolArrayForBigEndian(byteVal byte) []bool {
	boolArray := make([]bool, 8)

	for i := 0; i <= 7; i++ {
		bit := (byteVal >> i) & 1
		boolArray[i] = bit == 1
	}

	return boolArray
}

func BoolArrayToByteArray(boolVal []bool) []byte {
	ret := []byte{}
	idx := 0
	for {
		if idx > len(boolVal) {
			break
		}

		subArray := boolVal[idx:]
		byteVal := BoolArrayToByteForBigEndian(subArray)
		ret = append(ret, byteVal)
		idx += 9
	}

	return ret
}

func BoolArrayToByteArrayForLittleEndian(boolVal []bool) []byte {
	ret := []byte{}
	idx := 0
	for {
		if idx > len(boolVal) {
			break
		}

		subArray := boolVal[idx:]
		byteVal := BoolArrayToByteForLittleEndian(subArray)
		ret = append(ret, byteVal)
		idx += 9
	}

	return ret
}

func BoolArrayToByteForLittleEndian(boolArray []bool) byte {
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

func BoolArrayToByteForBigEndian(boolArray []bool) byte {
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

func ByteArrayToUint16ABArray(byteVal []byte) ([]uint16, error) {
	uint16Array := []uint16{}
	idx := 0
	totalCount := len(byteVal)
	for idx < totalCount {
		u16 := ByteToUint16AB(byteVal[idx : idx+2])
		uint16Array = append(uint16Array, u16)
		idx += 2
	}
	if idx > totalCount {
		return nil, fmt.Errorf("illegal byte stream")
	}

	return uint16Array, nil
}

func ByteToUint16AB(byteVal []byte) uint16 {
	_ = byteVal[1]
	return uint16(byteVal[1]) | uint16(byteVal[0])<<8
}

func AppendUint16AB(byteVal []byte, uVal uint16) []byte {
	return append(byteVal,
		byte(uVal>>8),
		byte(uVal))
}

func ByteToUint16BA(byteVal []byte) uint16 {
	_ = byteVal[1]
	return uint16(byteVal[0]) | uint16(byteVal[1])<<8
}

func AppendUint16BA(byteVal []byte, uVal uint16) []byte {
	return append(byteVal,
		byte(uVal),
		byte(uVal>>8))
}

func ByteToUint32ABCD(byteVal []byte) uint32 {
	_ = byteVal[3]
	return uint32(byteVal[3]) | uint32(byteVal[2])<<8 | uint32(byteVal[1])<<16 | uint32(byteVal[0])<<24
}

func AppendUint32ABCD(byteVal []byte, uVal uint32) []byte {
	return append(byteVal,
		byte(uVal>>24),
		byte(uVal>>16),
		byte(uVal>>8),
		byte(uVal))
}

func ByteToUint32CDAB(byteVal []byte) uint32 {
	_ = byteVal[3]
	return uint32(byteVal[1]) | uint32(byteVal[0])<<8 | uint32(byteVal[3])<<16 | uint32(byteVal[2])<<24
}

func AppendUint32CDAB(byteVal []byte, uVal uint32) []byte {
	return append(byteVal,
		byte(uVal>>8),
		byte(uVal),
		byte(uVal>>24),
		byte(uVal>>16),
	)
}

func ByteToUint32BADC(byteVal []byte) uint32 {
	_ = byteVal[3]
	return uint32(byteVal[2]) | uint32(byteVal[3])<<8 | uint32(byteVal[0])<<16 | uint32(byteVal[1])<<24
}

func AppendUint32BADC(byteVal []byte, uVal uint32) []byte {
	return append(byteVal,
		byte(uVal>>16),
		byte(uVal>>24),
		byte(uVal),
		byte(uVal>>8))
}

func ByteToUint32DCBA(byteVal []byte) uint32 {
	_ = byteVal[3]
	return uint32(byteVal[0]) | uint32(byteVal[1])<<8 | uint32(byteVal[2])<<16 | uint32(byteVal[3])<<24
}

func AppendUint32DCBA(byteVal []byte, uVal uint32) []byte {
	return append(byteVal,
		byte(uVal),
		byte(uVal>>8),
		byte(uVal>>16),
		byte(uVal>>24))
}
