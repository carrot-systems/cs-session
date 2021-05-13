package rest

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine, routesHandler RoutesHandler) {
	group := r.Group("/session")
	group.POST("/", routesHandler.CreateSessionHandler)
	group.DELETE("/:handle", routesHandler.RemoveSessionHandler)
}
