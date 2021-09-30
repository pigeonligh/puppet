package main

import (
	"fmt"

	"github.com/hpcloud/tail"
	"github.com/pigeonligh/puppet/common"
)

func main() {
	t, err := tail.TailFile(common.CommandPath, tail.Config{
		Follow: true,
	})
	if err != nil {
		fmt.Printf("failed to tail file, err: %v", err)
		return
	}

	for line := range t.Lines {
		cmd := common.NewFromJSON(line.Text)
		if cmd == nil {
			continue
		}
		cmd.Exec()
	}
}
