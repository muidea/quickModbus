package model

import (
	"encoding/binary"
)

type MBProtocol interface {
	FuncCode() byte
	Encode(buffVal []byte) (ret []byte, err byte)
	Decode(byteData []byte) (err byte)
}

type MBReadExceptionStatusReq struct {
}

func (s *MBReadExceptionStatusReq) FuncCode() byte {
	return ReadExceptionStatus
}

func (s *MBReadExceptionStatusReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadExceptionStatus)

	ret = buffVal
	return
}

func (s *MBReadExceptionStatusReq) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	funcCode := byteData[0]
	if funcCode != ReadExceptionStatus {
		err = IllegalFuncCode
		return
	}
	return
}

type MBReadExceptionStatusRsp struct {
	statusVal byte
}

func (s *MBReadExceptionStatusRsp) FuncCode() byte {
	return ReadExceptionStatus
}

func (s *MBReadExceptionStatusRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadExceptionStatus)
	buffVal = append(buffVal, s.statusVal)

	ret = buffVal
	return
}

func (s *MBReadExceptionStatusRsp) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < 2 {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != ReadExceptionStatus {
		err = IllegalFuncCode
		return
	}

	s.statusVal = byteData[1]
	return
}

func (s *MBReadExceptionStatusRsp) Status() byte {
	return s.statusVal
}

type MBDiagnosticsReq struct {
	dataFunction uint16
	dataVal      []byte
}

func (s *MBDiagnosticsReq) FuncCode() byte {
	return Diagnostics
}

func (s *MBDiagnosticsReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.dataFunction)
	buffVal = append(buffVal, s.dataVal...)

	ret = buffVal
	return
}

type MBDiagnosticsRsp struct {
}

func (s *MBDiagnosticsRsp) FuncCode() byte {
	return Diagnostics
}

type MBGetCommEventCounterReq struct {
}

func (s *MBGetCommEventCounterReq) FuncCode() byte {
	return GetCommEventCounter
}

type MBGetCommEventCounterRsp struct {
}

func (s *MBGetCommEventCounterRsp) FuncCode() byte {
	return GetCommEventCounter
}

type MBGetCommEventLogReq struct {
}

func (s *MBGetCommEventLogReq) FuncCode() byte {
	return GetCommEventLog
}

type MBGetCommEventLogRsp struct {
}

func (s *MBGetCommEventLogRsp) FuncCode() byte {
	return GetCommEventLog
}

type MBReportServerIDReq struct {
}

func (s *MBReportServerIDReq) FuncCode() byte {
	return ReportServerID
}

type MBReportServerIDRsp struct {
}

func (s *MBReportServerIDRsp) FuncCode() byte {
	return ReportServerID
}

func NewReadFileRecordReq() *MBReadFileRecordReq {
	return &MBReadFileRecordReq{}
}

func EmptyReadFileRecordReq() *MBReadFileRecordReq {
	return &MBReadFileRecordReq{}
}

type ReadRequestItem struct {
	referenceType byte   // 6
	fileNumber    uint16 // 2 bytes
	recordNumber  uint16 // 2 bytes
	recordLength  uint16 // 2 bytes
}

func (s *ReadRequestItem) FileNumber() uint16 {
	return s.fileNumber
}

func (s *ReadRequestItem) RecordNumber() uint16 {
	return s.recordNumber
}

func (s *ReadRequestItem) RecordLength() uint16 {
	return s.recordLength
}

func (s *ReadRequestItem) calcDataSize() byte {
	return 7
}

func (s *ReadRequestItem) encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, s.referenceType)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.fileNumber)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.recordNumber)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.recordLength)
	ret = buffVal
	return
}

func (s *ReadRequestItem) decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < 7 {
		err = IllegalData
		return
	}

	s.referenceType = byteData[0]
	s.fileNumber = binary.BigEndian.Uint16(byteData[1:3])
	s.recordNumber = binary.BigEndian.Uint16(byteData[3:5])
	s.recordLength = binary.BigEndian.Uint16(byteData[5:7])
	return
}

type MBReadFileRecordReq struct {
	items []*ReadRequestItem
}

func (s *MBReadFileRecordReq) FuncCode() byte {
	return ReadFileRecord
}

