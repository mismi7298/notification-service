package model

type CreateTemplateRequest struct {
	Type    string `json:"type"`
	Channel string `json:"Channel"`
	Content string `json:"content"`
	UserId  string `json:"user_id"` // Optional, if the template is user-specific
}

type UpdateTemplateRequest struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"Channel"`
	Content string `json:"content"`
}

type DeleteTemplateRequest struct {
	ID string `json:"id"`
}

type GetTemplateRequest struct {
	ID      string `json:"id"`
	Type    string `json:"type,omitempty"`
	Channel string `json:"Channel,omitempty"`
}
