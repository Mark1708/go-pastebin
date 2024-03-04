package config

import (
	"log"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server ServerConf
	DB     DatabaseConf
}

func New() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}
	return &c
}
