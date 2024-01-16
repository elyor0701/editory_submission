package handlers

import (
	"editory_submission/api/http"
	"editory_submission/api/models"
	"editory_submission/config"
	"editory_submission/genproto/submission_service"
	"editory_submission/pkg/util"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
)

// CreateUserDraft godoc
// @ID create_user_draft
// @Router /user/{user-id}/draft [POST]
// @Summary Create Draft
// @Description Create Draft
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft body models.CreateUserDraftReq true "CreateDraftRequestBody"
// @Success 201 {object} http.Response{data=submission_service.CreateArticleRes} "Draft data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUserDraft(c *gin.Context) {
	var (
		article   models.CreateUserDraftReq
		articlePB submission_service.CreateArticleReq
	)

	userId := c.Param("user-id")

	err := c.ShouldBindJSON(&article)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	artJSON, err := json.Marshal(article)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = json.Unmarshal(artJSON, &articlePB)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	articlePB.AuthorId = userId
	articlePB.Status = config.ARTICLE_STATUS_NEW

	resp, err := h.services.ArticleService().CreateArticle(
		c.Request.Context(),
		&articlePB,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetUserDraftList godoc
// @ID get_user_draft_list
// @Router /user/{user-id}/draft [GET]
// @Summary Get Draft List
// @Description  Get Draft List
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param status query string false "status"
// @Param group-id query string false "group id"
// @Success 200 {object} http.Response{data=submission_service.GetArticleListRes} "GetDraftListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserDraftList(c *gin.Context) {

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

	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		err = errors.New("invalid user id")
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.ArticleService().GetArticleList(
		c.Request.Context(),
		&submission_service.GetArticleListReq{
			Limit:    int32(limit),
			Offset:   int32(offset),
			Search:   c.Query("search"),
			AuthorId: userId,
			Status:   c.Query("status"),
			GroupId:  c.Query("group-id"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetUserDraftByID godoc
// @ID get_user_draft_by_id
// @Router /user/{user-id}/draft/{draft-id} [GET]
// @Summary Get Draft By ID
// @Description Get Draft By ID
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft-id path string true "draft-id"
// @Success 200 {object} http.Response{data=submission_service.Article} "DraftBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserDraftByID(c *gin.Context) {
	articleID := c.Param("draft-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	resp, err := h.services.ArticleService().GetArticle(
		c.Request.Context(),
		&submission_service.GetArticleReq{
			Id: articleID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateUserDraft godoc
// @ID update_user_draft
// @Router /user/{user-id}/draft [PUT]
// @Summary Update Draft
// @Description Update Draft
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft body submission_service.UpdateArticleReq true "UpdateDraftRequestBody"
// @Success 200 {object} http.Response{data=submission_service.UpdateArticleRes} "Draft data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUserDraft(c *gin.Context) {
	var article submission_service.UpdateArticleReq

	userId := c.Param("user-id")

	err := c.ShouldBindJSON(&article)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	article.AuthorId = userId

	resp, err := h.services.ArticleService().UpdateArticle(
		c.Request.Context(),
		&article,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteUserDraft godoc
// @ID delete_journal_draft
// @Router /user/{user-id}/draft/{draft-id} [DELETE]
// @Summary Delete Draft
// @Description Get Draft
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft-id path string true "draft-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteUserDraft(c *gin.Context) {
	articleID := c.Param("draft-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "draft id is an invalid uuid")
		return
	}

	_, err := h.services.ArticleService().DeleteArticle(
		c.Request.Context(),
		&submission_service.DeleteArticleReq{
			Id: articleID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}

// CreateJournalDraft godoc
// @ID create_journal_draft
// @Router /journal/{journal-id}/draft [POST]
// @Summary Create Draft
// @Description Create Draft
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param draft body submission_service.CreateArticleReq true "CreateDraftRequestBody"
// @Success 201 {object} http.Response{data=submission_service.CreateArticleRes} "Draft data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateJournalDraft(c *gin.Context) {
	var article submission_service.CreateArticleReq

	journalId := c.Param("journal-id")

	err := c.ShouldBindJSON(&article)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	article.JournalId = journalId
	article.Status = config.ARTICLE_STATUS_PUBLISHED

	resp, err := h.services.ArticleService().CreateArticle(
		c.Request.Context(),
		&article,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetJournalDraftList godoc
// @ID get_journal_draft_list
// @Router /journal/{journal-id}/draft [GET]
// @Summary Get Draft List
// @Description  Get Draft List
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param status query string false "status"
// @Param group-id query string false "group_id"
// @Success 200 {object} http.Response{data=submission_service.GetArticleListRes} "GetDraftListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalDraftList(c *gin.Context) {

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

	resp, err := h.services.ArticleService().GetArticleList(
		c.Request.Context(),
		&submission_service.GetArticleListReq{
			Limit:     int32(limit),
			Offset:    int32(offset),
			Search:    c.Query("search"),
			JournalId: journalId,
			Status:    c.Query("status"),
			GroupId:   c.Query("group-id"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetJournalDraftByID godoc
// @ID get_journal_draft_by_id
// @Router /journal/{journal-id}/draft/{draft-id} [GET]
// @Summary Get Draft By ID
// @Description Get Draft By ID
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param draft-id path string true "draft-id"
// @Success 200 {object} http.Response{data=submission_service.Article} "DraftBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalDraftByID(c *gin.Context) {
	articleID := c.Param("draft-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	resp, err := h.services.ArticleService().GetArticle(
		c.Request.Context(),
		&submission_service.GetArticleReq{
			Id: articleID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateJournalDraft godoc
// @ID update_journal_draft
// @Router /journal/{journal-id}/draft [PUT]
// @Summary Update Draft
// @Description Update Draft
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param draft body submission_service.UpdateArticleReq true "UpdateDraftRequestBody"
// @Success 200 {object} http.Response{data=submission_service.UpdateArticleRes} "Draft data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateJournalDraft(c *gin.Context) {
	var article submission_service.UpdateArticleReq

	journalId := c.Param("journal-id")

	err := c.ShouldBindJSON(&article)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	article.JournalId = journalId

	resp, err := h.services.ArticleService().UpdateArticle(
		c.Request.Context(),
		&article,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteJournalDraft godoc
// @ID delete_journal_draft
// @Router /journal/{journal-id}/draft/{draft-id} [DELETE]
// @Summary Delete Draft
// @Description Get Draft
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param draft-id path string true "draft-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteJournalDraft(c *gin.Context) {
	articleID := c.Param("draft-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	_, err := h.services.ArticleService().DeleteArticle(
		c.Request.Context(),
		&submission_service.DeleteArticleReq{
			Id: articleID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}

// GetAdminDraftList godoc
// @ID get_admin_draft_list
// @Router /admin/draft [GET]
// @Summary Get Draft List
// @Description  Get Draft List
// @Tags Admin
// @Accept json
// @Produce json
// @Param journal-id path string false "Journal Id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param status query string false "status"
// @Param group-id query string false "group id"
// @Success 200 {object} http.Response{data=submission_service.GetArticleListRes} "GetDraftListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminDraftList(c *gin.Context) {

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

	resp, err := h.services.ArticleService().GetArticleList(
		c.Request.Context(),
		&submission_service.GetArticleListReq{
			Limit:     int32(limit),
			Offset:    int32(offset),
			Search:    c.Query("search"),
			JournalId: c.Query("journal-id"),
			Status:    c.Query("status"),
			GroupId:   c.Query("group-id"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetAdminDraftByID godoc
// @ID get_admin_draft_by_id
// @Router /admin/draft/{draft-id} [GET]
// @Summary Get Draft By ID
// @Description Get Draft By ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param draft-id path string true "draft-id"
// @Success 200 {object} http.Response{data=submission_service.Article} "DraftBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminDraftByID(c *gin.Context) {
	articleID := c.Param("draft-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	resp, err := h.services.ArticleService().GetArticle(
		c.Request.Context(),
		&submission_service.GetArticleReq{
			Id: articleID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
