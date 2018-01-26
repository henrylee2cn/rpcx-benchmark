package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"

	tp "github.com/henrylee2cn/teleport"
)

type Hello struct {
	tp.PullCtx
}

func (h *Hello) Say(args *BenchmarkMessage) (*BenchmarkMessage, *tp.Rerror) {
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

	tp.SetLoggerLevel("error")
	tp.SetSocketNoDelay(false)

	var peer = tp.NewPeer(tp.PeerConfig{
		DefaultBodyCodec: "protobuf",
		ListenAddress:    *host,
	})
	defer peer.Close()

	peer.RoutePull(new(Hello))
	peer.Listen()
}
