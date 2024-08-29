package model

import (
	"encoding/binary"
	"io"
)

const MaxAduSize = 256

type MBTcpHeader interface {
	Transaction() uint16
	Protocol() uint16
	DataLen() uint16
	UnitID() byte
	Encode(writer io.Writer) (err byte)
	Decode(reader io.Reader) (err byte)
}

type MBSerialHeader interface {
	Address() byte
	Encode(writer io.Writer) (err byte)
	Decode(reader io.Reader) (err byte)
}

type mbTcpHeader struct {
	transaction uint16
	protocol    uint16
	packLen     uint16
	unitID      byte
}

func NewTcpHeader(transaction, dataLen uint16, unitID byte) MBTcpHeader {
	return &mbTcpHeader{
		transaction: transaction,
		packLen:     dataLen + 1,
		unitID:      unitID,
	}
}

func EmptyTcpHeader() MBTcpHeader {
	return &mbTcpHeader{}
}

func (s *mbTcpHeader) Transaction() uint16 {
	return s.transaction
}

func (s *mbTcpHeader) Protocol() uint16 {
	return s.protocol
}

func (s *mbTcpHeader) DataLen() uint16 {
	return s.packLen - 1
}

func (s *mbTcpHeader) UnitID() byte {
	return s.unitID
}

func (s *mbTcpHeader) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.transaction)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.protocol)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.packLen)
	buffVal = append(buffVal, s.unitID)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != aduTcpHeadLength {
		err = IllegalAddress
	}

	return
}

func (s *mbTcpHeader) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	dataVal := make([]byte, aduTcpHeadLength)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != aduTcpHeadLength {
		err = IllegalData
		return
	}

	s.transaction = binary.BigEndian.Uint16(dataVal[0:2])
	s.protocol = binary.BigEndian.Uint16(dataVal[2:4])
	s.packLen = binary.BigEndian.Uint16(dataVal[4:6])
	s.unitID = dataVal[6]
	return
}

func (s *mbTcpHeader) Same(ptr *mbTcpHeader) bool {
	if ptr == nil {
		return false
	}

	return s.transaction == ptr.transaction &&
		s.protocol == ptr.protocol &&
		s.packLen == ptr.packLen &&
		s.unitID == ptr.unitID
}

type mbSerialHeader struct {
	address byte
}

func NewSerialHeader(address byte) MBSerialHeader {
	return &mbSerialHeader{
		address: address,
	}
}

func EmptySerialHeader() MBSerialHeader {
	return &mbSerialHeader{}
}

func (s *mbSerialHeader) Address() byte {
	return s.address
}

func (s *mbSerialHeader) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	wSize, wErr := writer.Write([]byte{s.address})
	if wErr != nil || wSize != aduSerialHeadLength {
		err = IllegalAddress
		return
	}

	return
}

func (s *mbSerialHeader) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	dataVal := make([]byte, aduSerialHeadLength)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != aduSerialHeadLength {
		err = IllegalData
		return
	}

	s.address = dataVal[0]
	return
}

func EncodeMBTcpProtocol(header MBTcpHeader, pdu MBProtocol, writer io.Writer) (err byte) {
	err = header.Encode(writer)
	if err != SuccessCode {
		return
	}

	wSize, wErr := writer.Write([]byte{pdu.FuncCode()})
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}

	err = pdu.EncodePayload(writer)
	if err != SuccessCode {
		return
	}
	return
}

func DecodeMBTcpProtocol(reader io.Reader, actionType int) (MBTcpHeader, MBProtocol, byte) {
	if actionType == RequestAction {
		return decodeTcpRequestPDU(reader)
	}

	if actionType == ResponseAction {
		return decodeTcpResponsePDU(reader)
	}

	return nil, nil, IllegalData
}

