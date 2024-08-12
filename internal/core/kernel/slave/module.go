package slave

import (
	"github.com/muidea/magicCommon/event"
	"github.com/muidea/magicCommon/foundation/log"
	"github.com/muidea/magicCommon/module"
	"github.com/muidea/magicCommon/task"

	"github.com/muidea/quickModbus/internal/config"
	"github.com/muidea/quickModbus/pkg/common"
)

func init() {
	module.Register(New())
}

type Slave struct {
	slavePtr *MBSlave
}

func New() *Slave {
	return &Slave{
		slavePtr: &MBSlave{},
	}
}

func (s *Slave) ID() string {
	return common.SlaveModule
}

func (s *Slave) Setup(endpointName string, eventHub event.Hub, backgroundRoutine task.BackgroundRoutine) {
}

func (s *Slave) Run() {
	err := s.slavePtr.Run(config.BindAddr())
	if err != nil {
		log.Errorf("start modbus slave failed, bindAddr:%s, error:%s", config.BindAddr(), err.Error())
	}
}
