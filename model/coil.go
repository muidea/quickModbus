package model

import (
	"bytes"
	"encoding/binary"
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

type MBReadCoilsReq struct {
	address uint16
	count   uint16
}

func (s *MBReadCoilsReq) FuncCode() byte {
	return ReadCoils
}

func (s *MBReadCoilsReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D0 || s.count < 0x001 {
		err = IllegalCount
		return
	}

	buffVal = append(buffVal, ReadCoils)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)

	ret = buffVal
	return
}

func (s *MBReadCoilsReq) Decode(byteData []byte) (err byte) {
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
	if len(byteData) < minReqDataLength {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != ReadCoils {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.count = binary.BigEndian.Uint16(byteData[3:5])
	return
}

func (s *MBReadCoilsReq) Address() uint16 {
	return s.address
}

func (s *MBReadCoilsReq) Count() uint16 {
	return s.count
}

func NewReadCoilsRsp(count byte, data []byte) *MBReadCoilsRsp {
	return &MBReadCoilsRsp{
		count: count,
		data:  data,
	}
}

func EmptyReadCoilsRsp() *MBReadCoilsRsp {
	return &MBReadCoilsRsp{}
}

type MBReadCoilsRsp struct {
	count byte
	data  []byte
}

func (s *MBReadCoilsRsp) FuncCode() byte {
	return ReadCoils
}

func (s *MBReadCoilsRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadCoils)
	buffVal = append(buffVal, s.count)
	buffVal = append(buffVal, s.data...)

	ret = buffVal
	return
}

func (s *MBReadCoilsRsp) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}

	}()
	if len(byteData) < minRspDataLength {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != ReadCoils {
		err = IllegalFuncCode
		return
	}

	s.count = byteData[1]
	s.data = byteData[2:]
	return
}

func (s *MBReadCoilsRsp) Count() byte {
	return s.count
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

func (s *MBWriteSingleCoilReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if bytes.Compare(s.data, coilON) != 0 && bytes.Compare(s.data, coilOFF) != 0 {
		err = IllegalData
		return
	}

	buffVal = append(buffVal, WriteSingleCoil)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data...)

	ret = buffVal
	return
}

func (s *MBWriteSingleCoilReq) Decode(byteData []byte) (err byte) {
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
	if len(byteData) < minReqDataLength {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != WriteSingleCoil {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.data = byteData[3:5]
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

func (s *MBWriteSingleCoilRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if bytes.Compare(s.data, coilON) != 0 && bytes.Compare(s.data, coilOFF) != 0 {
		err = IllegalData
		return
	}

	buffVal = append(buffVal, WriteSingleCoil)
	buffVal = binary.BigEndian.AppendUint16(ret, s.address)
	buffVal = append(buffVal, s.data...)

	ret = buffVal
	return
}

func (s *MBWriteSingleCoilRsp) Decode(byteData []byte) (err byte) {
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
	if len(byteData) < minReqDataLength {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != WriteSingleCoil {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.data = byteData[3:5]
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
		address:  address,
		count:    count,
		dataSize: byte(len(data)),
		dataVal:  data[:],
	}
}

func EmptyWriteMultipleCoilsReq() *MBWriteMultipleCoilsReq {
	return &MBWriteMultipleCoilsReq{}
}

type MBWriteMultipleCoilsReq struct {
	address  uint16
	count    uint16
	dataSize byte
	dataVal  []byte
}

func (s *MBWriteMultipleCoilsReq) FuncCode() byte {
	return WriteMultipleCoils
}

func (s *MBWriteMultipleCoilsReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D0 || s.count < 0x001 || int(s.dataSize) > len(s.dataVal) || int(s.dataSize) < 0 {
		err = IllegalCount
		return
	}

	buffVal = append(buffVal, WriteMultipleCoils)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	buffVal = append(buffVal, s.dataSize)
	dataVal := s.dataVal[:int(s.dataSize)]
	buffVal = append(buffVal, dataVal...)

	ret = buffVal
	return
}

func (s *MBWriteMultipleCoilsReq) Decode(byteData []byte) (err byte) {
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
	if len(byteData) < minReqDataLength+1 {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != WriteMultipleCoils {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.count = binary.BigEndian.Uint16(byteData[3:5])
	s.dataSize = byteData[5]
	s.dataVal = byteData[6 : 6+int(s.dataSize)]
	return
}

func (s *MBWriteMultipleCoilsReq) Address() uint16 {
	return s.address
}

func (s *MBWriteMultipleCoilsReq) Count() uint16 {
	return s.count
}

func (s *MBWriteMultipleCoilsReq) DataSize() byte {
	return s.dataSize
}

func (s *MBWriteMultipleCoilsReq) Data() []byte {
	return s.dataVal[:int(s.dataSize)]
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

func (s *MBWriteMultipleCoilsRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D0 || s.count < 0x001 {
		err = IllegalCount
		return
	}

	buffVal = append(buffVal, WriteMultipleCoils)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)

	ret = buffVal
	return
}

func (s *MBWriteMultipleCoilsRsp) Decode(byteData []byte) (err byte) {
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
	if len(byteData) < minReqDataLength {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != WriteMultipleCoils {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.count = binary.BigEndian.Uint16(byteData[3:5])
	return
}

func (s *MBWriteMultipleCoilsRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteMultipleCoilsRsp) Count() uint16 {
	return s.count
}
