package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"

	"github.com/gin-gonic/gin"
	"github.com/gocoder2009/log-for-apiserver"
	"github.com/gocoder2009/log-for-apiserver/lager"
)

// Create
// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 在gopkg.in/go-playground/validator.v9这个包不能满足校验需求时，添加自己的校验逻辑
	if err := r.checkParam(); err != nil {
		SendResponse(c, err, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}

// 实际开发过程中，gopkg.in/go-playground/validator.v9这个包可能不能满足校验需求，这时候可在程序中加入自己的校验逻辑
func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}
