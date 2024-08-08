package slave

import (
	"github.com/muidea/magicCommon/event"
	"github.com/muidea/magicCommon/module"
	"github.com/muidea/magicCommon/task"

	"github.com/muidea/quickModbus/pkg/common"
)

func init() {
	module.Register(New())
}

type Slave struct {
}

func New() *Slave {
	return &Slave{}
}

func (s *Slave) ID() string {
	return common.SlaveModule
}

func (s *Slave) Setup(endpointName string, eventHub event.Hub, backgroundRoutine task.BackgroundRoutine) {
}
