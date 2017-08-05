package main

import (
	"flag"
	"net"
	"io"
	"log"
)

func usage() {
	log.Fatal("frwd [listener-ip]:<listener-port> [target-ip]:<target-port>\n")
}

func copy(i net.Conn, o net.Conn, ev chan error) {
	_, err := io.Copy(i, o)
	ev <- err
}

func forward( inc net.Conn, outadr string) {
	outc, err := net.Dial("tcp", outadr)
	if err != nil {		
		log.Printf("ERROR: could not open outgoing: %v\n", err)
		inc.Close()
	} else {
		ev := make(chan error)
		go copy(inc, outc, ev)
		go copy(outc, inc, ev)

		err := <- ev
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		
		err = <- ev // wait for the other endpoint
		inc.Close()
		outc.Close()
		close(ev)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		usage()
	}

	inadr := flag.Arg(0)
	outadr := flag.Arg(1)

	in, err := net.Listen("tcp", inadr)
	if err != nil {
		log.Fatal("Could not open incoming port: %v\n", err)
	}

	for {
		conn, err := in.Accept()
		if err != nil {
			log.Printf("ERROR: incoming failed: %v\n", err)
		} else {
			go forward(conn, outadr)
		}
	}
}
