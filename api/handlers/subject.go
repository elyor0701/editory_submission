package handlers

import (
	"editory_submission/api/http"
	"editory_submission/genproto/content_service"
	"editory_submission/pkg/util"
	"github.com/gin-gonic/gin"
)

// CreateAdminSubject godoc
// @ID create_subject
// @Router /admin/subject [POST]
// @Summary Create Subject
// @Description Create Subject
// @Tags Admin
// @Accept json
// @Produce json
// @Param subject body content_service.CreateSubjectReq true "CreateSubjectRequestBody"
// @Success 201 {object} http.Response{data=content_service.CreateSubjectRes} "Subject data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAdminSubject(c *gin.Context) {
	var subject content_service.CreateSubjectReq

	err := c.ShouldBindJSON(&subject)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SubjectService().CreateSubject(
		c.Request.Context(),
		&subject,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetAdminSubjectList godoc
// @ID get_admin_subject_list
// @Router /admin/subject [GET]
// @Summary Get Subject List
// @Description  Get Subject List
// @Tags Admin
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=content_service.GetSubjectListRes} "GetSubjectListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminSubjectList(c *gin.Context) {

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

	resp, err := h.services.SubjectService().GetSubjectList(
		c.Request.Context(),
		&content_service.GetSubjectListReq{
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

// GetAdminSubjectByID godoc
// @ID get_admin_subject_by_id
// @Router /admin/subject/{subject-id} [GET]
// @Summary Get Subject By ID
// @Description Get Subject By ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param subject-id path string true "subject-id"
// @Success 200 {object} http.Response{data=content_service.GetSubjectRes} "SubjectBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminSubjectByID(c *gin.Context) {
	subjectID := c.Param("subject-id")

	if !util.IsValidUUID(subjectID) {
		h.handleResponse(c, http.InvalidArgument, "subject id is an invalid uuid")
		return
	}

	resp, err := h.services.SubjectService().GetSubject(
		c.Request.Context(),
		&content_service.GetSubjectReq{
			Id: subjectID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateAdminSubject godoc
// @ID update_subject
// @Router /admin/subject [PUT]
// @Summary Update Subject
// @Description Update Subject
// @Tags Admin
// @Accept json
// @Produce json
// @Param subject body content_service.UpdateSubjectReq true "UpdateSubjectRequestBody"
// @Success 200 {object} http.Response{data=content_service.UpdateSubjectRes} "Subject data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAdminSubject(c *gin.Context) {
	var subject content_service.UpdateSubjectReq

	err := c.ShouldBindJSON(&subject)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SubjectService().UpdateSubject(
		c.Request.Context(),
		&subject,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteAdminSubject godoc
// @ID delete_subject
// @Router /admin/subject/{subject-id} [DELETE]
// @Summary Delete Subject
// @Description Get Subject
// @Tags Admin
// @Accept json
// @Produce json
// @Param subject-id path string true "subject-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteAdminSubject(c *gin.Context) {
	subjectID := c.Param("subject-id")

	if !util.IsValidUUID(subjectID) {
		h.handleResponse(c, http.InvalidArgument, "subject id is an invalid uuid")
		return
	}

	_, err := h.services.SubjectService().DeleteSubject(
		c.Request.Context(),
		&content_service.DeleteSubjectReq{
			Id: subjectID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}
