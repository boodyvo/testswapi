package main

import (
	"github.com/boodyvo/testswapi/fetcher/redis"
	"github.com/boodyvo/testswapi/repository/postgres"
	"github.com/boodyvo/testswapi/server"
	"github.com/sirupsen/logrus"
)

const (
	host = "0.0.0.0"
	port = "5678"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	loggerEntry := logrus.NewEntry(logger)

	db, err := postgres.New("postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		loggerEntry.WithError(err).Errorf("cannot create db")

		return
	}

	fetcher, err := redis.New("redis:6379")
	if err != nil {
		loggerEntry.WithError(err).Errorf("cannot create fetcher")

		return
	}

	serv, err := server.New(host, port, db, fetcher, loggerEntry)
	if err != nil {
		loggerEntry.WithError(err).Errorf("cannot create server")

		return
	}

	loggerEntry.Infof("starting server")
	err = serv.Start()
	loggerEntry.Infof("finished server with err: %s", err)
}
