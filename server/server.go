package main

import (
	"../config"
	"fmt"
	"net"
	"os"
	"rgoq"
)

var message = []byte("hello world")

func main() {
	service := ":1201"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	readLen, addr, err := conn.ReadFromUDP(buf[0:])
	checkError(err)
	fmt.Println("Read ", readLen, "bytesfrom ", addr)
	mTest := rgoq.CreateQueue(config.REDIS_OPTIONS, config.SERVER_KEY)
	errs := mTest.PushString(string(buf[:readLen]))
	checkError(errs)
	fmt.Println(string(buf[:readLen]))

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

