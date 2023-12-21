package handlers

import (
	"editory_submission/api/http"
	"editory_submission/genproto/content_service"
	"errors"

	"github.com/saidamir98/udevs_pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateArticle godoc
// @ID create_article
// @Router /journal/{journal-id}/article [POST]
// @Summary Create Article
// @Description Create Article
// @Tags Article
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param article body content_service.CreateArticleReq true "CreateArticleRequestBody"
// @Success 201 {object} http.Response{data=content_service.Article} "Article data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateArticle(c *gin.Context) {
	var article content_service.CreateArticleReq

	journalId := c.Param("journal-id")

	err := c.ShouldBindJSON(&article)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	article.JournalId = journalId

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

// GetArticleList godoc
// @ID get_article_list
// @Router /journal/{journal-id}/article [GET]
// @Summary Get Article List
// @Description  Get Article List
// @Tags Article
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=content_service.GetArticleListRes} "GetArticleListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetArticleList(c *gin.Context) {

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

	resp, err := h.services.ContentService().GetArticleList(
		c.Request.Context(),
		&content_service.GetArticleListReq{
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

// GetArticleByID godoc
// @ID get_article_by_id
// @Router /journal/{journal-id}/article/{article-id} [GET]
// @Summary Get Article By ID
// @Description Get Article By ID
// @Tags Article
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param article-id path string true "article-id"
// @Success 200 {object} http.Response{data=content_service.Article} "ArticleBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetArticleByID(c *gin.Context) {
	articleID := c.Param("article-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	resp, err := h.services.ContentService().GetArticle(
		c.Request.Context(),
		&content_service.PrimaryKey{
			Id: articleID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateArticle godoc
// @ID update_article
// @Router /journal/{journal-id}/article [PUT]
// @Summary Update Article
// @Description Update Article
// @Tags Article
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param article body content_service.Article true "UpdateArticleRequestBody"
// @Success 200 {object} http.Response{data=content_service.Article} "Article data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateArticle(c *gin.Context) {
	var article content_service.Article

	journalId := c.Param("journal-id")

	err := c.ShouldBindJSON(&article)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	article.JournalId = journalId

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

// DeleteArticle godoc
// @ID delete_article
// @Router /journal/{journal-id}/article/{article-id} [DELETE]
// @Summary Delete Article
// @Description Get Article
// @Tags Article
// @Accept json
// @Produce json
// @Param journal-id path string true "Journal Id"
// @Param article-id path string true "article-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteArticle(c *gin.Context) {
	articleID := c.Param("article-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	_, err := h.services.ContentService().DeleteArticle(
		c.Request.Context(),
		&content_service.PrimaryKey{
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
// @Router /admin/article [GET]
// @Summary Get Article List
// @Description  Get Article List
// @Tags Admin
// @Accept json
// @Produce json
// @Param journal-id path string false "Journal Id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=content_service.GetArticleListRes} "GetArticleListResponseBody"
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

	resp, err := h.services.ContentService().GetArticleList(
		c.Request.Context(),
		&content_service.GetArticleListReq{
			Limit:     int32(limit),
			Offset:    int32(offset),
			Search:    c.Query("search"),
			JournalId: c.Query("journal-id"),
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
// @Router /admin/article/{article-id} [GET]
// @Summary Get Article By ID
// @Description Get Article By ID
// @Tags Admin
// @Accept json
// @Produce json
// @Param article-id path string true "article-id"
// @Success 200 {object} http.Response{data=content_service.Article} "ArticleBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAdminArticleByID(c *gin.Context) {
	articleID := c.Param("article-id")

	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	resp, err := h.services.ContentService().GetArticle(
		c.Request.Context(),
		&content_service.PrimaryKey{
			Id: articleID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