func (s *MBReadFileRecordReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadFileRecord)
	buffVal = append(buffVal, s.calcDataSize())
	for _, val := range s.items {
		buffVal, err = val.encode(buffVal)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBReadFileRecordReq) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < 2 {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != ReadFileRecord {
		err = IllegalFuncCode
		return
	}
	dataSize := byteData[1]

	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &ReadRequestItem{}
		err = item.decode(byteData[offset : offset+7])
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBReadFileRecordReq) AppendItem(fileNumber uint16, recordNumber uint16, recordLength uint16) {
	s.items = append(s.items, &ReadRequestItem{
		referenceType: 6,
		fileNumber:    fileNumber,
		recordNumber:  recordNumber,
		recordLength:  recordLength,
	})
}

func (s *MBReadFileRecordReq) Items() []*ReadRequestItem {
	return s.items
}

func (s *MBReadFileRecordReq) calcDataSize() byte {
	dataSize := byte(0)
	for _, val := range s.items {
		dataSize += val.calcDataSize()
	}

	return dataSize
}

type ReadResponseItem struct {
	referenceType byte // 6
	recordData    []byte
}

func (s *ReadResponseItem) calcDataSize() byte {
	return byte(len(s.recordData)) + 1
}

func (s *ReadResponseItem) encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, s.calcDataSize())
	buffVal = append(buffVal, s.referenceType)
	buffVal = append(buffVal, s.recordData...)
	ret = buffVal
	return
}

func (s *ReadResponseItem) decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < 7 {
		err = IllegalData
		return
	}

	dataSize := byteData[0]
	referenceType := byteData[1]
	if referenceType != 6 {
		err = IllegalData
		return
	}

	s.recordData = byteData[2 : dataSize+1]
	return
}

type MBReadFileRecordRsp struct {
	items []*ReadResponseItem
}

func (s *MBReadFileRecordRsp) FuncCode() byte {
	return ReadFileRecord
}

func (s *MBReadFileRecordRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	buffVal = append(buffVal, ReadFileRecord)
	buffVal = append(buffVal, s.calcDataSize())
	for _, val := range s.items {
		buffVal, err = val.encode(buffVal)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBReadFileRecordRsp) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < 2 {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != ReadFileRecord {
		err = IllegalFuncCode
		return
	}
	dataSize := byteData[1]

	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &ReadResponseItem{}
		err = item.decode(byteData[offset : offset+7])
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBReadFileRecordRsp) AppendItem(dataVal []byte) {
	s.items = append(s.items, &ReadResponseItem{
		referenceType: 6,
		recordData:    dataVal,
	})
}

func (s *MBReadFileRecordRsp) Items() []*ReadResponseItem {
	return s.items
}

func (s *MBReadFileRecordRsp) calcDataSize() byte {
	dataSize := byte(0)
	for _, val := range s.items {
		dataSize += val.calcDataSize()
	}

	return dataSize
}

type WriteItem struct {
	referenceType byte   // 6
	fileNumber    uint16 // 2 bytes
	recordNumber  uint16 // 2 bytes
	recordData    []byte
}

func (s *WriteItem) FileNumber() uint16 {
	return s.fileNumber
}

func (s *WriteItem) RecordNumber() uint16 {
	return s.recordNumber
}

func (s *WriteItem) RecordData() []byte {
	return s.recordData
}

func (s *WriteItem) calcDataSize() byte {
	return 7 + byte(len(s.recordData))
}

func (s *WriteItem) encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	recordLength := uint16(len(s.recordData) / 2)
	buffVal = append(buffVal, s.referenceType)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.fileNumber)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.recordNumber)
	buffVal = binary.BigEndian.AppendUint16(buffVal, recordLength)
	buffVal = append(buffVal, s.recordData...)
	ret = buffVal
	return
}

func (s *WriteItem) decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < 7 {
		err = IllegalData
		return
	}

	s.referenceType = byteData[0]
	s.fileNumber = binary.BigEndian.Uint16(byteData[1:3])
	s.recordNumber = binary.BigEndian.Uint16(byteData[3:5])
	recordLength := binary.BigEndian.Uint16(byteData[5:7])
	s.recordData = byteData[7 : 7+recordLength*2]
	return
}

type MBWriteFileRecordReq struct {
	items []*WriteItem
}

func (s *MBWriteFileRecordReq) FuncCode() byte {
	return WriteFileRecord
}

