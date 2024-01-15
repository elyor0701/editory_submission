package handlers

import (
	"editory_submission/api/http"
	"editory_submission/config"
	"editory_submission/genproto/submission_service"
	"errors"

	"github.com/saidamir98/udevs_pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateJournalArticle godoc
// @ID create_journal_article
// @Router /journal/{journal-id}/draft [POST]
// @Summary Create Article
// @Description Create Article
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param article body submission_service.CreateArticleReq true "CreateArticleRequestBody"
// @Success 201 {object} http.Response{data=submission_service.CreateArticleRes} "Article data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateJournalArticle(c *gin.Context) {
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

// GetJournalArticleList godoc
// @ID get_journal_article_list
// @Router /journal/{journal-id}/draft [GET]
// @Summary Get Article List
// @Description  Get Article List
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param status query string false "status"
// @Success 200 {object} http.Response{data=submission_service.GetArticleListRes} "GetArticleListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalArticleList(c *gin.Context) {

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
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetJournalArticleByID godoc
// @ID get_journal_article_by_id
// @Router /journal/{journal-id}/draft/{draft-id} [GET]
// @Summary Get Article By ID
// @Description Get Article By ID
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param draft-id path string true "draft-id"
// @Success 200 {object} http.Response{data=submission_service.Article} "ArticleBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalArticleByID(c *gin.Context) {
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

// UpdateJournalArticle godoc
// @ID update_journal_article
// @Router /journal/{journal-id}/draft [PUT]
// @Summary Update Article
// @Description Update Article
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param article body submission_service.UpdateArticleReq true "UpdateArticleRequestBody"
// @Success 200 {object} http.Response{data=submission_service.UpdateArticleRes} "Article data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateJournalArticle(c *gin.Context) {
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

// DeleteJournalArticle godoc
// @ID delete_journal_draft
// @Router /journal/{journal-id}/draft/{draft-id} [DELETE]
// @Summary Delete Article
// @Description Get Article
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param draft-id path string true "draft-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteJournalArticle(c *gin.Context) {
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

// GetAdminArticleList godoc
// @ID get_admin_article_list
// @Router /admin/draft [GET]
// @Summary Get Article List
// @Description  Get Article List
// @Tags Admin
// @Accept json
// @Produce json
// @Param journal-id path string false "Journal Id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param status query string false "status"
// @Success 200 {object} http.Response{data=submission_service.GetArticleListRes} "GetArticleListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminArticleList(c *gin.Context) {

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
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetAdminArticleByID godoc
// @ID get_admin_article_by_id
// @Router /admin/draft/{draft-id} [GET]
// @Summary Get Article By ID
// @Description Get Article By ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param draft-id path string true "draft-id"
// @Success 200 {object} http.Response{data=submission_service.Article} "ArticleBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminArticleByID(c *gin.Context) {
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

// GetGeneralArticleList godoc
// @ID get_general_article_list
// @Router /article [GET]
// @Summary Get Article List
// @Description  Get Article List
// @Tags General
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=submission_service.GetArticleListRes} "GetArticleListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGeneralArticleList(c *gin.Context) {

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
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
			Status: config.ARTICLE_STATUS_PUBLISHED,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetGeneralArticleByID godoc
// @ID get_general_article_by_id
// @Router /article/{article-id} [GET]
// @Summary Get Article By ID
// @Description Get Article By ID
// @Tags General
// @Accept json
// @Produce json
// @Param article-id path string true "article-id"
// @Success 200 {object} http.Response{data=submission_service.Article} "ArticleBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGeneralArticleByID(c *gin.Context) {
	articleID := c.Param("article-id")

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
