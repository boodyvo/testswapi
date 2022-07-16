package models

import "time"

type Movie struct {
	EpisodeID     string    `json:"episode_id"`
	Title         string    `json:"title"`
	OpeningCrawl  string    `json:"opening_crawl"`
	ReleaseDate   time.Time `json:"release_date"`
	CommentsCount int       `json:"comments_count"`
}

func NewComment(movieID, text, ip string) *Comment {
	return &Comment{
		MovieID:   movieID,
		Text:      text,
		CreatorIP: ip,
	}
}

type Comment struct {
	ID        int       `json:"id" bun:",pk,autoincrement"`
	MovieID   string    `json:"movie_id" bun:"movie_id,notnull"`
	Text      string    `json:"text" bun:"text,notnull"`
	CreatorIP string    `json:"ip" bun:"ip,notnull"`
	CreatedAt time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
}

type CommentsCount struct {
	MovieID       string `json:"movie_id"`
	CommentsCount int    `json:"comments_count"`
}

type Character struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Height     int    `json:"height"`
	HeightFeet string `json:"height_feet"`
}
