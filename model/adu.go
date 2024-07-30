package model

import (
	"encoding/binary"
	"github.com/muidea/magicCommon/foundation/log"
)

const MaxAduSize = 256

type MBTcpHeader interface {
	Transaction() uint16
	Protocol() uint16
	DataLen() uint16
	UnitID() byte
	Encode(buffVal []byte) (ret []byte, err byte)
	Decode(byteData []byte) (err byte)
	Length() uint16
}

type MBSerialHeader interface {
	Address() byte
	Encode(buffVal []byte) (ret []byte, err byte)
	Decode(byteData []byte) (err byte)
	Length() uint16
}

type mbTcpHeader struct {
	transaction uint16
	protocol    uint16
	dataLen     uint16
	unitID      byte
}

func NewTcpHeader(transaction, dataLen uint16, unitID byte) MBTcpHeader {
	return &mbTcpHeader{
		transaction: transaction,
		dataLen:     dataLen,
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
	return s.dataLen
}

func (s *mbTcpHeader) UnitID() byte {
	return s.unitID
}

func (s *mbTcpHeader) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = binary.BigEndian.AppendUint16(buffVal, s.transaction)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.protocol)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.dataLen)
	buffVal = append(buffVal, s.unitID)

	ret = buffVal
	return
}

func (s *mbTcpHeader) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < aduTcpHeadLength {
		err = IllegalData
		return
	}

	s.transaction = binary.BigEndian.Uint16(byteData[0:2])
	s.protocol = binary.BigEndian.Uint16(byteData[2:4])
	s.dataLen = binary.BigEndian.Uint16(byteData[4:6])
	s.unitID = byteData[6]
	return
}

func (s *mbTcpHeader) Length() uint16 {
	return aduTcpHeadLength
}

func (s *mbTcpHeader) Same(ptr *mbTcpHeader) bool {
	if ptr == nil {
		return false
	}

	return s.transaction == ptr.transaction &&
		s.protocol == ptr.protocol &&
		s.dataLen == ptr.dataLen &&
		s.unitID == ptr.unitID
}

type mbSerialHeader struct {
	address byte
}

func NewSerialHeader() MBSerialHeader {
	return &mbSerialHeader{}
}

func EmptySerialHeader() MBSerialHeader {
	return &mbSerialHeader{}
}

func (s *mbSerialHeader) Address() byte {
	return s.address
}

func (s *mbSerialHeader) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, s.address)

	ret = buffVal
	return
}

func (s *mbSerialHeader) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < aduSerialHeadLength {
		err = IllegalData
		return
	}

	s.address = byteData[0]
	return
}

func (s *mbSerialHeader) Length() uint16 {
	return aduSerialHeadLength
}

func EncodeMBProtocol(header MBTcpHeader, pdu MBProtocol, buffVal []byte) (ret []byte, err byte) {
	buffVal, err = header.Encode(buffVal)
	if err != SuccessCode {
		return
	}

	buffVal, err = pdu.Encode(buffVal)
	if err != SuccessCode {
		return
	}

	buffVal = append(buffVal, buffVal...)
	buffVal = append(buffVal, buffVal...)
	ret = buffVal
	return
}

func DecodeMBProtocol(bytesData []byte, actionType int) (MBTcpHeader, MBProtocol, byte) {
	if len(bytesData) < aduTcpHeadLength+1 {
		log.Errorf("not enough data, data size:%d", len(bytesData))
		return nil, nil, IllegalData
	}

	if actionType == RequestAction {
		return decodeRequestPDU(bytesData)
	}

	if actionType == ResponseAction {
		return decodeResponsePDU(bytesData)
	}

	return nil, nil, IllegalData
}

func decodeRequestPDU(bytesData []byte) (MBTcpHeader, MBProtocol, byte) {
	var header MBTcpHeader
	var protocol MBProtocol
	var err byte
	switch bytesData[aduTcpHeadLength] {
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
	case WriteMultipleCoils:
		protocol = EmptyWriteMultipleCoilsReq()
	case WriteMultipleRegisters:
		protocol = EmptyWriteMultipleRegistersReq()
	default:
		err = IllegalFuncCode
	}

	if err != SuccessCode {
		return nil, nil, err
	}
	header = EmptyTcpHeader()
	err = header.Decode(bytesData)
	if err != SuccessCode {
		return nil, nil, err
	}
	err = protocol.Decode(bytesData[header.Length():])
	if err != SuccessCode {
		return nil, nil, err
	}

	return header, protocol, err
}

func decodeResponsePDU(bytesData []byte) (MBTcpHeader, MBProtocol, byte) {
	var header MBTcpHeader
	var protocol MBProtocol
	var err byte
	switch bytesData[aduTcpHeadLength] {
	case ReadCoils:
		protocol = EmptyReadCoilsRsp()
	case ReadDiscreteInputs:
		protocol = EmptyReadDiscreteInputsRsp()
	case ReadHoldingRegisters:
		protocol = EmptyReadHoldingRegistersRsp()
	case ReadInputRegisters:
		protocol = EmptyReadInputRegistersRsp()
	case WriteSingleCoil:
		protocol = EmptyWriteSingleCoilRsp()
	case WriteSingleRegister:
		protocol = EmptyWriteSingleRegisterRsp()
	case WriteMultipleCoils:
		protocol = EmptyWriteMultipleCoilsRsp()
	case WriteMultipleRegisters:
		protocol = EmptyWriteMultipleRegistersRsp()
	default:
		err = IllegalFuncCode
	}

	if err != SuccessCode {
		return nil, nil, err
	}
	header = EmptyTcpHeader()
	err = header.Decode(bytesData)
	if err != SuccessCode {
		return nil, nil, err
	}
	err = protocol.Decode(bytesData[header.Length():])
	if err != SuccessCode {
		return nil, nil, err
	}

	return header, protocol, err
}
