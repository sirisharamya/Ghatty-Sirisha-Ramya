package main

type Blog struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	Timestamp string `json:"timestamp"`
}
