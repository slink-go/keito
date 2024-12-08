package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func SaveKeyConfig(key []byte) {
	filename, err := getConfigFile()
	if err != nil {
		fmt.Printf("error saving key: %s", err)
	}
	if err = os.WriteFile(filename, key, 0600); err != nil {
		fmt.Printf("error saving key: %s", err)
	}
}
func ReadKeyConfig() []byte {
	filename, err := getConfigFile()
	if err != nil {
		//fmt.Printf("error reading saved key: %s", err)
		return nil
	}
	buff, err := os.ReadFile(filename)
	if err != nil {
		//fmt.Printf("error reading saved key: %s", err)
		return nil
	}
	return buff
}
func getConfigFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".keito"), nil
}
