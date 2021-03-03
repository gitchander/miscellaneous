package main

import (
	"fmt"
	"log"
	"os"
	"time"

	g "github.com/gosnmp/gosnmp"
)

func main() {
	//testGet()
	testSet()
}

func testGet() {
	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	g.Default.Target = "192.168.1.10"
	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	oids := []string{"1.3.6.1.2.1.1.4.0", "1.3.6.1.2.1.1.7.0"}
	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case g.OctetString:
			bytes := variable.Value.([]byte)
			fmt.Printf("string: %s\n", string(bytes))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
		}
	}
}

func testSet() {

	envTarget := "192.168.1.10"
	port := 161

	params := &g.GoSNMP{
		Target:    envTarget,
		Port:      uint16(port),
		Community: "public",
		Version:   g.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		Logger:    log.New(os.Stdout, "", 0),
	}
	err := params.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer params.Conn.Close()

	pdus := []g.SnmpPDU{
		g.SnmpPDU{
			Name:  ".1.3.6.1.4.9.27", // oid
			Type:  g.OctetString,
			Value: []byte("Hello, Snmp!"),
			//Logger: params.Logger,
		},
	}

	result, err := params.Set(pdus)
	if err != nil {
		log.Fatal("error -> ", err)
	}
	fmt.Println(result)
	fmt.Println("ok")
}
