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

// CreateUserArticle godoc
// @ID create_user_article
// @Router /user/{user-id}/draft [POST]
// @Summary Create Article
// @Description Create Article
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param article body models.CreateUserArticleReq true "CreateArticleRequestBody"
// @Success 201 {object} http.Response{data=submission_service.CreateArticleRes} "Article data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUserArticle(c *gin.Context) {
	var (
		article   models.CreateUserArticleReq
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

// GetUserArticleList godoc
// @ID get_user_article_list
// @Router /user/{user-id}/draft [GET]
// @Summary Get Article List
// @Description  Get Article List
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param status query string false "status"
// @Param group-id query string false "group id"
// @Success 200 {object} http.Response{data=submission_service.GetArticleListRes} "GetArticleListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserArticleList(c *gin.Context) {

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

// GetUserArticleByID godoc
// @ID get_user_article_by_id
// @Router /user/{user-id}/draft/{draft-id} [GET]
// @Summary Get Article By ID
// @Description Get Article By ID
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft-id path string true "draft-id"
// @Success 200 {object} http.Response{data=submission_service.Article} "ArticleBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserArticleByID(c *gin.Context) {
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

// UpdateUserArticle godoc
// @ID update_user_article
// @Router /user/{user-id}/draft [PUT]
// @Summary Update Article
// @Description Update Article
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param article body submission_service.UpdateArticleReq true "UpdateArticleRequestBody"
// @Success 200 {object} http.Response{data=submission_service.UpdateArticleRes} "Article data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUserArticle(c *gin.Context) {
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

// DeleteUserArticle godoc
// @ID delete_journal_article
// @Router /user/{user-id}/draft/{draft-id} [DELETE]
// @Summary Delete Article
// @Description Get Article
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft-id path string true "draft-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteUserArticle(c *gin.Context) {
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
// @Param group-id query string false "group_id"
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
			GroupId:   c.Query("group-id"),
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
// @Param group-id query string false "group id"
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
			GroupId:   c.Query("group-id"),
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
