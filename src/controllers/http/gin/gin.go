package gin

import (
	httpControllers "application/controllers/http"
	"application/exceptions"
	"application/services"
	"application/utils/dto"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpControllerLayer struct {
	serviceLayer services.IServiceLayer
}

func NewControllerLayer(serviceLayer services.IServiceLayer) httpControllers.IHttpControllerLayer {
	return &httpControllerLayer{serviceLayer}
}

func (httpControllerLayer *httpControllerLayer) GetHttpHandler() http.Handler {
	httpHandler := gin.Default()

	httpHandler.POST("/users/", httpControllerLayer.UserCreatePost)

	return httpHandler
}

func (httpControllerLayer *httpControllerLayer) UserCreatePost(c *gin.Context) {

	var userToCreate dto.UserCreateRequest

	if err := c.ShouldBindJSON(&userToCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userCreated, err := httpControllerLayer.serviceLayer.CreateUser(userToCreate)
	if err == exceptions.ErrUserEmailExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "emailExists"})
		return
	} else if err != nil {
		fmt.Println("InternalServerError:", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusCreated, userCreated)
	return

}