func (s *MBWriteFileRecordReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	buffVal = append(buffVal, WriteFileRecord)
	buffVal = append(buffVal, s.calcDataSize())
	for _, val := range s.items {
		buffVal, err = val.encode(buffVal)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBWriteFileRecordReq) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < 2 {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != WriteFileRecord {
		err = IllegalFuncCode
		return
	}
	dataSize := byteData[1]

	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &WriteItem{}
		err = item.decode(byteData[offset : offset+7])
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBWriteFileRecordReq) AppendItem(fileNumber, recordNumber uint16, recordData []byte) {
	s.items = append(s.items, &WriteItem{
		referenceType: 6,
		fileNumber:    fileNumber,
		recordNumber:  recordNumber,
		recordData:    recordData,
	})
}

func (s *MBWriteFileRecordReq) Items() []*WriteItem {
	return s.items
}

func (s *MBWriteFileRecordReq) calcDataSize() byte {
	dataSize := byte(0)
	for _, val := range s.items {
		dataSize += val.calcDataSize()
	}

	return dataSize
}

type MBWriteFileRecordRsp struct {
	items []*WriteItem
}

func (s *MBWriteFileRecordRsp) FuncCode() byte {
	return WriteFileRecord
}

func (s *MBWriteFileRecordRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	buffVal = append(buffVal, WriteFileRecord)
	buffVal = append(buffVal, s.calcDataSize())
	for _, val := range s.items {
		buffVal, err = val.encode(buffVal)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBWriteFileRecordRsp) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	if len(byteData) < 2 {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != WriteFileRecord {
		err = IllegalFuncCode
		return
	}
	dataSize := byteData[1]

	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &WriteItem{}
		err = item.decode(byteData[offset : offset+7])
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBWriteFileRecordRsp) AppendItem(fileNumber, recordNumber uint16, recordData []byte) {
	s.items = append(s.items, &WriteItem{
		referenceType: 6,
		fileNumber:    fileNumber,
		recordNumber:  recordNumber,
		recordData:    recordData,
	})
}

func (s *MBWriteFileRecordRsp) Items() []*WriteItem {
	return s.items
}

func (s *MBWriteFileRecordRsp) calcDataSize() byte {
	dataSize := byte(0)
	for _, val := range s.items {
		dataSize += val.calcDataSize()
	}

	return dataSize
}

type MBMaskWriteRegisterReq struct {
	address uint16
	andMask []byte
	orMask  []byte
}

func (s *MBMaskWriteRegisterReq) FuncCode() byte {
	return MaskWriteRegister
}

func (s *MBMaskWriteRegisterReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, MaskWriteRegister)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.andMask...)
	buffVal = append(buffVal, s.orMask...)

	ret = buffVal
	return
}

func (s *MBMaskWriteRegisterReq) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()
	if len(byteData) != 7 {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != MaskWriteRegister {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.andMask = byteData[3:5]
	s.orMask = byteData[5:7]
	return
}

func (s *MBMaskWriteRegisterReq) Address() uint16 {
	return s.address
}

func (s *MBMaskWriteRegisterReq) AndMask() []byte {
	return s.andMask
}

func (s *MBMaskWriteRegisterReq) OrMask() []byte {
	return s.orMask
}

type MBMaskWriteRegisterRsp struct {
	address uint16
	andMask []byte
	orMask  []byte
}

func (s *MBMaskWriteRegisterRsp) FuncCode() byte {
	return MaskWriteRegister
}

func (s *MBMaskWriteRegisterRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, MaskWriteRegister)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.andMask...)
	buffVal = append(buffVal, s.orMask...)

	ret = buffVal
	return
}

func (s *MBMaskWriteRegisterRsp) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()
	if len(byteData) != 7 {
		err = IllegalData
		return
	}

	funcCode := byteData[0]
	if funcCode != MaskWriteRegister {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	s.andMask = byteData[3:5]
	s.orMask = byteData[5:7]
	return
}

func (s *MBMaskWriteRegisterRsp) Address() uint16 {
	return s.address
}

func (s *MBMaskWriteRegisterRsp) AndMask() []byte {
	return s.andMask
}

func (s *MBMaskWriteRegisterRsp) OrMask() []byte {
	return s.orMask
}

type MBReadWriteMultipleRegistersReq struct {
	readAddress  uint16
	readCount    uint16
	writeAddress uint16
	writeCount   uint16
	writeData    []byte
}

