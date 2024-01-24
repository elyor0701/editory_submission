package handlers

import (
	"editory_submission/api/http"
	"editory_submission/api/models"
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"editory_submission/genproto/content_service"
	"editory_submission/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
)

// GetAdminAuthorList godoc
// @ID get_admin_author_list
// @Router /admin/author [GET]
// @Summary Get Author List
// @Description  Get Author List
// @Tags Admin
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=auth_service.GetUserListRes} "GetEditorListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminAuthorList(c *gin.Context) {
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
			Limit:    int32(limit),
			Offset:   int32(offset),
			Search:   c.DefaultQuery("search", ""),
			RoleType: config.AUTHOR,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// CreateAdminAuthor godoc
// @ID create_admin_author
// @Router /admin/author [POST]
// @Summary Create Author
// @Description Create Author
// @Tags Admin
// @Accept json
// @Produce json
// @Param user body models.CreateAuthorReq true "CreateEditorRequestBody"
// @Success 201 {object} http.Response{data=models.CreateAuthorRes} "Editor data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAdminAuthor(c *gin.Context) {
	var (
		user models.CreateAuthorReq
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
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

	role := auth_service.Role{
		RoleType: config.AUTHOR,
		UserId:   userPb.GetId(),
	}

	_, err = h.services.RoleService().CreateRole(
		c.Request.Context(),
		&role,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
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

// GetAdminAuthorByID godoc
// @ID get_admin_author_by_id
// @Router /admin/author/{author-id} [GET]
// @Summary Get Author By ID
// @Description Get Author By ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param author-id path string true "author-id"
// @Success 200 {object} http.Response{data=auth_service.User} "EditorBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminAuthorByID(c *gin.Context) {
	userID := c.Param("author-id")

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

// UpdateAdminAuthor godoc
// @ID update_admin_author
// @Router /admin/author [PUT]
// @Summary Update Editor
// @Description Update Editor
// @Tags Admin
// @Accept json
// @Produce json
// @Param editor body auth_service.UpdateUserReq true "UpdateUserRequestBody"
// @Success 200 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAdminAuthor(c *gin.Context) {
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

// GetJournalAuthorList godoc
// @ID get_journal_author_list
// @Router /journal/{journal-id}/author [GET]
// @Summary Get Author List
// @Description  Get Author List
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=content_service.GetJournalAuthorListRes} "GetEditorListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalAuthorList(c *gin.Context) {

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		h.handleResponse(c, http.InvalidArgument, errors.New("journal id is not valid"))
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

	resp, err := h.services.ContentService().GetJournalAuthorList(
		c.Request.Context(),
		&content_service.GetJournalAuthorListReq{
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

// CreateJournalAuthor godoc
// @ID create_journal_author
// @Router /journal/{journal-id}/author [POST]
// @Summary Create Author
// @Description Create Author
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param author body content_service.CreateJournalAuthorReq true "CreateJournalAuthorReq"
// @Success 201 {object} http.Response{data=content_service.CreateJournalAuthorRes} "Author data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateJournalAuthor(c *gin.Context) {
	var (
		author content_service.CreateJournalAuthorReq
	)

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		h.handleResponse(c, http.InvalidArgument, errors.New("journal id is not valid"))
		return
	}

	err := c.ShouldBindJSON(&author)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	author.JournalId = journalId

	resp, err := h.services.ContentService().CreateJournalAuthor(
		c.Request.Context(),
		&author,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetJournalAuthorByID godoc
// @ID get_journal_author_by_id
// @Router /journal/{journal-id}/author/{author-id} [GET]
// @Summary Get Author By ID
// @Description Get Author By ID
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param author-id path string true "author-id"
// @Success 200 {object} http.Response{data=content_service.GetJournalAuthorRes} "EditorBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalAuthorByID(c *gin.Context) {

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		h.handleResponse(c, http.InvalidArgument, errors.New("journal id is not valid"))
		return
	}

	authorID := c.Param("author-id")

	if !util.IsValidUUID(authorID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.services.ContentService().GetJournalAuthor(
		c.Request.Context(),
		&content_service.GetJournalAuthorReq{
			Id: authorID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateJournalAuthor godoc
// @ID update_journal_author
// @Router /journal/{journal-id}/author [PUT]
// @Summary Update Editor
// @Description Update Editor
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param author body content_service.UpdateJournalAuthorReq true "UpdateJournalAuthorReq"
// @Success 200 {object} http.Response{data=content_service.UpdateJournalAuthorRes} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateJournalAuthor(c *gin.Context) {
	var author content_service.UpdateJournalAuthorReq

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		h.handleResponse(c, http.InvalidArgument, errors.New("journal id is not valid"))
		return
	}

	err := c.ShouldBindJSON(&author)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	author.JournalId = journalId

	resp, err := h.services.ContentService().UpdateJournalAuthor(
		c.Request.Context(),
		&author,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteJournalAuthor godoc
// @ID delete_journal_author_by_id
// @Router /journal/{journal-id}/author/{author-id} [GET]
// @Summary Delete Author By ID
// @Description Delete Author By ID
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param author-id path string true "author-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteJournalAuthor(c *gin.Context) {

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		h.handleResponse(c, http.InvalidArgument, errors.New("journal id is not valid"))
		return
	}

	authorID := c.Param("author-id")
	if !util.IsValidUUID(authorID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.services.ContentService().DeleteJournalAuthor(
		c.Request.Context(),
		&content_service.DeleteJournalAuthorReq{
			Id: authorID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetGeneralJournalAuthorList godoc
// @ID get_general_journal_author_list
// @Router /general/journal/{journal-id}/author [GET]
// @Summary Get Author List
// @Description  Get Author List
// @Tags General
// @Accept json
// @Produce json
// @Param journal-id path integer true "journal-id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=content_service.GetJournalAuthorListRes} "GetEditorListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGeneralJournalAuthorList(c *gin.Context) {

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		h.handleResponse(c, http.InvalidArgument, errors.New("journal id is not valid"))
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

	resp, err := h.services.ContentService().GetJournalAuthorList(
		c.Request.Context(),
		&content_service.GetJournalAuthorListReq{
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

// GetGeneralJournalAuthorByID godoc
// @ID get_general_journal_author_by_id
// @Router /general/journal/{journal-id}/author/{author-id} [GET]
// @Summary Get Author By ID
// @Description Get Author By ID
// @Tags General
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param author-id path string true "author-id"
// @Success 200 {object} http.Response{data=content_service.GetJournalAuthorRes} "EditorBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGeneralJournalAuthorByID(c *gin.Context) {

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		h.handleResponse(c, http.InvalidArgument, errors.New("journal id is not valid"))
		return
	}

	authorID := c.Param("author-id")

	if !util.IsValidUUID(authorID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.services.ContentService().GetJournalAuthor(
		c.Request.Context(),
		&content_service.GetJournalAuthorReq{
			Id: authorID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