func decodeTcpRequestPDU(reader io.Reader) (MBTcpHeader, MBProtocol, byte) {
	header := EmptyTcpHeader()
	err := header.Decode(reader)
	if err != SuccessCode {
		return nil, nil, err
	}

	funcCode := make([]byte, 1)
	rSize, rErr := reader.Read(funcCode)
	if rErr != nil || rSize != 1 {
		err = IllegalAddress
		return nil, nil, err
	}

	var protocol MBProtocol
	switch funcCode[0] {
	case ReadCoils:
		protocol = EmptyReadCoilsReq()
	case ReadDiscreteInputs:
		protocol = EmptyReadDiscreteInputsReq()
	case ReadHoldingRegisters:
		protocol = EmptyReadHoldingRegistersReq()
	case ReadInputRegisters:
		protocol = EmptyReadInputRegistersReq()
	case WriteSingleCoil:
		protocol = EmptyWriteSingleCoilReq()
	case WriteSingleRegister:
		protocol = EmptyWriteSingleRegisterReq()
	case ReadExceptionStatus:
		protocol = EmptyReadExceptionStatusReq()
	case Diagnostics:
		protocol = EmptyDiagnosticsReq()
	case GetCommEventCounter:
		protocol = EmptyGetCommEventCounterReq()
	case GetCommEventLog:
		protocol = EmptyGetCommEventLogReq()
	case WriteMultipleCoils:
		protocol = EmptyWriteMultipleCoilsReq()
	case WriteMultipleRegisters:
		protocol = EmptyWriteMultipleRegistersReq()
	case ReportSlaveID:
		protocol = EmptyReportSlaveIDReq()
	case ReadFileRecord:
		protocol = EmptyReadFileRecordReq()
	case WriteFileRecord:
		protocol = EmptyWriteFileRecordReq()
	case MaskWriteRegister:
		protocol = EmptyMaskWriteRegisterReq()
	case ReadWriteMultipleRegisters:
		protocol = EmptyReadWriteMultipleRegistersReq()
	case ReadFIFOQueue:
		protocol = EmptyReadFIFOQueueReq()
	default:
		err = IllegalFuncCode
	}

	if err != SuccessCode {
		return nil, nil, err
	}
	if err != SuccessCode {
		return nil, nil, err
	}
	err = protocol.DecodePayload(reader)
	if err != SuccessCode {
		return nil, nil, err
	}

	return header, protocol, err
}

func decodeTcpResponsePDU(reader io.Reader) (MBTcpHeader, MBProtocol, byte) {
	header := EmptyTcpHeader()
	err := header.Decode(reader)
	if err != SuccessCode {
		return nil, nil, err
	}

	funcCode := make([]byte, 1)
	rSize, rErr := reader.Read(funcCode)
	if rErr != nil || rSize != 1 {
		err = IllegalAddress
		return nil, nil, err
	}
	var exceptionCode byte
	lCode := funcCode[0]
	if funcCode[0] > 0x80 {
		exceptionRsp := EmptyExceptionRsp()
		err = exceptionRsp.DecodePayload(reader)
		if err != SuccessCode {
			return nil, nil, err
		}
		exceptionCode = exceptionRsp.ExceptionCode()
		lCode = funcCode[0] - 0x80
	}
	var protocol MBProtocol
	switch lCode {
	case ReadCoils:
		protocol = EmptyReadCoilsRsp(exceptionCode)
	case ReadDiscreteInputs:
		protocol = EmptyReadDiscreteInputsRsp(exceptionCode)
	case ReadHoldingRegisters:
		protocol = EmptyReadHoldingRegistersRsp(exceptionCode)
	case ReadInputRegisters:
		protocol = EmptyReadInputRegistersRsp(exceptionCode)
	case WriteSingleCoil:
		protocol = EmptyWriteSingleCoilRsp(exceptionCode)
	case WriteSingleRegister:
		protocol = EmptyWriteSingleRegisterRsp(exceptionCode)
	case ReadExceptionStatus:
		protocol = EmptyReadExceptionStatusRsp(exceptionCode)
	case Diagnostics:
		protocol = EmptyDiagnosticsRsp(exceptionCode)
	case GetCommEventCounter:
		protocol = EmptyGetCommEventCounterRsp(exceptionCode)
	case GetCommEventLog:
		protocol = EmptyGetCommEventLogRsp(exceptionCode)
	case WriteMultipleCoils:
		protocol = EmptyWriteMultipleCoilsRsp(exceptionCode)
	case WriteMultipleRegisters:
		protocol = EmptyWriteMultipleRegistersRsp(exceptionCode)
	case ReportSlaveID:
		protocol = EmptyReportSlaveIDRsp(exceptionCode)
	case ReadFileRecord:
		protocol = EmptyReadFileRecordRsp(exceptionCode)
	case WriteFileRecord:
		protocol = EmptyWriteFileRecordRsp(exceptionCode)
	case MaskWriteRegister:
		protocol = EmptyMaskWriteRegisterRsp(exceptionCode)
	case ReadWriteMultipleRegisters:
		protocol = EmptyReadWriteMultipleRegistersRsp(exceptionCode)
	case ReadFIFOQueue:
		protocol = EmptyReadFIFOQueueRsp(exceptionCode)
	default:
		err = IllegalFuncCode
	}

	if err != SuccessCode {
		return nil, nil, err
	}
	if funcCode[0] < 0x80 {
		err = protocol.DecodePayload(reader)
		if err != SuccessCode {
			return nil, nil, err
		}
	}

	return header, protocol, err
}

