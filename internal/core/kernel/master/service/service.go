package service

import (
	"context"
	"encoding/json"
	"github.com/muidea/magicCommon/foundation/log"
	"net/http"
	"strings"

	cd "github.com/muidea/magicCommon/def"
	fn "github.com/muidea/magicCommon/foundation/net"

	engine "github.com/muidea/magicEngine/http"

	"github.com/muidea/quickModbus/internal/core/kernel/master/biz"
	"github.com/muidea/quickModbus/pkg/common"
)

const slaveIDContextKey = "_slaveID"

type Master struct {
	routeRegistry engine.RouteRegistry

	bizPtr *biz.Master
}

func New(bizPtr *biz.Master) *Master {
	return &Master{
		bizPtr: bizPtr,
	}
}

func (s *Master) BindRegistry(router engine.RouteRegistry) {
	s.routeRegistry = router
}

func (s *Master) RegisterRoute() {
	s.routeRegistry.AddHandler(common.ConnectSlave, engine.POST, s.ConnectSlave)
	s.routeRegistry.AddHandler(common.DisConnectSlave, engine.DELETE, s.DisConnectSlave, s)
	s.routeRegistry.AddHandler(common.ReadCoils, engine.POST, s.ReadCoils, s)
	s.routeRegistry.AddHandler(common.ReadDiscreteInputs, engine.POST, s.ReadDiscreteInputs, s)
	s.routeRegistry.AddHandler(common.ReadHoldingRegisters, engine.POST, s.ReadHoldingRegisters, s)
	s.routeRegistry.AddHandler(common.ReadInputRegisters, engine.POST, s.ReadInputRegisters, s)
	s.routeRegistry.AddHandler(common.WriteSingleCoil, engine.POST, s.WriteSingleCoil, s)
	s.routeRegistry.AddHandler(common.WriteSingleRegister, engine.POST, s.WriteSingleRegister, s)
	s.routeRegistry.AddHandler(common.ReadExceptionStatus, engine.POST, s.ReadExceptionStatus, s)
	s.routeRegistry.AddHandler(common.Diagnostics, engine.POST, s.Diagnostics, s)
	s.routeRegistry.AddHandler(common.GetCommEventCounter, engine.POST, s.GetCommEventCounter, s)
	s.routeRegistry.AddHandler(common.GetCommEventLog, engine.POST, s.GetCommEventLog, s)
	s.routeRegistry.AddHandler(common.WriteMultipleCoils, engine.POST, s.WriteMultipleCoils, s)
	s.routeRegistry.AddHandler(common.WriteMultipleRegisters, engine.POST, s.WriteMultipleRegisters, s)
	s.routeRegistry.AddHandler(common.ReportSlaveID, engine.POST, s.ReportSlaveID, s)
	s.routeRegistry.AddHandler(common.ReadFileRecord, engine.POST, s.ReadFileRecord, s)
	s.routeRegistry.AddHandler(common.WriteFileRecord, engine.POST, s.WriteFileRecord, s)
	s.routeRegistry.AddHandler(common.MaskWriteRegister, engine.POST, s.MaskWriteRegister, s)
	s.routeRegistry.AddHandler(common.ReadWriteMultipleRegisters, engine.POST, s.ReadWriteMultipleRegisters, s)
	s.routeRegistry.AddHandler(common.ReadFIFOQueue, engine.POST, s.ReadFIFOQueue, s)
}

func (s *Master) MiddleWareHandle(ctx engine.RequestContext, res http.ResponseWriter, req *http.Request) {
	pathItems := strings.Split(req.URL.Path, "/")
	if len(pathItems) < 3 {
		return
	}

	ctx.Update(context.WithValue(ctx.Context(), slaveIDContextKey, pathItems[2]))
}

func (s *Master) ConnectSlave(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	result := &common.ConnectSlaveResponse{}
	for {
		param := &common.ConnectSlaveRequest{}
		err := fn.ParseJSONBody(req, nil, param)
		if err != nil {
			result.ErrorCode = cd.IllegalParam
			result.Reason = "invalid param"
			break
		}

		slaveID, slaveErr := s.bizPtr.ConnectSlave(param.SlaveAddr, param.DeviceID)
		if slaveErr != nil {
			log.Errorf("connect slave failed, slaveAddr:%s, deviceID:%v, error:%s", param.SlaveAddr, param.DeviceID, slaveErr.Error())
			result.Result = *slaveErr
			break
		}

		result.SlaveID = slaveID
		result.ErrorCode = cd.Succeeded
		break
	}

	block, err := json.Marshal(result)
	if err == nil {
		_, _ = res.Write(block)
		return
	}

	res.WriteHeader(http.StatusExpectationFailed)
}

func (s *Master) DisConnectSlave(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	result := &cd.Result{}
	for {
		slaveID := ctx.Value(slaveIDContextKey).(string)
		disconnectResult := s.bizPtr.DisConnectSlave(slaveID)
		if disconnectResult != nil {
			log.Errorf("disconnect slave failed, slaveIDContextKey:%v, error:%s", slaveID, disconnectResult.Error())
			result = disconnectResult
			break
		}

		result.ErrorCode = cd.Succeeded
		break
	}

	block, err := json.Marshal(result)
	if err == nil {
		_, _ = res.Write(block)
		return
	}

	res.WriteHeader(http.StatusExpectationFailed)
}

