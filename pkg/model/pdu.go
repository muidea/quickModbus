package model

import (
	"encoding/binary"
	"io"
)

type MBProtocol interface {
	FuncCode() byte
	Encode(writer io.Writer) (err byte)
	Decode(reader io.Reader) (err byte)
	CalcLen() uint16
	EncodePayload(writer io.Writer) (err byte)
	DecodePayload(reader io.Reader) (err byte)
	CalcPayloadLen() uint16
}

func NewReadExceptionStatusReq() *MBReadExceptionStatusReq {
	return &MBReadExceptionStatusReq{}
}

func EmptyReadExceptionStatusReq() *MBReadExceptionStatusReq {
	return &MBReadExceptionStatusReq{}
}

type MBReadExceptionStatusReq struct {
}

func (s *MBReadExceptionStatusReq) FuncCode() byte {
	return ReadExceptionStatus
}

func (s *MBReadExceptionStatusReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	wSize, wErr := writer.Write([]byte{s.FuncCode()})
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadExceptionStatusReq) Decode(reader io.Reader) (err byte) {
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

	return
}

func (s *MBReadExceptionStatusReq) CalcLen() uint16 {
	return 1
}

func (s *MBReadExceptionStatusReq) EncodePayload(_ io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	return
}

func (s *MBReadExceptionStatusReq) DecodePayload(_ io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	return
}

func (s *MBReadExceptionStatusReq) CalcPayloadLen() uint16 {
	return 0
}

func NewReadExceptionStatusRsp() *MBReadExceptionStatusRsp {
	return &MBReadExceptionStatusRsp{}
}

func EmptyReadExceptionStatusRsp(exceptionCode byte) *MBReadExceptionStatusRsp {
	return &MBReadExceptionStatusRsp{
		exceptionCode: exceptionCode,
	}
}

type MBReadExceptionStatusRsp struct {
	exceptionCode byte
	statusVal     byte
}

func (s *MBReadExceptionStatusRsp) FuncCode() byte {
	return ReadExceptionStatus
}

func (s *MBReadExceptionStatusRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBReadExceptionStatusRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	wSize, wErr := writer.Write([]byte{s.FuncCode(), s.statusVal})
	if wErr != nil || wSize != 2 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBReadExceptionStatusRsp) Decode(reader io.Reader) (err byte) {
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

	s.statusVal = dataVal[1]
	return
}

func (s *MBReadExceptionStatusRsp) CalcLen() uint16 {
	return 2
}

func (s *MBReadExceptionStatusRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	wSize, wErr := writer.Write([]byte{s.statusVal})
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBReadExceptionStatusRsp) DecodePayload(reader io.Reader) (err byte) {
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

	s.statusVal = dataVal[0]
	return
}

func (s *MBReadExceptionStatusRsp) CalcPayloadLen() uint16 {
	return 1
}

func (s *MBReadExceptionStatusRsp) Status() byte {
	return s.statusVal
}

func NewDiagnosticsReq(subFuncCode uint16, data []byte) *MBDiagnosticsReq {
	return &MBDiagnosticsReq{
		subFuncCode: subFuncCode,
		dataVal:     data,
	}
}

func EmptyDiagnosticsReq() *MBDiagnosticsReq {
	return &MBDiagnosticsReq{}
}

type MBDiagnosticsReq struct {
	subFuncCode uint16
	dataVal     []byte
}

func (s *MBDiagnosticsReq) FuncCode() byte {
	return Diagnostics
}

