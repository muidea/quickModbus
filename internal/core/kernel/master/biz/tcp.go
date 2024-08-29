package biz

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/muidea/magicEngine/tcp"

	"github.com/muidea/magicCommon/foundation/log"
	"github.com/muidea/magicCommon/foundation/signal"

	"github.com/muidea/quickModbus/pkg/common"
	"github.com/muidea/quickModbus/pkg/model"
)

func NewTCPMaster(deviceID byte) MBMaster {
	return &mbTCPMaster{
		deviceID: deviceID,
	}
}

type mbTCPMaster struct {
	serverAddr string
	signalGard signal.Gard

	tcpClient tcp.Client
	serialNo  int
	deviceID  byte
}

func (s *mbTCPMaster) transaction() uint16 {
	s.serialNo++
	return uint16(s.serialNo)
}

func (s *mbTCPMaster) reset() {
	s.serialNo = 0
	s.signalGard.Reset()
	s.tcpClient = nil
}

func (s *mbTCPMaster) connect(serverAddr string) (ret tcp.Client, err error) {
	err = s.signalGard.PutSignal(s.serialNo)
	if err != nil {
		return
	}

	client := tcp.NewClient(s)
	err = client.Connect(serverAddr)
	if err != nil {
		s.signalGard.CleanSignal(s.serialNo)
		return
	}

	addrVal, addrErr := s.signalGard.WaitSignal(s.serialNo, defaultTimeOut)
	if addrErr != nil {
		client.Close()
		err = addrErr
		return
	}

	log.Infof("connect slave %s ok", addrVal)
	ret = client
	return
}

func (s *mbTCPMaster) Start(serverAddr string) (err error) {
	connClient, connErr := s.connect(serverAddr)
	if connErr != nil {
		err = connErr
		log.Errorf("start master %s failed, error:%s", serverAddr, connErr.Error())
		return
	}

	s.tcpClient = connClient
	s.serverAddr = serverAddr
	return
}

func (s *mbTCPMaster) Stop() {
	if s.tcpClient == nil {
		return
	}

	s.tcpClient.Close()
}

func (s *mbTCPMaster) IsConnect() bool {
	return s.tcpClient != nil
}

func (s *mbTCPMaster) ReConnect() (err error) {
	connClient, connErr := s.connect(s.serverAddr)
	if connErr != nil {
		err = connErr
		log.Errorf("reconnect master %s failed, error:%s", s.serverAddr, connErr.Error())
		return
	}

	s.tcpClient = connClient
	return
}

func (s *mbTCPMaster) OnConnect(ep tcp.Endpoint) {
	err := s.signalGard.TriggerSignal(s.serialNo, ep.RemoteAddr().String())
	if err != nil {
		log.Errorf("onConnect triggerSignal failed, error:%s", err.Error())
		return
	}
}

func (s *mbTCPMaster) OnDisConnect(ep tcp.Endpoint) {
	log.Warnf("onDisConnect from %s", ep.RemoteAddr().String())
	s.reset()
}

func (s *mbTCPMaster) OnRecvData(ep tcp.Endpoint, data []byte) {
	dataVal := bytes.NewBuffer(data)
	protocolHeader, protocolVal, protocolErr := model.DecodeMBTcpProtocol(dataVal, model.ResponseAction)
	if protocolErr != model.SuccessCode {
		log.Errorf("decode mbprotocol failed, remoteAddr:%s, error:%v", ep.RemoteAddr().String(), protocolErr)
		return
	}

	err := s.signalGard.TriggerSignal(int(protocolHeader.Transaction()), protocolVal)
	if err != nil {
		log.Errorf("onRecvData triggerSignal failed, remoteAddr:%s, error:%s", ep.RemoteAddr(), err.Error())
	}
}

