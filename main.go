package main

import (
	"github.com/Marco98/ytpodproxy/pkg/server"
	log "github.com/sirupsen/logrus"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	log.WithFields(log.Fields{
		"version": version,
		"commit":  commit,
		"date":    date,
	}).Info("starting ytpodproxy")
	if err := server.Run(); err != nil {
		log.WithError(err).Fatal("fatal exception")
	}
}
