package main

import (
	"flag"
	"log"

	"github.com/EUDAT-GEF/DEX/def"
	"github.com/EUDAT-GEF/DEX/server"
)

var configFilePath = "config.json"

func main() {
	flag.StringVar(&configFilePath, "config", configFilePath, "configuration file")
	flag.Parse()

	config, err := def.ReadConfigFile(configFilePath)
	if err != nil {
		log.Fatal("FATAL: ", err)
	}

	server, err := server.NewServer(config.Server, config.TmpDir)
	if err != nil {
		log.Fatal("FATAL: ", def.Err(err, "Cannot create API server"))
	}

	log.Println("Starting DEX server at: ", config.Server.Address)
	err = server.Start()
	if err != nil {
		log.Fatal("FATAL: ", def.Err(err, "Cannot start API server"))
	}
}
