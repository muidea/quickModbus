package model

import "encoding/binary"

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

	buffVal = append(buffVal, s.FuncCode())
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
	if funcCode != s.FuncCode() {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.count = binary.BigEndian.Uint16(byteData[3:5])
	return
}

func (s *MBReadDiscreteInputsReq) Address() uint16 {
	return s.address
}

func (s *MBReadDiscreteInputsReq) Count() uint16 {
	return s.count
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

func (s *MBReadDiscreteInputsRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, s.FuncCode())
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
	if funcCode != s.FuncCode() {
		err = IllegalFuncCode
		return
	}

	s.count = byteData[1]
	s.data = byteData[2:]
	return
}

func (s *MBReadDiscreteInputsRsp) Count() byte {
	return s.count
}

func (s *MBReadDiscreteInputsRsp) Data() []byte {
	return s.data
}
