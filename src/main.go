package main

import (
	"fmt"
	"keito/cmd"
	_ "keito/cmd/key"
	_ "keito/cmd/parse"
	_ "keito/cmd/token"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	cmd.Execute()
}
