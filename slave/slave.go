package slave

import "github.com/muidea/magicEngine/tcp"

type MBSlave struct {
	tcpServer tcp.Server
}

func (s *MBSlave) Run(bindAddr string) (err error) {
	server := tcp.NewServer(s, 100)
	s.tcpServer = server
	err = server.Run(bindAddr)
	if err != nil {
		return
	}

	return
}

func (s *MBSlave) OnConnect(ep tcp.Endpoint) {

}

func (s *MBSlave) OnDisConnect(ep tcp.Endpoint) {

}

func (s *MBSlave) OnRecvData(ep tcp.Endpoint, data []byte) {

}
