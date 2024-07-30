package master

import "github.com/muidea/magicEngine/tcp"

type MBMaster struct {
	tcpClient tcp.Client
}

func (s *MBMaster) Run(serverAddr string) (err error) {
	client := tcp.NewClient(s)
	err = client.Connect(serverAddr)
	if err != nil {
		return
	}

	s.tcpClient = client
	return
}

func (s *MBMaster) OnConnect(ep tcp.Endpoint) {

}

func (s *MBMaster) OnDisConnect(ep tcp.Endpoint) {

}

func (s *MBMaster) OnRecvData(ep tcp.Endpoint, data []byte) {

}

func (s *MBMaster) ReadCoils(address, count uint16) (ret []byte, err error) {
	return
}

func (s *MBMaster) ReadDiscreteInputs(address, count uint16) (ret []byte, err error) {
	return
}

func (s *MBMaster) ReadHoldingRegisters(address, count uint16) (ret []byte, err error) {
	return
}

func (s *MBMaster) ReadInputRegisters(address, count uint16) (ret []byte, err error) {
	return
}

func (s *MBMaster) WriteSingleCoil(address uint16, data []byte) (ret []byte, err error) {
	return
}

func (s *MBMaster) WriteMultipleCoils(address, count uint16, data []byte) (ret []byte, err error) {
	return
}

func (s *MBMaster) WriteSingleRegister(address uint16, data []byte) (ret []byte, err error) {
	return
}

func (s *MBMaster) WriteMultipleRegisters(address, count uint16, data []byte) (ret []byte, err error) {
	return
}

func (s *MBMaster) ReadExceptionStatus() (ret []byte, err error) {
	return
}

func (s *MBMaster) Diagnostics() (ret []byte, err error) {
	return
}

func (s *MBMaster) GetCommEventCounter() (ret []byte, err error) {
	return
}

func (s *MBMaster) GetCommEventLog() (ret []byte, err error) {
	return
}

func (s *MBMaster) ReportServerID() (ret []byte, err error) {
	return
}

func (s *MBMaster) ReadFileRecord() (ret []byte, err error) {
	return
}

func (s *MBMaster) WriteFileRecord() (ret []byte, err error) {
	return
}

func (s *MBMaster) MaskWriteRegister() (ret []byte, err error) {
	return
}

func (s *MBMaster) ReadWriteMultipleRegisters() (ret []byte, err error) {
	return
}

func (s *MBMaster) ReadFIFOQueue() (ret []byte, err error) {
	return
}
