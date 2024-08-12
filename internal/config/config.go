package config

const cfgPath = "/var/app/config/cfg.json"

func init() {
	err := LoadConfig(cfgPath)
	if err == nil {
		return
	}
}

func LoadConfig(cfgFile string) (err error) {
	if cfgFile == "" {
		return
	}
	return
}

func BindAddr() string {
	return "0.0.0.0:502"
}

func SlaveAddr() string {
	return ""
}

type config struct {
	ModbusBindPort string `json:"modbusBindPort"`
}
