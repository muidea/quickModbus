package model

import (
	"encoding/binary"
	"io"
)

func NewReadDiscreteInputsReq(address, count uint16) *MBReadDiscreteInputsReq {
	return &MBReadDiscreteInputsReq{
		address: address,
		count:   count,
	}
}

func EmptyReadDiscreteInputsReq() *MBReadDiscreteInputsReq {
	return &MBReadDiscreteInputsReq{}
}

type MBReadDiscreteInputsReq struct {
	address uint16
	count   uint16
}

func (s *MBReadDiscreteInputsReq) FuncCode() byte {
	return ReadDiscreteInputs
}

func (s *MBReadDiscreteInputsReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D0 || s.count < 0x001 {
		err = IllegalCount
		return
	}

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 5 {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadDiscreteInputsReq) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

		if s.count > 0x7D0 || s.count < 0x001 {
			err = IllegalCount
			return
		}
	}()
	dataVal := make([]byte, 5)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 5 {
		err = IllegalAddress
		return
	}
	funcCode := dataVal[0]
	if funcCode != s.FuncCode() {
		err = IllegalData
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[1:3])
	s.count = binary.BigEndian.Uint16(dataVal[3:5])
	return
}

func (s *MBReadDiscreteInputsReq) CalcLen() uint16 {
	return 5
}

func (s *MBReadDiscreteInputsReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D0 || s.count < 0x001 {
		err = IllegalCount
		return
	}

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 4 {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadDiscreteInputsReq) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

		if s.count > 0x7D0 || s.count < 0x001 {
			err = IllegalCount
			return
		}
	}()
	dataVal := make([]byte, 4)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 4 {
		err = IllegalAddress
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[0:2])
	s.count = binary.BigEndian.Uint16(dataVal[2:4])
	return
}

func (s *MBReadDiscreteInputsReq) CalcPayloadLen() uint16 {
	return 4
}

func (s *MBReadDiscreteInputsReq) Address() uint16 {
	return s.address
}

func (s *MBReadDiscreteInputsReq) Count() uint16 {
	return s.count
}

func NewReadDiscreteInputsRsp(count byte, data []byte) *MBReadDiscreteInputsRsp {
	return &MBReadDiscreteInputsRsp{
		data: data,
	}
}

func EmptyReadDiscreteInputsRsp(exceptionCode byte) *MBReadDiscreteInputsRsp {
	return &MBReadDiscreteInputsRsp{
		exceptionCode: exceptionCode,
	}
}

type MBReadDiscreteInputsRsp struct {
	exceptionCode byte
	data          []byte
}

func (s *MBReadDiscreteInputsRsp) FuncCode() byte {
	return ReadDiscreteInputs
}

func (s *MBReadDiscreteInputsRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBReadDiscreteInputsRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = append(buffVal, byte(len(s.data)))
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 2 {
		err = IllegalAddress
		return
	}
	wSize, wErr = writer.Write(s.data)
	if wErr != nil || wSize != len(s.data) {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBReadDiscreteInputsRsp) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

	}()
	dataVal := make([]byte, 2)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 1 {
		err = IllegalAddress
		return
	}
	funcCode := dataVal[0]
	if funcCode != s.FuncCode() {
		err = IllegalData
		return
	}

	dataSize := int(dataVal[1])
	dataVal = make([]byte, dataSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
		err = IllegalAddress
		return
	}

	s.data = dataVal
	return
}

func (s *MBReadDiscreteInputsRsp) CalcLen() uint16 {
	return uint16(len(s.data)) + 2
}

func (s *MBReadDiscreteInputsRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	wSize, wErr := writer.Write([]byte{byte(len(s.data))})
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}
	wSize, wErr = writer.Write(s.data)
	if wErr != nil || wSize != len(s.data) {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBReadDiscreteInputsRsp) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

	}()
	dataVal := make([]byte, 1)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 1 {
		err = IllegalAddress
		return
	}

	dataSize := int(dataVal[0])
	dataVal = make([]byte, dataSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
		err = IllegalAddress
		return
	}

	s.data = dataVal
	return
}

func (s *MBReadDiscreteInputsRsp) CalcPayloadLen() uint16 {
	return uint16(len(s.data)) + 1
}

func (s *MBReadDiscreteInputsRsp) Data() []byte {
	return s.data
}
