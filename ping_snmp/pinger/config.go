package main

type Config struct {
	Ping PingParams
	Snmp SnmpParams
}

type PingParams struct {
	From    string
	Target  string
	Timeout int // in seconds
}

type SnmpParams struct {
	Target  string
	Port    int
	OID     string
	Timeout int // in seconds
}

var defaultConfig = Config{
	Ping: PingParams{
		From:    "127.0.0.1",
		Target:  "192.168.0.1",
		Timeout: 3,
	},
	Snmp: SnmpParams{
		Target:  "192.168.1.10",
		Port:    161,
		OID:     ".1.3.6.1.4.9.27",
		Timeout: 2,
	},
}

func saveDefaultConfig() {
	config := defaultConfig
	err := configSaveToml(config, "config.toml")
	checkError(err)
}