func (s *mbTCPMaster) ReadCoils(address, count uint16) (retData []byte, exCode byte, err error) {
	protocol := model.NewReadCoilsReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("ReadCoils,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("ReadCoils,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("ReadCoils,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("ReadCoils failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("ReadCoils failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBReadCoilsRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read coils response")
		log.Errorf("ReadCoils failed, error:%s", err.Error())
		return
	}

	retData = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) ReadDiscreteInputs(address, count uint16) (retData []byte, exCode byte, err error) {
	protocol := model.NewReadDiscreteInputsReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("ReadDiscreteInputs,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("ReadDiscreteInputs,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("ReadDiscreteInputs,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("ReadDiscreteInputs failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("ReadDiscreteInputs failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBReadDiscreteInputsRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read discrete inputs response")
		log.Errorf("ReadDiscreteInputs failed, error:%s", err.Error())
		return
	}

	retData = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) ReadHoldingRegisters(address, count uint16) (retData []byte, exCode byte, err error) {
	protocol := model.NewReadHoldingRegistersReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("ReadHoldingRegisters,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("ReadHoldingRegisters,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("ReadHoldingRegisters,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("ReadHoldingRegisters failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("ReadHoldingRegisters failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBReadHoldingRegistersRsp)
	if !readOK {
		err = fmt.Errorf("ReadHoldingRegisters,recv illegal read holding registers response")
		log.Errorf(err.Error())
		return
	}

	retData = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) ReadInputRegisters(address, count uint16) (retData []byte, exCode byte, err error) {
	protocol := model.NewReadInputRegistersReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("ReadInputRegisters,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("ReadInputRegisters,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("ReadInputRegisters,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("ReadInputRegisters failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("ReadInputRegisters failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBReadInputRegistersRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read input registers response")
		log.Errorf("ReadInputRegisters failed, error:%s", err.Error())
		return
	}

	retData = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) WriteSingleCoil(address uint16, data []byte) (retAddr uint16, retData []byte, exCode byte, err error) {
	protocol := model.NewWriteSingleCoilReq(address, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("WriteSingleCoil,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("WriteSingleCoil,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("WriteSingleCoil,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("WriteSingleCoil failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("WriteSingleCoil failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBWriteSingleCoilRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write single coil response")
		log.Errorf("WriteSingleCoil failed, error:%s", err.Error())
		return
	}

	retAddr = readVal.Address()
	retData = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) WriteMultipleCoils(address, count uint16, data []byte) (retAddr, retCount uint16, exCode byte, err error) {
	protocol := model.NewWriteMultipleCoilsReq(address, count, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("WriteMultipleCoils, encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("WriteMultipleCoils, signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("WriteMultipleCoils, tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("WriteMultipleCoils failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("WriteMultipleCoils failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBWriteMultipleCoilsRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write multiple coils response")
		log.Errorf("WriteMultipleCoils failed, error:%s", err.Error())
		return
	}

	retAddr = readVal.Address()
	retCount = readVal.Count()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) WriteSingleRegister(address uint16, data []byte) (retAddr uint16, retData []byte, exCode byte, err error) {
	protocol := model.NewWriteSingleRegisterReq(address, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("WriteSingleRegister,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("WriteSingleRegister,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("WriteSingleRegister,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("WriteSingleRegister failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("WriteSingleRegister failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBWriteSingleRegisterRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write single register response")
		log.Errorf("WriteSingleRegister failed, error:%s", err.Error())
		return
	}

	retAddr = readVal.Address()
	retData = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) WriteMultipleRegisters(address, count uint16, data []byte) (retAddr, retCount uint16, exCode byte, err error) {
	protocol := model.NewWriteMultipleRegistersReq(address, count, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("WriteMultipleRegisters,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("WriteMultipleRegisters,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("WriteMultipleRegisters,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("WriteMultipleRegisters failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("WriteMultipleRegisters failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBWriteMultipleRegistersRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write multiple registers response")
		log.Errorf("WriteMultipleRegisters failed, error:%s", err.Error())
		return
	}

	retAddr = readVal.Address()
	retCount = readVal.Count()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) ReadExceptionStatus() (retStatus, exCode byte, err error) {
	protocol := model.NewReadExceptionStatusReq()
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("ReadExceptionStatus,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("ReadExceptionStatus,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("ReadExceptionStatus,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("ReadExceptionStatus failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("ReadExceptionStatus failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBReadExceptionStatusRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read exception status response")
		log.Errorf("ReadExceptionStatus failed, error:%s", err.Error())
		return
	}

	exCode = readVal.ExceptionCode()
	retStatus = readVal.Status()
	return
}

func (s *mbTCPMaster) Diagnostics(subFuncCode uint16, data []byte) (retSubFuncCode uint16, retData []byte, exCode byte, err error) {
	protocol := model.NewDiagnosticsReq(subFuncCode, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("diagnostics, encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("diagnostics,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("diagnostics,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("diagnostics failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("diagnostics failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBDiagnosticsRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write multiple registers response")
		log.Errorf("diagnostics failed, error:%s", err.Error())
		return
	}

	retSubFuncCode = readVal.SubFunctionCode()
	retData = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) GetCommEventCounter() (status uint16, eventCount uint16, exCode byte, err error) {
	protocol := model.NewGetCommEventCounterReq()
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("GetCommEventCounter encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("GetCommEventCounter signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("GetCommEventCounter tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("GetCommEventCounter failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("GetCommEventCounter failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBGetCommEventCounterRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write multiple registers response")
		log.Errorf("GetCommEventCounter failed, error:%s", err.Error())
		return
	}

	status = readVal.CommStatus()
	eventCount = readVal.EventCount()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) GetCommEventLog() (status uint16, eventCount, messageCount uint16, events []byte, exCode byte, err error) {
	protocol := model.NewGetCommEventLogReq()
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("GetCommEventCounter encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("GetCommEventCounter signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("GetCommEventCounter tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("GetCommEventCounter failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("GetCommEventCounter failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBGetCommEventLogRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write multiple registers response")
		log.Errorf("GetCommEventCounter failed, error:%s", err.Error())
		return
	}

	status = readVal.CommStatus()
	eventCount = readVal.EventCount()
	messageCount = readVal.MessageCount()
	events = readVal.CommonEvents()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) ReportSlaveID() (ret []byte, exCode byte, err error) {
	protocol := model.NewReportSlaveIDReq()
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("ReportSlaveID encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("ReportSlaveID signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("ReportSlaveID tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("ReportSlaveID failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("ReportSlaveID failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBReportSlaveIDRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write multiple registers response")
		log.Errorf("ReportSlaveID failed, error:%s", err.Error())
		return
	}

	exCode = readVal.ExceptionCode()
	ret = readVal.SlaveIDInfo()
	return
}

func (s *mbTCPMaster) ReadFileRecord(items []*common.ReadItem) (ret [][]byte, exCode byte, err error) {
	reqItems := []*model.ReadRequestItem{}
	for _, val := range items {
		reqItems = append(reqItems, model.NewReadRequestItem(val.FileNumber, val.RecordNumber, val.RecordLength))
	}

	protocol := model.NewReadFileRecordReq(reqItems)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("ReadFileRecord encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("ReadFileRecord signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("ReadFileRecord tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("ReadFileRecord failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("ReadFileRecord failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBReadFileRecordRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read file record response")
		log.Errorf("ReadFileRecord failed, error:%s", err.Error())
		return
	}

	exCode = readVal.ExceptionCode()
	for _, val := range readVal.Items() {
		ret = append(ret, val.Data())
	}
	return
}

func (s *mbTCPMaster) WriteFileRecord(items []*common.WriteItem) (exCode byte, err error) {
	reqItems := []*model.WriteItem{}
	for _, val := range items {
		byteVal, byteErr := hex.DecodeString(val.RecordData)
		if byteErr != nil {
			err = byteErr
			return
		}

		reqItems = append(reqItems, model.NewWriteItem(val.FileNumber, val.RecordNumber, byteVal))
	}

	protocol := model.NewWriteFileRecordReq(reqItems)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("WriteFileRecord encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("WriteFileRecord signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("WriteFileRecord tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("WriteFileRecord failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("WriteFileRecord failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBWriteFileRecordRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write file record response")
		log.Errorf("WriteFileRecord failed, error:%s", err.Error())
		return
	}

	exCode = readVal.ExceptionCode()
	if len(items) != len(readVal.Items()) {
		err = fmt.Errorf("mismatch write file record item size")
		log.Errorf("WriteFileRecord failed, error:%s", err.Error())
		return
	}
	return
}

func (s *mbTCPMaster) MaskWriteRegister(address uint16, andBytes []byte, orBytes []byte) (retAddr uint16, retAnd []byte, retOr []byte, exCode byte, err error) {
	protocol := model.NewMaskWriteRegisterReq(address, andBytes, orBytes)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("WriteMultipleRegisters,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("WriteMultipleRegisters,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("WriteMultipleRegisters,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("WriteMultipleRegisters failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("WriteMultipleRegisters failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBMaskWriteRegisterRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write multiple registers response")
		log.Errorf("WriteMultipleRegisters failed, error:%s", err.Error())
		return
	}

	retAddr = readVal.Address()
	retAnd = readVal.AndMask()
	retOr = readVal.OrMask()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) ReadWriteMultipleRegisters(readAddr, readCount uint16, writeAddr, writeCount uint16, writeData []byte) (retData []byte, exCode byte, err error) {
	protocol := model.NewReadWriteMultipleRegistersReq(readAddr, readCount, writeAddr, writeCount, writeData)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("ReadWriteMultipleRegisters,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("ReadWriteMultipleRegisters,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("ReadWriteMultipleRegisters,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("ReadWriteMultipleRegisters failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("ReadWriteMultipleRegisters failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBReadWriteMultipleRegistersRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read&write multiple registers response")
		log.Errorf("ReadWriteMultipleRegisters failed, error:%s", err.Error())
		return
	}

	retData = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *mbTCPMaster) ReadFIFOQueue(address uint16) (retDataCount uint16, retDataVal []byte, exCode byte, err error) {
	protocol := model.NewReadFIFOQueueReq(address)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBTcpProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("ReadFIFOQueue,encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("ReadFIFOQueue,signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("ReadFIFOQueue,tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, defaultTimeOut)
	if recvErr != nil {
		err = recvErr
		log.Errorf("ReadFIFOQueue failed, error:%s", err.Error())
		return
	}
	if recvVal == nil {
		err = fmt.Errorf("recv illegal data")
		log.Errorf("ReadFIFOQueue failed, error:%s", err.Error())
		return
	}

	readVal, readOK := recvVal.(*model.MBReadFIFOQueueRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read fifo queue response")
		log.Errorf("ReadFIFOQueue failed, error:%s", err.Error())
		return
	}

	retDataCount = readVal.DataCount()
	retDataVal = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}