func (s *MBDiagnosticsReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.subFuncCode)
	buffVal = append(buffVal, s.dataVal[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 5 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBDiagnosticsReq) Decode(reader io.Reader) (err byte) {
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
	funCode := dataVal[0]
	if funCode != s.FuncCode() {
		err = IllegalData
		return
	}

	s.subFuncCode = binary.BigEndian.Uint16(dataVal[1:3])
	s.dataVal = dataVal[3:5]
	return
}

func (s *MBDiagnosticsReq) CalcLen() uint16 {
	return 5
}

func (s *MBDiagnosticsReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.subFuncCode)
	buffVal = append(buffVal, s.dataVal[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 4 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBDiagnosticsReq) DecodePayload(reader io.Reader) (err byte) {
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
	s.subFuncCode = binary.BigEndian.Uint16(dataVal[0:2])
	s.dataVal = dataVal[2:4]
	return
}

func (s *MBDiagnosticsReq) CalcPayloadLen() uint16 {
	return 4
}

func NewDiagnosticsRsp(subFuncCode uint16, data []byte) *MBDiagnosticsRsp {
	return &MBDiagnosticsRsp{
		subFuncCode: subFuncCode,
		dataVal:     data,
	}
}

func EmptyDiagnosticsRsp(exceptionCode byte) *MBDiagnosticsRsp {
	return &MBDiagnosticsRsp{
		exceptionCode: exceptionCode,
	}
}

type MBDiagnosticsRsp struct {
	exceptionCode byte
	subFuncCode   uint16
	dataVal       []byte
}

func (s *MBDiagnosticsRsp) FuncCode() byte {
	return Diagnostics
}

func (s *MBDiagnosticsRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBDiagnosticsRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.subFuncCode)
	buffVal = append(buffVal, s.dataVal[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 5 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBDiagnosticsRsp) Decode(reader io.Reader) (err byte) {
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
	funCode := dataVal[0]
	if funCode != s.FuncCode() {
		err = IllegalData
		return
	}

	s.subFuncCode = binary.BigEndian.Uint16(dataVal[1:3])
	s.dataVal = dataVal[3:5]
	return
}

func (s *MBDiagnosticsRsp) CalcLen() uint16 {
	return 5
}

func (s *MBDiagnosticsRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.subFuncCode)
	buffVal = append(buffVal, s.dataVal[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 4 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBDiagnosticsRsp) DecodePayload(reader io.Reader) (err byte) {
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
	s.subFuncCode = binary.BigEndian.Uint16(dataVal[0:2])
	s.dataVal = dataVal[2:4]
	return
}

func (s *MBDiagnosticsRsp) CalcPayloadLen() uint16 {
	return 4
}

func (s *MBDiagnosticsRsp) SubFunctionCode() uint16 {
	return s.subFuncCode
}

func (s *MBDiagnosticsRsp) Data() []byte {
	return s.dataVal
}

func NewGetCommEventCounterReq() *MBGetCommEventCounterReq {
	return &MBGetCommEventCounterReq{}
}

func EmptyGetCommEventCounterReq() *MBGetCommEventCounterReq {
	return &MBGetCommEventCounterReq{}
}

type MBGetCommEventCounterReq struct {
}

func (s *MBGetCommEventCounterReq) FuncCode() byte {
	return GetCommEventCounter
}

func (s *MBGetCommEventCounterReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	wSize, wErr := writer.Write([]byte{s.FuncCode()})
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBGetCommEventCounterReq) Decode(reader io.Reader) (err byte) {
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

	return
}

func (s *MBGetCommEventCounterReq) CalcLen() uint16 {
	return 1
}

func (s *MBGetCommEventCounterReq) EncodePayload(_ io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	return
}

func (s *MBGetCommEventCounterReq) DecodePayload(_ io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	return
}

func (s *MBGetCommEventCounterReq) CalcPayloadLen() uint16 {
	return 0
}

func NewGetCommEventCounterRsp() *MBGetCommEventCounterRsp {
	return &MBGetCommEventCounterRsp{}
}

func EmptyGetCommEventCounterRsp(exceptionCode byte) *MBGetCommEventCounterRsp {
	return &MBGetCommEventCounterRsp{
		exceptionCode: exceptionCode,
	}
}

type MBGetCommEventCounterRsp struct {
	exceptionCode byte
	commStatus    uint16
	eventCount    uint16
}

func (s *MBGetCommEventCounterRsp) FuncCode() byte {
	return GetCommEventCounter
}

func (s *MBGetCommEventCounterRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBGetCommEventCounterRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	buffSize := 5
	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.commStatus)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.eventCount)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != buffSize {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBGetCommEventCounterRsp) Decode(reader io.Reader) (err byte) {
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

	s.commStatus = binary.BigEndian.Uint16(dataVal[1:3])
	s.eventCount = binary.BigEndian.Uint16(dataVal[3:5])
	return
}

func (s *MBGetCommEventCounterRsp) CalcLen() uint16 {
	return 5
}

func (s *MBGetCommEventCounterRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	buffSize := 4
	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.commStatus)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.eventCount)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != buffSize {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBGetCommEventCounterRsp) DecodePayload(reader io.Reader) (err byte) {
	dataSize := 4
	dataVal := make([]byte, dataSize)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
		err = IllegalAddress
		return
	}

	s.commStatus = binary.BigEndian.Uint16(dataVal[0:2])
	s.eventCount = binary.BigEndian.Uint16(dataVal[2:4])
	return
}

func (s *MBGetCommEventCounterRsp) CalcPayloadLen() uint16 {
	return 4
}

func (s *MBGetCommEventCounterRsp) CommStatus() uint16 {
	return s.commStatus
}

func (s *MBGetCommEventCounterRsp) EventCount() uint16 {
	return s.eventCount
}

func NewGetCommEventLogReq() *MBGetCommEventLogReq {
	return &MBGetCommEventLogReq{}
}

func EmptyGetCommEventLogReq() *MBGetCommEventLogReq {
	return &MBGetCommEventLogReq{}
}

type MBGetCommEventLogReq struct {
}

func (s *MBGetCommEventLogReq) FuncCode() byte {
	return GetCommEventLog
}

func (s *MBGetCommEventLogReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	wSize, wErr := writer.Write([]byte{s.FuncCode()})
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBGetCommEventLogReq) Decode(reader io.Reader) (err byte) {
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

	return
}

func (s *MBGetCommEventLogReq) CalcLen() uint16 {
	return 1
}

func (s *MBGetCommEventLogReq) EncodePayload(_ io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	return
}

func (s *MBGetCommEventLogReq) DecodePayload(_ io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	return
}

func (s *MBGetCommEventLogReq) CalcPayloadLen() uint16 {
	return 0
}

func NewGetCommEventLogRsp() *MBGetCommEventLogRsp {
	return &MBGetCommEventLogRsp{}
}

func EmptyGetCommEventLogRsp(exceptionCode byte) *MBGetCommEventLogRsp {
	return &MBGetCommEventLogRsp{
		exceptionCode: exceptionCode,
	}
}

type MBGetCommEventLogRsp struct {
	exceptionCode byte
	commStatus    uint16
	eventCount    uint16
	messageCount  uint16
	commonEvents  []byte
}

func (s *MBGetCommEventLogRsp) FuncCode() byte {
	return GetCommEventLog
}

func (s *MBGetCommEventLogRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBGetCommEventLogRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	totalSize := byte(6 + len(s.commonEvents))
	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = append(buffVal, totalSize)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.commStatus)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.eventCount)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.messageCount)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 8 {
		err = IllegalAddress
		return
	}
	wSize, wErr = writer.Write(s.commonEvents)
	if wErr != nil || wSize != len(s.commonEvents) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBGetCommEventLogRsp) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	dataVal := make([]byte, 8)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 8 {
		err = IllegalAddress
		return
	}
	funcCode := dataVal[0]
	if funcCode != s.FuncCode() {
		err = IllegalData
		return
	}

	totalSize := int(dataVal[1])
	if totalSize < 6 {
		err = IllegalData
		return
	}

	s.commStatus = binary.BigEndian.Uint16(dataVal[2:4])
	s.eventCount = binary.BigEndian.Uint16(dataVal[4:6])
	s.messageCount = binary.BigEndian.Uint16(dataVal[6:8])
	dataVal = make([]byte, totalSize-6)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != (totalSize-6) {
		err = IllegalAddress
		return
	}

	s.commonEvents = dataVal
	return
}

