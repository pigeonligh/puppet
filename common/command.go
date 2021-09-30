package common

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	CommandSleep = "sleep"
	CommandExit  = "exit"
)

type Command struct {
	Caller string        `json:"caller"`
	Args   []interface{} `json:"args"`
}

func (c *Command) String() string {
	args := make([]string, len(c.Args))
	for i, arg := range c.Args {
		args[i] = fmt.Sprint(arg)
	}
	return fmt.Sprintf("%s %s", c.Caller, strings.Join(args, " "))
}

func (c *Command) JSON() string {
	data, err := json.Marshal(c)
	if err != nil {
		Logger.Errorf("failed to marshal command, err: %v", err)
		return ""
	}
	return string(data)
}

func (c *Command) Exec() {
	Logger.Infof("run: %s", c.String())
	switch c.Caller {
	case CommandSleep:
		d, err := time.ParseDuration(c.Args[0].(string))
		if err != nil {
			Logger.Errorf("failed to parse duration, err: %v", err)
			return
		}
		time.Sleep(d)
	case CommandExit:
		v, err := strconv.Atoi(c.Args[0].(string))
		if err != nil {
			Logger.Errorf("failed to parse code, err: %v", err)
			return
		}
		os.Exit(v)
	}
}

func NewFromJSON(cmd string) *Command {
	c := &Command{}
	err := json.Unmarshal([]byte(cmd), c)
	if err != nil {
		Logger.Errorf("failed to unmarshal command, err: %v", err)
	}
	return c
}

func New(caller string, args []string) *Command {
	c := &Command{}
	c.Caller = caller
	switch c.Caller {
	case CommandSleep:
		if len(args) < 1 {
			Logger.Error("command format error")
			return nil
		}
		_, err := time.ParseDuration(args[0])
		if err != nil {
			Logger.Errorf("failed to parse duration, err: %v", err)
			return nil
		}
		c.Args = []interface{}{args[0]}
	case CommandExit:
		if len(args) < 1 {
			Logger.Error("command format error")
			return nil
		}
		_, err := strconv.Atoi(args[0])
		if err != nil {
			Logger.Errorf("failed to parse code, err: %v", err)
			return nil
		}
		c.Args = []interface{}{args[0]}
	}
	return c
}
