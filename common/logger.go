package common

import (
	log "gopkg.pigeonligh.com/easygo/elog"
)

var Logger *log.Logger

func init() {
	log.Default()
	Logger = log.With(map[string]string{
		"source": "",
	})
}
