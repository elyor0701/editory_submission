package handlers

import (
	"editory_submission/api/http"
	"editory_submission/config"
	"editory_submission/genproto/submission_service"
	"github.com/saidamir98/udevs_pkg/util"

	"github.com/gin-gonic/gin"
)

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
