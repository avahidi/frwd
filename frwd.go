package main

import (
	"flag"
	"io"
	"log"
	"net"
)

type connectionData struct {
	err error
	n   int64
}

func usage() {
	log.Fatal("frwd [listener-ip]:<listener-port> [target-ip]:<target-port>\n")
}

func copy(i net.Conn, o net.Conn, ev chan connectionData) {
	n, err := io.Copy(i, o)
	ev <- connectionData{err: err, n: n}
}

func forward(inc net.Conn, outadr string) {
	inadr := inc.LocalAddr().String()
	outc, err := net.Dial("tcp", outadr)
	if err != nil {
		log.Printf("ERROR: failed to open %s -> %s: %v\n", inadr, outadr, err)
		inc.Close()
	} else {
		log.Printf("INFO: Connected %s -> %s\n", inadr, outadr)
		ev := make(chan connectionData)
		go copy(inc, outc, ev)
		go copy(outc, inc, ev)

		cd1 := <-ev
		if cd1.err != nil {
			log.Printf("ERROR: %v\n", err)
		}

		inc.Close()
		outc.Close()
		cd2 := <-ev // wait for the other endpoint
		close(ev)
		log.Printf("INFO: Closing %s -> %s, data %d/%d\n", inadr, outadr, cd1.n, cd2.n)
	}
}

func main() {
	if flag.Parse(); flag.NArg() != 2 {
		usage()
	}

	inadr, outadr := flag.Arg(0), flag.Arg(1)

	in, err := net.Listen("tcp", inadr)
	if err != nil {
		log.Fatal("Could not open incoming port: %v\n", err)
	}

	for {
		conn, err := in.Accept()
		if err != nil {
			log.Printf("ERROR: incoming failed: %v\n", err)
		} else {
			log.Printf("INFO: Connection from %s...\n", conn.RemoteAddr())
			go forward(conn, outadr)
		}
	}
}
