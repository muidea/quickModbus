package model

import (
	"bytes"
	"encoding/binary"
	"io"
)

func NewReadCoilsReq(address, count uint16) *MBReadCoilsReq {
	return &MBReadCoilsReq{
		address: address,
		count:   count,
	}
}

func EmptyReadCoilsReq() *MBReadCoilsReq {
	return &MBReadCoilsReq{}
}

const payloadSize = 4

type MBReadCoilsReq struct {
	address uint16
	count   uint16
}

func (s *MBReadCoilsReq) FuncCode() byte {
	return ReadCoils
}

func (s *MBReadCoilsReq) Encode(writer io.Writer) (err byte) {
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
	if wErr != nil || wSize != (payloadSize+1) {
		err = IllegalAddress
	}

	return
}

func (s *MBReadCoilsReq) EncodePayload(writer io.Writer) (err byte) {
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
	if wErr != nil || wSize != payloadSize {
		err = IllegalAddress
	}

	return
}

func (s *MBReadCoilsReq) Decode(reader io.Reader) (err byte) {
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

	dataVal := make([]byte, payloadSize+1)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != len(dataVal) {
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

func (s *MBReadCoilsReq) DecodePayload(reader io.Reader) (err byte) {
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

	dataVal := make([]byte, payloadSize)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != payloadSize {
		err = IllegalAddress
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[0:2])
	s.count = binary.BigEndian.Uint16(dataVal[2:4])
	return
}

func (s *MBReadCoilsReq) Address() uint16 {
	return s.address
}

func (s *MBReadCoilsReq) Count() uint16 {
	return s.count
}

func NewReadCoilsRsp(data []byte) *MBReadCoilsRsp {
	return &MBReadCoilsRsp{
		data: data,
	}
}

func EmptyReadCoilsRsp() *MBReadCoilsRsp {
	return &MBReadCoilsRsp{}
}

type MBReadCoilsRsp struct {
	data []byte
}

func (s *MBReadCoilsRsp) FuncCode() byte {
	return ReadCoils
}

func (s *MBReadCoilsRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = append(buffVal, byte(len(s.data)))
	buffVal = append(buffVal, s.data...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBReadCoilsRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, byte(len(s.data)))
	buffVal = append(buffVal, s.data...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBReadCoilsRsp) Decode(reader io.Reader) (err byte) {
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

func (s *MBReadCoilsRsp) DecodePayload(reader io.Reader) (err byte) {
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

func (s *MBReadCoilsRsp) Data() []byte {
	return s.data
}

func NewWriteSingleCoilReq(address uint16, data []byte) *MBWriteSingleCoilReq {
	return &MBWriteSingleCoilReq{
		address: address,
		data:    data,
	}
}

func EmptyWriteSingleCoilReq() *MBWriteSingleCoilReq {
	return &MBWriteSingleCoilReq{}
}

type MBWriteSingleCoilReq struct {
	address uint16
	data    []byte
}

func (s *MBWriteSingleCoilReq) FuncCode() byte {
	return WriteSingleCoil
}

func (s *MBWriteSingleCoilReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if bytes.Compare(s.data, coilON) != 0 && bytes.Compare(s.data, coilOFF) != 0 {
		err = IllegalData
		return
	}

	buffSize := 5
	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != buffSize {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBWriteSingleCoilReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if bytes.Compare(s.data, coilON) != 0 && bytes.Compare(s.data, coilOFF) != 0 {
		err = IllegalData
		return
	}

	buffSize := 4
	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != buffSize {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBWriteSingleCoilReq) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

		if bytes.Compare(s.data, coilON) != 0 && bytes.Compare(s.data, coilOFF) != 0 {
			err = IllegalData
			return
		}
	}()

	dataSize := 5
	dataVal := make([]byte, dataSize)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
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

func (s *MBWriteSingleCoilReq) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

		if bytes.Compare(s.data, coilON) != 0 && bytes.Compare(s.data, coilOFF) != 0 {
			err = IllegalData
			return
		}
	}()

	dataSize := 4
	dataVal := make([]byte, dataSize)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
		err = IllegalAddress
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[0:2])
	s.data = dataVal[2:4]
	return
}

func (s *MBWriteSingleCoilReq) Address() uint16 {
	return s.address
}

func (s *MBWriteSingleCoilReq) Data() []byte {
	return s.data
}

func NewWriteSingleCoilRsp(address uint16, data []byte) *MBWriteSingleCoilRsp {
	return &MBWriteSingleCoilRsp{
		address: address,
		data:    data,
	}
}

func EmptyWriteSingleCoilRsp() *MBWriteSingleCoilRsp {
	return &MBWriteSingleCoilRsp{}
}

type MBWriteSingleCoilRsp struct {
	address uint16
	data    []byte
}

func (s *MBWriteSingleCoilRsp) FuncCode() byte {
	return WriteSingleCoil
}

func (s *MBWriteSingleCoilRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if bytes.Compare(s.data, coilON) != 0 && bytes.Compare(s.data, coilOFF) != 0 {
		err = IllegalData
		return
	}

	buffSize := 5
	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != buffSize {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBWriteSingleCoilRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if bytes.Compare(s.data, coilON) != 0 && bytes.Compare(s.data, coilOFF) != 0 {
		err = IllegalData
		return
	}

	buffSize := 4
	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != buffSize {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBWriteSingleCoilRsp) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

		if bytes.Compare(s.data, coilON) != 0 && bytes.Compare(s.data, coilOFF) != 0 {
			err = IllegalData
			return
		}
	}()
	dataSize := 5
	dataVal := make([]byte, dataSize)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
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

func (s *MBWriteSingleCoilRsp) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

		if bytes.Compare(s.data, coilON) != 0 && bytes.Compare(s.data, coilOFF) != 0 {
			err = IllegalData
			return
		}
	}()
	dataSize := 4
	dataVal := make([]byte, dataSize)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
		err = IllegalAddress
		return
	}
	s.address = binary.BigEndian.Uint16(dataVal[0:2])
	s.data = dataVal[2:4]
	return
}

func (s *MBWriteSingleCoilRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteSingleCoilRsp) Data() []byte {
	return s.data
}

func NewWriteMultipleCoilsReq(address, count uint16, data []byte) *MBWriteMultipleCoilsReq {
	return &MBWriteMultipleCoilsReq{
		address: address,
		count:   count,
		dataVal: data[:],
	}
}

func EmptyWriteMultipleCoilsReq() *MBWriteMultipleCoilsReq {
	return &MBWriteMultipleCoilsReq{}
}

type MBWriteMultipleCoilsReq struct {
	address uint16
	count   uint16
	dataVal []byte
}

func (s *MBWriteMultipleCoilsReq) FuncCode() byte {
	return WriteMultipleCoils
}

func (s *MBWriteMultipleCoilsReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D0 || s.count < 0x001 {
		err = IllegalCount
		return
	}

	dataSize := 5 + len(s.dataVal)
	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	buffVal = append(buffVal, byte(len(s.dataVal)))
	buffVal = append(buffVal, s.dataVal...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != dataSize {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBWriteMultipleCoilsReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D0 || s.count < 0x001 {
		err = IllegalCount
		return
	}

	dataSize := 5 + len(s.dataVal)
	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	buffVal = append(buffVal, byte(len(s.dataVal)))
	buffVal = append(buffVal, s.dataVal...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != dataSize {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBWriteMultipleCoilsReq) Decode(reader io.Reader) (err byte) {
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
	dataVal = make([]byte, 1)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != 1 {
		err = IllegalAddress
		return
	}
	dataSize := dataVal[0]
	dataVal = make([]byte, dataSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != int(dataSize) {
		err = IllegalAddress
		return
	}
	s.dataVal = dataVal
	return
}

func (s *MBWriteMultipleCoilsReq) DecodePayload(reader io.Reader) (err byte) {
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
	dataVal = make([]byte, 1)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != 1 {
		err = IllegalAddress
		return
	}
	dataSize := dataVal[0]
	dataVal = make([]byte, dataSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != int(dataSize) {
		err = IllegalAddress
		return
	}
	s.dataVal = dataVal
	return
}

func (s *MBWriteMultipleCoilsReq) Address() uint16 {
	return s.address
}

func (s *MBWriteMultipleCoilsReq) Count() uint16 {
	return s.count
}

func (s *MBWriteMultipleCoilsReq) Data() []byte {
	return s.dataVal
}

func NewWriteMultipleCoilsRsp(address, count uint16) *MBWriteMultipleCoilsRsp {
	return &MBWriteMultipleCoilsRsp{
		address: address,
		count:   count,
	}
}

func EmptyWriteMultipleCoilsRsp() *MBWriteMultipleCoilsRsp {
	return &MBWriteMultipleCoilsRsp{}
}

type MBWriteMultipleCoilsRsp struct {
	address uint16
	count   uint16
}

func (s *MBWriteMultipleCoilsRsp) FuncCode() byte {
	return WriteMultipleCoils
}

func (s *MBWriteMultipleCoilsRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D0 || s.count < 0x001 {
		err = IllegalCount
		return
	}

	dataSize := 5
	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != dataSize {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBWriteMultipleCoilsRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D0 || s.count < 0x001 {
		err = IllegalCount
		return
	}

	dataSize := 4
	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != dataSize {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBWriteMultipleCoilsRsp) Decode(reader io.Reader) (err byte) {
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

func (s *MBWriteMultipleCoilsRsp) DecodePayload(reader io.Reader) (err byte) {
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

func (s *MBWriteMultipleCoilsRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteMultipleCoilsRsp) Count() uint16 {
	return s.count
}