func (s *MBGetCommEventLogRsp) CalcLen() uint16 {
	return 8 + uint16(len(s.commonEvents))
}

func (s *MBGetCommEventLogRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	totalSize := byte(6 + len(s.commonEvents))
	buffVal := make([]byte, 0)
	buffVal = append(buffVal, totalSize)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.commStatus)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.eventCount)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.messageCount)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 7 {
		err = IllegalAddress
		return
	}
	wSize, wErr = writer.Write(s.commonEvents)
	if wErr != nil || wSize != len(s.commonEvents) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBGetCommEventLogRsp) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	dataVal := make([]byte, 7)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 7 {
		err = IllegalAddress
		return
	}
	totalSize := int(dataVal[0])
	if totalSize < 6 {
		err = IllegalData
		return
	}

	s.commStatus = binary.BigEndian.Uint16(dataVal[1:3])
	s.eventCount = binary.BigEndian.Uint16(dataVal[3:5])
	s.messageCount = binary.BigEndian.Uint16(dataVal[5:7])
	dataVal = make([]byte, totalSize-6)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != (totalSize-6) {
		err = IllegalAddress
		return
	}

	s.commonEvents = dataVal
	return
}

func (s *MBGetCommEventLogRsp) CalcPayloadLen() uint16 {
	return 7 + uint16(len(s.commonEvents))
}

func (s *MBGetCommEventLogRsp) CommStatus() uint16 {
	return s.commStatus
}

func (s *MBGetCommEventLogRsp) EventCount() uint16 {
	return s.eventCount
}

func (s *MBGetCommEventLogRsp) MessageCount() uint16 {
	return s.messageCount
}

func (s *MBGetCommEventLogRsp) CommonEvents() []byte {
	return s.commonEvents
}

func NewReportSlaveIDReq() *MBReportSlaveIDReq {
	return &MBReportSlaveIDReq{}
}

func EmptyReportSlaveIDReq() *MBReportSlaveIDReq {
	return &MBReportSlaveIDReq{}
}

type MBReportSlaveIDReq struct {
}

func (s *MBReportSlaveIDReq) FuncCode() byte {
	return ReportSlaveID
}

func (s *MBReportSlaveIDReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	wSize, wErr := writer.Write([]byte{s.FuncCode()})
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReportSlaveIDReq) Decode(reader io.Reader) (err byte) {
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

	return
}

func (s *MBReportSlaveIDReq) CalcLen() uint16 {
	return 1
}

func (s *MBReportSlaveIDReq) EncodePayload(_ io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	return
}

func (s *MBReportSlaveIDReq) DecodePayload(_ io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	return
}

func (s *MBReportSlaveIDReq) CalcPayloadLen() uint16 {
	return 0
}

func NewReportSlaveIDRsp(info []byte) *MBReportSlaveIDRsp {
	return &MBReportSlaveIDRsp{
		slaveIDInfo: info,
	}
}

func EmptyReportSlaveIDRsp(exceptionCode byte) *MBReportSlaveIDRsp {
	return &MBReportSlaveIDRsp{
		exceptionCode: exceptionCode,
	}
}

