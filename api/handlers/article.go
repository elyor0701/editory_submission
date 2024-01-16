package handlers

import (
	"editory_submission/api/http"
	"editory_submission/genproto/content_service"
	"github.com/saidamir98/udevs_pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateJournalArticle godoc
// @ID create_journal_article
// @Router /journal/{journal-id}/article [POST]
// @Summary Create Article
// @Description Create Article
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param article body content_service.CreateArticleReq true "CreateArticleRequestBody"
// @Success 201 {object} http.Response{data=content_service.CreateArticleRes} "Article data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateJournalArticle(c *gin.Context) {
	var article content_service.CreateArticleReq

	err := c.ShouldBindJSON(&article)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	article.JournalId = c.Param("journal-id")

	resp, err := h.services.ContentService().CreateArticle(
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
// @Router /journal/{journal-id}/article [GET]
// @Summary Get Article List
// @Description  Get Article List
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param date-from query string false "date-from"
// @Param date-to query string false "date-to"
// @Param sort query string false "sort"
// @Success 200 {object} http.Response{data=content_service.GetArticleListRes} "GetArticleListResponseBody"
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

	resp, err := h.services.ContentService().GetArticleList(
		c.Request.Context(),
		&content_service.GetArticleListReq{
			Limit:     int32(limit),
			Offset:    int32(offset),
			Search:    c.Query("search"),
			JournalId: c.Param("journal-id"),
			DateFrom:  c.Query("date-from"),
			DateTo:    c.Query("date-to"),
			Sort:      c.Query("sort"),
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
// @Router /journal/{journal-id}/article/{article-id} [GET]
// @Summary Get Article By ID
// @Description Get Article By ID
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param article-id path string true "article-id"
// @Success 200 {object} http.Response{data=content_service.GetArticleRes} "ArticleBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetJournalArticleByID(c *gin.Context) {
	articleID := c.Param("article-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	resp, err := h.services.ContentService().GetArticle(
		c.Request.Context(),
		&content_service.GetArticleReq{
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
// @Router /journal/{journal-id}/article [PUT]
// @Summary Update Article
// @Description Update Article
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param article body content_service.UpdateArticleReq true "UpdateArticleRequestBody"
// @Success 200 {object} http.Response{data=content_service.UpdateArticleRes} "Article data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateJournalArticle(c *gin.Context) {
	var article content_service.UpdateArticleReq

	err := c.ShouldBindJSON(&article)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	article.JournalId = c.Param("journal-id")

	resp, err := h.services.ContentService().UpdateArticle(
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
// @ID delete_journal_article
// @Router /journal/{journal-id}/article/{article-id} [DELETE]
// @Summary Delete Article
// @Description Get Article
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param article-id path string true "article-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteJournalArticle(c *gin.Context) {
	articleID := c.Param("article-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	_, err := h.services.ContentService().DeleteArticle(
		c.Request.Context(),
		&content_service.DeleteArticleReq{
			Id: articleID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
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
// @Param date-from query string false "date-from"
// @Param date-to query string false "date-to"
// @Param sort query string false "sort"
// @Success 200 {object} http.Response{data=content_service.GetArticleListRes} "GetArticleListResponseBody"
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

	resp, err := h.services.ContentService().GetArticleList(
		c.Request.Context(),
		&content_service.GetArticleListReq{
			Limit:    int32(limit),
			Offset:   int32(offset),
			Search:   c.Query("search"),
			DateFrom: c.Query("date-from"),
			DateTo:   c.Query("date-to"),
			Sort:     c.Query("sort"),
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
// @Success 200 {object} http.Response{data=content_service.GetArticleRes} "ArticleBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetGeneralArticleByID(c *gin.Context) {
	articleID := c.Param("article-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	resp, err := h.services.ContentService().GetArticle(
		c.Request.Context(),
		&content_service.GetArticleReq{
			Id: articleID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
