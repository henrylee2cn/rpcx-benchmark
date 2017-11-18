package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	tp "github.com/henrylee2cn/teleport"
)

type Hello struct {
	tp.PullCtx
}

func (h *Hello) Say(args *BenchmarkMessage) (*BenchmarkMessage, tp.Xerror) {
	args.Field1 = "OK"
	args.Field2 = 100
	return args, nil
}

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

	tp.SetRawlogLevel("error")

	var peer = tp.NewPeer(&tp.PeerConfig{
		TlsCertFile:          "",
		TlsKeyFile:           "",
		SlowCometDuration:    time.Millisecond * 500,
		DefaultHeaderCodec:   "protobuf",
		DefaultBodyCodec:     "protobuf",
		DefaultBodyGzipLevel: 0,
		PrintBody:            false,
		ListenAddrs:          []string{*host},
	})
	defer peer.Close()

	peer.PullRouter.Reg(new(Hello))
	peer.Listen()
}
