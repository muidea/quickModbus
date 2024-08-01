package model

import "encoding/binary"

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

func (s *MBReadHoldingRegistersReq) Address() uint16 {
	return s.address
}

func (s *MBReadHoldingRegistersReq) Count() uint16 {
	return s.count
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

func (s *MBReadHoldingRegistersRsp) Count() byte {
	return s.count
}

func (s *MBReadHoldingRegistersRsp) Data() []byte {
	return s.data
}