func (s *Master) ReadCoils(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	result := &common.ReadCoilsResponse{}
	for {
		param := &common.ReadCoilsRequest{}
		err := fn.ParseJSONBody(req, nil, param)
		if err != nil {
			result.ErrorCode = cd.IllegalParam
			result.Reason = "invalid param"
			break
		}
		slaveID := ctx.Value(slaveIDContextKey).(string)
		readVal, readExCode, readErr := s.bizPtr.ReadCoils(slaveID, param.Address, param.Count, param.EndianType)
		result.ExceptionCode = readExCode
		if readErr != nil {
			log.Errorf("read coils failed, slaveID:%s, address:%d, count:%d, endianType:%d, exCode:%v, error:%s", slaveID, param.Address, param.Count, param.EndianType, readExCode, readErr.Error())
			result.Result = *readErr
			break
		}

		result.Values = readVal
		result.ErrorCode = cd.Succeeded
		break
	}

	block, err := json.Marshal(result)
	if err == nil {
		_, _ = res.Write(block)
		return
	}

	res.WriteHeader(http.StatusExpectationFailed)
}

func (s *Master) ReadDiscreteInputs(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	result := &common.ReadDiscreteInputsResponse{}
	for {
		param := &common.ReadDiscreteInputsRequest{}
		err := fn.ParseJSONBody(req, nil, param)
		if err != nil {
			result.ErrorCode = cd.IllegalParam
			result.Reason = "invalid param"
			break
		}
		slaveID := ctx.Value(slaveIDContextKey).(string)
		readVal, readExCode, readErr := s.bizPtr.ReadDiscreteInputs(slaveID, param.Address, param.Count, param.EndianType)
		result.ExceptionCode = readExCode
		if readErr != nil {
			log.Errorf("read discrete inputs failed, slaveID:%s, address:%d, count:%d, endianType:%d, exCode:%v, error:%s", slaveID, param.Address, param.Count, param.EndianType, readExCode, readErr.Error())
			result.Result = *readErr
			break
		}

		result.Values = readVal
		result.ErrorCode = cd.Succeeded
		break
	}

	block, err := json.Marshal(result)
	if err == nil {
		_, _ = res.Write(block)
		return
	}

	res.WriteHeader(http.StatusExpectationFailed)
}

func (s *Master) ReadHoldingRegisters(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	result := &common.ReadHoldingRegistersResponse{}
	for {
		param := &common.ReadHoldingRegistersRequest{}
		err := fn.ParseJSONBody(req, nil, param)
		if err != nil {
			result.ErrorCode = cd.IllegalParam
			result.Reason = "invalid param"
			break
		}
		slaveID := ctx.Value(slaveIDContextKey).(string)
		readVal, readExCode, readErr := s.bizPtr.ReadHoldingRegisters(slaveID, param.Address, param.Count, param.ValueType, param.EndianType)
		result.ExceptionCode = readExCode

		if readErr != nil {
			log.Errorf("read holding registers failed, slaveID:%s, address:%d, count:%d, valueType:%d, endianType:%d, exCode:%v, error:%s", slaveID, param.Address, param.Count, param.ValueType, param.EndianType, readExCode, readErr.Error())
			result.Result = *readErr
			break
		}

		result.Values = readVal
		result.ErrorCode = cd.Succeeded
		break
	}

	block, err := json.Marshal(result)
	if err == nil {
		_, _ = res.Write(block)
		return
	}

	res.WriteHeader(http.StatusExpectationFailed)
}

func (s *Master) ReadInputRegisters(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	result := &common.ReadReadInputRegistersResponse{}
	for {
		param := &common.ReadReadInputRegistersRequest{}
		err := fn.ParseJSONBody(req, nil, param)
		if err != nil {
			result.ErrorCode = cd.IllegalParam
			result.Reason = "invalid param"
			break
		}
		slaveID := ctx.Value(slaveIDContextKey).(string)
		readVal, readExCode, readErr := s.bizPtr.ReadInputRegisters(slaveID, param.Address, param.Count, param.ValueType, param.EndianType)
		result.ExceptionCode = readExCode
		if readErr != nil {
			log.Errorf("read input registers failed, slaveID:%s, address:%d, count:%d, valueType:%d, endianType:%d, exCode:%v, error:%s", slaveID, param.Address, param.Count, param.ValueType, param.EndianType, readExCode, readErr.Error())
			result.Result = *readErr
			break
		}

		result.Values = readVal
		result.ErrorCode = cd.Succeeded
		break
	}

	block, err := json.Marshal(result)
	if err == nil {
		_, _ = res.Write(block)
		return
	}

	res.WriteHeader(http.StatusExpectationFailed)
}

func (s *Master) WriteSingleCoil(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) WriteSingleRegister(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) ReadExceptionStatus(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) Diagnostics(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) GetCommEventCounter(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) GetCommEventLog(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) WriteMultipleCoils(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) WriteMultipleRegisters(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) ReportSlaveID(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) ReadFileRecord(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) WriteFileRecord(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) MaskWriteRegister(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) ReadWriteMultipleRegisters(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) ReadFIFOQueue(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}
