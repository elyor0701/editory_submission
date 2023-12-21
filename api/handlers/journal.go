package handlers

import (
	"editory_submission/api/http"
	"editory_submission/genproto/content_service"

	"github.com/saidamir98/udevs_pkg/util"

	"github.com/gin-gonic/gin"
)

//// CreateJournal godoc
//// @ID create_journal
//// @Router /journal [POST]
//// @Summary Create Journal
//// @Description Create Journal types: ['EDITOR_SPOTLIGHT', 'SPECIAL_ISSUE', 'ABOUT_JOURNAL', 'EDITORIAL_BARD', 'PEER_REVIEW_PROCESS', 'PUBLICATION_ETHICS', 'ABSTRACTING_INDEXING', 'ARTICLE_PROCESSING_CHARGES']
//// @Tags Journal
//// @Accept json
//// @Produce json
//// @Param journal body content_service.CreateJournalReq true "CreateJournalRequestBody"
//// @Success 201 {object} http.Response{data=content_service.Journal} "Journal data"
//// @Response 400 {object} http.Response{data=string} "Bad Request"
//// @Failure 500 {object} http.Response{data=string} "Server Error"
//func (h *Handler) CreateJournal(c *gin.Context) {
//	var journal content_service.CreateJournalReq
//
//	err := c.ShouldBindJSON(&journal)
//	if err != nil {
//		h.handleResponse(c, http.BadRequest, err.Error())
//		return
//	}
//
//	resp, err := h.services.ContentService().CreateJournal(
//		c.Request.Context(),
//		&journal,
//	)
//
//	if err != nil {
//		h.handleResponse(c, http.GRPCError, err.Error())
//		return
//	}
//
//	h.handleResponse(c, http.Created, resp)
//}

//// GetJournalList godoc
//// @ID get_journal_list
//// @Router /journal [GET]
//// @Summary Get Journal List
//// @Description  Get Journal List
//// @Tags Journal
//// @Accept json
//// @Produce json
//// @Param offset query integer false "offset"
//// @Param limit query integer false "limit"
//// @Param search query string false "search"
//// @Success 200 {object} http.Response{data=content_service.GetJournalListRes} "GetJournalListResponseBody"
//// @Response 400 {object} http.Response{data=string} "Invalid Argument"
//// @Failure 500 {object} http.Response{data=string} "Server Error"
//func (h *Handler) GetJournalList(c *gin.Context) {
//
//	offset, err := h.getOffsetParam(c)
//	if err != nil {
//		h.handleResponse(c, http.InvalidArgument, err.Error())
//		return
//	}
//
//	limit, err := h.getLimitParam(c)
//	if err != nil {
//		h.handleResponse(c, http.InvalidArgument, err.Error())
//		return
//	}
//
//	resp, err := h.services.ContentService().GetJournalList(
//		c.Request.Context(),
//		&content_service.GetList{
//			Limit:  int32(limit),
//			Offset: int32(offset),
//			Search: c.DefaultQuery("search", ""),
//		},
//	)
//
//	if err != nil {
//		h.handleResponse(c, http.GRPCError, err.Error())
//		return
//	}
//
//	h.handleResponse(c, http.OK, resp)
//}

