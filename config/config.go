package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"redirect-tls/handler"
)

type (
	RawConfig struct {
		Listen                                string
		RedirectHttps                         string
		InboundBufferSize, OutboundBufferSize int
		Http                                  RawHandler
	}
	RawHandler struct {
		Handler string
		Args    string
	}
)

type (
	Config struct {
		Listen        string
		RedirectHttps string
		Http          handler.Handler
	}
)

func readRawConfig(path string) (conf RawConfig, err error) {
	conf = RawConfig{InboundBufferSize: 4, OutboundBufferSize: 32}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return
	}
	return
}

func ReadConfig(path string) (conf Config, err error) {
	rawConf, err := readRawConfig(path)
	if err != nil {
		return
	}

	handler.InitBufferPools(rawConf.InboundBufferSize*1024, rawConf.OutboundBufferSize*1024)

	conf.Listen = rawConf.Listen
	conf.RedirectHttps = rawConf.RedirectHttps

	conf.Http = newHandler(rawConf.Http.Handler, rawConf.Http.Args)
	return
}
func newHandler(name, args string) handler.Handler {
	switch name {
	case "":
		return handler.NoopHandler
	case "proxyPass":
		return handler.NewProxyPassHandler(args)
	case "fileServer":
		return handler.NewFileServerHandler(args)
	default:
		log.Fatalf("handler %s not supported\n", name)
	}
	return nil
}
