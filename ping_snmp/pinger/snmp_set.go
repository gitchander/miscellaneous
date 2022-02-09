package main

import (
	"fmt"
	"log"
	"os"
	"time"

	g "github.com/gosnmp/gosnmp"
)

func setSnmp(config Config, state bool) error {

	params := &g.GoSNMP{
		Target:    config.Snmp.Target,
		Port:      uint16(config.Snmp.Port),
		Community: "public",
		Version:   g.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		Logger:    g.NewLogger(log.New(os.Stdout, "", 0)),
	}
	err := params.Connect()
	if err != nil {
		return fmt.Errorf("snmp Connect() err: %v", err)
	}
	defer params.Conn.Close()

	val := fmt.Sprintf("[%s] -> [%s]: %v", config.Ping.From, config.Ping.Target, state)
	log.Println(val)

	pdus := []g.SnmpPDU{
		g.SnmpPDU{
			Name:  config.Snmp.OID,
			Type:  g.OctetString,
			Value: []byte(val),
			//Logger: params.Logger,
		},
	}

	result, err := params.Set(pdus)
	if err != nil {
		return fmt.Errorf("snmp Set() err: %s", err)
	}
	fmt.Println("snmp Set() result:", result)
	fmt.Println("snmp Set() ok")
	return nil
}
