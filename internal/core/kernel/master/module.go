package master

import (
	"github.com/muidea/magicCommon/event"
	"github.com/muidea/magicCommon/module"
	"github.com/muidea/magicCommon/task"
	engine "github.com/muidea/magicEngine/http"

	"github.com/muidea/quickModbus/pkg/common"
)

func init() {
	module.Register(New())
}

type Master struct {
	routeRegistry     engine.Router
	eventHub          event.Hub
	backgroundRoutine task.BackgroundRoutine
}

func New() *Master {
	return &Master{}
}

func (s *Master) ID() string {
	return common.MasterModule
}

func (s *Master) BindRegistry(router engine.Router) {
	s.routeRegistry = router
}

func (s *Master) Setup(endpointName string, eventHub event.Hub, backgroundRoutine task.BackgroundRoutine) {
	s.eventHub = eventHub
	s.backgroundRoutine = backgroundRoutine

}

func (s *Master) Run() {

}
