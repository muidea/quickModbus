package biz

import (
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

func (s *Master) ReadCoils(slaveID string, address, count, endianType uint16) (ret []bool, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("readCoils failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	readVal, readExCode, readErr := vVal.(*MBMaster).ReadCoils(address, count)
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

	boolVal, boolErr := common.BytesToBoolArray(readVal, endianType)
	if boolErr != nil {
		log.Errorf("readCoils failed, error:%s", boolErr.Error())
		err = cd.NewError(cd.UnExpected, boolErr.Error())
		return
	}

	ret = boolVal
	return
}

func (s *Master) ReadDiscreteInputs(slaveID string, address, count, endianType uint16) (ret []bool, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("readDiscreteInputs failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	readVal, readExCode, readErr := vVal.(*MBMaster).ReadDiscreteInputs(address, count)
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

	boolVal, boolErr := common.BytesToBoolArray(readVal, endianType)
	if boolErr != nil {
		log.Errorf("readDiscreteInputs failed, error:%s", boolErr.Error())
		err = cd.NewError(cd.UnExpected, boolErr.Error())
		return
	}

	ret = boolVal
	return
}

func (s *Master) ReadHoldingRegisters(slaveID string, address, count, valueType, endianType uint16) (ret interface{}, exCode byte, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("readHoldingRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
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

	readVal, readExCode, readErr := vVal.(*MBMaster).ReadHoldingRegisters(address, u16Count)
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

	var itemVal interface{}
	var itemErr error
	switch valueType {
	case common.Int16Value:
		itemVal, itemErr = common.BytesToInt16Array(readVal, endianType)
	case common.UInt16Value:
		itemVal, itemErr = common.BytesToUint16Array(readVal, endianType)
	case common.Int32Value:
		itemVal, itemErr = common.BytesToInt32Array(readVal, endianType)
	case common.UInt32Value:
		itemVal, itemErr = common.BytesToUint32Array(readVal, endianType)
	case common.Float32Value:
		itemVal, itemErr = common.BytesToFloat32Array(readVal, endianType)
	case common.Int64Value:
		itemVal, itemErr = common.BytesToInt64Array(readVal, endianType)
	case common.UInt64Value:
		itemVal, itemErr = common.BytesToUint64Array(readVal, endianType)
	case common.Float64Value:
		itemVal, itemErr = common.BytesToFloat64Array(readVal, endianType)
	default:
	}
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

	readVal, readExCode, readErr := vVal.(*MBMaster).ReadInputRegisters(address, u16Count)
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

	var itemVal interface{}
	var itemErr error
	switch valueType {
	case common.Int16Value:
		itemVal, itemErr = common.BytesToInt16Array(readVal, endianType)
	case common.UInt16Value:
		itemVal, itemErr = common.BytesToUint16Array(readVal, endianType)
	case common.Int32Value:
		itemVal, itemErr = common.BytesToInt32Array(readVal, endianType)
	case common.UInt32Value:
		itemVal, itemErr = common.BytesToUint32Array(readVal, endianType)
	case common.Float32Value:
		itemVal, itemErr = common.BytesToFloat32Array(readVal, endianType)
	case common.Int64Value:
		itemVal, itemErr = common.BytesToInt64Array(readVal, endianType)
	case common.UInt64Value:
		itemVal, itemErr = common.BytesToUint64Array(readVal, endianType)
	case common.Float64Value:
		itemVal, itemErr = common.BytesToFloat64Array(readVal, endianType)
	default:
	}
	if itemErr != nil {
		log.Errorf("readInputRegisters failed, decode failed error:%s", itemErr.Error())
		err = cd.NewError(cd.UnExpected, itemErr.Error())
		return
	}

	ret = itemVal
	return
}