type MBReportSlaveIDRsp struct {
	exceptionCode byte
	slaveIDInfo   []byte
}

func (s *MBReportSlaveIDRsp) FuncCode() byte {
	return ReportSlaveID
}

func (s *MBReportSlaveIDRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBReportSlaveIDRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	totalSize := byte(len(s.slaveIDInfo))
	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = append(buffVal, totalSize)
	buffVal = append(buffVal, s.slaveIDInfo...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBReportSlaveIDRsp) Decode(reader io.Reader) (err byte) {
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

	totalSize := int(dataVal[1])
	if totalSize < 1 {
		err = IllegalData
		return
	}

	dataVal = make([]byte, totalSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != totalSize {
		err = IllegalAddress
		return
	}

	s.slaveIDInfo = dataVal
	return
}

func (s *MBReportSlaveIDRsp) CalcLen() uint16 {
	return 2 + uint16(len(s.slaveIDInfo))
}

func (s *MBReportSlaveIDRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	totalSize := byte(len(s.slaveIDInfo))
	buffVal := make([]byte, 0)
	buffVal = append(buffVal, totalSize)
	buffVal = append(buffVal, s.slaveIDInfo...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReportSlaveIDRsp) DecodePayload(reader io.Reader) (err byte) {
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
	totalSize := int(dataVal[0])
	if totalSize < 1 {
		err = IllegalData
		return
	}

	dataVal = make([]byte, totalSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != totalSize {
		err = IllegalAddress
		return
	}

	s.slaveIDInfo = dataVal
	return
}

func (s *MBReportSlaveIDRsp) CalcPayloadLen() uint16 {
	return 1 + uint16(len(s.slaveIDInfo))
}

func (s *MBReportSlaveIDRsp) SlaveIDInfo() []byte {
	return s.slaveIDInfo
}

func NewReadFileRecordReq(items []*ReadRequestItem) *MBReadFileRecordReq {
	return &MBReadFileRecordReq{
		items: items,
	}
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

func NewReadRequestItem(fileNumber, recordNumber, recordLength uint16) *ReadRequestItem {
	return &ReadRequestItem{
		referenceType: byte(6),
		fileNumber:    fileNumber,
		recordNumber:  recordNumber,
		recordLength:  recordLength,
	}
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

func (s *ReadRequestItem) encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 7)
	buffVal = append(buffVal, s.referenceType)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.fileNumber)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.recordNumber)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.recordLength)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 7 {
		err = IllegalAddress
	}
	return
}

func (s *ReadRequestItem) decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	dataVal := make([]byte, 7)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 7 {
		err = IllegalAddress
		return
	}

	s.referenceType = dataVal[0]
	s.fileNumber = binary.BigEndian.Uint16(dataVal[1:3])
	s.recordNumber = binary.BigEndian.Uint16(dataVal[3:5])
	s.recordLength = binary.BigEndian.Uint16(dataVal[5:7])
	return
}

type MBReadFileRecordReq struct {
	items []*ReadRequestItem
}

func (s *MBReadFileRecordReq) FuncCode() byte {
	return ReadFileRecord
}

