package biz

import (
	"fmt"
	cd "github.com/muidea/magicCommon/def"
	"github.com/muidea/magicCommon/foundation/cache"
)

type Master struct {
	slaveInfoCache cache.KVCache
}

func (s *Master) ConnectSlave(slaveAddr string, devID int) (ret string, err *cd.Result) {
	mbID := fmt.Sprintf("mb%03d", devID)
	val := s.slaveInfoCache.Fetch(mbID)
	if val != nil {
		//err = cd.NewError(cd.)
	}
	return
}

func (s *Master) DisConnectSlave(slaveID string) (err *cd.Result) {
	return
}
