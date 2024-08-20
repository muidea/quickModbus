package common

import cd "github.com/muidea/magicCommon/def"

const MasterModule = "/kernel/master"

const (
	RawValue    = 0
	BoolValue   = 1
	Int16Value  = 2
	UInt16Value = 3
	Int32Value  = 4
	UInt32Value = 5
	Int64Value  = 6
	UInt64Value = 7
	FloatValue  = 8
	DoubleValue = 9
)

const (
	ABEndian   = 0
	BAEndian   = 1
	ABCDEndian = 2
	CDABEndian = 3
	BADCEndian = 4
	DCBAEndian = 5
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
	SlaveAddr string `json:"slaveAddr"`
	DeviceID  byte   `json:"deviceID"`
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
	Values []bool `json:"values"`
}

type ReadDiscreteInputsRequest struct {
	Address uint16 `json:"address"`
	Count   uint16 `json:"count"`
}

type ReadDiscreteInputsResponse struct {
	cd.Result
	Values []bool `json:"values"`
}

type ReadHoldingRegistersRequest struct {
	Address    uint16 `json:"address"`
	Count      uint16 `json:"count"`
	ValueType  uint16 `json:"valueType"`
	EndianType uint16 `json:"endianType"`
}

type ReadHoldingRegistersResponse struct {
	cd.Result
	Values []interface{} `json:"values"`
}

type ReadReadInputRegistersRequest struct {
	Address    uint16 `json:"address"`
	Count      uint16 `json:"count"`
	ValueType  uint16 `json:"valueType"`
	EndianType uint16 `json:"endianType"`
}

type ReadReadInputRegistersResponse struct {
	cd.Result
	Values []interface{} `json:"values"`
}

type WriteSingleCoilRequest struct {
	Address uint16 `json:"address"`
	Value   bool   `json:"value"`
}

type WriteSingleCoilResponse struct {
	cd.Result
}

type WriteSingleRegisterRequest struct {
	Address    uint16      `json:"address"`
	Value      interface{} `json:"value"`
	ValueType  uint16      `json:"valueType"`
	EndianType uint16      `json:"endianType"`
}

type WriteSingleRegisterResponse struct {
	cd.Result
}

type ReadExceptionStatusRequest struct {
}

type ReadExceptionStatusResponse struct {
	cd.Result
	Value interface{} `json:"value"`
}

type DiagnosticsRequest struct {
	Function uint16 `json:"function"`
	Value    string `json:"value"`
}

type DiagnosticsResponse struct {
	cd.Result
	Function uint16 `json:"function"`
	Value    string `json:"value"`
}

type GetCommEventCounterRequest struct {
}

type GetCommEventCounterResponse struct {
	cd.Result
	CommStatus uint16 `json:"commStatus"`
	EventCount uint16 `json:"eventCount"`
}

type GetCommEventLogRequest struct {
}

type GetCommEventLogResponse struct {
	cd.Result
	CommStatus   int `json:"commStatus"`
	EventCount   int `json:"eventCount"`
	MessageCount int `json:"messageCount"`
	CommEvents   int `json:"commEvents"`
}

type WriteMultipleCoilsRequest struct {
	Address int    `json:"address"`
	Value   []bool `json:"value"`
}

type WriteMultipleCoilsResponse struct {
	cd.Result
	Address int `json:"address"`
	Count   int `json:"count"`
}

type WriteMultipleRegistersRequest struct {
	Address   int           `json:"address"`
	Values    []interface{} `json:"values"`
	ValueType int           `json:"valueType"`
}

type WriteMultipleRegistersResponse struct {
	cd.Result
	Address int `json:"address"`
	Count   int `json:"count"`
}

type ReportSlaveIDRequest struct {
}

type ReportSlaveIDResponse struct {
	cd.Result
	SlaveID         string `json:"slaveID"`
	IndicatorStatus int    `json:"indicatorStatus"`
}

type ReadItem struct {
	FileNumber   int `json:"fileNumber"`
	RecordNumber int `json:"recordNumber"`
	RecordLength int `json:"recordLength"`
}

type ReadFileRecordRequest struct {
	Items []*ReadItem `json:"items"`
}

type ReadFileRecordResponse struct {
	cd.Result
	ItemData []string `json:"itemData"`
}

type WriteItem struct {
	FileNumber   int    `json:"fileNumber"`
	RecordNumber int    `json:"recordNumber"`
	RecordData   string `json:"recordData"`
}

type WriteFileRecordRequest struct {
	Items []*WriteItem `json:"items"`
}

type WriteFileRecordResponse struct {
	cd.Result
	Items []*WriteItem `json:"items"`
}

type MaskWriteRegisterRequest struct {
	Address int    `json:"address"`
	AndMask []bool `json:"andMask"`
	OrMask  []bool `json:"orMask"`
}

type MaskWriteRegisterResponse struct {
	cd.Result
	Address int    `json:"address"`
	AndMask []bool `json:"andMask"`
	OrMask  []bool `json:"orMask"`
}

type ReadWriteMultipleRegistersRequest struct {
	ReadAddress  int           `json:"readAddress"`
	ReadCount    int           `json:"readCount"`
	WriteAddress int           `json:"writeAddress"`
	WriteValues  []interface{} `json:"writeValues"`
	ValueType    int           `json:"valueType"`
}

type ReadWriteMultipleRegistersResponse struct {
	cd.Result
}

type ReadFIFOQueueRequest struct {
}

type ReadFIFOQueueResponse struct {
	cd.Result
}
