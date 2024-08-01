package model

import "encoding/binary"

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

func (s *MBReadInputRegistersReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D || s.count < 0x001 {
		err = IllegalCount
		return
	}

	buffVal = append(buffVal, ReadInputRegisters)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	ret = buffVal
	return
}

func (s *MBReadInputRegistersReq) Decode(byteData []byte) (err byte) {
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
	if funcCode != ReadInputRegisters {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.count = binary.BigEndian.Uint16(byteData[3:5])
	return
}

func (s *MBReadInputRegistersReq) Address() uint16 {
	return s.address
}

func (s *MBReadInputRegistersReq) Count() uint16 {
	return s.count
}

func NewReadInputRegistersRsp(count byte, data []byte) *MBReadInputRegistersRsp {
	return &MBReadInputRegistersRsp{
		count: count,
		data:  data,
	}
}

func EmptyReadInputRegistersRsp() *MBReadInputRegistersRsp {
	return &MBReadInputRegistersRsp{}
}

type MBReadInputRegistersRsp struct {
	count byte
	data  []byte
}

func (s *MBReadInputRegistersRsp) FuncCode() byte {
	return ReadInputRegisters
}

func (s *MBReadInputRegistersRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadInputRegisters)
	buffVal = append(buffVal, s.count)
	buffVal = append(buffVal, s.data...)

	ret = buffVal
	return
}

func (s *MBReadInputRegistersRsp) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < minRspDataLength {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != ReadInputRegisters {
		err = IllegalFuncCode
		return
	}

	s.count = byteData[1]
	s.data = byteData[2 : s.count+2]

	return
}

func (s *MBReadInputRegistersRsp) Count() byte {
	return s.count
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

func (s *MBWriteSingleRegisterReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, WriteSingleRegister)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data...)

	ret = buffVal
	return
}

func (s *MBWriteSingleRegisterReq) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()
	if len(byteData) < minReqDataLength {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != WriteSingleRegister {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.data = byteData[3:5]
	return
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

func (s *MBWriteSingleRegisterRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, WriteSingleRegister)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.data...)

	ret = buffVal
	return
}

func (s *MBWriteSingleRegisterRsp) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < minRspDataLength {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != WriteSingleRegister {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.data = byteData[3:5]
	return
}

func (s *MBWriteSingleRegisterRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteSingleRegisterRsp) Data() []byte {
	return s.data
}

func NewWriteMultipleRegistersReq(address, count uint16, data []byte) *MBWriteMultipleRegistersReq {
	return &MBWriteMultipleRegistersReq{
		address:  address,
		count:    count,
		dataSize: byte(len(data)),
		dataVal:  data[:],
	}
}

func EmptyWriteMultipleRegistersReq() *MBWriteMultipleRegistersReq {
	return &MBWriteMultipleRegistersReq{}
}

type MBWriteMultipleRegistersReq struct {
	address  uint16
	count    uint16
	dataSize byte
	dataVal  []byte
}

func (s *MBWriteMultipleRegistersReq) FuncCode() byte {
	return WriteMultipleRegisters
}

func (s *MBWriteMultipleRegistersReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x0078 || s.count < 0x0001 || int(s.dataSize) > len(s.dataVal) || int(s.dataSize) < 0 {
		err = IllegalCount
		return
	}

	buffVal = append(buffVal, WriteMultipleRegisters)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)
	buffVal = append(buffVal, s.dataSize)
	dataVal := s.dataVal[:int(s.dataSize)]
	buffVal = append(buffVal, dataVal...)

	ret = buffVal
	return
}

func (s *MBWriteMultipleRegistersReq) Decode(byteData []byte) (err byte) {
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
	if len(byteData) < minReqDataLength+1 {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != WriteMultipleRegisters {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.count = binary.BigEndian.Uint16(byteData[3:5])
	s.dataSize = byteData[5]
	s.dataVal = byteData[6 : 6+s.count]
	return
}

func (s *MBWriteMultipleRegistersReq) Address() uint16 {
	return s.address
}

func (s *MBWriteMultipleRegistersReq) Count() uint16 {
	return s.count
}

func (s *MBWriteMultipleRegistersReq) DataSize() byte {
	return s.dataSize
}

func (s *MBWriteMultipleRegistersReq) Data() []byte {
	return s.dataVal[:int(s.dataSize)]
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

func (s *MBWriteMultipleRegistersRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x0078 || s.count < 0x0001 {
		err = IllegalCount
		return
	}

	buffVal = append(buffVal, WriteMultipleRegisters)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)

	ret = buffVal
	return
}

func (s *MBWriteMultipleRegistersRsp) Decode(byteData []byte) (err byte) {
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
	if funcCode != WriteMultipleRegisters {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.count = binary.BigEndian.Uint16(byteData[3:5])
	return
}

func (s *MBWriteMultipleRegistersRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteMultipleRegistersRsp) Count() uint16 {
	return s.count
}
