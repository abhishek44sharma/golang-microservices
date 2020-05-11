package oauth

import (
	"golang-microservices/oauth_api/src/api/domain/oauth"
	"golang-microservices/oauth_api/src/api/services"
	"golang-microservices/src/api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Get_Status(), apiErr)
		return
	}
	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Get_Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)

}

func GetAccessToken(c *gin.Context) {
	tokenId := c.Param("token_id")

	token, err := services.OauthService.GetAccessToken(tokenId)
	if err != nil {
		c.JSON(err.Get_Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}