func (s *MBReadFileRecordReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	wSize, wErr := writer.Write([]byte{s.FuncCode(), s.calcDataSize()})
	if wErr != nil || wSize != 2 {
		err = IllegalAddress
		return
	}

	for _, val := range s.items {
		err = val.encode(writer)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBReadFileRecordReq) Decode(reader io.Reader) (err byte) {
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

	dataSize := dataVal[1]
	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &ReadRequestItem{}
		err = item.decode(reader)
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBReadFileRecordReq) CalcLen() uint16 {
	return uint16(s.calcDataSize()) + 2
}

func (s *MBReadFileRecordReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	wSize, wErr := writer.Write([]byte{s.calcDataSize()})
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}

	for _, val := range s.items {
		err = val.encode(writer)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBReadFileRecordReq) DecodePayload(reader io.Reader) (err byte) {
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

	dataSize := dataVal[0]
	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &ReadRequestItem{}
		err = item.decode(reader)
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBReadFileRecordReq) CalcPayloadLen() uint16 {
	return uint16(s.calcDataSize()) + 1
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

func NewReadFileRecordRsp() *MBReadFileRecordRsp {
	return &MBReadFileRecordRsp{}
}

func EmptyReadFileRecordRsp(exceptionCode byte) *MBReadFileRecordRsp {
	return &MBReadFileRecordRsp{
		exceptionCode: exceptionCode,
	}
}

type ReadResponseItem struct {
	referenceType byte // 6
	recordData    []byte
}

func (s *ReadResponseItem) calcDataSize() byte {
	return byte(len(s.recordData)) + 1
}

func (s *ReadResponseItem) encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffSize := s.calcDataSize() + 1
	buffVal := make([]byte, buffSize)
	buffVal = append(buffVal, s.calcDataSize())
	buffVal = append(buffVal, s.referenceType)
	buffVal = append(buffVal, s.recordData...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != int(buffSize) {
		err = IllegalAddress
	}

	return
}

func (s *ReadResponseItem) decode(reader io.Reader) (err byte) {
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

	dataSize := dataVal[0]
	referenceType := dataVal[1]
	if referenceType != 6 {
		err = IllegalData
		return
	}

	dataVal = make([]byte, dataSize-1)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != 2 {
		err = IllegalAddress
		return
	}
	s.recordData = dataVal
	return
}

func (s *ReadResponseItem) Data() []byte {
	return s.recordData
}

type MBReadFileRecordRsp struct {
	exceptionCode byte
	items         []*ReadResponseItem
}

func (s *MBReadFileRecordRsp) FuncCode() byte {
	return ReadFileRecord
}

func (s *MBReadFileRecordRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBReadFileRecordRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	wSize, wErr := writer.Write([]byte{s.FuncCode(), s.calcDataSize()})
	if wErr != nil || wSize != 2 {
		err = IllegalAddress
		return
	}

	for _, val := range s.items {
		err = val.encode(writer)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBReadFileRecordRsp) Decode(reader io.Reader) (err byte) {
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
	dataSize := dataVal[1]

	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &ReadResponseItem{}
		err = item.decode(reader)
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBReadFileRecordRsp) CalcLen() uint16 {
	return uint16(s.calcDataSize()) + 2
}

func (s *MBReadFileRecordRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	wSize, wErr := writer.Write([]byte{s.FuncCode(), s.calcDataSize()})
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}

	for _, val := range s.items {
		err = val.encode(writer)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBReadFileRecordRsp) DecodePayload(reader io.Reader) (err byte) {
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
	dataSize := dataVal[0]

	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &ReadResponseItem{}
		err = item.decode(reader)
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBReadFileRecordRsp) CalcPayloadLen() uint16 {
	return uint16(s.calcDataSize()) + 1
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

func NewWriteFileRecordReq(items []*WriteItem) *MBWriteFileRecordReq {
	return &MBWriteFileRecordReq{
		items: items,
	}
}

func EmptyWriteFileRecordReq() *MBWriteFileRecordReq {
	return &MBWriteFileRecordReq{}
}

type WriteItem struct {
	referenceType byte   // 6
	fileNumber    uint16 // 2 bytes
	recordNumber  uint16 // 2 bytes
	recordData    []byte
}

func NewWriteItem(fileNumber, recordNumber uint16, recordData []byte) *WriteItem {
	return &WriteItem{
		referenceType: 6,
		fileNumber:    fileNumber,
		recordNumber:  recordNumber,
		recordData:    recordData,
	}
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

func (s *WriteItem) encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 7)
	recordLength := uint16(len(s.recordData) / 2)
	buffVal = append(buffVal, s.referenceType)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.fileNumber)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.recordNumber)
	buffVal = binary.BigEndian.AppendUint16(buffVal, recordLength)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 7 {
		err = IllegalAddress
		return
	}
	wSize, wErr = writer.Write(s.recordData)
	if wErr != nil || wSize != len(s.recordData) {
		err = IllegalAddress
		return
	}
	return
}

func (s *WriteItem) decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	dataVal := make([]byte, 7)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 7 {
		err = IllegalAddress
		return
	}

	s.referenceType = dataVal[0]
	s.fileNumber = binary.BigEndian.Uint16(dataVal[1:3])
	s.recordNumber = binary.BigEndian.Uint16(dataVal[3:5])
	dataSize := int(binary.BigEndian.Uint16(dataVal[5:7]) * 2)
	dataVal = make([]byte, dataSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
		err = IllegalAddress
		return
	}

	s.recordData = dataVal
	return
}

type MBWriteFileRecordReq struct {
	items []*WriteItem
}

func (s *MBWriteFileRecordReq) FuncCode() byte {
	return WriteFileRecord
}

func (s *MBWriteFileRecordReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	buffVal := make([]byte, 2)
	buffVal[0] = s.FuncCode()
	buffVal[1] = s.calcDataSize()
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 2 {
		err = IllegalAddress
		return
	}

	for _, val := range s.items {
		err = val.encode(writer)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBWriteFileRecordReq) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
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

	dataSize := dataVal[1]
	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &WriteItem{}
		err = item.decode(reader)
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBWriteFileRecordReq) CalcLen() uint16 {
	return uint16(s.calcDataSize()) + 2
}

