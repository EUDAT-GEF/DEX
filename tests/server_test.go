package tests

import (
	"log"
	"testing"

	"github.com/EUDAT-GEF/DEX/def"
)

var configFilePath = "../config.json"

func TestServer(t *testing.T) {
	config, err := def.ReadConfigFile(configFilePath)
	checkMsg(t, err, "reading config files")

	log.Println(config.Server.Address)
}
