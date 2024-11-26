package models

type Lesson struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Language    string `json:"language"`
	Difficulty  string `json:"difficulty"`
}
