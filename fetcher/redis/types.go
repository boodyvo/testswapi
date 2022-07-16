package redis

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/boodyvo/testswapi/models"
)

type response struct {
	Count   int               `json:"count"`
	Results []*moviesResponse `json:"results"`
}

type moviesResponse struct {
	EpisodeID    int      `json:"episode_id"`
	OpeningCrawl string   `json:"opening_crawl"`
	Characters   []string `json:"characters"`
	ReleaseDate  string   `json:"release_date"`
	Title        string   `json:"title"`
}

func (f *moviesResponse) ToMovie() *models.Movie {
	tm, _ := time.Parse(layout, f.ReleaseDate)

	return &models.Movie{
		EpisodeID:    strconv.Itoa(f.EpisodeID),
		Title:        f.Title,
		OpeningCrawl: f.OpeningCrawl,
		ReleaseDate:  tm,
	}
}

type characterResponse struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Height string `json:"height"`
}

func (f *characterResponse) ToCharacter() *models.Character {
	height, err := strconv.Atoi(f.Height)
	if err != nil {
		height = 0
	}

	return &models.Character{
		Name:       f.Name,
		Gender:     f.Gender,
		Height:     height,
		HeightFeet: toFeet(float64(height)),
	}
}

func toFeet(height float64) string {
	heightFeet, heightInches := math.Modf(0.0328 * height)

	return fmt.Sprintf(
		"%.0f %.2f",
		heightFeet,
		heightInches*12,
	)
}
