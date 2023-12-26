package handlers

import (
	"context"
	"editory_submission/api/http"
	"editory_submission/api/models"
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"editory_submission/genproto/notification_service"
	"editory_submission/pkg/logger"
	"editory_submission/pkg/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CreateJournalUser godoc
// @ID create_journal_user
// @Router /journal/{journal-id}/user [POST]
// @Summary Create User
// @Description Create User
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param user body models.CreateJournalUserReq true "CreateUserRequestBody"
// @Success 201 {object} http.Response{data=models.CreateJournalUserRes} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateJournalUser(c *gin.Context) {
	var user models.CreateJournalUserReq

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		err = errors.New("not valid journal id")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	userPb, err := h.services.UserService().GetUser(
		c.Request.Context(),
		&auth_service.GetUserReq{
			Email: user.Email,
		},
	)
	if err != nil {
		if util.IsErrNoRows(err) {
			userPb, err = h.services.UserService().CreateUser(
				c.Request.Context(),
				&auth_service.User{
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Email:     user.Email,
					Password:  config.DEFAULT_PASSWORD,
				},
			)

			if err != nil {
				h.handleResponse(c, http.GRPCError, err.Error())
				return
			}
		} else {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
	}

	validRoleTypes := map[string]bool{
		config.EDITOR:   true,
		config.REVIEWER: true,
	}

	if _, ok := validRoleTypes[user.RoleType]; ok {
		role := auth_service.Role{
			RoleType:  user.RoleType,
			JournalId: journalId,
			UserId:    userPb.GetId(),
		}

		_, err = h.services.RoleService().CreateRole(
			c.Request.Context(),
			&role,
		)
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
	}

	res := models.CreateEditorRes{
		Id:        userPb.GetId(),
		FirstName: userPb.GetFirstName(),
		LastName:  userPb.GetLastName(),
		Email:     userPb.GetEmail(),
		Password:  userPb.GetPassword(),
	}

	h.handleResponse(c, http.Created, res)
}

// GetJournalUserList godoc
// @ID get_journal_user_list
// @Router /journal/{journal-id}/user [GET]
// @Summary Get User List
// @Description  Get User List
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=auth_service.GetUserListRes} "GetUserListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalUserList(c *gin.Context) {

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		err := errors.New("not valid journal id")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

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

	resp, err := h.services.UserService().GetUserListByRole(
		c.Request.Context(),
		&auth_service.GetUserListByRoleReq{
			Limit:     int32(limit),
			Offset:    int32(offset),
			Search:    c.DefaultQuery("search", ""),
			JournalId: journalId,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetJournalUserByID godoc
// @ID get_journal_user_by_id
// @Router /journal/{journal-id}/user/{user-id} [GET]
// @Summary Get User By ID
// @Description Get User By ID
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal id"
// @Param user-id path string true "user-id"
// @Success 200 {object} http.Response{data=auth_service.User} "UserBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalUserByID(c *gin.Context) {
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

// UpdateJournalUser godoc
// @ID update_journal_user
// @Router /journal/{journal-id}/user [PUT]
// @Summary Update User
// @Description Update User
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal id"
// @Param user body auth_service.UpdateUserReq true "UpdateUserRequestBody"
// @Success 200 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateJournalUser(c *gin.Context) {
	var user auth_service.UpdateUserReq

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

// DeleteJournalUser godoc
// @ID delete_journal_user
// @Router /journal/{journal-id}/user/{user-id} [DELETE]
// @Summary Delete User
// @Description Get User
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal id"
// @Param user-id path string true "user-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteJournalUser(c *gin.Context) {
	userID := c.Param("user-id")
	if !util.IsValidUUID(userID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	userRoles, err := h.services.RoleService().GetRoleList(
		c.Request.Context(),
		&auth_service.GetRoleListReq{
			UserId:    userID,
			JournalId: journalId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	fmt.Println("userRoles---->", userRoles)

	for _, v := range userRoles.GetRoles() {
		_, err := h.services.RoleService().DeleteRole(
			c.Request.Context(),
			&auth_service.DeleteRoleReq{
				Id: v.GetId(),
			},
		)
		if err != nil {
			h.log.Error("cant delete the user role", logger.Any("err", err.Error()))
			continue
		}
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

	//err = h.SendVerificationMessage(
	//	c.Request.Context(),
	//	&models.SendVerificationMessageReq{
	//		UserId:      data.UserId,
	//		RedirectUrl: data.RedirectUrl,
	//	},
	//)
	//if err != nil {
	//	h.handleResponse(c, http.InternalServerError, err)
	//	return
	//}

	_, err = h.services.NotificationService().GenerateMailMessage(c.Request.Context(), &notification_service.GenerateMailMessageReq{
		UserId:       data.UserId,
		RedirectLink: data.RedirectUrl,
		Type:         config.REGISTRATION,
	})
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

// GetProfileByID godoc
// @ID get_profile_by_id
// @Router /profile/{profile-id} [GET]
// @Summary Get User By ID
// @Description Get User By ID
// @Tags User
// @Accept json
// @Produce json
// @Param profile-id path string true "profile-id"
// @Success 200 {object} http.Response{data=models.GetAdminUserRes} "UserBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetProfileByID(c *gin.Context) {
	userID := c.Param("profile-id")

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

// UpdateProfile godoc
// @ID update_profile
// @Router /profile [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UpdateAdminUserReq true "UpdateUserRequestBody"
// @Success 200 {object} http.Response{data=models.UpdateAdminUserRes} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateProfile(c *gin.Context) {
	var user models.UpdateAdminUserReq

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().UpdateUser(
		c.Request.Context(),
		&auth_service.UpdateUserReq{
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

	res.MessageStatus = true
	_, err = h.services.NotificationService().GenerateMailMessage(c.Request.Context(), &notification_service.GenerateMailMessageReq{
		UserId:       resp.GetId(),
		RedirectLink: user.RedirectUrl,
		Type:         config.REGISTRATION,
	})
	if err != nil {
		h.log.Error("cant send verification message", logger.String("err", err.Error()))
		res.MessageStatus = false
	}

	//err = h.SendVerificationMessage(
	//	c.Request.Context(),
	//	&models.SendVerificationMessageReq{
	//		UserId:      resp.GetId(),
	//		RedirectUrl: user.RedirectUrl,
	//	},
	//)
	//if err != nil {
	//	h.log.Error("cant send verification message", logger.String("err", err.Error()))
	//	res.MessageStatus = false
	//}
	res.Email = resp.GetEmail()
	res.UserId = resp.GetId()

	h.handleResponse(c, http.Created, res)
}

func (h *Handler) SendVerificationMessage(ctx context.Context, req *models.SendVerificationMessageReq) error {
	//if !util.IsValidUUID(req.UserId) {
	//	err := errors.New("invalid user id")
	//	return err
	//}
	//
	//user, err := h.services.UserService().GetUser(
	//	ctx,
	//	&auth_service.GetUserReq{
	//		Id: req.UserId,
	//	},
	//)
	//if err != nil {
	//	return err
	//}
	//
	//res, err := h.services.UserService().GenerateEmailVerificationToken(ctx, &auth_service.GenerateEmailVerificationTokenReq{
	//	Email:  user.GetEmail(),
	//	UserId: user.GetId(),
	//})
	//if err != nil {
	//	return err
	//}
	//
	//message := helper.MakeEmailMessage(
	//	map[string]string{
	//		"first_name": user.GetFirstName(),
	//		"verification_link": fmt.Sprintf("%s?email=%s&token=%s",
	//			req.RedirectUrl,
	//			res.GetEmail(),
	//			res.GetToken(),
	//		),
	//		"user_id": req.UserId,
	//	},
	//)
	//
	//err = helper.GoMessageSend(helper.SendMessageByEmail{
	//	From: helper.EmailInfo{
	//		Username: h.cfg.EmailUsername,
	//		Password: h.cfg.EmailPassword,
	//	},
	//	To:      res.GetEmail(),
	//	Subject: "Email Verification",
	//	Message: message,
	//})
	//if err != nil {
	//	return err
	//}

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
		userPb auth_service.UpdateUserReq
		res    models.RegisterDetailRes
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	userPb = auth_service.UpdateUserReq{
		Id:         user.Id,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Phone:      user.Phone,
		ExtraPhone: user.ExtraPhone,
		CountryId:  user.CountryId,
		CityId:     user.CityId,
		ProfSphere: user.ProfSphere,
		Degree:     user.Degree,
		Address:    user.Address,
		PostCode:   user.PostCode,
		Gender:     user.Gender,
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