func (s *MBWriteFileRecordReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	buffVal := make([]byte, 1)
	buffVal[0] = s.calcDataSize()
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}

	for _, val := range s.items {
		err = val.encode(writer)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBWriteFileRecordReq) DecodePayload(reader io.Reader) (err byte) {
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
	dataSize := dataVal[0]

	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &WriteItem{}
		err = item.decode(reader)
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBWriteFileRecordReq) CalcPayloadLen() uint16 {
	return uint16(s.calcDataSize()) + 1
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

func NewWriteFileRecordRsp() *MBWriteFileRecordRsp {
	return &MBWriteFileRecordRsp{}
}

func EmptyWriteFileRecordRsp(exceptionCode byte) *MBWriteFileRecordRsp {
	return &MBWriteFileRecordRsp{
		exceptionCode: exceptionCode,
	}
}

type MBWriteFileRecordRsp struct {
	exceptionCode byte
	items         []*WriteItem
}

func (s *MBWriteFileRecordRsp) FuncCode() byte {
	return WriteFileRecord
}

func (s *MBWriteFileRecordRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBWriteFileRecordRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	buffVal := make([]byte, 2)
	buffVal[0] = s.FuncCode()
	buffVal[1] = s.calcDataSize()
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 2 {
		err = IllegalAddress
		return
	}

	for _, val := range s.items {
		err = val.encode(writer)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBWriteFileRecordRsp) Decode(reader io.Reader) (err byte) {
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

	dataSize := dataVal[1]
	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &WriteItem{}
		err = item.decode(reader)
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBWriteFileRecordRsp) CalcLen() uint16 {
	return 0
}

func (s *MBWriteFileRecordRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()
	buffVal := make([]byte, 1)
	buffVal[0] = s.calcDataSize()
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
		return
	}

	for _, val := range s.items {
		err = val.encode(writer)
		if err != SuccessCode {
			return
		}
	}

	return
}

func (s *MBWriteFileRecordRsp) DecodePayload(reader io.Reader) (err byte) {
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
	dataSize := dataVal[0]

	offset := byte(0)
	for {
		if offset >= dataSize {
			break
		}

		item := &WriteItem{}
		err = item.decode(reader)
		if err != SuccessCode {
			return
		}

		s.items = append(s.items, item)
		offset += item.calcDataSize()
	}

	return
}

func (s *MBWriteFileRecordRsp) CalcPayloadLen() uint16 {
	return 0
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

func NewMaskWriteRegisterReq(address uint16, andBytes []byte, orBytes []byte) *MBMaskWriteRegisterReq {
	return &MBMaskWriteRegisterReq{
		address: address,
		andMask: andBytes,
		orMask:  orBytes,
	}
}

func EmptyMaskWriteRegisterReq() *MBMaskWriteRegisterReq {
	return &MBMaskWriteRegisterReq{}
}

type MBMaskWriteRegisterReq struct {
	address uint16
	andMask []byte
	orMask  []byte
}

func (s *MBMaskWriteRegisterReq) FuncCode() byte {
	return MaskWriteRegister
}

func (s *MBMaskWriteRegisterReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.andMask[0:2]...)
	buffVal = append(buffVal, s.orMask[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 7 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBMaskWriteRegisterReq) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()
	dataVal := make([]byte, 7)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 5 {
		err = IllegalAddress
		return
	}

	funcCode := dataVal[0]
	if funcCode != s.FuncCode() {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[1:3])
	s.andMask = dataVal[3:5]
	s.orMask = dataVal[5:7]
	return
}

func (s *MBMaskWriteRegisterReq) CalcLen() uint16 {
	return 7
}

func (s *MBMaskWriteRegisterReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.andMask[0:2]...)
	buffVal = append(buffVal, s.orMask[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 6 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBMaskWriteRegisterReq) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()
	dataVal := make([]byte, 6)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 6 {
		err = IllegalAddress
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[0:2])
	s.andMask = dataVal[2:4]
	s.orMask = dataVal[4:6]
	return
}

func (s *MBMaskWriteRegisterReq) CalcPayloadLen() uint16 {
	return 6
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

func NewMaskWriteRegisterRsp() *MBMaskWriteRegisterRsp {
	return &MBMaskWriteRegisterRsp{}
}

func EmptyMaskWriteRegisterRsp(exceptionCode byte) *MBMaskWriteRegisterRsp {
	return &MBMaskWriteRegisterRsp{
		exceptionCode: exceptionCode,
	}
}

type MBMaskWriteRegisterRsp struct {
	exceptionCode byte
	address       uint16
	andMask       []byte
	orMask        []byte
}

func (s *MBMaskWriteRegisterRsp) FuncCode() byte {
	return MaskWriteRegister
}

func (s *MBMaskWriteRegisterRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBMaskWriteRegisterRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.andMask[0:2]...)
	buffVal = append(buffVal, s.orMask[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 7 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBMaskWriteRegisterRsp) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()
	dataVal := make([]byte, 7)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 5 {
		err = IllegalAddress
		return
	}

	funcCode := dataVal[0]
	if funcCode != s.FuncCode() {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[1:3])
	s.andMask = dataVal[3:5]
	s.orMask = dataVal[5:7]
	return
}

func (s *MBMaskWriteRegisterRsp) CalcLen() uint16 {
	return 7
}

func (s *MBMaskWriteRegisterRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	buffVal = append(buffVal, s.andMask[0:2]...)
	buffVal = append(buffVal, s.orMask[0:2]...)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 6 {
		err = IllegalAddress
		return
	}
	return
}

func (s *MBMaskWriteRegisterRsp) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()
	dataVal := make([]byte, 6)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 6 {
		err = IllegalAddress
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[0:2])
	s.andMask = dataVal[2:4]
	s.orMask = dataVal[4:6]
	return
}

func (s *MBMaskWriteRegisterRsp) CalcPayloadLen() uint16 {
	return 6
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

func NewReadWriteMultipleRegistersReq(readAddr, readCount uint16, writeAddr, writeCount uint16, writeData []byte) *MBReadWriteMultipleRegistersReq {
	return &MBReadWriteMultipleRegistersReq{
		readAddress:  readAddr,
		readCount:    readCount,
		writeAddress: writeAddr,
		writeCount:   writeCount,
		writeData:    writeData,
	}
}

func EmptyReadWriteMultipleRegistersReq() *MBReadWriteMultipleRegistersReq {
	return &MBReadWriteMultipleRegistersReq{}
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

func (s *MBReadWriteMultipleRegistersReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.readAddress)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.readCount)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.writeAddress)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.writeCount)
	buffVal = append(buffVal, byte(len(s.writeData)))
	buffVal = append(buffVal, s.writeData...)

	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadWriteMultipleRegistersReq) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	dataVal := make([]byte, 10)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 10 {
		err = IllegalAddress
		return
	}

	funcCode := dataVal[0]
	if funcCode != s.FuncCode() {
		err = IllegalFuncCode
		return
	}

	s.readAddress = binary.BigEndian.Uint16(dataVal[1:3])
	s.readCount = binary.BigEndian.Uint16(dataVal[3:5])
	s.writeAddress = binary.BigEndian.Uint16(dataVal[5:7])
	s.writeCount = binary.BigEndian.Uint16(dataVal[7:9])
	dataSize := int(dataVal[9])
	dataVal = make([]byte, dataSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
		err = IllegalAddress
		return
	}

	s.writeData = dataVal
	return
}

