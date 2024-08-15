package service

import (
	"context"
	"github.com/muidea/quickModbus/pkg/common"
	"net/http"

	engine "github.com/muidea/magicEngine/http"
)

type Master struct {
	routeRegistry engine.RouteRegistry
}

func (s *Master) BindRegistry(router engine.RouteRegistry) {
	s.routeRegistry = router
}

func (s *Master) RegisterRoute() {
	s.routeRegistry.AddHandler(common.ConnectSlave, engine.POST, s.ConnectSlave)
	s.routeRegistry.AddHandler(common.DisConnectSlave, engine.DELETE, s.DisConnectSlave)
	s.routeRegistry.AddHandler(common.ReadCoils, engine.POST, s.ReadCoils)
	s.routeRegistry.AddHandler(common.ReadDiscreteInputs, engine.POST, s.ReadDiscreteInputs)
	s.routeRegistry.AddHandler(common.ReadHoldingRegisters, engine.POST, s.ReadHoldingRegisters)
	s.routeRegistry.AddHandler(common.ReadInputRegisters, engine.POST, s.ReadInputRegisters)
	s.routeRegistry.AddHandler(common.WriteSingleCoil, engine.POST, s.WriteSingleCoil)
	s.routeRegistry.AddHandler(common.WriteSingleRegister, engine.POST, s.WriteSingleRegister)
	s.routeRegistry.AddHandler(common.ReadExceptionStatus, engine.POST, s.ReadExceptionStatus)
	s.routeRegistry.AddHandler(common.Diagnostics, engine.POST, s.Diagnostics)
	s.routeRegistry.AddHandler(common.GetCommEventCounter, engine.POST, s.GetCommEventCounter)
	s.routeRegistry.AddHandler(common.GetCommEventLog, engine.POST, s.GetCommEventLog)
	s.routeRegistry.AddHandler(common.WriteMultipleCoils, engine.POST, s.WriteMultipleCoils)
	s.routeRegistry.AddHandler(common.WriteMultipleRegisters, engine.POST, s.WriteMultipleRegisters)
	s.routeRegistry.AddHandler(common.ReportSlaveID, engine.POST, s.ReportSlaveID)
	s.routeRegistry.AddHandler(common.ReadFileRecord, engine.POST, s.ReadFileRecord)
	s.routeRegistry.AddHandler(common.WriteFileRecord, engine.POST, s.WriteFileRecord)
	s.routeRegistry.AddHandler(common.MaskWriteRegister, engine.POST, s.MaskWriteRegister)
	s.routeRegistry.AddHandler(common.ReadWriteMultipleRegisters, engine.POST, s.ReadWriteMultipleRegisters)
	s.routeRegistry.AddHandler(common.ReadFIFOQueue, engine.POST, s.ReadFIFOQueue)
}

func (s *Master) ConnectSlave(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) DisConnectSlave(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) ReadCoils(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) ReadDiscreteInputs(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) ReadHoldingRegisters(ctx context.Context, res http.ResponseWriter, req *http.Request) {

}

func (s *Master) ReadInputRegisters(ctx context.Context, res http.ResponseWriter, req *http.Request) {

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
