package user

import (
	"github.com/gin-gonic/gin"
	. "github.com/porcorosso/restGo/handler"
	"github.com/porcorosso/restGo/model"
	"strconv"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, nil)
}
