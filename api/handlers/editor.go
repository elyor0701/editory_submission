package handlers

import (
	"editory_submission/api/http"
	"editory_submission/api/models"
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"editory_submission/pkg/util"
	"github.com/gin-gonic/gin"
)

// CreateEditor godoc
// @ID create_admin_editor
// @Router /admin/editor [POST]
// @Summary Create Editor
// @Description Create Editor
// @Tags Admin
// @Accept json
// @Produce json
// @Param user body models.CreateEditorReq true "CreateEditorRequestBody"
// @Success 201 {object} http.Response{data=models.CreateEditorRes} "Editor data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateEditor(c *gin.Context) {
	var (
		user models.CreateEditorReq
	)

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	role := []*auth_service.Role{
		{
			RoleType:  config.EDITOR,
			JournalId: user.JournalId,
		},
	}

	resp, err := h.services.UserService().CreateUser(
		c.Request.Context(),
		&auth_service.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  user.Password,
			Role:      role,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	res := models.CreateEditorRes{
		Id:        resp.GetId(),
		FirstName: resp.GetFirstName(),
		LastName:  resp.GetLastName(),
		Email:     resp.GetEmail(),
		Password:  resp.GetPassword(),
	}

	h.handleResponse(c, http.Created, res)
}

// GetEditorList godoc
// @ID get_editor_list
// @Router /admin/editor [GET]
// @Summary Get Editor List
// @Description  Get Editor List
// @Tags Admin
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param journal-id query string false "journal-id"
// @Success 200 {object} http.Response{data=auth_service.GetUserListRes} "GetEditorListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetEditorList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	journalId := c.Query("journal-id")

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
			RoleType:  config.EDITOR,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetEditorByID godoc
// @ID get_editor_by_id
// @Router /admin/editor/{editor-id} [GET]
// @Summary Get Editor By ID
// @Description Get Editor By ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param editor-id path string true "editor-id"
// @Success 200 {object} http.Response{data=auth_service.User} "EditorBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetEditorByID(c *gin.Context) {
	userID := c.Param("editor-id")

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

// UpdateEditor godoc
// @ID update_editor
// @Router /admin/editor [PUT]
// @Summary Update Editor
// @Description Update Editor
// @Tags Admin
// @Accept json
// @Produce json
// @Param editor body auth_service.User true "UpdateUserRequestBody"
// @Success 200 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateEditor(c *gin.Context) {
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

// DeleteEditor godoc
// @ID delete_editor
// @Router /admin/editor/{editor-id} [DELETE]
// @Summary Delete Editor
// @Description Get Editor
// @Tags Admin
// @Accept json
// @Produce json
// @Param editor-id path string true "editor-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteEditor(c *gin.Context) {
	h.handleResponse(c, http.NoContent, "")
	return

	userID := c.Param("editor-id")

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
