package handlers

import (
	"editory_submission/api/http"
	"editory_submission/api/models"
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	"editory_submission/genproto/content_service"
	"editory_submission/genproto/notification_service"
	"editory_submission/pkg/logger"
	"editory_submission/pkg/util"

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
// @Param journal body models.JournalUpdateReq true "UpdateJournalRequestBody"
// @Success 200 {object} http.Response{data=content_service.Journal} "Journal data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateJournal(c *gin.Context) {
	var (
		journal     models.JournalUpdateReq
		journalData []*content_service.JournalData
		subjects    []*content_service.Subject
	)

	err := c.ShouldBindJSON(&journal)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	for _, val := range journal.JournalData {
		journalData = append(journalData, &content_service.JournalData{
			JournalId: val.JournalId,
			Text:      val.Text,
			Type:      val.Type,
			ShortText: val.ShortText,
		})
	}

	for _, val := range journal.Subjects {
		subjects = append(subjects, &content_service.Subject{
			Id:    val.Id,
			Title: val.Title,
		})
	}

	resp, err := h.services.ContentService().UpdateJournal(
		c.Request.Context(),
		&content_service.Journal{
			Id:                        journal.Id,
			Title:                     journal.Title,
			CoverPhoto:                journal.CoverPhoto,
			Description:               journal.Description,
			AcceptanceToPublication:   journal.AcceptanceToPublication,
			AcceptanceRate:            journal.AcceptanceRate,
			SubmissionToFinalDecision: journal.SubmissionToFinalDecision,
			CitationIndicator:         journal.CitationIndicator,
			ImpactFactor:              journal.ImpactFactor,
			JournalData:               journalData,
			Subjects:                  subjects,
		},
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
// @Param journal body models.AdminJournalCreateReq true "CreateJournalRequestBody"
// @Success 201 {object} http.Response{data=models.AdminJournalCreateRes} "Journal data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateAdminJournal(c *gin.Context) {
	var journal models.AdminJournalCreateReq

	err := c.ShouldBindJSON(&journal)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	userPb, err := h.services.UserService().GetUser(
		c.Request.Context(),
		&auth_service.GetUserReq{
			Email: journal.Email,
		},
	)
	if err != nil {
		if util.IsErrNoRows(err) {
			userPb, err = h.services.UserService().CreateUser(
				c.Request.Context(),
				&auth_service.User{
					FirstName: journal.Author,
					Email:     journal.Email,
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

	resp, err := h.services.ContentService().CreateJournal(
		c.Request.Context(),
		&content_service.CreateJournalReq{
			Title:       journal.Title,
			Description: journal.Description,
			Isbn:        journal.Isbn,
			AuthorId:    userPb.Id,
			Status:      `ACTIVE`,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	role := auth_service.Role{
		RoleType:  config.EDITOR,
		JournalId: resp.Id,
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

	_, err = h.services.NotificationService().GenerateMailMessage(c.Request.Context(), &notification_service.GenerateMailMessageReq{
		UserId: resp.GetId(),
		Type:   config.NEW_JOURNAL_USER,
	})
	if err != nil {
		h.log.Error("cant send verification message", logger.String("err", err.Error()))
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
// @Param journal body models.AdminJournalUpdateReq true "UpdateJournalRequestBody"
// @Success 200 {object} http.Response{data=models.AdminJournalUpdateRes} "Journal data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateAdminJournal(c *gin.Context) {
	var (
		journal   models.AdminJournalUpdateReq
		journalPb content_service.Journal
	)

	err := c.ShouldBindJSON(&journal)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if util.IsValidEmail(journal.Email) {
		user, err := h.services.UserService().GetUser(
			c.Request.Context(),
			&auth_service.GetUserReq{
				Email: journal.Email,
			},
		)
		if err != nil {
			if util.IsErrNoRows(err) {
				user, err = h.services.UserService().CreateUser(
					c.Request.Context(),
					&auth_service.User{
						FirstName: journal.Author,
						Email:     journal.Email,
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
			RoleType:  config.EDITOR,
			JournalId: journal.Id,
			UserId:    user.Id,
		}

		_, err = h.services.RoleService().CreateRole(
			c.Request.Context(),
			&role,
		)
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}

		journalPb.AuthorId = user.Id
	}

	journalPb.Id = journal.Id
	journalPb.Title = journal.Title
	journalPb.Description = journal.Description
	journalPb.Isbn = journal.Isbn

	resp, err := h.services.ContentService().UpdateJournal(
		c.Request.Context(),
		&journalPb,
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

// GetGeneralJournalList godoc
// @ID get_general_journal_list
// @Router /general/journal [GET]
// @Summary Get Journal List
// @Description  Get Journal List
// @Tags General
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param date-from query string false "date-from"
// @Param date-to query string false "date-to"
// @Param sort query string false "sort"
// @Success 200 {object} http.Response{data=content_service.GetJournalListRes} "GetJournalListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGeneralJournalList(c *gin.Context) {

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
			Status:   "ACTIVE",
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

// GetGeneralJournalByID godoc
// @ID get_general_journal_by_id
// @Router /general/journal/{journal-id} [GET]
// @Summary Get Journal By ID
// @Description Get Journal By ID
// @Tags General
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Success 200 {object} http.Response{data=content_service.Journal} "JournalBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGeneralJournalByID(c *gin.Context) {
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
