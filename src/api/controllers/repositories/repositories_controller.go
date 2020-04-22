package repositories

import (
	"golang-microservices/src/api/domain/repositories"
	"golang-microservices/src/api/services"
	"golang-microservices/src/api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Get_Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Get_Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
