package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	goPath "path"
)

type Config struct {
	File string `json:"storage"`
}

var (
	configPath = ".reminder"
	Dir        = ""
	App        = &Config{File: "storage/output.csv"}
)

func init() {
	Set("")
}

func Set(path string) {
	binPath, pathErr := os.Executable()
	if pathErr != nil {
		fmt.Println(pathErr)
		os.Exit(1)
	}

	Dir = goPath.Dir(binPath)

	if path != "" {
		configPath = path

	} else {
		configPath = goPath.Join(Dir, configPath)
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

func Db() string {
	return goPath.Join(Dir, App.File)
}
