package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service"

	"github.com/gin-gonic/gin"
	"github.com/gocoder2009/log-for-apiserver"
)

// List godoc
//
//	@Security		JWT
//	@Summary		List the users in the database
//	@Description	get users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string						false																									"username"
//	@Param			offset		query		int							true																									"offset"
//	@Param			limit		query		int							true																									"limit"
//	@Success		200			{object}	user.SwaggerListResponse	"{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user	'admin'	get	random	string	'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28	00:25:33","updatedAt":"2018-05-28	00:25:33"}]}}"
//	@Router			/v1/user [get]
func List(c *gin.Context) {
	log.Info("List function called.")
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
