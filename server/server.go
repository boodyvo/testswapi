package server

import (
	"github.com/boodyvo/testswapi/fetcher"
	"github.com/boodyvo/testswapi/repository"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	defaultTimout = 300 * time.Second

	healthcheckEndpoint = "/health"
	baseApiEndpoint     = "/api/v1"
	moviesEndpoint      = baseApiEndpoint + "/movies"
	commentsEndpoint    = moviesEndpoint + "/:id/comments"
	charactersEndpoint  = moviesEndpoint + "/:id/characters"
)

type Service struct {
	host       string
	port       string
	server     *http.Server
	repository repository.CommentsRepository
	fetcher    fetcher.StarWarsAPI
	router     *gin.Engine
	logger     *logrus.Entry
}

type ServiceOption func(*Service)

func New(
	host string,
	port string,
	repository repository.CommentsRepository,
	fetcher fetcher.StarWarsAPI,
	logger *logrus.Entry,
) (*Service, error) {
	router := gin.Default()

	return newService(router, host, port, repository, fetcher, logger)
}

func newService(
	router *gin.Engine,
	host string,
	port string,
	repository repository.CommentsRepository,
	fetcher fetcher.StarWarsAPI,
	logger *logrus.Entry,
) (*Service, error) {
	s := &Service{
		host:       host,
		port:       port,
		router:     router,
		repository: repository,
		fetcher:    fetcher,
		logger:     logger,
	}

	router.GET(healthcheckEndpoint, s.HealthHandler)
	router.GET(moviesEndpoint, s.ListMoviesHandler)
	router.GET(commentsEndpoint, s.ListCommentsHandler)
	router.POST(commentsEndpoint, s.CreateCommentHandler)
	router.GET(charactersEndpoint, s.ListCharactersHandler)

	router.Use(gin.Recovery())

	return s, nil
}

func (s *Service) Start() error {
	address := net.JoinHostPort(s.host, s.port)

	s.server = &http.Server{
		Addr:         address,
		Handler:      s.router,
		ReadTimeout:  defaultTimout,
		WriteTimeout: defaultTimout,
	}

	s.logger.Infof("listening on %s", address)

	return s.server.ListenAndServe()
}

func (s *Service) Stop() error {
	s.logger.Infof("stopping server listening")

	if s.server != nil {
		return s.server.Close()
	}

	return nil
}
