package models

type Task struct {
	Id int `json:"id"`
	Version int `json:"version"`
	Title string `json:"title"`
	Description string `json:"description"`
	StatusID string `json:"status_id"`
	UserID string `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type TaskComment struct {
	Id int `json:"id"`
	Version int `json:"version"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