func EncodeMBSerialProtocol(header MBSerialHeader, pdu MBProtocol, writer io.Writer) (err byte) {
	err = header.Encode(writer)
	if err != SuccessCode {
		return
	}

	wSize, wErr := writer.Write([]byte{pdu.FuncCode()})
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}

	err = pdu.EncodePayload(writer)
	if err != SuccessCode {
		return
	}
	return
}

func DecodeMBSerialProtocol(reader io.Reader, actionType int) (MBSerialHeader, MBProtocol, byte) {
	if actionType == RequestAction {
		return decodeSerialRequestPDU(reader)
	}

	if actionType == ResponseAction {
		return decodeSerialResponsePDU(reader)
	}

	return nil, nil, IllegalData
}

func decodeSerialRequestPDU(reader io.Reader) (MBSerialHeader, MBProtocol, byte) {
	header := EmptySerialHeader()
	err := header.Decode(reader)
	if err != SuccessCode {
		return nil, nil, err
	}

	funcCode := make([]byte, 1)
	rSize, rErr := reader.Read(funcCode)
	if rErr != nil || rSize != 1 {
		err = IllegalAddress
		return nil, nil, err
	}

	var protocol MBProtocol
	switch funcCode[0] {
	case ReadCoils:
		protocol = EmptyReadCoilsReq()
	case ReadDiscreteInputs:
		protocol = EmptyReadDiscreteInputsReq()
	case ReadHoldingRegisters:
		protocol = EmptyReadHoldingRegistersReq()
	case ReadInputRegisters:
		protocol = EmptyReadInputRegistersReq()
	case WriteSingleCoil:
		protocol = EmptyWriteSingleCoilReq()
	case WriteSingleRegister:
		protocol = EmptyWriteSingleRegisterReq()
	case ReadExceptionStatus:
		protocol = EmptyReadExceptionStatusReq()
	case Diagnostics:
		protocol = EmptyDiagnosticsReq()
	case GetCommEventCounter:
		protocol = EmptyGetCommEventCounterReq()
	case GetCommEventLog:
		protocol = EmptyGetCommEventLogReq()
	case WriteMultipleCoils:
		protocol = EmptyWriteMultipleCoilsReq()
	case WriteMultipleRegisters:
		protocol = EmptyWriteMultipleRegistersReq()
	case ReportSlaveID:
		protocol = EmptyReportSlaveIDReq()
	case ReadFileRecord:
		protocol = EmptyReadFileRecordReq()
	case WriteFileRecord:
		protocol = EmptyWriteFileRecordReq()
	case MaskWriteRegister:
		protocol = EmptyMaskWriteRegisterReq()
	case ReadWriteMultipleRegisters:
		protocol = EmptyReadWriteMultipleRegistersReq()
	case ReadFIFOQueue:
		protocol = EmptyReadFIFOQueueReq()
	default:
		err = IllegalFuncCode
	}

	if err != SuccessCode {
		return nil, nil, err
	}
	if err != SuccessCode {
		return nil, nil, err
	}
	err = protocol.DecodePayload(reader)
	if err != SuccessCode {
		return nil, nil, err
	}

	return header, protocol, err
}

