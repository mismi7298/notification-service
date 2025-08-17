package repository

import (
	"context"
	"database/sql"
	"notification-service/src/templates/handler/model"
)

// template Repository stores templates in a database

type TemplateRepository struct {
	db *sql.DB
}

type TemplateRepositoryInterface interface {
	CreateTemplate(ctx context.Context, template *model.CreateTemplateRequest) (*model.TemplateResponse, error)
	UpdateTemplate(ctx context.Context, template *model.UpdateTemplateRequest) (*model.TemplateResponse, error)
	DeleteTemplate(ctx context.Context, template *model.DeleteTemplateRequest) error
	GetTemplateByID(ctx context.Context, id string) (*model.TemplateResponse, error)
	GetTemplateByTypeAndChannel(ctx context.Context, template *model.GetTemplateRequest) (*model.TemplateResponse, error)
}

func NewTemplateRepository(db *sql.DB) TemplateRepositoryInterface {
	return &TemplateRepository{db: db}
}

func (r *TemplateRepository) CreateTemplate(ctx context.Context, template *model.CreateTemplateRequest) (*model.TemplateResponse, error) {

	// create
	query := "INSERT INTO templates (type, Channel, content, userId) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at"
	row := r.db.QueryRowContext(ctx, query, template.Type, template.Channel, template.Content, template.UserId)

	var createdTemplate model.TemplateResponse
	if err := row.Scan(&createdTemplate.ID, &createdTemplate.CreatedAt, &createdTemplate.UpdatedAt); err != nil {
		return nil, err
	}
	return &createdTemplate, nil
}

func (r *TemplateRepository) UpdateTemplate(ctx context.Context, template *model.UpdateTemplateRequest) (*model.TemplateResponse, error) {
	// update
	query := "UPDATE templates SET content = $1 WHERE id = $2 RETURNING id, type, Channel, content, created_at, updated_at"
	row := r.db.QueryRowContext(ctx, query, template.Content, template.ID)

	var updatedTemplate model.TemplateResponse
	if err := row.Scan(&updatedTemplate.ID, &updatedTemplate.Type, &updatedTemplate.Channel, &updatedTemplate.Content, &updatedTemplate.CreatedAt, &updatedTemplate.UpdatedAt); err != nil {
		return nil, err
	}
	return &updatedTemplate, nil
}

func (r *TemplateRepository) DeleteTemplate(ctx context.Context, template *model.DeleteTemplateRequest) error {
	// delete
	query := "DELETE FROM templates WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, template.ID)
	return err
}

func (r *TemplateRepository) GetTemplateByID(ctx context.Context, id string) (*model.TemplateResponse, error) {
	// get
	query := "SELECT * FROM templates WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	var retrievedTemplate model.TemplateResponse
	if err := row.Scan(&retrievedTemplate.ID, &retrievedTemplate.Type, &retrievedTemplate.Channel, &retrievedTemplate.Content, &retrievedTemplate.CreatedAt, &retrievedTemplate.UpdatedAt); err != nil {
		return nil, err
	}
	return &retrievedTemplate, nil
}

func (r *TemplateRepository) GetTemplateByTypeAndChannel(ctx context.Context, template *model.GetTemplateRequest) (*model.TemplateResponse, error) {
	// get
	query := "SELECT * FROM templates WHERE type = $1 AND Channel = $2"
	row := r.db.QueryRowContext(ctx, query, template.Type, template.Channel)

	var retrievedTemplate model.TemplateResponse
	if err := row.Scan(&retrievedTemplate.ID, &retrievedTemplate.Type, &retrievedTemplate.Channel, &retrievedTemplate.Content, &retrievedTemplate.CreatedAt, &retrievedTemplate.UpdatedAt); err != nil {
		return nil, err
	}
	return &retrievedTemplate, nil
}
