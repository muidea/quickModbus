package biz

import (
	"github.com/muidea/magicEngine/tcp"
	"github.com/muidea/quickModbus/pkg/common"
)

const defaultTimeOut = 5
const connectID = 0

type MBMaster interface {
	Start(serverAddr string) (err error)
	Stop()
	IsConnect() bool
	ReConnect() (err error)
	EndianType() byte
	OnConnect(ep tcp.Endpoint)
	OnDisConnect(ep tcp.Endpoint)
	OnRecvData(ep tcp.Endpoint, data []byte)
	ReadCoils(address, count uint16) (retData []byte, exCode byte, err error)
	ReadDiscreteInputs(address, count uint16) (retData []byte, exCode byte, err error)
	ReadHoldingRegisters(address, count uint16) (retData []byte, exCode byte, err error)
	ReadInputRegisters(address, count uint16) (retData []byte, exCode byte, err error)
	WriteSingleCoil(address uint16, data []byte) (retAddr uint16, retData []byte, exCode byte, err error)
	WriteMultipleCoils(address, count uint16, data []byte) (retAddr, retCount uint16, exCode byte, err error)
	WriteSingleRegister(address uint16, data []byte) (retAddr uint16, retData []byte, exCode byte, err error)
	WriteMultipleRegisters(address, count uint16, data []byte) (retAddr, retCount uint16, exCode byte, err error)
	ReadExceptionStatus() (retStatus, exCode byte, err error)
	Diagnostics(subFuncCode uint16, data []byte) (retSubFuncCode uint16, retData []byte, exCode byte, err error)
	GetCommEventCounter() (status uint16, eventCount uint16, exCode byte, err error)
	GetCommEventLog() (status uint16, eventCount, messageCount uint16, events []byte, exCode byte, err error)
	ReportSlaveID() (ret []byte, exCode byte, err error)
	ReadFileRecord(items []*common.ReadItem) (ret [][]byte, exCode byte, err error)
	WriteFileRecord(items []*common.WriteItem) (exCode byte, err error)
	MaskWriteRegister(address uint16, andBytes []byte, orBytes []byte) (retAddr uint16, retAnd []byte, retOr []byte, exCode byte, err error)
	ReadWriteMultipleRegisters(readAddr, readCount uint16, writeAddr, writeCount uint16, writeData []byte) (retData []byte, exCode byte, err error)
	ReadFIFOQueue(address uint16) (retDataCount uint16, retDataVal []byte, exCode byte, err error)
}
