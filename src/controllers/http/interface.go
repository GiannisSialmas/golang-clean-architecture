package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHttpControllerLayer interface {
	GetHttpHandler() http.Handler
	IHttpControllerGin //In a better world, this should be a generic
}

type IHttpControllerGin interface {
	UserCreatePost(c *gin.Context)
}
