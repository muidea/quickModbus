package biz

import (
	"bytes"
	"fmt"
	"github.com/muidea/magicCommon/foundation/log"
	"github.com/muidea/magicCommon/foundation/signal"
	"github.com/muidea/magicEngine/tcp"
	"github.com/muidea/quickModbus/pkg/model"
)

type MBMaster struct {
	serverAddr string
	signalGard signal.Gard

	tcpClient tcp.Client
	serialNo  int
}

func (s *MBMaster) transaction() uint16 {
	s.serialNo++
	return uint16(s.serialNo)
}

func (s *MBMaster) Start(serverAddr string) (err error) {
	s.signalGard.PutSignal(s.serialNo)

	client := tcp.NewClient(s)
	err = client.Connect(serverAddr)
	if err != nil {
		s.signalGard.CleanSignal(s.serialNo)
		return
	}

	s.tcpClient = client

	addrVal, addrErr := s.signalGard.WaitSignal(s.serialNo, 10)
	if addrErr != nil {
		client.Close()
		err = addrErr
		return
	}

	log.Infof("connect slave %s ok", addrVal)
	return
}

func (s *MBMaster) Stop() {
	if s.tcpClient == nil {
		return
	}

	s.tcpClient.Close()
}

func (s *MBMaster) OnConnect(ep tcp.Endpoint) {
	err := s.signalGard.TriggerSignal(s.serialNo, ep.RemoteAddr().String())
	if err != nil {
		log.Errorf("onConnect triggerSignal failed, error:%s", err.Error())
		return
	}
}

func (s *MBMaster) OnDisConnect(ep tcp.Endpoint) {

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

func (s *MBMaster) ReadCoils(address, count uint16) (ret []byte, err error) {
	protocol := model.NewReadCoilsReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), 0)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, 5)
	if recvErr != nil || recvVal == nil {
		err = recvErr
		return
	}

	readVal, readOK := recvVal.(*model.MBReadCoilsRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read coils response")
		log.Errorf(err.Error())
		return
	}

	ret = readVal.Data()
	return
}

func (s *MBMaster) ReadDiscreteInputs(address, count uint16) (ret []byte, err error) {
	protocol := model.NewReadDiscreteInputsReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), 0)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, 5)
	if recvErr != nil || recvVal == nil {
		err = recvErr
		return
	}

	readVal, readOK := recvVal.(*model.MBReadDiscreteInputsRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read discrete inputs response")
		log.Errorf(err.Error())
		return
	}

	ret = readVal.Data()
	return
}

func (s *MBMaster) ReadHoldingRegisters(address, count uint16) (ret []byte, err error) {
	protocol := model.NewReadHoldingRegistersReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), 0)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, 5)
	if recvErr != nil || recvVal == nil {
		err = recvErr
		return
	}

	readVal, readOK := recvVal.(*model.MBReadHoldingRegistersRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read holding registers response")
		log.Errorf(err.Error())
		return
	}

	ret = readVal.Data()
	return
}

func (s *MBMaster) ReadInputRegisters(address, count uint16) (ret []byte, err error) {
	protocol := model.NewReadInputRegistersReq(address, count)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), 0)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, 5)
	if recvErr != nil || recvVal == nil {
		err = recvErr
		return
	}

	readVal, readOK := recvVal.(*model.MBReadInputRegistersRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal read input registers response")
		log.Errorf(err.Error())
		return
	}

	ret = readVal.Data()
	return
}

func (s *MBMaster) WriteSingleCoil(address uint16, data []byte) (retAddr uint16, retData []byte, err error) {
	protocol := model.NewWriteSingleCoilReq(address, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), 0)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, 5)
	if recvErr != nil {
		err = recvErr
		return
	}

	readVal, readOK := recvVal.(*model.MBWriteSingleCoilRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write single coil response")
		log.Errorf(err.Error())
		return
	}

	retAddr = readVal.Address()
	retData = readVal.Data()
	return
}

func (s *MBMaster) WriteMultipleCoils(address, count uint16, data []byte) (retAddr, retCount uint16, err error) {
	protocol := model.NewWriteMultipleCoilsReq(address, count, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), 0)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, 5)
	if recvErr != nil || recvVal == nil {
		err = recvErr
		return
	}

	readVal, readOK := recvVal.(*model.MBWriteMultipleCoilsRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write multiple coils response")
		log.Errorf(err.Error())
		return
	}

	retAddr = readVal.Address()
	retCount = readVal.Count()
	return
}

func (s *MBMaster) WriteSingleRegister(address uint16, data []byte) (retAddr uint16, retData []byte, err error) {
	protocol := model.NewWriteSingleRegisterReq(address, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), 0)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, 5)
	if recvErr != nil || recvVal == nil {
		err = recvErr
		return
	}

	readVal, readOK := recvVal.(*model.MBWriteSingleRegisterRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write single register response")
		log.Errorf(err.Error())
		return
	}

	retAddr = readVal.Address()
	retData = readVal.Data()
	return
}

func (s *MBMaster) WriteMultipleRegisters(address, count uint16, data []byte) (retAddr, retCount uint16, err error) {
	protocol := model.NewWriteMultipleRegistersReq(address, count, data)
	header := model.NewTcpHeader(s.transaction(), protocol.CalcLen(), 0)

	buffVal := bytes.NewBuffer(nil)
	eErr := model.EncodeMBProtocol(header, protocol, buffVal)
	if eErr != model.SuccessCode {
		err = fmt.Errorf("encode mbprotocol failed, error:%v", eErr)
		log.Errorf(err.Error())
		return
	}

	signalID := int(header.Transaction())
	err = s.signalGard.PutSignal(signalID)
	if err != nil {
		log.Errorf("signalGard.PutSignal failed, error:%s", err.Error())
		return
	}
	byteVal := buffVal.Bytes()
	err = s.tcpClient.SendData(byteVal)
	if err != nil {
		log.Errorf("tcpClient.SendData failed, error:%s", err.Error())
		return
	}

	recvVal, recvErr := s.signalGard.WaitSignal(signalID, 5)
	if recvErr != nil || recvVal == nil {
		err = recvErr
		return
	}

	readVal, readOK := recvVal.(*model.MBWriteMultipleRegistersRsp)
	if !readOK {
		err = fmt.Errorf("recv illegal write multiple registers response")
		log.Errorf(err.Error())
		return
	}

	retAddr = readVal.Address()
	retCount = readVal.Count()
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
