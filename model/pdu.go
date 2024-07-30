package model

import (
	"bytes"
	"encoding/binary"
)

type MBProtocol interface {
	Encode(buffVal []byte) (ret []byte, err byte)
	Decode(byteData []byte) (err byte)
	Length() uint16
	FuncCode() byte
}

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

func (s *MBReadCoilsReq) Address() uint16 {
	return s.address
}

func (s *MBReadCoilsReq) Count() uint16 {
	return s.count
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

func (s *MBReadCoilsReq) Length() uint16 {
	return pduReqHeadLength
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

func (s *MBReadCoilsRsp) Count() byte {
	return s.count
}

func (s *MBReadCoilsRsp) Data() []byte {
	return s.data
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

func (s *MBReadCoilsRsp) Length() uint16 {
	return 2 + uint16(s.count)
}

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

func (s *MBReadDiscreteInputsReq) Address() uint16 {
	return s.address
}

func (s *MBReadDiscreteInputsReq) Count() uint16 {
	return s.count
}

func (s *MBReadDiscreteInputsReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D0 || s.count < 0x001 {
		err = IllegalCount
		return
	}

	buffVal = append(buffVal, ReadDiscreteInputs)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)

	ret = buffVal
	return
}

func (s *MBReadDiscreteInputsReq) Decode(byteData []byte) (err byte) {
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
	if funcCode != ReadDiscreteInputs {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.count = binary.BigEndian.Uint16(byteData[3:5])
	return
}

func (s *MBReadDiscreteInputsReq) Length() uint16 {
	return pduReqHeadLength
}

func NewReadDiscreteInputsRsp(count byte, data []byte) *MBReadDiscreteInputsRsp {
	return &MBReadDiscreteInputsRsp{
		count: count,
		data:  data,
	}
}

func EmptyReadDiscreteInputsRsp() *MBReadDiscreteInputsRsp {
	return &MBReadDiscreteInputsRsp{}
}

type MBReadDiscreteInputsRsp struct {
	count byte
	data  []byte
}

func (s *MBReadDiscreteInputsRsp) FuncCode() byte {
	return ReadDiscreteInputs
}

func (s *MBReadDiscreteInputsRsp) Count() byte {
	return s.count
}

func (s *MBReadDiscreteInputsRsp) Data() []byte {
	return s.data
}

func (s *MBReadDiscreteInputsRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadDiscreteInputs)
	buffVal = append(buffVal, s.count)
	buffVal = append(buffVal, s.data...)

	ret = buffVal
	return
}

func (s *MBReadDiscreteInputsRsp) Decode(byteData []byte) (err byte) {
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
	if funcCode != ReadDiscreteInputs {
		err = IllegalFuncCode
		return
	}

	s.count = byteData[1]
	s.data = byteData[2:]
	return
}

func (s *MBReadDiscreteInputsRsp) Length() uint16 {
	return 2 + uint16(s.count)
}

func NewReadHoldingRegistersReq(address, count uint16) *MBReadHoldingRegistersReq {
	return &MBReadHoldingRegistersReq{
		address: address,
		count:   count,
	}
}

func EmptyReadHoldingRegistersReq() *MBReadHoldingRegistersReq {
	return &MBReadHoldingRegistersReq{}
}

type MBReadHoldingRegistersReq struct {
	address uint16
	count   uint16
}

func (s *MBReadHoldingRegistersReq) FuncCode() byte {
	return ReadHoldingRegisters
}

func (s *MBReadHoldingRegistersReq) Address() uint16 {
	return s.address
}

func (s *MBReadHoldingRegistersReq) Count() uint16 {
	return s.count
}

func (s *MBReadHoldingRegistersReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	if s.count > 0x7D || s.count < 0x001 {
		err = IllegalCount
		return
	}

	buffVal = append(buffVal, ReadHoldingRegisters)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.count)

	ret = buffVal
	return
}

func (s *MBReadHoldingRegistersReq) Decode(byteData []byte) (err byte) {
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
	if funcCode != ReadHoldingRegisters {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.count = binary.BigEndian.Uint16(byteData[3:5])
	return
}

func (s *MBReadHoldingRegistersReq) Length() uint16 {
	return pduReqHeadLength
}

func NewReadHoldingRegistersRsp(count byte, data []byte) *MBReadHoldingRegistersRsp {
	return &MBReadHoldingRegistersRsp{
		count: count,
		data:  data,
	}
}

func EmptyReadHoldingRegistersRsp() *MBReadHoldingRegistersRsp {
	return &MBReadHoldingRegistersRsp{}
}

type MBReadHoldingRegistersRsp struct {
	count byte
	data  []byte
}

func (s *MBReadHoldingRegistersRsp) FuncCode() byte {
	return ReadHoldingRegisters
}

func (s *MBReadHoldingRegistersRsp) Count() byte {
	return s.count
}

func (s *MBReadHoldingRegistersRsp) Data() []byte {
	return s.data
}

func (s *MBReadHoldingRegistersRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadHoldingRegisters)
	buffVal = append(buffVal, s.count)
	buffVal = append(buffVal, s.data...)

	ret = buffVal
	return
}

func (s *MBReadHoldingRegistersRsp) Decode(byteData []byte) (err byte) {
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
	if funcCode != ReadHoldingRegisters {
		err = IllegalFuncCode
		return
	}

	s.count = byteData[1]
	s.data = byteData[2 : s.count+2]
	return
}

func (s *MBReadHoldingRegistersRsp) Length() uint16 {
	return 2 + uint16(s.count)
}

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

func (s *MBReadInputRegistersReq) Address() uint16 {
	return s.address
}

func (s *MBReadInputRegistersReq) Count() uint16 {
	return s.count
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

func (s *MBReadInputRegistersReq) Length() uint16 {
	return pduReqHeadLength
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

func (s *MBReadInputRegistersRsp) Count() byte {
	return s.count
}

func (s *MBReadInputRegistersRsp) Data() []byte {
	return s.data
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

func (s *MBReadInputRegistersRsp) Length() uint16 {
	return 2 + uint16(s.count)
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

func (s *MBWriteSingleCoilReq) Address() uint16 {
	return s.address
}

func (s *MBWriteSingleCoilReq) Data() []byte {
	return s.data
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

func (s *MBWriteSingleCoilReq) Length() uint16 {
	return pduReqHeadLength
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

func (s *MBWriteSingleCoilRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteSingleCoilRsp) Data() []byte {
	return s.data
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

func (s *MBWriteSingleCoilRsp) Length() uint16 {
	return pduReqHeadLength
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
	s.dataVal = byteData[5 : 5+int(s.dataSize)]
	return
}

func (s *MBWriteMultipleCoilsReq) Length() uint16 {
	return pduReqHeadLength + 1 + uint16(s.dataSize)
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

func (s *MBWriteMultipleCoilsRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteMultipleCoilsRsp) Count() uint16 {
	return s.count
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

func (s *MBWriteMultipleCoilsRsp) Length() uint16 {
	return pduReqHeadLength
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

func (s *MBWriteSingleRegisterReq) Address() uint16 {
	return s.address
}

func (s *MBWriteSingleRegisterReq) Data() []byte {
	return s.data
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

func (s *MBWriteSingleRegisterReq) Length() uint16 {
	return pduReqHeadLength
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

func (s *MBWriteSingleRegisterRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteSingleRegisterRsp) Data() []byte {
	return s.data
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

func (s *MBWriteSingleRegisterRsp) Length() uint16 {
	return pduReqHeadLength
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

func (s *MBWriteMultipleRegistersReq) Length() uint16 {
	return pduReqHeadLength + 1 + uint16(s.dataSize)
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

func (s *MBWriteMultipleRegistersRsp) Address() uint16 {
	return s.address
}

func (s *MBWriteMultipleRegistersRsp) Count() uint16 {
	return s.count
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

func (s *MBWriteMultipleRegistersRsp) Length() uint16 {
	return pduReqHeadLength
}

func NewExceptionRsp(funcCode, exceptionCode byte) *MBExceptionRsp {
	return &MBExceptionRsp{
		funcCode:      funcCode,
		exceptionCode: exceptionCode,
	}
}

func EmptyExceptionRsp() *MBExceptionRsp {
	return &MBExceptionRsp{}
}

type MBExceptionRsp struct {
	funcCode      byte
	exceptionCode byte
}

func (s *MBExceptionRsp) FuncCode() byte {
	return s.funcCode
}

func (s *MBExceptionRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBExceptionRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, s.funcCode)
	buffVal = append(buffVal, s.exceptionCode)

	ret = buffVal
	return
}

func (s *MBExceptionRsp) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < minRspDataLength {
		err = IllegalData
		return
	}

	s.funcCode = byteData[0]
	s.exceptionCode = byteData[1]
	return
}

func (s *MBExceptionRsp) Length() uint16 {
	return 2
}
