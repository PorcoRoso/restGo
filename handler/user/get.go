package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/porcorosso/restGo/handler"
	"github.com/porcorosso/restGo/model"
	"github.com/porcorosso/restGo/pkg/errno"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	userName := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(userName)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
