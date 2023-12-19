package handlers

import (
	"context"
	"editory_submission/api/http"
	"editory_submission/api/models"
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"editory_submission/pkg/helper"
	"editory_submission/pkg/logger"
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

// ResendVerificationMessage godoc
// @ID send_verification_message
// @Router /register/resend [POST]
// @Summary Resend Verification Message
// @Description ReSendVerificationMessage
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body models.SendVerificationMessageReq true "data"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) ResendVerificationMessage(c *gin.Context) {
	var (
		data models.SendVerificationMessageReq
	)

	err := c.ShouldBindJSON(&data)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = h.SendVerificationMessage(
		c.Request.Context(),
		&models.SendVerificationMessageReq{
			UserId:      data.UserId,
			RedirectUrl: data.RedirectUrl,
		},
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err)
		return
	}

	h.handleResponse(c, http.NoContent, "")
}

// EmailVerification godoc
// @ID email_verification
// @Router /register/verify [POST]
// @Summary Verification
// @Description Verification
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body models.EmailVerificationReq true "Data"
// @Success 200  {object} http.Response{data=models.EmailVerificationRes} "Status"
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

	resp, err := h.services.UserService().EmailVerification(
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

	res := models.EmailVerificationRes{
		Status: resp.GetStatus(),
		UserId: resp.GetUserId(),
	}

	h.handleResponse(c, http.OK, res)
}

// GetAdminUserByID godoc
// @ID get_admin_user_by_id
// @Router /admin/user/{user-id} [GET]
// @Summary Get User By ID
// @Description Get User By ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param user-id path string true "user-id"
// @Success 200 {object} http.Response{data=models.GetAdminUserRes} "UserBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminUserByID(c *gin.Context) {
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

	user := models.GetAdminUserRes{
		Gender:    resp.GetGender(),
		FirstName: resp.GetFirstName(),
		LastName:  resp.GetLastName(),
		Email:     resp.GetEmail(),
	}

	h.handleResponse(c, http.OK, user)
}

// UpdateAdminUser godoc
// @ID update_admin_user
// @Router /admin/user [PUT]
// @Summary Update User
// @Description Update User
// @Tags Admin
// @Accept json
// @Produce json
// @Param user body models.UpdateAdminUserReq true "UpdateUserRequestBody"
// @Success 200 {object} http.Response{data=models.UpdateAdminUserRes} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAdminUser(c *gin.Context) {
	var user models.UpdateAdminUserReq

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().UpdateUser(
		c.Request.Context(),
		&auth_service.User{
			Gender:    user.Gender,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.LastName,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	res := models.UpdateAdminUserRes{
		Gender:    resp.GetGender(),
		FirstName: resp.GetFirstName(),
		LastName:  resp.GetLastName(),
		Email:     resp.GetEmail(),
		Password:  resp.GetPassword(),
	}

	h.handleResponse(c, http.OK, res)
}

// RegistrationEmail godoc
// @ID registration_email
// @Router /register/email [POST]
// @Summary Create User
// @Description Create User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.RegistrationEmailReq true "Registration body"
// @Success 201 {object} http.Response{data=models.RegistrationEmailRes} "Email"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RegistrationEmail(c *gin.Context) {
	var (
		user models.RegistrationEmailReq
		res  models.RegistrationEmailRes
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if user.Password != user.ConfirmPassword {
		err = errors.New("password and confirm password are not matched")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().CreateUser(
		c.Request.Context(),
		&auth_service.User{
			Email:       user.Email,
			Password:    user.Password,
			IsCompleted: false,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	err = h.SendVerificationMessage(
		c.Request.Context(),
		&models.SendVerificationMessageReq{
			UserId:      resp.GetId(),
			RedirectUrl: user.RedirectUrl,
		},
	)
	res.MessageStatus = true
	if err != nil {
		h.log.Error("cant send verification message", logger.String("err", err.Error()))
		res.MessageStatus = false
	}
	res.Email = resp.GetEmail()
	res.UserId = resp.GetId()

	h.handleResponse(c, http.Created, res)
}

func (h *Handler) SendVerificationMessage(ctx context.Context, req *models.SendVerificationMessageReq) error {
	if !util.IsValidUUID(req.UserId) {
		err := errors.New("invalid user id")
		return err
	}

	user, err := h.services.UserService().GetUser(
		ctx,
		&auth_service.GetUserReq{
			Id: req.UserId,
		},
	)
	if err != nil {
		return err
	}

	res, err := h.services.UserService().GenerateEmailVerificationToken(ctx, &auth_service.GenerateEmailVerificationTokenReq{
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
				req.RedirectUrl,
				res.GetEmail(),
				res.GetToken(),
			),
			"user_id": req.UserId,
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

// RegisterDetail godoc
// @ID register_details
// @Router /register/details [POST]
// @Summary Update User
// @Description Update User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.RegisterDetailReq true "Register User Details"
// @Success 200 {object} http.Response{data=models.RegisterDetailRes} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RegisterDetail(c *gin.Context) {
	var (
		user   models.RegisterDetailReq
		userPb auth_service.User
		res    models.RegisterDetailRes
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	userPb = auth_service.User{
		Id:          user.Id,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Phone:       user.Phone,
		ExtraPhone:  user.ExtraPhone,
		CountryId:   user.CountryId,
		CityId:      user.CityId,
		ProfSphere:  user.ProfSphere,
		Degree:      user.Degree,
		Address:     user.Address,
		PostCode:    user.PostCode,
		Gender:      user.Gender,
		IsCompleted: true,
	}

	resp, err := h.services.UserService().UpdateUser(
		c.Request.Context(),
		&userPb,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	if user.IsReviewer {
		_, err := h.services.RoleService().CreateRole(
			c.Request.Context(),
			&auth_service.Role{
				RoleType: config.REVIEWER,
				UserId:   user.Id,
			},
		)

		if err != nil {
			h.log.Error("cant create role for user", logger.String("err", err.Error()))
		}
	}

	res = models.RegisterDetailRes{
		Id:         resp.Id,
		Username:   resp.Username,
		FirstName:  resp.FirstName,
		LastName:   resp.LastName,
		Phone:      resp.Phone,
		ExtraPhone: resp.ExtraPhone,
		CountryId:  resp.CountryId,
		CityId:     resp.CityId,
		ProfSphere: resp.ProfSphere,
		Degree:     resp.Degree,
		Address:    resp.Address,
		PostCode:   resp.PostCode,
		Gender:     resp.Gender,
		IsReviewer: user.IsReviewer,
	}

	h.handleResponse(c, http.OK, res)
}
