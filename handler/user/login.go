package user

import (
	. "github.com/porcorosso/restGo/handler"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/porcorosso/restGo/model"
	"github.com/porcorosso/restGo/pkg/auth"
	"github.com/porcorosso/restGo/pkg/errno"
	"github.com/porcorosso/restGo/pkg/token"
)

// @Summary Login generates the authentication token
// @Produce  json
// @Param token header string true "token"
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /login [post]
func Login(c *gin.Context) {
	// Login login by username & password, verify name exist, auth password, if all pass, return JWT
	var u model.UserModel

	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Debugf("username is %s, password is %s", u.Username, u.Password)

	// Get the user infomation by the login username.
	d, err := model.GetUser(u.Username)
	if err != nil {
		log.Errorf(err, "query user by username : %s", u.Username)
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the jwt
	t, err := token.Sign(c,
		token.Context{
			ID:       d.Id,
			Username: d.Username,
		},
		"",
	)

	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
