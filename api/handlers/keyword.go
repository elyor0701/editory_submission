package handlers

import (
	"editory_submission/api/http"
	"editory_submission/genproto/auth_service"
	"editory_submission/pkg/util"
	"github.com/gin-gonic/gin"
)

// CreateAdminKeyword godoc
// @ID create_keyword
// @Router /admin/keyword [POST]
// @Summary Create Keyword
// @Description Create Keyword
// @Tags Admin
// @Accept json
// @Produce json
// @Param keyword body auth_service.CreateKeywordReq true "CreateKeywordRequestBody"
// @Success 201 {object} http.Response{data=auth_service.CreateKeywordRes} "Keyword data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAdminKeyword(c *gin.Context) {
	var keyword auth_service.CreateKeywordReq

	err := c.ShouldBindJSON(&keyword)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.KeywordService().CreateKeyword(
		c.Request.Context(),
		&keyword,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetAdminKeywordList godoc
// @ID get_admin_keyword_list
// @Router /admin/keyword [GET]
// @Summary Get Keyword List
// @Description  Get Keyword List
// @Tags Admin
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=auth_service.GetKeywordListRes} "GetKeywordListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminKeywordList(c *gin.Context) {

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

	resp, err := h.services.KeywordService().GetKeywordList(
		c.Request.Context(),
		&auth_service.GetKeywordListReq{
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

// GetAdminKeywordByID godoc
// @ID get_admin_keyword_by_id
// @Router /admin/keyword/{keyword-id} [GET]
// @Summary Get Keyword By ID
// @Description Get Keyword By ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param keyword-id path string true "keyword-id"
// @Success 200 {object} http.Response{data=auth_service.GetKeywordRes} "KeywordBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminKeywordByID(c *gin.Context) {
	keywordID := c.Param("keyword-id")

	if !util.IsValidUUID(keywordID) {
		h.handleResponse(c, http.InvalidArgument, "keyword id is an invalid uuid")
		return
	}

	resp, err := h.services.KeywordService().GetKeyword(
		c.Request.Context(),
		&auth_service.GetKeywordReq{
			Id: keywordID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateAdminKeyword godoc
// @ID update_keyword
// @Router /admin/keyword [PUT]
// @Summary Update Keyword
// @Description Update Keyword
// @Tags Admin
// @Accept json
// @Produce json
// @Param keyword body auth_service.UpdateKeywordReq true "UpdateKeywordRequestBody"
// @Success 200 {object} http.Response{data=auth_service.UpdateKeywordRes} "Keyword data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAdminKeyword(c *gin.Context) {
	var keyword auth_service.UpdateKeywordReq

	err := c.ShouldBindJSON(&keyword)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.KeywordService().UpdateKeyword(
		c.Request.Context(),
		&keyword,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteAdminKeyword godoc
// @ID delete_keyword
// @Router /admin/keyword/{keyword-id} [DELETE]
// @Summary Delete Keyword
// @Description Get Keyword
// @Tags Admin
// @Accept json
// @Produce json
// @Param keyword-id path string true "keyword-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteAdminKeyword(c *gin.Context) {
	keywordID := c.Param("keyword-id")

	if !util.IsValidUUID(keywordID) {
		h.handleResponse(c, http.InvalidArgument, "keyword id is an invalid uuid")
		return
	}

	_, err := h.services.KeywordService().DeleteKeyword(
		c.Request.Context(),
		&auth_service.DeleteKeywordReq{
			Id: keywordID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}
