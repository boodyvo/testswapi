package server

import (
	"fmt"
	"github.com/boodyvo/testswapi/fetcher"
	"github.com/boodyvo/testswapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) HealthHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/plain")
	ctx.String(
		http.StatusOK,
		"OK",
	)
}

func (s *Service) ListMoviesHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	s.logger.Info("got list movies request")

	movies, err := s.fetcher.ListMovies(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

		return
	}
	movieIDs := make([]string, 0, len(movies))
	for _, movie := range movies {
		movieIDs = append(movieIDs, movie.EpisodeID)
	}

	counts, err := s.repository.ListCommentsCount(ctx, movieIDs)
	for i, count := range counts {
		movies[i].CommentsCount = count
	}

	s.logger.WithField("movies", movies).Info("sending all movies")

	ctx.JSON(http.StatusOK, movies)
}

func (s *Service) CreateCommentHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	s.logger.Info("got create comment request")

	var request *CommentRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, err)

		return
	}

	movieID := ctx.Param("id")
	ip := ctx.ClientIP()
	comment := models.NewComment(movieID, request.Text, ip)

	s.logger.WithField("comment", comment).Info("will save comment")
	err := s.repository.CreateComment(ctx, comment)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (s *Service) ListCommentsHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	s.logger.Info("got list comments request")

	movieID := ctx.Param("id")
	comments, err := s.repository.ListComments(ctx, movieID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

		return
	}
	s.logger.
		WithField("movie_id", movieID).
		WithField("comments", comments).
		Info("sending comments")

	ctx.JSON(http.StatusOK, comments)
}

func (s *Service) ListCharactersHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	s.logger.Info("got list comments request")

	movieID := ctx.Param("id")
	request := fetcher.CharacterRequest{
		MovieID: movieID,
	}
	gender := ctx.Query("gender")
	if gender != "" {
		request.Filter = &fetcher.CharacterFilter{
			Gender: gender,
		}
	}
	sortName := ctx.Query("sort")
	order := ctx.Query("order")
	if sortName != "" {
		isAsc := true
		if order == "desc" {
			isAsc = false
		}

		request.Sort = &fetcher.CharacterSort{
			SortName: sortName,
			IsAsc:    isAsc,
		}
	}
	characters, count, height, err := s.fetcher.ListCharacters(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

		return
	}
	s.logger.WithField("characters", characters).Info("sending characters")

	ctx.JSON(http.StatusOK, gin.H{
		"height":     height,
		"count":      count,
		"characters": characters,
	})
}
