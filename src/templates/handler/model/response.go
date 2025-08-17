package model

type TemplateResponse struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Channel   string `json:"Channel"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
