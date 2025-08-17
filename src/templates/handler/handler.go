package handler

import (
	"net/http"
	"notification-service/src/templates/handler/model"
	"notification-service/src/templates/service"

	"github.com/gin-gonic/gin"
)

// template handler supports management of templates via APIs
// user can create, update, delete, get templates

type TemplateHandler struct {
	service service.TemplateServiceInterface
}

func NewTemplateHandler(service service.TemplateServiceInterface) *TemplateHandler {
	return &TemplateHandler{service: service}
}

func (th *TemplateHandler) AddTemplateRoutes(router *gin.Engine) {
	router.POST("/templates", th.createTemplate)
	router.DELETE("/templates/:id", th.DeleteTemplate)
	router.GET("/templates/:id", th.GetTemplateByID)
	router.PUT("/templates/:id", th.UpdateTemplate)
}

// endpoint for creating a new template using
func (h *TemplateHandler) createTemplate(c *gin.Context) {
	var req model.CreateTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateTemplate(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *TemplateHandler) DeleteTemplate(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteTemplate(c.Request.Context(), &model.DeleteTemplateRequest{ID: id}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *TemplateHandler) GetTemplateByID(c *gin.Context) {
	id := c.Param("id")
	template, err := h.service.GetTemplateByID(c.Request.Context(), &model.GetTemplateRequest{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, template)
}

func (h *TemplateHandler) UpdateTemplate(c *gin.Context) {
	var req model.UpdateTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	req.ID = id // Set the ID from the URL parameter
	resp, err := h.service.UpdateTemplate(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
