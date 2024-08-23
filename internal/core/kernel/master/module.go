package master

import (
	"github.com/muidea/magicCommon/event"
	"github.com/muidea/magicCommon/module"
	"github.com/muidea/magicCommon/task"
	engine "github.com/muidea/magicEngine/http"

	"github.com/muidea/quickModbus/internal/core/kernel/master/biz"
	"github.com/muidea/quickModbus/internal/core/kernel/master/service"
	"github.com/muidea/quickModbus/pkg/common"
)

func init() {
	module.Register(New())
}

type Master struct {
	routeRegistry     engine.RouteRegistry
	eventHub          event.Hub
	backgroundRoutine task.BackgroundRoutine

	bizPtr     *biz.Master
	servicePtr *service.Master
}

func New() *Master {
	return &Master{}
}

func (s *Master) ID() string {
	return common.MasterModule
}

func (s *Master) BindRegistry(router engine.RouteRegistry) {
	s.routeRegistry = router
}

func (s *Master) Setup(endpointName string, eventHub event.Hub, backgroundRoutine task.BackgroundRoutine) {
	s.eventHub = eventHub
	s.backgroundRoutine = backgroundRoutine

	s.bizPtr = biz.New(eventHub, backgroundRoutine)
	s.servicePtr = service.New(s.bizPtr)
	s.servicePtr.BindRegistry(s.routeRegistry)
}

func (s *Master) Run() {
	s.servicePtr.RegisterRoute()
}