// GetJournalByID godoc
// @ID get_journal_by_id
// @Router /journal/{journal-id} [GET]
// @Summary Get Journal By ID
// @Description Get Journal By ID
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Success 200 {object} http.Response{data=content_service.Journal} "JournalBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalByID(c *gin.Context) {
	journalID := c.Param("journal-id")

	if !util.IsValidUUID(journalID) {
		h.handleResponse(c, http.InvalidArgument, "journal id is an invalid uuid")
		return
	}

	resp, err := h.services.ContentService().GetJournal(
		c.Request.Context(),
		&content_service.PrimaryKey{
			Id: journalID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateJournal godoc
// @ID update_journal
// @Router /journal [PUT]
// @Summary Update Journal
// @Description Update Journal
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal body content_service.Journal true "UpdateJournalRequestBody"
// @Success 200 {object} http.Response{data=content_service.Journal} "Journal data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateJournal(c *gin.Context) {
	var journal content_service.Journal

	err := c.ShouldBindJSON(&journal)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ContentService().UpdateJournal(
		c.Request.Context(),
		&journal,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteJournal godoc
// @ID delete_journal
// @Router /journal/{journal-id} [DELETE]
// @Summary Delete Journal
// @Description Get Journal
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteJournal(c *gin.Context) {
	h.handleResponse(c, http.NoContent, "")
	return
	journalID := c.Param("journal-id")

	if !util.IsValidUUID(journalID) {
		h.handleResponse(c, http.InvalidArgument, "journal id is an invalid uuid")
		return
	}

	_, err := h.services.ContentService().DeleteJournal(
		c.Request.Context(),
		&content_service.PrimaryKey{
			Id: journalID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}

// CreateAdminJournal godoc
// @ID create_admin_journal
// @Router /admin/journal [POST]
// @Summary Create Journal
// @Description Create Journal types: ['EDITOR_SPOTLIGHT', 'SPECIAL_ISSUE', 'ABOUT_JOURNAL', 'EDITORIAL_BARD', 'PEER_REVIEW_PROCESS', 'PUBLICATION_ETHICS', 'ABSTRACTING_INDEXING', 'ARTICLE_PROCESSING_CHARGES']
// @Tags Admin
// @Accept json
// @Produce json
// @Param journal body content_service.CreateJournalReq true "CreateJournalRequestBody"
// @Success 201 {object} http.Response{data=content_service.Journal} "Journal data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAdminJournal(c *gin.Context) {
	var journal content_service.CreateJournalReq

	err := c.ShouldBindJSON(&journal)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ContentService().CreateJournal(
		c.Request.Context(),
		&journal,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetAdminJournalList godoc
// @ID get_admin_journal_list
// @Router /admin/journal [GET]
// @Summary Get Journal List
// @Description  Get Journal List
// @Tags Admin
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param status query string false "status"
// @Param date-from query string false "date-from"
// @Param date-to query string false "date-to"
// @Param sort query string false "sort"
// @Success 200 {object} http.Response{data=content_service.GetJournalListRes} "GetJournalListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminJournalList(c *gin.Context) {

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

	resp, err := h.services.ContentService().GetJournalList(
		c.Request.Context(),
		&content_service.GetList{
			Limit:    int32(limit),
			Offset:   int32(offset),
			Search:   c.DefaultQuery("search", ""),
			Status:   c.DefaultQuery("status", ""),
			DateFrom: c.DefaultQuery("date-from", ""),
			DateTo:   c.DefaultQuery("date-to", ""),
			Sort:     c.DefaultQuery("sort", ""),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetAdminJournalByID godoc
// @ID get_admin_journal_by_id
// @Router /admin/journal/{journal-id} [GET]
// @Summary Get Journal By ID
// @Description Get Journal By ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Success 200 {object} http.Response{data=content_service.Journal} "JournalBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminJournalByID(c *gin.Context) {
	journalID := c.Param("journal-id")

	if !util.IsValidUUID(journalID) {
		h.handleResponse(c, http.InvalidArgument, "journal id is an invalid uuid")
		return
	}

	resp, err := h.services.ContentService().GetJournal(
		c.Request.Context(),
		&content_service.PrimaryKey{
			Id: journalID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateAdminJournal godoc
// @ID update_admin_journal
// @Router /admin/journal [PUT]
// @Summary Update Journal
// @Description Update Journal
// @Tags Admin
// @Accept json
// @Produce json
// @Param journal body content_service.Journal true "UpdateJournalRequestBody"
// @Success 200 {object} http.Response{data=content_service.Journal} "Journal data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAdminJournal(c *gin.Context) {
	var journal content_service.Journal

	err := c.ShouldBindJSON(&journal)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ContentService().UpdateJournal(
		c.Request.Context(),
		&journal,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteAdminJournal godoc
// @ID delete_admin_journal
// @Router /admin/journal/{journal-id} [DELETE]
// @Summary Delete Journal
// @Description Get Journal
// @Tags Admin
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteAdminJournal(c *gin.Context) {
	journalID := c.Param("journal-id")

	if !util.IsValidUUID(journalID) {
		h.handleResponse(c, http.InvalidArgument, "journal id is an invalid uuid")
		return
	}

	_, err := h.services.ContentService().DeleteJournal(
		c.Request.Context(),
		&content_service.PrimaryKey{
			Id: journalID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}
