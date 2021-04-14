package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"redirect-tls/config"
	"redirect-tls/handler"
)

const version = "0.1.0"

var conf config.Config

func main() {
	fmt.Println("redirect-tls version", version)

	configPath := flag.String("config", "./config.yaml", "Path to config file")
	flag.Parse()

	var err error
	conf, err = config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to read config %s: %v", *configPath, err)
	}

	if conf.RedirectHttps != "" {
		handler.ServeRedirectHttps(conf.RedirectHttps)
	}

	if conf.Listen != "" {
		listenAndServe()
	}
}

func listenAndServe() {
	ln, err := net.Listen("tcp", conf.Listen)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", conf.Listen, err)
	}
	defer func() { _ = ln.Close() }()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("fail to establish conn: %v\n", err)
			continue
		}
		go conf.Http.Handle(conn)
	}
}
