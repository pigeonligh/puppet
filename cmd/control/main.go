package main

import (
	"fmt"
	"os"

	"github.com/pigeonligh/puppet/common"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("empty command")
		return
	}

	cmd := common.New(args[1], args[2:])
	if cmd != nil {
		file, err := os.OpenFile(common.CommandPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Printf("failed to open file, err: %v", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(cmd.JSON() + "\n")
		if err != nil {
			fmt.Printf("failed to write file, err: %v", err)
			return
		}
	}
}