func (s *MBReadWriteMultipleRegistersReq) CalcLen() uint16 {
	return 10 + uint16(len(s.writeData))
}

func (s *MBReadWriteMultipleRegistersReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.readAddress)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.readCount)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.writeAddress)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.writeCount)
	buffVal = append(buffVal, byte(len(s.writeData)))
	buffVal = append(buffVal, s.writeData...)

	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadWriteMultipleRegistersReq) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	dataVal := make([]byte, 9)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 9 {
		err = IllegalAddress
		return
	}

	s.readAddress = binary.BigEndian.Uint16(dataVal[0:2])
	s.readCount = binary.BigEndian.Uint16(dataVal[2:4])
	s.writeAddress = binary.BigEndian.Uint16(dataVal[4:6])
	s.writeCount = binary.BigEndian.Uint16(dataVal[6:8])
	dataSize := int(dataVal[8])
	dataVal = make([]byte, dataSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
		err = IllegalAddress
		return
	}

	s.writeData = dataVal
	return
}

func (s *MBReadWriteMultipleRegistersReq) CalcPayloadLen() uint16 {
	return 9 + uint16(len(s.writeData))
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

func NewReadWriteMultipleRegistersRsp() *MBReadWriteMultipleRegistersRsp {
	return &MBReadWriteMultipleRegistersRsp{}
}

func EmptyReadWriteMultipleRegistersRsp(exceptionCode byte) *MBReadWriteMultipleRegistersRsp {
	return &MBReadWriteMultipleRegistersRsp{
		exceptionCode: exceptionCode,
	}
}

type MBReadWriteMultipleRegistersRsp struct {
	exceptionCode byte
	dataVal       []byte
}

func (s *MBReadWriteMultipleRegistersRsp) FuncCode() byte {
	return ReadWriteMultipleRegisters
}

func (s *MBReadWriteMultipleRegistersRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBReadWriteMultipleRegistersRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = append(buffVal, byte(len(s.dataVal)))
	buffVal = append(buffVal, s.dataVal...)

	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadWriteMultipleRegistersRsp) Decode(reader io.Reader) (err byte) {
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
		err = IllegalFuncCode
		return
	}

	dataSize := int(dataVal[1])
	dataVal = make([]byte, dataSize)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != dataSize {
		err = IllegalAddress
		return
	}

	s.dataVal = dataVal
	return
}

func (s *MBReadWriteMultipleRegistersRsp) CalcLen() uint16 {
	return uint16(len(s.dataVal)) + 2
}

func (s *MBReadWriteMultipleRegistersRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, byte(len(s.dataVal)))
	buffVal = append(buffVal, s.dataVal...)

	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadWriteMultipleRegistersRsp) DecodePayload(reader io.Reader) (err byte) {
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

	s.dataVal = dataVal
	return
}

func (s *MBReadWriteMultipleRegistersRsp) CalcPayloadLen() uint16 {
	return uint16(len(s.dataVal)) + 1
}

func (s *MBReadWriteMultipleRegistersRsp) Data() []byte {
	return s.dataVal
}

func NewReadFIFOQueueReq(address uint16) *MBReadFIFOQueueReq {
	return &MBReadFIFOQueueReq{
		address: address,
	}
}

func EmptyReadFIFOQueueReq() *MBReadFIFOQueueReq {
	return &MBReadFIFOQueueReq{}
}

