package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"github.com/tquach/zk-rest-api/strflag"
)

const (
	DefaultZKTimeout = 10 * time.Second
)

var (
	addr    = flag.String("addr", ":8080", "address and port to listen on")
	zkHosts strflag.StringSlice
)

func main() {
	flag.Parse()

	if len(zkHosts) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	conn, _, err := zk.Connect(zkHosts, DefaultZKTimeout)
	if err != nil {
		log.Fatalf("zk fail: %s", err)
	}

	ctrl := &Controller{zkConn: conn}
	log.Printf("Starting service on %s...", *addr)
	log.Fatal(http.ListenAndServe(*addr, ctrl))
}

func init() {
	flag.Var(&zkHosts, "zk", "Comma-separated list of hosts to zookeeper.")
}
