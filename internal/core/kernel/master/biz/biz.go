package biz

import (
	"bytes"
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

func (s *Master) ConnectSlave(slaveAddr string, devID byte) (ret string, err *cd.Result) {
	slaveID := fmt.Sprintf("mb%03d", devID)
	val := s.slaveInfoCache.Fetch(slaveID)
	if val != nil {
		errMsg := fmt.Sprintf("duplicate slave device %d", devID)
		log.Errorf("connectSlave failed, error:%s", errMsg)
		err = cd.NewError(cd.Duplicated, errMsg)
		return
	}

	masterPtr := &MBMaster{}
	errInfo := masterPtr.Start(slaveAddr, devID)
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

	vVal.(*MBMaster).Stop()
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

	mbMasterPtr := vVal.(*MBMaster)
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
	mbMasterPtr := vVal.(*MBMaster)
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
		log.Errorf("readHoldingRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(*MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("readHoldingRegisters failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	u16Count := uint16(0)
	switch valueType {
	case common.Int16Value, common.UInt16Value:
		u16Count = count
	case common.Int32Value, common.UInt32Value, common.Float32Value:
		u16Count = count * 2
	case common.Int64Value, common.UInt64Value, common.Float64Value:
		u16Count = count * 4
	default:
		errMsg := fmt.Sprintf("illegal valueType, type:%v", valueType)
		log.Errorf("readHoldingRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
	}
	if err != nil {
		return
	}

	readVal, readExCode, readErr := mbMasterPtr.ReadHoldingRegisters(address, u16Count)
	if readErr != nil {
		log.Errorf("readHoldingRegisters failed, error:%s", readErr.Error())
		err = cd.NewError(cd.UnExpected, readErr.Error())
		return
	}
	if readExCode != model.SuccessCode {
		exCode = readExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", readExCode)
		log.Errorf("readHoldingRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	if len(readVal) != int(u16Count*2) {
		errMsg := fmt.Sprintf("illegal read value count")
		log.Errorf("readHoldingRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	itemVal, itemErr := s.decodeReadVal(readVal, valueType, endianType, count)
	if itemErr != nil {
		log.Errorf("readHoldingRegisters failed, decode failed error:%s", itemErr.Error())
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
		log.Errorf("readInputRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}
	mbMasterPtr := vVal.(*MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("readInputRegisters failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	u16Count := uint16(0)
	switch valueType {
	case common.Int16Value, common.UInt16Value:
		u16Count = count
	case common.Int32Value, common.UInt32Value, common.Float32Value:
		u16Count = count * 2
	case common.Int64Value, common.UInt64Value, common.Float64Value:
		u16Count = count * 4
	default:
		errMsg := fmt.Sprintf("illegal valueType, type:%v", valueType)
		log.Errorf("readInputRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
	}
	if err != nil {
		return
	}

	readVal, readExCode, readErr := mbMasterPtr.ReadInputRegisters(address, u16Count)
	if readErr != nil {
		log.Errorf("readInputRegisters failed, error:%s", readErr.Error())
		err = cd.NewError(cd.UnExpected, readErr.Error())
		return
	}
	if readExCode != model.SuccessCode {
		exCode = readExCode
		errMsg := fmt.Sprintf("modbus exception code:%v", readExCode)
		log.Errorf("readInputRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	if len(readVal) != int(u16Count*2) {
		errMsg := fmt.Sprintf("illegal read value count")
		log.Errorf("readInputRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	itemVal, itemErr := s.decodeReadVal(readVal, valueType, endianType, count)
	if itemErr != nil {
		log.Errorf("readInputRegisters failed, decode failed error:%s", itemErr.Error())
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

	mbMasterPtr := vVal.(*MBMaster)
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

	mbMasterPtr := vVal.(*MBMaster)
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

func (s *Master) WriteSingleRegister(slaveID string, address uint16, value float64, valueType uint16) (exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("writeSingleRegister failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	mbMasterPtr := vVal.(*MBMaster)
	if !mbMasterPtr.IsConnect() {
		connErr := mbMasterPtr.ReConnect()
		if connErr != nil {
			log.Errorf("writeSingleRegister failed, reconnect slave error:%s", connErr.Error())
			err = cd.NewError(cd.UnExpected, connErr.Error())
			return
		}
	}

	cVal, cErr := common.ConvertFloat64To(value, valueType)
	if cErr != nil {
		log.Errorf("writeSingleRegister failed, reconnect slave error:%s", cErr.Error())
		err = cd.NewError(cd.UnExpected, cErr.Error())
		return
	}

	var byteVal []byte
	switch valueType {
	case common.Int16Value:
		byteVal, cErr = common.AppendInt16(byteVal, cVal.(int16))
	case common.UInt16Value:
		byteVal, cErr = common.AppendUint16(byteVal, cVal.(uint16))
	default:
		errMsg := fmt.Sprintf("illegal valueType, type:%v", valueType)
		log.Errorf("WriteSingleRegister failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
	}
	if err != nil {
		return
	}

	writeAddr, writeData, writeExCode, writeErr := mbMasterPtr.WriteSingleRegister(address, byteVal)
	if writeErr != nil {
		log.Errorf("writeCoils failed, error:%s", writeErr.Error())
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

	mbMasterPtr := vVal.(*MBMaster)
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