type MBReadFIFOQueueReq struct {
	address uint16
}

func (s *MBReadFIFOQueueReq) FuncCode() byte {
	return ReadFIFOQueue
}

func (s *MBReadFIFOQueueReq) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadFIFOQueueReq) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()
	dataVal := make([]byte, 3)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 3 {
		err = IllegalAddress
		return
	}

	funcCode := dataVal[0]
	if funcCode != s.FuncCode() {
		err = IllegalFuncCode
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[1:3])

	return
}

func (s *MBReadFIFOQueueReq) CalcLen() uint16 {
	return 3
}

func (s *MBReadFIFOQueueReq) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.address)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadFIFOQueueReq) DecodePayload(reader io.Reader) (err byte) {
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
	if rErr != nil || rSize != 2 {
		err = IllegalAddress
		return
	}

	s.address = binary.BigEndian.Uint16(dataVal[0:2])

	return
}

func (s *MBReadFIFOQueueReq) CalcPayloadLen() uint16 {
	return 2
}

func (s *MBReadFIFOQueueReq) Address() uint16 {
	return s.address
}

func NewReadFIFOQueueRsp() *MBReadFIFOQueueRsp {
	return &MBReadFIFOQueueRsp{}
}

func EmptyReadFIFOQueueRsp(exceptionCode byte) *MBReadFIFOQueueRsp {
	return &MBReadFIFOQueueRsp{
		exceptionCode: exceptionCode,
	}
}

type MBReadFIFOQueueRsp struct {
	exceptionCode byte
	dataCount     uint16
	dataVal       []byte
}

func (s *MBReadFIFOQueueRsp) FuncCode() byte {
	return ReadFIFOQueue
}

func (s *MBReadFIFOQueueRsp) ExceptionCode() byte {
	return s.exceptionCode
}

func (s *MBReadFIFOQueueRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = append(buffVal, s.FuncCode())
	buffVal = binary.BigEndian.AppendUint16(buffVal, uint16(len(s.dataVal))+2)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.dataCount)
	buffVal = append(buffVal, s.dataVal...)

	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadFIFOQueueRsp) Decode(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
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
		err = IllegalFuncCode
		return
	}

	byteSize := int(binary.BigEndian.Uint16(dataVal[1:3]))
	s.dataCount = binary.BigEndian.Uint16(dataVal[3:5])
	dataVal = make([]byte, byteSize-2)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != (byteSize-2) {
		err = IllegalAddress
		return
	}

	s.dataVal = dataVal
	return
}

func (s *MBReadFIFOQueueRsp) CalcLen() uint16 {
	return 5 + uint16(len(s.dataVal))
}

func (s *MBReadFIFOQueueRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 0)
	buffVal = binary.BigEndian.AppendUint16(buffVal, uint16(len(s.dataVal))+2)
	buffVal = binary.BigEndian.AppendUint16(buffVal, s.dataCount)
	buffVal = append(buffVal, s.dataVal...)

	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != len(buffVal) {
		err = IllegalAddress
		return
	}

	return
}

func (s *MBReadFIFOQueueRsp) DecodePayload(reader io.Reader) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
		if err != SuccessCode {
			return
		}
	}()
	dataVal := make([]byte, 4)
	rSize, rErr := reader.Read(dataVal)
	if rErr != nil || rSize != 4 {
		err = IllegalAddress
		return
	}

	byteSize := int(binary.BigEndian.Uint16(dataVal[0:2]))
	s.dataCount = binary.BigEndian.Uint16(dataVal[2:4])
	dataVal = make([]byte, byteSize-2)
	rSize, rErr = reader.Read(dataVal)
	if rErr != nil || rSize != (byteSize-2) {
		err = IllegalAddress
		return
	}

	s.dataVal = dataVal
	return
}

func (s *MBReadFIFOQueueRsp) CalcPayloadLen() uint16 {
	return 4 + uint16(len(s.dataVal))
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

func (s *MBExceptionRsp) Encode(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 2)
	buffVal = append(buffVal, s.funcCode)
	buffVal = append(buffVal, s.exceptionCode)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 2 {
		err = IllegalAddress
	}

	return
}

func (s *MBExceptionRsp) Decode(reader io.Reader) (err byte) {
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
	s.funcCode = dataVal[0]
	s.exceptionCode = dataVal[1]
	return
}

func (s *MBExceptionRsp) EncodePayload(writer io.Writer) (err byte) {
	defer func() {
		if errInfo := recover(); errInfo != nil {
			err = IllegalData
		}
	}()

	buffVal := make([]byte, 1)
	buffVal = append(buffVal, s.exceptionCode)
	wSize, wErr := writer.Write(buffVal)
	if wErr != nil || wSize != 1 {
		err = IllegalAddress
	}

	return
}

func (s *MBExceptionRsp) DecodePayload(reader io.Reader) (err byte) {
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
	s.exceptionCode = dataVal[0]
	return
}

func (s *MBExceptionRsp) ExceptionCode() byte {
	return s.exceptionCode
}
