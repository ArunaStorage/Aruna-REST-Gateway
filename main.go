package main

import (
	"log"

	"github.com/ScienceObjectsDB/CORE-API-Gateway/config"
	"github.com/ScienceObjectsDB/CORE-API-Gateway/gateway"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	ConfigFile string `short:"c" long:"configfile" description:"File of the config file" default:"config/config-local.yaml"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.HandleConfigFile()

	err = gateway.StartGateway()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
