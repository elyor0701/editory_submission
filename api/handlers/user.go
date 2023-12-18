package handlers

import (
	"editory_submission/api/http"
	"editory_submission/api/models"
	"editory_submission/genproto/auth_service"
	"editory_submission/pkg/helper"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/util"
)

// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body auth_service.User true "CreateUserRequestBody"
// @Success 201 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	var user auth_service.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().CreateUser(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	// @TODO create function to work with queries

	//reqURL := c.Request.URL
	//
	//queryParams, _ := url.ParseQuery(reqURL.RawQuery)
	//
	//queryParams.Add("user-id", resp.GetId())
	//
	//reqURL.RawQuery = queryParams.Encode()
	//
	//c.Request.URL = reqURL
	//
	//err = h.SendVerificationMessageShared(c)
	//if err != nil {
	//	h.log.Error("send verification message is failed", logger.Any("error", err.Error()))
	//}

	h.handleResponse(c, http.Created, resp)
}

// GetUserList godoc
// @ID get_user_list
// @Router /user [GET]
// @Summary Get User List
// @Description  Get User List
// @Tags User
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=auth_service.GetUserListRes} "GetUserListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserList(c *gin.Context) {

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.UserService().GetUserList(
		c.Request.Context(),
		&auth_service.GetUserListReq{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.DefaultQuery("search", ""),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetUserByID godoc
// @ID get_user_by_id
// @Router /user/{user-id} [GET]
// @Summary Get User By ID
// @Description Get User By ID
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user-id"
// @Success 200 {object} http.Response{data=auth_service.User} "UserBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserByID(c *gin.Context) {
	userID := c.Param("user-id")

	if !util.IsValidUUID(userID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.services.UserService().GetUser(
		c.Request.Context(),
		&auth_service.GetUserReq{
			Id: userID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateUser godoc
// @ID update_user
// @Router /user [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param user body auth_service.User true "UpdateUserRequestBody"
// @Success 200 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {
	var user auth_service.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().UpdateUser(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteUser godoc
// @ID delete_user
// @Router /user/{user-id} [DELETE]
// @Summary Delete User
// @Description Get User
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("user-id")

	if !util.IsValidUUID(userID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	_, err := h.services.UserService().DeleteUser(
		c.Request.Context(),
		&auth_service.DeleteUserReq{
			Id: userID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}

// SendVerificationMessage godoc
// @ID send_verification_message
// @Router /send-verification-message [POST]
// @Summary Send Verification Message
// @Description SendVerificationMessage
// @Tags User
// @Accept json
// @Produce json
// @Param data body models.SendVerificationMessageReq true "data"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) SendVerificationMessage(c *gin.Context) {

	err := h.SendVerificationMessageShared(c)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err)
		return
	}

	h.handleResponse(c, http.NoContent, "")
}

func (h *Handler) SendVerificationMessageShared(c *gin.Context) error {
	var (
		data models.SendVerificationMessageReq
	)

	err := c.ShouldBindJSON(&data)
	if err != nil {
		return err
	}

	if !util.IsValidUUID(data.UserId) {
		err := errors.New("invalid user id")
		return err
	}

	user, err := h.services.UserService().GetUser(
		c.Request.Context(),
		&auth_service.GetUserReq{
			Id: data.UserId,
		},
	)
	if err != nil {
		return err
	}

	res, err := h.services.UserService().GenerateEmailVerificationToken(c, &auth_service.GenerateEmailVerificationTokenReq{
		Email:  user.GetEmail(),
		UserId: user.GetId(),
	})
	if err != nil {
		return err
	}

	message := helper.MakeEmailMessage(
		map[string]string{
			"first_name": user.GetFirstName(),
			"verification_link": fmt.Sprintf("%s?email=%s&token=%s",
				data.RedirectUrl,
				res.GetEmail(),
				res.GetToken(),
			),
		},
	)

	err = helper.GoMessageSend(helper.SendMessageByEmail{
		From: helper.EmailInfo{
			Username: h.cfg.EmailUsername,
			Password: h.cfg.EmailPassword,
		},
		To:      res.GetEmail(),
		Subject: "Email Verification",
		Message: message,
	})
	if err != nil {
		return err
	}

	return nil
}

// EmailVerification godoc
// @ID email_verification
// @Router /verification [PUT]
// @Summary Verification
// @Description Verification
// @Tags User
// @Accept json
// @Produce json
// @Param data body models.EmailVerificationReq true "Data"
// @Success 200  {object} http.Response{data=auth_service.EmailVerificationRes} "Status"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) EmailVerification(c *gin.Context) {
	var (
		data models.EmailVerificationReq
	)

	err := c.ShouldBindJSON(&data)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	res, err := h.services.UserService().EmailVerification(
		c.Request.Context(),
		&auth_service.EmailVerificationReq{
			Email: data.Email,
			Token: data.Token,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, res)
}
