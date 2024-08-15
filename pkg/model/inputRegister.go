package model

import (
	"encoding/binary"
	"io"
)

func NewReadInputRegistersReq(address, count uint16) *MBReadInputRegistersReq {
	return &MBReadInputRegistersReq{
		address: address,
		count:   count,
	}
}

func EmptyReadInputRegistersReq() *MBReadInputRegistersReq {
	return &MBReadInputRegistersReq{}
}

type MBReadInputRegistersReq struct {
	address uint16
	count   uint16
}

func (s *MBReadInputRegistersReq) FuncCode() byte {
	return ReadInputRegisters
}

func (s *MBReadInputRegistersReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D || s.count < 0x001 {
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

func (s *MBReadInputRegistersReq) Decode(reader io.Reader) (err byte) {
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

func (s *MBReadInputRegistersReq) CalcLen() uint16 {
	return 5
}

func (s *MBReadInputRegistersReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D || s.count < 0x001 {
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

func (s *MBReadInputRegistersReq) DecodePayload(reader io.Reader) (err byte) {
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

func (s *MBReadInputRegistersReq) CalcPayloadLen() uint16 {
	return 4
}

func (s *MBReadInputRegistersReq) Address() uint16 {
	return s.address
}

func (s *MBReadInputRegistersReq) Count() uint16 {
	return s.count
}

func NewReadInputRegistersRsp(data []byte) *MBReadInputRegistersRsp {
	return &MBReadInputRegistersRsp{
		data: data,
	}
}

func EmptyReadInputRegistersRsp() *MBReadInputRegistersRsp {
	return &MBReadInputRegistersRsp{}
}

type MBReadInputRegistersRsp struct {
	data []byte
}

func (s *MBReadInputRegistersRsp) FuncCode() byte {
	return ReadInputRegisters
}

func (s *MBReadInputRegistersRsp) Encode(writer io.Writer) (err byte) {
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

func (s *MBReadInputRegistersRsp) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	dataVal := make([]byte, 2)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 2 {
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

func (s *MBReadInputRegistersRsp) CalcLen() uint16 {
	return uint16(len(s.data)) + 2
}

func (s *MBReadInputRegistersRsp) EncodePayload(writer io.Writer) (err byte) {
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

func (s *MBReadInputRegistersRsp) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
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

func (s *MBReadInputRegistersRsp) CalcPayloadLen() uint16 {
	return uint16(len(s.data)) + 1
}

func (s *MBReadInputRegistersRsp) Data() []byte {
	return s.data
}

func NewWriteSingleRegisterReq(address uint16, data []byte) *MBWriteSingleRegisterReq {
	return &MBWriteSingleRegisterReq{
		address: address,
		data:    data,
	}
}

func EmptyWriteSingleRegisterReq() *MBWriteSingleRegisterReq {
	return &MBWriteSingleRegisterReq{}
}

type MBWriteSingleRegisterReq struct {
	address uint16
	data    []byte
}

func (s *MBWriteSingleRegisterReq) FuncCode() byte {
	return WriteSingleRegister
}

func (s *MBWriteSingleRegisterReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 5 {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBWriteSingleRegisterReq) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
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
	s.data = dataVal[3:5]
	return
}

func (s *MBWriteSingleRegisterReq) CalcLen() uint16 {
	return 5
}

func (s *MBWriteSingleRegisterReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 4 {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBWriteSingleRegisterReq) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	dataVal := make([]byte, 4)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 4 {
		err = IllegalAddress
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[0:2])
	s.data = dataVal[2:4]
	return
}

func (s *MBWriteSingleRegisterReq) CalcPayloadLen() uint16 {
	return 4
}

func (s *MBWriteSingleRegisterReq) Address() uint16 {
	return s.address
}

func (s *MBWriteSingleRegisterReq) Data() []byte {
	return s.data
}

func NewWriteSingleRegisterRsp(address uint16, data []byte) *MBWriteSingleRegisterRsp {
	return &MBWriteSingleRegisterRsp{
		address: address,
		data:    data,
	}
}

func EmptyWriteSingleRegisterRsp() *MBWriteSingleRegisterRsp {
	return &MBWriteSingleRegisterRsp{}
}

type MBWriteSingleRegisterRsp struct {
	address uint16
	data    []byte
}

func (s *MBWriteSingleRegisterRsp) FuncCode() byte {
	return WriteSingleRegister
}

func (s *MBWriteSingleRegisterRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 5 {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBWriteSingleRegisterRsp) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
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
	s.data = dataVal[3:5]
	return
}

func (s *MBWriteSingleRegisterRsp) CalcLen() uint16 {
	return 5
}

func (s *MBWriteSingleRegisterRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 4 {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBWriteSingleRegisterRsp) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	dataVal := make([]byte, 4)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 4 {
		err = IllegalAddress
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[0:2])
	s.data = dataVal[2:4]
	return
}

func (s *MBWriteSingleRegisterRsp) CalcPayloadLen() uint16 {
	return 4
}

func (s *MBWriteSingleRegisterRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteSingleRegisterRsp) Data() []byte {
	return s.data
}

func NewWriteMultipleRegistersReq(address, count uint16, data []byte) *MBWriteMultipleRegistersReq {
	return &MBWriteMultipleRegistersReq{
		address: address,
		count:   count,
		dataVal: data[:],
	}
}

func EmptyWriteMultipleRegistersReq() *MBWriteMultipleRegistersReq {
	return &MBWriteMultipleRegistersReq{}
}

type MBWriteMultipleRegistersReq struct {
	address uint16
	count   uint16
	dataVal []byte
}

func (s *MBWriteMultipleRegistersReq) FuncCode() byte {
	return WriteMultipleRegisters
}

func (s *MBWriteMultipleRegistersReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x0078 || s.count < 0x0001 {
		err = IllegalCount
		return
	}

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	buffVal = append(buffVal, byte(len(s.dataVal)))
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 6 {
		err = IllegalAddress
		return
	}
	wSize, wErr = writer.Write(s.dataVal)
	if wErr != nil || wSize != len(s.dataVal) {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBWriteMultipleRegistersReq) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

		if s.count > 0x0078 || s.count < 0x001 {
			err = IllegalCount
			return
		}
	}()
	dataVal := make([]byte, 6)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 6 {
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
	dataSize := dataVal[5]
	dataVal = make([]byte, dataSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != int(dataSize) {
		err = IllegalAddress
		return
	}

	s.dataVal = dataVal
	return
}

func (s *MBWriteMultipleRegistersReq) CalcLen() uint16 {
	return uint16(len(s.dataVal)) + 6
}

func (s *MBWriteMultipleRegistersReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x0078 || s.count < 0x0001 {
		err = IllegalCount
		return
	}

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	buffVal = append(buffVal, byte(len(s.dataVal)))
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 5 {
		err = IllegalAddress
		return
	}
	wSize, wErr = writer.Write(s.dataVal)
	if wErr != nil || wSize != 5 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBWriteMultipleRegistersReq) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

		if s.count > 0x0078 || s.count < 0x001 {
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

	s.address = binary.BigEndian.Uint16(dataVal[0:2])
	s.count = binary.BigEndian.Uint16(dataVal[2:4])
	dataSize := dataVal[4]
	dataVal = make([]byte, dataSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != int(dataSize) {
		err = IllegalAddress
		return
	}

	s.dataVal = dataVal
	return
}

func (s *MBWriteMultipleRegistersReq) CalcPayloadLen() uint16 {
	return uint16(len(s.dataVal)) + 5
}

func (s *MBWriteMultipleRegistersReq) Address() uint16 {
	return s.address
}

func (s *MBWriteMultipleRegistersReq) Count() uint16 {
	return s.count
}

func (s *MBWriteMultipleRegistersReq) Data() []byte {
	return s.dataVal
}

func NewWriteMultipleRegistersRsp(address, count uint16) *MBWriteMultipleRegistersRsp {
	return &MBWriteMultipleRegistersRsp{
		address: address,
		count:   count,
	}
}

func EmptyWriteMultipleRegistersRsp() *MBWriteMultipleRegistersRsp {
	return &MBWriteMultipleRegistersRsp{}
}

type MBWriteMultipleRegistersRsp struct {
	address uint16
	count   uint16
}

func (s *MBWriteMultipleRegistersRsp) FuncCode() byte {
	return WriteMultipleRegisters
}

func (s *MBWriteMultipleRegistersRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x0078 || s.count < 0x0001 {
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

func (s *MBWriteMultipleRegistersRsp) Decode(reader io.Reader) (err byte) {
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

func (s *MBWriteMultipleRegistersRsp) CalcLen() uint16 {
	return 5
}

func (s *MBWriteMultipleRegistersRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x0078 || s.count < 0x0001 {
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

func (s *MBWriteMultipleRegistersRsp) DecodePayload(reader io.Reader) (err byte) {
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

func (s *MBWriteMultipleRegistersRsp) CalcPayloadLen() uint16 {
	return 4
}

func (s *MBWriteMultipleRegistersRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteMultipleRegistersRsp) Count() uint16 {
	return s.count
}