func (s *MBReadWriteMultipleRegistersReq) FuncCode() byte {
	return ReadWriteMultipleRegisters
}

func (s *MBReadWriteMultipleRegistersReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadWriteMultipleRegisters)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.readAddress)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.readCount)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.writeAddress)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.writeCount)
	buffVal = append(buffVal, byte(len(s.writeData)))
	buffVal = append(buffVal, s.writeData...)

	ret = buffVal
	return
}

func (s *MBReadWriteMultipleRegistersReq) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	funcCode := byteData[0]
	if funcCode != ReadWriteMultipleRegisters {
		err = IllegalFuncCode
		return
	}

	s.readAddress = binary.BigEndian.Uint16(byteData[1:3])
	s.readCount = binary.BigEndian.Uint16(byteData[3:5])
	s.writeAddress = binary.BigEndian.Uint16(byteData[5:7])
	s.writeCount = binary.BigEndian.Uint16(byteData[7:9])
	byteSize := byteData[9]
	s.writeData = byteData[10 : 10+byteSize]
	return
}

func (s *MBReadWriteMultipleRegistersReq) ReadAddress() uint16 {
	return s.readAddress
}

func (s *MBReadWriteMultipleRegistersReq) ReadCount() uint16 {
	return s.readCount
}

func (s *MBReadWriteMultipleRegistersReq) WriteAddress() uint16 {
	return s.writeAddress
}

func (s *MBReadWriteMultipleRegistersReq) WriteCount() uint16 {
	return s.writeCount
}

func (s *MBReadWriteMultipleRegistersReq) WriteData() []byte {
	return s.writeData
}

type MBReadWriteMultipleRegistersRsp struct {
	dataVal []byte
}

func (s *MBReadWriteMultipleRegistersRsp) FuncCode() byte {
	return ReadWriteMultipleRegisters
}

func (s *MBReadWriteMultipleRegistersRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadWriteMultipleRegisters)
	buffVal = append(buffVal, byte(len(s.dataVal)))
	buffVal = append(buffVal, s.dataVal...)

	ret = buffVal
	return
}

func (s *MBReadWriteMultipleRegistersRsp) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	funcCode := byteData[0]
	if funcCode != ReadWriteMultipleRegisters {
		err = IllegalFuncCode
		return
	}

	dataSize := byteData[1]
	s.dataVal = byteData[2 : 2+dataSize]
	return
}

type MBReadFIFOQueueReq struct {
	address uint16
}

func (s *MBReadFIFOQueueReq) FuncCode() byte {
	return ReadFIFOQueue
}

func (s *MBReadFIFOQueueReq) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadFIFOQueue)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)

	ret = buffVal
	return
}

func (s *MBReadFIFOQueueReq) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()

	funcCode := byteData[0]
	if funcCode != ReadFIFOQueue {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(byteData[1:3])
	return
}

func (s *MBReadFIFOQueueReq) Address() uint16 {
	return s.address
}

type MBReadFIFOQueueRsp struct {
	dataCount uint16
	dataVal   []byte
}

func (s *MBReadFIFOQueueRsp) FuncCode() byte {
	return ReadFIFOQueue
}

func (s *MBReadFIFOQueueRsp) Encode(buffVal []byte) (ret []byte, err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal = append(buffVal, ReadFIFOQueue)
	buffVal = binary.BigEndian.AppendUint16(buffVal, uint16(len(s.dataVal))+2)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.dataCount)
	buffVal = append(buffVal, s.dataVal...)

	ret = buffVal
	return
}

func (s *MBReadFIFOQueueRsp) Decode(byteData []byte) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()

	funcCode := byteData[0]
	if funcCode != ReadFIFOQueue {
		err = IllegalFuncCode
		return
	}

	byteSize := binary.BigEndian.Uint16(byteData[1:3])
	s.dataCount = binary.BigEndian.Uint16(byteData[3:5])
	s.dataVal = byteData[5 : 5+byteSize-2]
	return
}

func (s *MBReadFIFOQueueRsp) DataCount() uint16 {
	return s.dataCount
}

func (s *MBReadFIFOQueueRsp) Data() []byte {
	return s.dataVal
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

func (s *MBExceptionRsp) ExceptionCode() byte {
	return s.exceptionCode
}
