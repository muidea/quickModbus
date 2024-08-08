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

type config struct {
	ModbusBindPort string `json:"modbusBindPort"`
}
