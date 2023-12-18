package handlers

import (
	"editory_submission/api/http"
	"editory_submission/genproto/content_service"
	"editory_submission/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
)

// CreateEdition godoc
// @ID create_edition
// @Router /journal/{journal-id}/edition [POST]
// @Summary Create Edition
// @Description Create Edition
// @Tags Edition
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param edition body content_service.CreateEditionReq true "CreateEditionRequestBody"
// @Success 201 {object} http.Response{data=content_service.Edition} "Edition data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateEdition(c *gin.Context) {
	var edition content_service.CreateEditionReq

	journalId := c.Param("journal-id")

	err := c.ShouldBindJSON(&edition)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	edition.JournalId = journalId

	resp, err := h.services.ContentService().CreateEdition(
		c.Request.Context(),
		&edition,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetEditionList godoc
// @ID get_edition_list
// @Router /journal/{journal-id}/edition [GET]
// @Summary Get Edition List
// @Description  Get Edition List
// @Tags Edition
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=content_service.GetEditionListRes} "GetEditionListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetEditionList(c *gin.Context) {

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

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		err = errors.New("invalid journal id")
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.ContentService().GetEditionList(
		c.Request.Context(),
		&content_service.GetEditionListReq{
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

// GetEditionByID godoc
// @ID get_edition_by_id
// @Router /journal/{journal-id}/edition/{edition-id} [GET]
// @Summary Get Edition By ID
// @Description Get Edition By ID
// @Tags Edition
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param edition-id path string true "edition-id"
// @Success 200 {object} http.Response{data=content_service.Edition} "EditionBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetEditionByID(c *gin.Context) {
	editionID := c.Param("edition-id")

	if !util.IsValidUUID(editionID) {
		h.handleResponse(c, http.InvalidArgument, "edition id is an invalid uuid")
		return
	}

	resp, err := h.services.ContentService().GetEdition(
		c.Request.Context(),
		&content_service.PrimaryKey{
			Id: editionID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateEdition godoc
// @ID update_edition
// @Router /journal/{journal-id}/edition [PUT]
// @Summary Update Edition
// @Description Update Edition
// @Tags Edition
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param edition body content_service.Edition true "UpdateEditionRequestBody"
// @Success 200 {object} http.Response{data=content_service.Edition} "Edition data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateEdition(c *gin.Context) {
	var edition content_service.Edition

	journalId := c.Param("journal-id")

	err := c.ShouldBindJSON(&edition)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	edition.JournalId = journalId

	resp, err := h.services.ContentService().UpdateEdition(
		c.Request.Context(),
		&edition,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteEdition godoc
// @ID delete_edition
// @Router /journal/{journal-id}/edition/{edition-id} [DELETE]
// @Summary Delete Edition
// @Description Get Edition
// @Tags Edition
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param edition-id path string true "edition-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteEdition(c *gin.Context) {
	editionID := c.Param("edition-id")

	if !util.IsValidUUID(editionID) {
		h.handleResponse(c, http.InvalidArgument, "edition id is an invalid uuid")
		return
	}

	_, err := h.services.ContentService().DeleteEdition(
		c.Request.Context(),
		&content_service.PrimaryKey{
			Id: editionID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}
