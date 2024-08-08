package master

import (
	"github.com/muidea/magicCommon/event"
	"github.com/muidea/magicCommon/module"
	"github.com/muidea/magicCommon/task"

	"github.com/muidea/quickModbus/pkg/common"
)

func init() {
	module.Register(New())
}

type Master struct {
}

func New() *Master {
	return &Master{}
}

func (s *Master) ID() string {
	return common.MasterModule
}

func (s *Master) Setup(endpointName string, eventHub event.Hub, backgroundRoutine task.BackgroundRoutine) {
}
