package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	atDomain "github.com/james077/bookstore_oauth-api/src/domain/access_token"
	"github.com/james077/bookstore_oauth-api/src/services/access_token"
	"github.com/james077/bookstore_utils-go/rest_errors"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var request atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(&request)
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		fmt.Println(err)
		return
	}

	accessToken, err := handler.service.Create(request)
	fmt.Println(&request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	
	c.JSON(http.StatusCreated, accessToken)
}
