package handlers

import (
	"editory_submission/api/http"
	"editory_submission/genproto/content_service"

	"github.com/saidamir98/udevs_pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateUniversity godoc
// @ID create_university
// @Router /university [POST]
// @Summary Create University
// @Description Create University
// @Tags University
// @Accept json
// @Produce json
// @Param university body content_service.CreateUniversityReq true "CreateUniversityRequestBody"
// @Success 201 {object} http.Response{data=content_service.CreateUniversityRes} "University data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUniversity(c *gin.Context) {
	var university content_service.CreateUniversityReq

	err := c.ShouldBindJSON(&university)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UniversityService().CreateUniversity(
		c.Request.Context(),
		&university,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetUniversityList godoc
// @ID get_university_list
// @Router /university [GET]
// @Summary Get University List
// @Description  Get University List
// @Tags University
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=content_service.GetUniversityListRes} "GetUniversityListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUniversityList(c *gin.Context) {

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

	resp, err := h.services.UniversityService().GetUniversityList(
		c.Request.Context(),
		&content_service.GetUniversityListReq{
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

// GetUniversityByID godoc
// @ID get_university_by_id
// @Router /university/{university-id} [GET]
// @Summary Get University By ID
// @Description Get University By ID
// @Tags University
// @Accept json
// @Produce json
// @Param university-id path string true "university-id"
// @Success 200 {object} http.Response{data=content_service.GetUniversityRes} "UniversityBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUniversityByID(c *gin.Context) {
	universityID := c.Param("university-id")

	if !util.IsValidUUID(universityID) {
		h.handleResponse(c, http.InvalidArgument, "university id is an invalid uuid")
		return
	}

	resp, err := h.services.UniversityService().GetUniversity(
		c.Request.Context(),
		&content_service.GetUniversityReq{
			Id: universityID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateUniversity godoc
// @ID update_university
// @Router /university [PUT]
// @Summary Update University
// @Description Update University
// @Tags University
// @Accept json
// @Produce json
// @Param university body content_service.UpdateUniversityReq true "UpdateUniversityRequestBody"
// @Success 200 {object} http.Response{data=content_service.UpdateUniversityRes} "University data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUniversity(c *gin.Context) {
	var university content_service.UpdateUniversityReq

	err := c.ShouldBindJSON(&university)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UniversityService().UpdateUniversity(
		c.Request.Context(),
		&university,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteUniversity godoc
// @ID delete_university
// @Router /university/{university-id} [DELETE]
// @Summary Delete University
// @Description Get University
// @Tags University
// @Accept json
// @Produce json
// @Param university-id path string true "university-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteUniversity(c *gin.Context) {
	universityID := c.Param("university-id")

	if !util.IsValidUUID(universityID) {
		h.handleResponse(c, http.InvalidArgument, "university id is an invalid uuid")
		return
	}

	_, err := h.services.UniversityService().DeleteUniversity(
		c.Request.Context(),
		&content_service.DeleteUniversityReq{
			Id: universityID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}
