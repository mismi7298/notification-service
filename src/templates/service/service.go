package service

import (
	"context"
	"notification-service/src/templates/handler/model"
	"notification-service/src/templates/repository"
)

type TemplateServiceInterface interface {
	CreateTemplate(ctx context.Context, template *model.CreateTemplateRequest) (*model.TemplateResponse, error)
	UpdateTemplate(ctx context.Context, template *model.UpdateTemplateRequest) (*model.TemplateResponse, error)
	GetTemplateByID(ctx context.Context, template *model.GetTemplateRequest) (*model.TemplateResponse, error)
	DeleteTemplate(ctx context.Context, template *model.DeleteTemplateRequest) error
	GetTemplateByTypeAndChannel(ctx context.Context, template *model.GetTemplateRequest) (*model.TemplateResponse, error)
}

type TemplateService struct {
	repository repository.TemplateRepositoryInterface
}

func NewTemplateService(repo repository.TemplateRepositoryInterface) TemplateServiceInterface {
	return &TemplateService{
		repository: repo,
	}
}

func (s *TemplateService) CreateTemplate(ctx context.Context, template *model.CreateTemplateRequest) (*model.TemplateResponse, error) {
	// Call the repository layer to create the template
	createdTemplate, err := s.repository.CreateTemplate(ctx, template)
	if err != nil {
		return nil, err
	}
	return createdTemplate, nil
}

func (s *TemplateService) UpdateTemplate(ctx context.Context, template *model.UpdateTemplateRequest) (*model.TemplateResponse, error) {
	// Call the repository layer to update the template
	updatedTemplate, err := s.repository.UpdateTemplate(ctx, template)
	if err != nil {
		return nil, err
	}
	return updatedTemplate, nil
}

func (s *TemplateService) GetTemplateByID(ctx context.Context, template *model.GetTemplateRequest) (*model.TemplateResponse, error) {
	// Call the repository layer to get the template by ID
	getTemplate, err := s.repository.GetTemplateByID(ctx, template.ID)
	if err != nil {
		return nil, err
	}
	return getTemplate, nil
}

func (s *TemplateService) GetTemplateByTypeAndChannel(ctx context.Context, template *model.GetTemplateRequest) (*model.TemplateResponse, error) {
	// Call the repository layer to get the template by type and Channel
	getTemplate, err := s.repository.GetTemplateByTypeAndChannel(ctx, template)
	if err != nil {
		return nil, err
	}
	return getTemplate, nil
}

func (s *TemplateService) DeleteTemplate(ctx context.Context, template *model.DeleteTemplateRequest) error {
	// Call the repository layer to delete the template
	err := s.repository.DeleteTemplate(ctx, template)
	if err != nil {
		return err
	}
	return nil
}
