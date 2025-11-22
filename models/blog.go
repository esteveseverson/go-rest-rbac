package models

type Blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}