func decodeSerialResponsePDU(reader io.Reader) (MBSerialHeader, MBProtocol, byte) {
	header := EmptySerialHeader()
	err := header.Decode(reader)
	if err != SuccessCode {
		return nil, nil, err
	}

	funcCode := make([]byte, 1)
	rSize, rErr := reader.Read(funcCode)
	if rErr != nil || rSize != 1 {
		err = IllegalAddress
		return nil, nil, err
	}
	var exceptionCode byte
	lCode := funcCode[0]
	if funcCode[0] > 0x80 {
		exceptionRsp := EmptyExceptionRsp()
		err = exceptionRsp.DecodePayload(reader)
		if err != SuccessCode {
			return nil, nil, err
		}
		exceptionCode = exceptionRsp.ExceptionCode()
		lCode = funcCode[0] - 0x80
	}
	var protocol MBProtocol
	switch lCode {
	case ReadCoils:
		protocol = EmptyReadCoilsRsp(exceptionCode)
	case ReadDiscreteInputs:
		protocol = EmptyReadDiscreteInputsRsp(exceptionCode)
	case ReadHoldingRegisters:
		protocol = EmptyReadHoldingRegistersRsp(exceptionCode)
	case ReadInputRegisters:
		protocol = EmptyReadInputRegistersRsp(exceptionCode)
	case WriteSingleCoil:
		protocol = EmptyWriteSingleCoilRsp(exceptionCode)
	case WriteSingleRegister:
		protocol = EmptyWriteSingleRegisterRsp(exceptionCode)
	case ReadExceptionStatus:
		protocol = EmptyReadExceptionStatusRsp(exceptionCode)
	case Diagnostics:
		protocol = EmptyDiagnosticsRsp(exceptionCode)
	case GetCommEventCounter:
		protocol = EmptyGetCommEventCounterRsp(exceptionCode)
	case GetCommEventLog:
		protocol = EmptyGetCommEventLogRsp(exceptionCode)
	case WriteMultipleCoils:
		protocol = EmptyWriteMultipleCoilsRsp(exceptionCode)
	case WriteMultipleRegisters:
		protocol = EmptyWriteMultipleRegistersRsp(exceptionCode)
	case ReportSlaveID:
		protocol = EmptyReportSlaveIDRsp(exceptionCode)
	case ReadFileRecord:
		protocol = EmptyReadFileRecordRsp(exceptionCode)
	case WriteFileRecord:
		protocol = EmptyWriteFileRecordRsp(exceptionCode)
	case MaskWriteRegister:
		protocol = EmptyMaskWriteRegisterRsp(exceptionCode)
	case ReadWriteMultipleRegisters:
		protocol = EmptyReadWriteMultipleRegistersRsp(exceptionCode)
	case ReadFIFOQueue:
		protocol = EmptyReadFIFOQueueRsp(exceptionCode)
	default:
		err = IllegalFuncCode
	}

	if err != SuccessCode {
		return nil, nil, err
	}
	if funcCode[0] < 0x80 {
		err = protocol.DecodePayload(reader)
		if err != SuccessCode {
			return nil, nil, err
		}
	}

	return header, protocol, err
}
