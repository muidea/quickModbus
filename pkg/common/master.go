package common

import cd "github.com/muidea/magicCommon/def"

const MasterModule = "/kernel/master"

const (
	RawValue     = 0
	BoolValue    = 1
	Int16Value   = 2
	UInt16Value  = 3
	Int32Value   = 4
	UInt32Value  = 5
	Int64Value   = 6
	UInt64Value  = 7
	Float32Value = 8
	Float64Value = 9
)

/*
Default 0 不调整字节序，以PLC返回为准
ABCD 1 Big-endian 按照顺序排序
BADC 2 Big-endian byte swap 按照单字反转
CDAB 3 Little-endian byte swap 按照双字反转
DCBA 4 Little-endian 按照倒序排序
AB 两字节有效，顺序排列
BA 两字节有效，倒序排列
*/
const (
	DefaultEndian = 0
	ABCDEndian    = 1
	BADCEndian    = 2
	CDABEndian    = 3
	DCBAEndian    = 4
	ABEndian      = 5
	BAEndian      = 6
)

const (
	ModbusTcp          = 0
	ModbusRTUOverTcp   = 1
	ModbusASCIIOverTcp = 2
)

const (
	ConnectSlave               = "/slave/connect"
	DisConnectSlave            = "/slave/:id/disconnect"
	ReadCoils                  = "/slave/:id/coils/read"
	ReadDiscreteInputs         = "/slave/:id/discrete/input/read"
	ReadHoldingRegisters       = "/slave/:id/holding/register/read"
	ReadInputRegisters         = "/slave/:id/input/register/read"
	WriteSingleCoil            = "/slave/:id/coil/write"
	WriteSingleRegister        = "/slave/:id/register/write"
	ReadExceptionStatus        = "/slave/:id/exception/status/read"
	Diagnostics                = "/slave/:id/diagnostics"
	GetCommEventCounter        = "/slave/:id/event/counter/read"
	GetCommEventLog            = "/slave/:id/event/log/read"
	WriteMultipleCoils         = "/slave/:id/coils/write"
	WriteMultipleRegisters     = "/slave/:id/registers/write"
	ReportSlaveID              = "/slave/:id/slave/report"
	ReadFileRecord             = "/slave/:id/file/record/read"
	WriteFileRecord            = "/slave/:id/file/record/write"
	MaskWriteRegister          = "/slave/:id/register/write/mask"
	ReadWriteMultipleRegisters = "/slave/:id/registers/rw"
	ReadFIFOQueue              = "/slave/:id/queue/read"
)

type ConnectSlaveRequest struct {
	SlaveAddr  string `json:"slaveAddr"`
	DeviceID   byte   `json:"deviceID"`
	DeviceType byte   `json:"deviceType"`
	EndianType byte   `json:"endianType"`
}

type ConnectSlaveResponse struct {
	cd.Result
	SlaveID string
}

type ReadCoilsRequest struct {
	Address uint16 `json:"address"`
	Count   uint16 `json:"count"`
}

type ReadCoilsResponse struct {
	cd.Result
	ExceptionCode byte        `json:"exceptionCode"`
	Values        interface{} `json:"values"`
}

type ReadDiscreteInputsRequest struct {
	Address uint16 `json:"address"`
	Count   uint16 `json:"count"`
}

type ReadDiscreteInputsResponse struct {
	cd.Result
	ExceptionCode byte        `json:"exceptionCode"`
	Values        interface{} `json:"values"`
}

type ReadHoldingRegistersRequest struct {
	Address    uint16 `json:"address"`
	Count      uint16 `json:"count"`
	ValueType  uint16 `json:"valueType"`
	EndianType byte   `json:"endianType"`
}

type ReadHoldingRegistersResponse struct {
	cd.Result
	ExceptionCode byte        `json:"exceptionCode"`
	Values        interface{} `json:"values"`
}

type ReadReadInputRegistersRequest struct {
	Address    uint16 `json:"address"`
	Count      uint16 `json:"count"`
	ValueType  uint16 `json:"valueType"`
	EndianType byte   `json:"endianType"`
}

type ReadReadInputRegistersResponse struct {
	cd.Result
	ExceptionCode byte        `json:"exceptionCode"`
	Values        interface{} `json:"values"`
}

type WriteSingleCoilRequest struct {
	Address uint16 `json:"address"`
	Value   bool   `json:"value"`
}

