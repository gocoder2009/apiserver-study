package user

import (
	"strconv"

	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete godoc
//
//	@Security		JWT
//	@Summary		Delete a user by the user identifier
//	@Description	Delete user by ID
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint64				true	"The user's database id index num"
//	@Success		200	{object}	handler.Response	"{"code":0,"message":"OK","data":null}"
//	@Router			/v1/user/{id} [delete]
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
