package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var (
	server = ""
)

func init() {
	flag.StringVar(&server, "s", "127.0.0.1:80", "server's ip address and port ext: 127.0.0.1:80")
	flag.Parse()

}

func main() {
	fmt.Println(dialCost())
}
func dialCost() string {

	startTime := time.Now()
	timeout := time.NewTicker(time.Second * 5)
	go func() {
		for {
			select {
			case <-timeout.C:
				log.Fatal("time out")
			default:
				continue
			}
		}
	}()
	conn, err := net.DialTCP("tcp", nil, resolveTCPAddr(server))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cost := time.Since(startTime).String()
	return cost
}

func resolveTCPAddr(addr string) *net.TCPAddr {
	resolved, error := net.ResolveTCPAddr("tcp", addr)
	if nil != error {
		log.Fatalf("Error occured while resolving TCP address \"%v\": %v\n", addr, error)
	}

	return resolved
}