type WriteSingleCoilResponse struct {
	cd.Result
	ExceptionCode byte `json:"exceptionCode"`
}

type WriteSingleRegisterRequest struct {
	Address    uint16 `json:"address"`
	Value      uint16 `json:"value"`
	EndianType byte   `json:"endianType"`
}

type WriteSingleRegisterResponse struct {
	cd.Result
	ExceptionCode byte `json:"exceptionCode"`
}

type ReadExceptionStatusRequest struct {
}

type ReadExceptionStatusResponse struct {
	cd.Result
	ExceptionCode byte `json:"exceptionCode"`
	Status        byte `json:"status"`
}

type DiagnosticsRequest struct {
	Function uint16 `json:"function"`
	Value    string `json:"value"`
}

type DiagnosticsResponse struct {
	cd.Result
	ExceptionCode byte   `json:"exceptionCode"`
	Value         string `json:"value"`
}

type GetCommEventCounterRequest struct {
}

type GetCommEventCounterResponse struct {
	cd.Result
	ExceptionCode byte   `json:"exceptionCode"`
	CommStatus    uint16 `json:"commStatus"`
	EventCount    uint16 `json:"eventCount"`
}

type GetCommEventLogRequest struct {
}

type GetCommEventLogResponse struct {
	cd.Result
	ExceptionCode byte   `json:"exceptionCode"`
	CommStatus    uint16 `json:"commStatus"`
	EventCount    uint16 `json:"eventCount"`
	MessageCount  uint16 `json:"messageCount"`
	CommEvents    string `json:"commEvents"`
}

type WriteMultipleCoilsRequest struct {
	Address uint16 `json:"address"`
	Values  []bool `json:"values"`
}

type WriteMultipleCoilsResponse struct {
	cd.Result
	ExceptionCode byte `json:"exceptionCode"`
}

type WriteMultipleRegistersRequest struct {
	Address    uint16    `json:"address"`
	Values     []float64 `json:"values"`
	ValueType  uint16    `json:"valueType"`
	EndianType byte      `json:"endianType"`
}

type WriteMultipleRegistersResponse struct {
	cd.Result
	ExceptionCode byte `json:"exceptionCode"`
}

type ReportSlaveIDRequest struct {
}

type ReportSlaveIDResponse struct {
	cd.Result
	ExceptionCode byte   `json:"exceptionCode"`
	SlaveID       string `json:"slaveID"`
}

type ReadItem struct {
	FileNumber   uint16 `json:"fileNumber"`
	RecordNumber uint16 `json:"recordNumber"`
	RecordLength uint16 `json:"recordLength"`
}

type ReadFileRecordRequest struct {
	Items []*ReadItem `json:"items"`
}

type ReadFileRecordResponse struct {
	cd.Result
	ExceptionCode byte     `json:"exceptionCode"`
	ItemData      []string `json:"itemData"`
}

type WriteItem struct {
	FileNumber   uint16 `json:"fileNumber"`
	RecordNumber uint16 `json:"recordNumber"`
	RecordData   string `json:"recordData"`
}

type WriteFileRecordRequest struct {
	Items []*WriteItem `json:"items"`
}

type WriteFileRecordResponse struct {
	cd.Result
	ExceptionCode byte `json:"exceptionCode"`
}

type MaskWriteRegisterRequest struct {
	Address uint16 `json:"address"`
	AndMask uint16 `json:"andMask"`
	OrMask  uint16 `json:"orMask"`
}

type MaskWriteRegisterResponse struct {
	cd.Result
	ExceptionCode byte `json:"exceptionCode"`
}

type ReadWriteMultipleRegistersRequest struct {
	ReadAddress    uint16    `json:"readAddress"`
	ReadCount      uint16    `json:"readCount"`
	ReadValueType  uint16    `json:"readValueType"`
	WriteAddress   uint16    `json:"writeAddress"`
	WriteValues    []float64 `json:"writeValues"`
	WriteValueType uint16    `json:"writeValueType"`
	EndianType     byte      `json:"endianType"`
}

type ReadWriteMultipleRegistersResponse struct {
	cd.Result
	ExceptionCode byte        `json:"exceptionCode"`
	Values        interface{} `json:"values"`
}

type ReadFIFOQueueRequest struct {
	Address uint16 `json:"address"`
}

type ReadFIFOQueueResponse struct {
	cd.Result
	ExceptionCode byte     `json:"exceptionCode"`
	Data          []string `json:"data"`
}
