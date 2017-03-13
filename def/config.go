package def

import (
	"encoding/json"
	"os"
)

// Configuration keeps the configs for the entire application
type Configuration struct {
	Server ServerConfig

	// TmpDir is the directory to keep session files in
	// If the path is relative, it will be used as a subfolder of the system temporary directory
	TmpDir string
}

// ServerConfig keeps the configuration options needed to make a Server
type ServerConfig struct {
	Address          string
	ReadTimeoutSecs  int
	WriteTimeoutSecs int
}

// ReadConfigFile reads a configuration file
func ReadConfigFile(configFilepath string) (Configuration, error) {
	var config Configuration

	file, err := os.Open(configFilepath)
	if err != nil {
		return config, Err(err, "Cannot open config file %s", configFilepath)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, Err(err, "Cannot read config file %s", configFilepath)
	}

	return config, nil
}
