package model

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

var CoilON = []byte{0xFF, 0x00}
var CoilOFF = []byte{0x00, 0x00}

const (
	SuccessCode     = 0x00
	IllegalFuncCode = 0x01
	IllegalAddress  = 0x02
	IllegalCount    = 0x03
	IllegalData     = 0x04
)
