package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	File string `json:"storage"`
}

var (
	configPath = ".reminder"
	App        = &Config{File: "storage/output.csv"}
)

func init() {
	Set("")
}

func Set(path string) {
	if path != "" {
		configPath = path
	}

	_, err := os.Stat(configPath)
	if err != nil {
		return
	}

	jsonFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byte, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byte, &App)
}
