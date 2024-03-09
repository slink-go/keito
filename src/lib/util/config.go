package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func SaveKeyConfig(key string) {
	filename, err := getConfigFile()
	if err != nil {
		fmt.Printf("error saving key: %s", err)
	}
	if err = os.WriteFile(filename, []byte(key), 0600); err != nil {
		fmt.Printf("error saving key: %s", err)
	}
}
func ReadKeyConfig() string {
	filename, err := getConfigFile()
	if err != nil {
		//fmt.Printf("error reading saved key: %s", err)
		return ""
	}
	buff, err := os.ReadFile(filename)
	if err != nil {
		//fmt.Printf("error reading saved key: %s", err)
		return ""
	}
	return string(buff)
}
func getConfigFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".keito"), nil
}
