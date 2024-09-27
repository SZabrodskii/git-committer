package config

import "embed"

//go:embed config.json
var ConfigFS embed.FS

func GetConfigExample() ([]byte, error) {
	return ConfigFS.ReadFile("config.json")
}
