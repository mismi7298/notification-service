package server

import (
	"notification-service/server/bootup"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db, err := bootup.InitDB()
	if err != nil {
		panic(err)
	}
	handlers := bootup.InitHandlers(db)
	bootup.InitRouter(router, handlers)
	router.Run(":8080")
}
