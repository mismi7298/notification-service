package bootup

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, handler *Handlers) {
	handler.NotificationHandler.AddNotificationRoutes(router)
	handler.TemplateHandler.AddTemplateRoutes(router)
}
