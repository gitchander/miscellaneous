package main

import (
	"fmt"
	"net"
	"os"
	"time"

	ping1 "github.com/go-ping/ping"
	ping2 "github.com/paulstuart/ping"
	ping3 "github.com/tatsushid/go-fastping"
)

func main() {
	testPing1_V1()
}

func testPing1_V1() {
	if len(os.Args) < 2 {
		fmt.Println("no args")
		return
	}
	err := ping2.Pinger(os.Args[1], 5)
	if err == nil {
		fmt.Println("ok")
		return
	}
	fmt.Println(err)
}

func testPing1_V2() {
	fmt.Println(ping2.Ping(os.Args[1], 5))
}

func testPing2() {
	p := ping3.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func testPing3_V1() {
	pinger, err := ping1.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Println(stats)
}

func testPing3_V2() {
	pinger, err := ping1.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}

	pinger.OnRecv = func(pkt *ping1.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *ping1.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Run()
}
