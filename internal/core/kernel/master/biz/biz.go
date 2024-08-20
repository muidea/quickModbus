package biz

import (
	"fmt"
	cd "github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/foundation/cache"
	"github.com/muidea/magicCommon/foundation/log"
	"github.com/muidea/quickModbus/pkg/common"
	"github.com/muidea/quickModbus/pkg/model"
)

type Master struct {
	slaveInfoCache cache.KVCache
}

func New() *Master {
	return &Master{
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

func (s *Master) ReadCoils(slaveID string, address, count uint16) (ret []bool, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("readCoils failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	readVal, readErr := vVal.(*MBMaster).ReadCoils(address, count)
	if readErr != nil {
		log.Errorf("readCoils failed, error:%s", readErr.Error())
		err = cd.NewError(cd.UnExpected, readErr.Error())
		return
	}

	ret = model.ByteArrayToBoolArray(readVal)
	return
}

func (s *Master) ReadDiscreteInputs(slaveID string, address, count uint16) (ret []bool, err *cd.Result) {
	vVal := s.slaveInfoCache.Fetch(slaveID)
	if vVal == nil {
		errMsg := fmt.Sprintf("no exist slave device %s", slaveID)
		log.Errorf("readDiscreteInputs failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	readVal, readErr := vVal.(*MBMaster).ReadDiscreteInputs(address, count)
	if readErr != nil {
		log.Errorf("readDiscreteInputs failed, error:%s", readErr.Error())
		err = cd.NewError(cd.UnExpected, readErr.Error())
		return
	}

	ret = model.ByteArrayToBoolArray(readVal)
	return
}

func (s *Master) ReadHoldingRegisters(slaveID string, address, count, valueType, endianType uint16) (ret []interface{}, err *cd.Result) {
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
	case common.Int32Value, common.UInt32Value, common.FloatValue:
		u16Count = count * 2
	case common.Int64Value, common.UInt64Value, common.DoubleValue:
		u16Count = count * 4
	default:
	}

	readVal, readErr := vVal.(*MBMaster).ReadHoldingRegisters(address, u16Count)
	if readErr != nil {
		log.Errorf("readHoldingRegisters failed, error:%s", readErr.Error())
		err = cd.NewError(cd.UnExpected, readErr.Error())
		return
	}
	if len(readVal) != int(u16Count*2) {
		errMsg := fmt.Sprintf("illegal read value count")
		log.Errorf("readHoldingRegisters failed, error:%s", errMsg)
		err = cd.NewError(cd.UnExpected, errMsg)
		return
	}

	return
}
