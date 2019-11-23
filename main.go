package main

import (
	"flag"
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

	nowt := time.Now()
	timeout := time.NewTicker(time.Second * 5)
	conn, err := net.DialTCP("tcp", nil, resolveTCPAddr(server))
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

	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
	log.Printf("cost : %s", time.Since(nowt).String())

}

func resolveTCPAddr(addr string) *net.TCPAddr {
	resolved, error := net.ResolveTCPAddr("tcp", addr)
	if nil != error {
		log.Fatalf("Error occured while resolving TCP address \"%v\": %v\n", addr, error)
	}

	return resolved
}
