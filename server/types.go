package server

type CommentRequest struct {
	Text string `json:"text" binding:"required,max=501"`
}
