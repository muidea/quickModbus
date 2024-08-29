package biz

import (
	"bytes"
	"encoding/hex"
	"fmt"

	cd "github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/event"
	"github.com/muidea/magicCommon/foundation/cache"
	"github.com/muidea/magicCommon/foundation/log"
	"github.com/muidea/magicCommon/task"

	"github.com/muidea/quickModbus/internal/core/base/biz"
	"github.com/muidea/quickModbus/pkg/common"
	"github.com/muidea/quickModbus/pkg/model"
)

type Master struct {
	biz.Base

	slaveInfoCache cache.KVCache
}

func New(
	eventHub event.Hub,
	backgroundRoutine task.BackgroundRoutine,
) *Master {
	return &Master{
		Base:           biz.New(common.MasterModule, eventHub, backgroundRoutine),
		slaveInfoCache: cache.NewKVCache(nil),
	}
}

func (s *Master) ConnectSlave(slaveAddr string, devID, devType byte) (ret string, err *cd.Result) {
	slaveID := fmt.Sprintf("mb%03d", devID)
	val := s.slaveInfoCache.Fetch(slaveID)
	if val != nil {
		errMsg := fmt.Sprintf("duplicate slave device %d", devID)
		log.Errorf("connectSlave failed, error:%s", errMsg)
		err = cd.NewError(cd.Duplicated, errMsg)
		return
	}

	var masterPtr MBMaster
	if devType == common.ModbusTcp {
		masterPtr = NewTCPMaster(devID)
	} else if devType == common.ModbusRTUOverTcp {
		masterPtr = NewRTUMaster(devID)
	} else {
		errMsg := fmt.Sprintf("illegal slave device type, id:%v, type:%v", devID, devType)
		log.Errorf("connectSlave failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	errInfo := masterPtr.Start(slaveAddr)
	if errInfo != nil {
		log.Errorf("connectSlave failed, error:%s", errInfo.Error())
		err = cd.NewError(cd.UnExpected, errInfo.Error())
		return
	}

	s.slaveInfoCache.Put(slaveID, masterPtr, cache.ForeverAgeValue)
	ret = slaveID
	return
}

func (s *Master) DisConnectSlave(slaveID string) (err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("disconnectSlave failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	vVal.(MBMaster).Stop()
	s.slaveInfoCache.Remove(slaveID)
	return
}

func (s *Master) ReadCoils(slaveID string, address, count uint16) (ret []bool, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("readCoils failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("readCoils failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	readVal, readExCode, readErr := mbMasterPtr.ReadCoils(address, count)
	if readErr != nil {
		log.Errorf("readCoils failed, error:%s", readErr.Error())
		err = cd.NewError(cd.UnExpected, readErr.Error())
		return
	}
	if readExCode != model.SuccessCode {
		exCode = readExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", readExCode)
		log.Errorf("readCoils failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	boolVal, boolErr := common.BytesToBoolArray(readVal)
	if boolErr != nil {
		log.Errorf("readCoils failed, error:%s", boolErr.Error())
		err = cd.NewError(cd.UnExpected, boolErr.Error())
		return
	}

	ret = boolVal[:count]
	return
}

func (s *Master) ReadDiscreteInputs(slaveID string, address, count uint16) (ret []bool, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("readDiscreteInputs failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("readDiscreteInputs failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	readVal, readExCode, readErr := mbMasterPtr.ReadDiscreteInputs(address, count)
	if readErr != nil {
		log.Errorf("readDiscreteInputs failed, error:%s", readErr.Error())
		err = cd.NewError(cd.UnExpected, readErr.Error())
		return
	}
	if readExCode != model.SuccessCode {
		exCode = readExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", readExCode)
		log.Errorf("readDiscreteInputs failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	boolVal, boolErr := common.BytesToBoolArray(readVal)
	if boolErr != nil {
		log.Errorf("readDiscreteInputs failed, error:%s", boolErr.Error())
		err = cd.NewError(cd.UnExpected, boolErr.Error())
		return
	}

	ret = boolVal[:count]
	return
}

func (s *Master) decodeReadVal(readVal []byte, valueType, endianType, count uint16) (interface{}, error) {
	var itemVal interface{}
	var itemErr error
	switch valueType {
	case common.Int16Value:
		iVal, iErr := common.BytesToInt16Array(readVal)
		if iErr == nil {
			itemVal = iVal[:count]
		}
		itemErr = iErr
	case common.UInt16Value:
		uVal, uErr := common.BytesToUint16Array(readVal)
		if uErr == nil {
			itemVal = uVal[:count]
		}
		itemErr = uErr
	case common.Int32Value:
		iVal, iErr := common.BytesToInt32Array(readVal, endianType)
		if iErr == nil {
			itemVal = iVal[:count]
		}
		itemErr = iErr
	case common.UInt32Value:
		uVal, uErr := common.BytesToUint32Array(readVal, endianType)
		if uErr == nil {
			itemVal = uVal[:count]
		}
		itemErr = uErr
	case common.Float32Value:
		fVal, fErr := common.BytesToFloat32Array(readVal, endianType)
		if fErr == nil {
			itemVal = fVal[:count]
		}
		itemErr = fErr
	case common.Int64Value:
		iVal, iErr := common.BytesToInt64Array(readVal, endianType)
		if iErr == nil {
			itemVal = iVal[:count]
		}
		itemErr = iErr
	case common.UInt64Value:
		uVal, uErr := common.BytesToUint64Array(readVal, endianType)
		if uErr == nil {
			itemVal = uVal[:count]
		}
		itemErr = uErr
	case common.Float64Value:
		fVal, fErr := common.BytesToFloat64Array(readVal, endianType)
		if fErr == nil {
			itemVal = fVal[:count]
		}
		itemErr = fErr
	default:
	}

	return itemVal, itemErr
}

func (s *Master) ReadHoldingRegisters(slaveID string, address, count, valueType, endianType uint16) (ret interface{}, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("ReadHoldingRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("ReadHoldingRegisters failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	dataCount, dataErr := s.prepareReadData(count, valueType)
	if dataErr != nil {
		err = cd.NewError(cd.UnExpected, dataErr.Error())
		return
	}

	readVal, readExCode, readErr := mbMasterPtr.ReadHoldingRegisters(address, dataCount)
	if readErr != nil {
		log.Errorf("ReadHoldingRegisters failed, error:%s", readErr.Error())
		err = cd.NewError(cd.UnExpected, readErr.Error())
		return
	}
	if readExCode != model.SuccessCode {
		exCode = readExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", readExCode)
		log.Errorf("ReadHoldingRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	if len(readVal) != int(dataCount*2) {
		errMsg := fmt.Sprintf("illegal read value count")
		log.Errorf("ReadHoldingRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	itemVal, itemErr := s.decodeReadVal(readVal, valueType, endianType, count)
	if itemErr != nil {
		log.Errorf("ReadHoldingRegisters failed, decode failed error:%s", itemErr.Error())
		err = cd.NewError(cd.UnExpected, itemErr.Error())
		return
	}

	ret = itemVal
	return
}

func (s *Master) ReadInputRegisters(slaveID string, address, count, valueType, endianType uint16) (ret interface{}, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("ReadInputRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("ReadInputRegisters failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	dataCount, dataErr := s.prepareReadData(count, valueType)
	if dataErr != nil {
		log.Errorf("ReadInputRegisters failed, prepareReadData error:%s", dataErr.Error())
		err = cd.NewError(cd.UnExpected, dataErr.Error())
		return
	}

	readVal, readExCode, readErr := mbMasterPtr.ReadInputRegisters(address, dataCount)
	if readErr != nil {
		log.Errorf("ReadInputRegisters failed, error:%s", readErr.Error())
		err = cd.NewError(cd.UnExpected, readErr.Error())
		return
	}
	if readExCode != model.SuccessCode {
		exCode = readExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", readExCode)
		log.Errorf("ReadInputRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	if len(readVal) != int(dataCount*2) {
		errMsg := fmt.Sprintf("illegal read value count")
		log.Errorf("ReadInputRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	itemVal, itemErr := s.decodeReadVal(readVal, valueType, endianType, count)
	if itemErr != nil {
		log.Errorf("ReadInputRegisters failed, decode failed error:%s", itemErr.Error())
		err = cd.NewError(cd.UnExpected, itemErr.Error())
		return
	}

	ret = itemVal
	return
}

func (s *Master) WriteSingleCoil(slaveID string, address uint16, value bool) (exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("readCoils failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("readCoils failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	var byteVal []byte
	if value {
		byteVal = model.CoilON
	} else {
		byteVal = model.CoilOFF
	}

	writeAddr, writeData, writeExCode, writeErr := mbMasterPtr.WriteSingleCoil(address, byteVal)
	if writeErr != nil {
		log.Errorf("writeCoils failed, error:%s", writeErr.Error())
		err = cd.NewError(cd.UnExpected, writeErr.Error())
		return
	}
	if writeExCode != model.SuccessCode {
		exCode = writeExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", writeExCode)
		log.Errorf("writeSingleCoil failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	if writeAddr != address || bytes.Compare(byteVal, writeData) != 0 {
		errMsg := fmt.Sprintf("mismatch write single coil value")
		log.Errorf("writeSingleCoil failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	return
}

func (s *Master) WriteMultipleCoils(slaveID string, address uint16, value []bool) (exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("writeMultipleCoils failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("writeMultipleCoils failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	valCount := uint16(len(value))
	var byteVal []byte
	var byteErr error
	byteVal, byteErr = common.AppendBoolArray(byteVal, value)
	if byteErr != nil {
		log.Errorf("writeMultipleCoils failed, common.AppendBoolArray error:%s", byteErr.Error())
		err = cd.NewError(cd.UnExpected, byteErr.Error())
		return
	}

	writeAddr, writeCount, writeExCode, writeErr := mbMasterPtr.WriteMultipleCoils(address, valCount, byteVal)
	if writeErr != nil {
		log.Errorf("writeMultipleCoils failed, error:%s", writeErr.Error())
		err = cd.NewError(cd.UnExpected, writeErr.Error())
		return
	}
	if writeExCode != model.SuccessCode {
		exCode = writeExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", writeExCode)
		log.Errorf("writeMultipleCoils failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	if writeAddr != address || valCount != writeCount {
		errMsg := fmt.Sprintf("mismatch write multiple coil value")
		log.Errorf("writeMultipleCoils failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	return
}

func (s *Master) WriteSingleRegister(slaveID string, address uint16, value uint16) (exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("WriteSingleRegister failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("WriteSingleRegister failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	var byteVal []byte
	var byteErr error
	byteVal, byteErr = common.AppendUint16(byteVal, value)
	if byteErr != nil {
		log.Errorf("WriteSingleRegister failed, AppendUint16 error:%s", byteErr.Error())
		err = cd.NewError(cd.UnExpected, byteErr.Error())
		return
	}

	writeAddr, writeData, writeExCode, writeErr := mbMasterPtr.WriteSingleRegister(address, byteVal)
	if writeErr != nil {
		log.Errorf("WriteSingleRegister failed, error:%s", writeErr.Error())
		err = cd.NewError(cd.UnExpected, writeErr.Error())
		return
	}
	if writeExCode != model.SuccessCode {
		exCode = writeExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", writeExCode)
		log.Errorf("WriteSingleRegister failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	if writeAddr != address || bytes.Compare(byteVal, writeData) != 0 {
		errMsg := fmt.Sprintf("mismatch write single register value")
		log.Errorf("WriteSingleRegister failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	return
}

func (s *Master) WriteMultipleRegisters(slaveID string, address uint16, values []float64, valueTyp, endianType uint16) (exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("writeMultipleRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("writeMultipleRegisters failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	var byteVal []byte
	var byteErr error
	valCount := uint16(0)
	for _, val := range values {
		cVal, cErr := common.ConvertFloat64To(val, valueTyp)
		if cErr != nil {
			log.Errorf("writeMultipleRegisters failed, common.ConvertFloat64To error:%s", cErr.Error())
			err = cd.NewError(cd.UnExpected, cErr.Error())
			return
		}

		switch valueTyp {
		case common.Int16Value:
			byteVal, byteErr = common.AppendInt16(byteVal, cVal.(int16))
			valCount++
		case common.UInt16Value:
			byteVal, byteErr = common.AppendUint16(byteVal, cVal.(uint16))
			valCount++
		case common.Int32Value:
			byteVal, byteErr = common.AppendInt32(byteVal, cVal.(int32), endianType)
			valCount += 2
		case common.UInt32Value:
			byteVal, byteErr = common.AppendUint32(byteVal, cVal.(uint32), endianType)
			valCount += 2
		case common.Float32Value:
			byteVal, byteErr = common.AppendFloat32(byteVal, cVal.(float32), endianType)
			valCount += 2
		case common.Int64Value:
			byteVal, byteErr = common.AppendInt64(byteVal, cVal.(int64), endianType)
			valCount += 4
		case common.UInt64Value:
			byteVal, byteErr = common.AppendUint64(byteVal, cVal.(uint64), endianType)
			valCount += 4
		case common.Float64Value:
			byteVal, byteErr = common.AppendFloat64(byteVal, cVal.(float64), endianType)
			valCount += 4
		}
		if byteErr != nil {
			log.Errorf("writeMultipleRegisters failed, AppendValueToArray error:%s", byteErr.Error())
			err = cd.NewError(cd.UnExpected, cErr.Error())
			return
		}
	}

	writeAddr, writeCount, writeExCode, writeErr := mbMasterPtr.WriteMultipleRegisters(address, valCount, byteVal)
	if writeErr != nil {
		log.Errorf("writeMultipleRegisters failed, error:%s", writeErr.Error())
		err = cd.NewError(cd.UnExpected, writeErr.Error())
		return
	}
	if writeExCode != model.SuccessCode {
		exCode = writeExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", writeExCode)
		log.Errorf("writeMultipleRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	if writeAddr != address || valCount != writeCount {
		errMsg := fmt.Sprintf("mismatch write multiple register values")
		log.Errorf("writeMultipleRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	return
}

func (s *Master) MaskWriteRegister(slaveID string, address uint16, andMask uint16, orMask uint16) (exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("MaskWriteRegister failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("readCoils failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	var andByteVal []byte
	var andErr error
	andByteVal, andErr = common.AppendUint16(andByteVal, andMask)
	if andErr != nil {
		log.Errorf("MaskWriteRegister failed, andMask AppendUint16 error:%s", andErr)
		err = cd.NewError(cd.UnExpected, andErr.Error())
		return
	}

	var orByteVal []byte
	var orErr error
	orByteVal, orErr = common.AppendUint16(orByteVal, orMask)
	if orErr != nil {
		log.Errorf("MaskWriteRegister failed, orMask AppendBoolArray error:%s", orErr)
		err = cd.NewError(cd.UnExpected, orErr.Error())
		return
	}

	maskAddr, maskAnd, maskOr, maskExCode, maskErr := mbMasterPtr.MaskWriteRegister(address, andByteVal, orByteVal)
	if maskErr != nil {
		log.Errorf("MaskWriteRegister failed, error:%s", maskErr.Error())
		err = cd.NewError(cd.UnExpected, maskErr.Error())
		return
	}
	if maskExCode != model.SuccessCode {
		exCode = maskExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", maskExCode)
		log.Errorf("MaskWriteRegister failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	if address != maskAddr || bytes.Compare(andByteVal, maskAnd) != 0 || bytes.Compare(orByteVal, maskOr) != 0 {
		errMsg := fmt.Sprintf("mismatch mask write register")
		log.Errorf("MaskWriteRegister failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	return
}

func (s *Master) ReadWriteMultipleRegisters(slaveID string, readAddr, readCount, readValueType uint16, writeAddr uint16, writeValues []float64, writeValueType, endianType uint16) (ret interface{}, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("ReadWriteMultipleRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("ReadWriteMultipleRegisters failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	readValCount, readValErr := s.prepareReadData(readCount, readValueType)
	if readValErr != nil {
		log.Errorf("ReadWriteMultipleRegisters failed, prepareReadData error:%s", readValErr.Error())
		err = cd.NewError(cd.UnExpected, readValErr.Error())
		return
	}
	writeByteVal, writeCount, writeErr := s.prepareWriteData(writeValues, writeValueType, endianType)
	if writeErr != nil {
		log.Errorf("ReadWriteMultipleRegisters failed, prepareWriteData error:%s", writeErr.Error())
		err = cd.NewError(cd.UnExpected, writeErr.Error())
		return
	}

	retVal, retExCode, retErr := mbMasterPtr.ReadWriteMultipleRegisters(readAddr, readValCount, writeAddr, writeCount, writeByteVal)
	if retErr != nil {
		log.Errorf("ReadWriteMultipleRegisters failed, error:%s", retErr.Error())
		err = cd.NewError(cd.UnExpected, retErr.Error())
		return
	}
	if retExCode != model.SuccessCode {
		exCode = retExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", retExCode)
		log.Errorf("ReadWriteMultipleRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	if len(retVal) != int(readValCount*2) {
		errMsg := fmt.Sprintf("illegal read value count")
		log.Errorf("ReadWriteMultipleRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	itemVal, itemErr := s.decodeReadVal(retVal, readValueType, endianType, readCount)
	if itemErr != nil {
		log.Errorf("ReadWriteMultipleRegisters failed, decode failed error:%s", itemErr.Error())
		err = cd.NewError(cd.UnExpected, itemErr.Error())
		return
	}

	ret = itemVal
	return
}

func (s *Master) prepareReadData(readNum, valueType uint16) (uint16, error) {
	var readCount = uint16(0)
	var readErr error
	switch valueType {
	case common.Int16Value, common.UInt16Value:
		readCount = readNum
	case common.Int32Value, common.UInt32Value, common.Float32Value:
		readCount = readNum * 2
	case common.Int64Value, common.UInt64Value, common.Float64Value:
		readCount = readNum * 4
	default:
		readErr = fmt.Errorf("illegal valueType, type:%v", valueType)
	}
	if readErr != nil {
		log.Errorf("prepareReadData failed, error:%s", readErr.Error())
		return 0, readErr
	}

	return readCount, nil
}

func (s *Master) prepareWriteData(values []float64, valueType uint16, endianType uint16) ([]byte, uint16, error) {
	var writeByteVal []byte
	var writeCount = uint16(0)
	var writeByteErr error
	for _, val := range values {
		cVal, cErr := common.ConvertFloat64To(val, valueType)
		if cErr != nil {
			log.Errorf("common.ConvertFloat64To error:%s", cErr.Error())
			return nil, 0, cErr
		}

		switch valueType {
		case common.Int16Value:
			writeByteVal, writeByteErr = common.AppendInt16(writeByteVal, cVal.(int16))
			writeCount++
		case common.UInt16Value:
			writeByteVal, writeByteErr = common.AppendUint16(writeByteVal, cVal.(uint16))
			writeCount++
		case common.Int32Value:
			writeByteVal, writeByteErr = common.AppendInt32(writeByteVal, cVal.(int32), endianType)
			writeCount += 2
		case common.UInt32Value:
			writeByteVal, writeByteErr = common.AppendUint32(writeByteVal, cVal.(uint32), endianType)
			writeCount += 2
		case common.Float32Value:
			writeByteVal, writeByteErr = common.AppendFloat32(writeByteVal, cVal.(float32), endianType)
			writeCount += 2
		case common.Int64Value:
			writeByteVal, writeByteErr = common.AppendInt64(writeByteVal, cVal.(int64), endianType)
			writeCount += 4
		case common.UInt64Value:
			writeByteVal, writeByteErr = common.AppendUint64(writeByteVal, cVal.(uint64), endianType)
			writeCount += 4
		case common.Float64Value:
			writeByteVal, writeByteErr = common.AppendFloat64(writeByteVal, cVal.(float64), endianType)
			writeCount += 4
		default:
			writeByteErr = fmt.Errorf("illegal valueType, type:%v", valueType)
		}

		if writeByteErr != nil {
			log.Errorf("prepareWriteData failed, AppendValueToArray error:%s", writeByteErr.Error())
			return nil, 0, writeByteErr
		}
	}

	return writeByteVal, writeCount, nil
}

func (s *Master) ReadExceptionStatus(slaveID string) (status, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("ReadExceptionStatus failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("ReadExceptionStatus failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	retVal, retExCode, retErr := mbMasterPtr.ReadExceptionStatus()
	if retErr != nil {
		log.Errorf("ReadExceptionStatus failed, error:%s", retErr.Error())
		err = cd.NewError(cd.UnExpected, retErr.Error())
		return
	}
	if retExCode != model.SuccessCode {
		exCode = retExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", retExCode)
		log.Errorf("ReadExceptionStatus failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	status = retVal
	return
}

func (s *Master) Diagnostics(slaveID string, subFuncCode uint16, dataVal string) (ret string, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("Diagnostics failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("Diagnostics failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	byteVal, byteErr := hex.DecodeString(dataVal)
	if byteErr != nil {
		log.Errorf("Diagnostics failed, hex.DecodeString error:%s", byteErr.Error())
		err = cd.NewError(cd.UnExpected, byteErr.Error())
		return
	}

	retSubFuncCode, retDataVal, retExCode, retErr := mbMasterPtr.Diagnostics(subFuncCode, byteVal)
	if retErr != nil {
		log.Errorf("ReadExceptionStatus failed, error:%s", retErr.Error())
		err = cd.NewError(cd.UnExpected, retErr.Error())
		return
	}
	if retExCode != model.SuccessCode {
		exCode = retExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", retExCode)
		log.Errorf("ReadExceptionStatus failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	if retSubFuncCode != subFuncCode {
		errMsg := fmt.Sprintf("mismatch subfunction code, request:%v response:%v", subFuncCode, retSubFuncCode)
		log.Errorf("ReadExceptionStatus failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	ret = hex.EncodeToString(retDataVal)
	return
}

func (s *Master) GetCommEventCounter(slaveID string) (status, eventCount uint16, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("GetCommEventCounter failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("GetCommEventCounter failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	retStatus, retEventCount, retExCode, retErr := mbMasterPtr.GetCommEventCounter()
	if retErr != nil {
		log.Errorf("GetCommEventCounter failed, error:%s", retErr.Error())
		err = cd.NewError(cd.UnExpected, retErr.Error())
		return
	}
	if retExCode != model.SuccessCode {
		exCode = retExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", retExCode)
		log.Errorf("GetCommEventCounter failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	status = retStatus
	eventCount = retEventCount
	return
}

func (s *Master) GetCommEventLog(slaveID string) (status, eventCount, messageCount uint16, events string, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("GetCommEventLog failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("GetCommEventLog failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	retStatus, retEventCount, retMessageCount, retEvents, retExCode, retErr := mbMasterPtr.GetCommEventLog()
	if retErr != nil {
		log.Errorf("GetCommEventLog failed, error:%s", retErr.Error())
		err = cd.NewError(cd.UnExpected, retErr.Error())
		return
	}
	if retExCode != model.SuccessCode {
		exCode = retExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", retExCode)
		log.Errorf("GetCommEventLog failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	status = retStatus
	eventCount = retEventCount
	messageCount = retMessageCount
	events = hex.EncodeToString(retEvents)
	return
}

func (s *Master) ReportSlaveID(slaveID string) (ret string, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("ReportSlaveID failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("ReportSlaveID failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	retSlaveInfo, retExCode, retErr := mbMasterPtr.ReportSlaveID()
	if retErr != nil {
		log.Errorf("ReportSlaveID failed, error:%s", retErr.Error())
		err = cd.NewError(cd.UnExpected, retErr.Error())
		return
	}
	if retExCode != model.SuccessCode {
		exCode = retExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", retExCode)
		log.Errorf("ReportSlaveID failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	ret = hex.EncodeToString(retSlaveInfo)
	return
}

func (s *Master) ReadFileRecord(slaveID string, items []*common.ReadItem) (ret []string, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("ReadFileRecord failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("ReadFileRecord failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	retFileContent, retExCode, retErr := mbMasterPtr.ReadFileRecord(items)
	if retErr != nil {
		log.Errorf("ReadFileRecord failed, error:%s", retErr.Error())
		err = cd.NewError(cd.UnExpected, retErr.Error())
		return
	}
	if retExCode != model.SuccessCode {
		exCode = retExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", retExCode)
		log.Errorf("ReadFileRecord failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	for _, val := range retFileContent {
		ret = append(ret, hex.EncodeToString(val))
	}
	return
}

func (s *Master) WriteFileRecord(slaveID string, items []*common.WriteItem) (exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("WriteFileRecord failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("WriteFileRecord failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	retExCode, retErr := mbMasterPtr.WriteFileRecord(items)
	if retErr != nil {
		log.Errorf("WriteFileRecord failed, error:%s", retErr.Error())
		err = cd.NewError(cd.UnExpected, retErr.Error())
		return
	}
	if retExCode != model.SuccessCode {
		exCode = retExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", retExCode)
		log.Errorf("WriteFileRecord failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	return
}

func (s *Master) ReadFIFOQueue(slaveID string, address uint16) (retData []string, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("ReadFIFOQueue failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("ReadFIFOQueue failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	readDataCount, readDataVal, readExCode, readErr := mbMasterPtr.ReadFIFOQueue(address)
	if readErr != nil {
		log.Errorf("ReadFIFOQueue failed, error:%s", readErr.Error())
		err = cd.NewError(cd.UnExpected, readErr.Error())
		return
	}
	if readExCode != model.SuccessCode {
		exCode = readExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", readExCode)
		log.Errorf("ReadFIFOQueue failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	for idx := 0; idx < int(readDataCount); idx += 2 {
		subByte := hex.EncodeToString(readDataVal[idx : idx+2])
		retData = append(retData, subByte)
	}
	return
}
