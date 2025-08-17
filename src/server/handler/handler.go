package handler

// this is notification server : handles incoming requests for notifications
// create, update, delete, get notifications

import (
	"net/http"
	"notification-service/src/server/handler/model"
	"notification-service/src/server/service"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service service.ServiceInterface
}

func NewNotificationHandler(service service.ServiceInterface) *NotificationHandler {
	return &NotificationHandler{service: service}
}

func (nh *NotificationHandler) AddNotificationRoutes(router *gin.Engine) {
	router.POST("/notifications", nh.createNotification)
	// router.DELETE("/notifications/:id", nh.deleteNotification)
	router.GET("/notifications/:id", nh.getNotificationByID)
	router.PUT("/notifications/:id", nh.updateNotification)
}

func (nh *NotificationHandler) createNotification(c *gin.Context) {
	var req model.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := nh.service.CreateNotification(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// func (nh *NotificationHandler) deleteNotification(c *gin.Context) {
// 	id := c.Param("id")
// 	if err := nh.service.DeleteNotification(c.Request.Context(), &model.DeleteNotificationRequest{ID: id}); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.Status(http.StatusNoContent)
// }

func (nh *NotificationHandler) getNotificationByID(c *gin.Context) {
	id := c.Param("id")
	resp, err := nh.service.GetNotificationByID(c.Request.Context(), &model.GetNotificationRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (nh *NotificationHandler) updateNotification(c *gin.Context) {
	var req model.UpdateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	req.Id = id // Set the ID from the URL parameter
	resp, err := nh.service.UpdateNotification(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
