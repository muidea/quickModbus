package biz

import (
	"bytes"
	"fmt"
	"github.com/muidea/magicCommon/foundation/log"
	"github.com/muidea/magicCommon/foundation/signal"
	"github.com/muidea/magicEngine/tcp"
	"github.com/muidea/quickModbus/pkg/model"
)

const defaultTimeOut = 5

type MBMaster struct {
	serverAddr string
	signalGard signal.Gard

	tcpClient tcp.Client
	serialNo  int
	deviceID  byte
}

func (s *MBMaster) transaction() uint16 {
	s.serialNo++
	return uint16(s.serialNo)
}

func (s *MBMaster) reset() {
	s.serialNo = 0
	s.signalGard.Reset()
	s.tcpClient = nil
}

func (s *MBMaster) connect(serverAddr string) (ret tcp.Client, err error) {
	s.signalGard.PutSignal(s.serialNo)

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

func (s *MBMaster) Start(serverAddr string, deviceID byte) (err error) {
	connClient, connErr := s.connect(serverAddr)
	if connErr != nil {
		err = connErr
		log.Errorf("start master %s failed, error:%s", serverAddr, connErr.Error())
		return
	}

	s.tcpClient = connClient
	s.serverAddr = serverAddr
	s.deviceID = deviceID
	return
}

func (s *MBMaster) Stop() {
	if s.tcpClient == nil {
		return
	}

	s.tcpClient.Close()
}

func (s *MBMaster) IsConnect() bool {
	return s.tcpClient != nil
}

func (s *MBMaster) ReConnect() (err error) {
	connClient, connErr := s.connect(s.serverAddr)
	if connErr != nil {
		err = connErr
		log.Errorf("reconnect master %s failed, error:%s", s.serverAddr, connErr.Error())
		return
	}

	s.tcpClient = connClient
	return
}

func (s *MBMaster) OnConnect(ep tcp.Endpoint) {
	err := s.signalGard.TriggerSignal(s.serialNo, ep.RemoteAddr().String())
	if err != nil {
		log.Errorf("onConnect triggerSignal failed, error:%s", err.Error())
		return
	}
}

func (s *MBMaster) OnDisConnect(ep tcp.Endpoint) {
	log.Warnf("onDisConnect from %s", ep.RemoteAddr().String())
	s.reset()
}

func (s *MBMaster) OnRecvData(ep tcp.Endpoint, data []byte) {
	dataVal := bytes.NewBuffer(data)
	protocolHeader, protocolVal, protocolErr := model.DecodeMBProtocol(dataVal, model.ResponseAction)
	if protocolErr != model.SuccessCode {
		log.Errorf("decode mbprotocol failed, error:%v", protocolErr)
		return
	}

	err := s.signalGard.TriggerSignal(int(protocolHeader.Transaction()), protocolVal)
	if err != nil {
		log.Errorf("onRecvData triggerSignal failed, error:%s", err.Error())
	}
}

func (s *MBMaster) ReadCoils(address, count uint16) (ret []byte, exCode byte, err error) {
	protocol := model.NewReadCoilsReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
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

	ret = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *MBMaster) ReadDiscreteInputs(address, count uint16) (ret []byte, exCode byte, err error) {
	protocol := model.NewReadDiscreteInputsReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
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

	ret = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *MBMaster) ReadHoldingRegisters(address, count uint16) (ret []byte, exCode byte, err error) {
	protocol := model.NewReadHoldingRegistersReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
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

	ret = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *MBMaster) ReadInputRegisters(address, count uint16) (ret []byte, exCode byte, err error) {
	protocol := model.NewReadInputRegistersReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
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

	ret = readVal.Data()
	exCode = readVal.ExceptionCode()
	return
}

func (s *MBMaster) WriteSingleCoil(address uint16, data []byte) (retAddr uint16, retData []byte, exCode byte, err error) {
	protocol := model.NewWriteSingleCoilReq(address, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
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

func (s *MBMaster) WriteMultipleCoils(address, count uint16, data []byte) (retAddr, retCount uint16, exCode byte, err error) {
	protocol := model.NewWriteMultipleCoilsReq(address, count, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
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

func (s *MBMaster) WriteSingleRegister(address uint16, data []byte) (retAddr uint16, retData []byte, exCode byte, err error) {
	protocol := model.NewWriteSingleRegisterReq(address, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
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

func (s *MBMaster) WriteMultipleRegisters(address, count uint16, data []byte) (retAddr, retCount uint16, exCode byte, err error) {
	protocol := model.NewWriteMultipleRegistersReq(address, count, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), s.deviceID)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
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

func (s *MBMaster) ReadExceptionStatus() (ret []byte, err error) {
	return
}

func (s *MBMaster) Diagnostics() (ret []byte, err error) {
	return
}

func (s *MBMaster) GetCommEventCounter() (ret []byte, err error) {
	return
}

func (s *MBMaster) GetCommEventLog() (ret []byte, err error) {
	return
}

func (s *MBMaster) ReportServerID() (ret []byte, err error) {
	return
}

func (s *MBMaster) ReadFileRecord() (ret []byte, err error) {
	return
}

func (s *MBMaster) WriteFileRecord() (ret []byte, err error) {
	return
}

func (s *MBMaster) MaskWriteRegister() (ret []byte, err error) {
	return
}

func (s *MBMaster) ReadWriteMultipleRegisters() (ret []byte, err error) {
	return
}

func (s *MBMaster) ReadFIFOQueue() (ret []byte, err error) {
	return
}
