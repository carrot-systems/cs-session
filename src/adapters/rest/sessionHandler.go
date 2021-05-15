package rest

import (
	"github.com/carrot-systems/cs-session/src/core/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) handleError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(codeForError(err), domain.Status{
		Success: false,
		Message: err.Error(),
	})
}

func (rH RoutesHandler) CreateSessionHandler(c *gin.Context) {
	var creationRequest domain.SessionCreationRequest

	err := c.ShouldBindJSON(&creationRequest)

	if err != nil {
		rH.handleError(c, ErrFormValidation)
		return
	}

	session, err := rH.Usecases.CreateSession(&creationRequest)

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, domain.Status{
		Success: true,
		Message: domain.StatusSessionCreated,
		Data:    session,
	})
}

func (rH RoutesHandler) RemoveSessionHandler(c *gin.Context) {
	err := rH.Usecases.DeleteSession("")

	if err != nil {
		c.AbortWithStatusJSON(codeForError(err), domain.Status{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, domain.Status{
		Success: true,
		Message: domain.StatusSessionDeleted,
	})
}
