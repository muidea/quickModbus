package service

import (
	"context"
	"net/http"

	engine "github.com/muidea/magicEngine/http"
)

type Master struct {
	routeRegistry engine.RouteRegistry
}

func (s *Master) BindRegistry(router engine.RouteRegistry) {
	s.routeRegistry = router
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
