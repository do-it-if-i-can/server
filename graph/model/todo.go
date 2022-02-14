package model

type NewTodo struct {
	Text     string   `json:"text"`
	UserID   string   `json:"userId"`
	Category Category `json:"category"`
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}
