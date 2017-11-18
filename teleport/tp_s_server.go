package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	"github.com/henrylee2cn/teleport/socket"
)

var (
	host       = flag.String("s", "127.0.0.1:8972", "listened ip and port")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	debugAddr  = flag.String("d", "127.0.0.1:9981", "server ip and port")
)

func main() {
	flag.Parse()

	go func() {
		log.Println(http.ListenAndServe(*debugAddr, nil))
	}()

	lis, err := net.Listen("tcp", *host)
	if err != nil {
		log.Fatalf("[SVR] listen err: %v", err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("[SVR] accept err: %v", err)
		}
		go func(s socket.Socket) {
			defer s.Close()
			var args = new(BenchmarkMessage)
			for {
				// read request
				var packet = socket.GetPacket(func(_ *socket.Header) interface{} {
					*args = BenchmarkMessage{}
					return args
				})
				err = s.ReadPacket(packet)
				if err != nil {
					// log.Printf("[SVR] read request err: %v", err)
					return
				} else {
					// log.Printf("[SVR] read request: %v", packet)
				}

				// write response
				packet.Header.StatusCode = 200
				packet.Header.Status = "ok"

				args.Field1 = "OK"
				args.Field2 = 100

				err = s.WritePacket(packet)
				if err != nil {
					log.Printf("[SVR] write response err: %v", err)
				} else {
					// log.Printf("[SVR] write response: %v", packet)
				}
				socket.PutPacket(packet)
			}
		}(socket.GetSocket(conn))
	}
}
