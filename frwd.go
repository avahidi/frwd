package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

type connectionData struct {
	err error
	n   int64
}

var udp = flag.Bool("u", false, "forward UDP")
var verbose = flag.Bool("v", false, "verbose")

func usage() {
	log.Fatal("frwd [local-ip]:<local-port> [target-ip]:<target-port>\n")
}

func info(format string, a ...interface{}) {
	if *verbose {
		fmt.Printf("INFO "+format, a...)
	}
}
func fail(format string, a ...interface{}) {
	fmt.Printf("ERROR "+format, a...)
}

func copy(i net.Conn, o net.Conn, ev chan connectionData) {
	n, err := io.Copy(i, o)
	ev <- connectionData{err: err, n: n}
}

func tcpForward(inc net.Conn, outadr string) {
	info("TCP connection from %v...\n", inc.RemoteAddr())

	inadr := inc.LocalAddr().String()
	outc, err := net.Dial("tcp", outadr)
	if err != nil {
		fail("failed to open %s -> %s: %v\n", inadr, outadr, err)
		inc.Close()
	} else {
		info("Connected %s -> %s\n", inadr, outadr)
		ev := make(chan connectionData)
		go copy(inc, outc, ev)
		go copy(outc, inc, ev)

		cd1 := <-ev
		if cd1.err != nil {
			fail("%v\n", err)
		}

		inc.Close()
		outc.Close()
		cd2 := <-ev // wait for the other endpoint
		close(ev)
		info("TCP %s -> %s, data %d/%d\n", inadr, outadr, cd1.n, cd2.n)
	}
}

func tcpServe(inadr, outadr string) error {
	in, err := net.Listen("tcp", inadr)
	if err != nil {
		return err
	}

	for {
		conn, err := in.Accept()
		if err != nil {
			fail("incoming failed: %v\n", err)
		} else {
			go tcpForward(conn, outadr)
		}
	}
}

func udpServe(inadrstr, outadrstr string) error {
	inadr, err := net.ResolveUDPAddr("udp", inadrstr)
	if err != nil {
		return err
	}

	outadr, err := net.ResolveUDPAddr("udp", outadrstr)
	if err != nil {
		return err
	}

	buffer := make([]byte, 2048)
	conn, err := net.ListenUDP("udp", inadr)
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		n, adrin, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fail("when receiving UDP: %v\n", err)
		} else {
			info("Connection from %v...\n", adrin)
			_, err = conn.WriteToUDP(buffer[:n], outadr)
			if err != nil {
				fail("when sending UDP: %v\n", err)
			}
			info("TCP %v -> %v, data %d\n", adrin, outadr, n)
		}
	}
}

func main() {
	if flag.Parse(); flag.NArg() != 2 {
		usage()
	}

	inadr, outadr := flag.Arg(0), flag.Arg(1)
	var err error

	if *udp {
		err = udpServe(inadr, outadr)
	} else {
		err = tcpServe(inadr, outadr)
	}

	if err != nil {
		log.Fatal(err)
	}
}
