package handlers

import (
	"editory_submission/api/http"
	"editory_submission/api/models"
	"editory_submission/config"
	pb "editory_submission/genproto/submission_service"
	"editory_submission/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
)

// GetUserDraftCheckList godoc
// @ID get_user_draft_check_list
// @Router /user/{user-id}/draft/{draft-id}/check [GET]
// @Summary Get user draft check list
// @Description  Get user draft check list
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user-id"
// @Param draft-id path string true "draft-id"
// @Success 200 {object} http.Response{data=pb.GetArticleCheckerListRes} "GetArticleReviewerListRes"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserDraftCheckList(c *gin.Context) {

	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	draftId := c.Param("draft-id")
	if !util.IsValidUUID(draftId) {
		h.handleResponse(c, http.InvalidArgument, "draft id is an invalid uuid")
		return
	}

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

	resp, err := h.services.CheckerService().GetArticleCheckerList(
		c.Request.Context(),
		&pb.GetArticleCheckerListReq{
			Limit:     int32(limit),
			Offset:    int32(offset),
			Search:    c.DefaultQuery("search", ""),
			ArticleId: draftId,
			Type:      config.EDITOR,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetUserDraftCheckByID godoc
// @ID get_user_draft_check_by_id
// @Router /user/{user-id}/draft/{draft-id}/check/{check-id} [GET]
// @Summary Get user draft check
// @Description  Get user draft check
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user-id"
// @Param draft-id path string true "draft-id"
// @Param check-id path string true "check-id"
// @Success 200 {object} http.Response{data=pb.GetArticleCheckerRes} "GetArticleReviewerRes"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserDraftCheckByID(c *gin.Context) {

	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	draftId := c.Param("draft-id")
	if !util.IsValidUUID(draftId) {
		h.handleResponse(c, http.InvalidArgument, "draft id is an invalid uuid")
		return
	}

	checkId := c.Param("check-id")
	if !util.IsValidUUID(checkId) {
		h.handleResponse(c, http.InvalidArgument, "check id is an invalid uuid")
		return
	}

	resp, err := h.services.CheckerService().GetArticleChecker(
		c.Request.Context(),
		&pb.GetArticleCheckerReq{
			Id: checkId,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// CreateArticleCheck godoc
// @ID create_article_check
// @Router /journal/{journal-id}/draft/{draft-id}/check [POST]
// @Summary Create Article check
// @Description Create Article check
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal Id"
// @Param draft-id path string true "draft Id"
// @Param check body models.CreateArticleCheckReq true "CreateArticleCheckReq"
// @Success 201 {object} http.Response{data=pb.CreateArticleCheckerRes} "CreateArticleReviewerRes"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateArticleCheck(c *gin.Context) {
	var (
		fileComments []*pb.FileComment
		check        models.CreateArticleCheckReq
	)

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		err := errors.New("journal-id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	articleId := c.Param("draft-id")
	if !util.IsValidUUID(articleId) {
		err := errors.New("draft-id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err := c.ShouldBindJSON(&check)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	for _, val := range check.FileComments {
		fileComments = append(fileComments, &pb.FileComment{
			Id:      val.Id,
			Type:    val.Type,
			FileId:  val.FileId,
			Comment: val.Comment,
		})
	}

	resp, err := h.services.CheckerService().CreateArticleChecker(
		c.Request.Context(),
		&pb.CreateArticleCheckerReq{
			CheckerId: check.EditorId,
			ArticleId: articleId,
			Status:    check.Status,
			Type:      config.EDITOR,
			Comments:  fileComments,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	draftStatus := ""
	draftStep := ""

	switch check.Status {
	case config.ARTICLE_EDITOR_STATUS_REJECTED:
		draftStatus = config.ARTICLE_STATUS_DENIED
		draftStep = config.DRAFT_STEP_AUTHOR
	case config.ARTICLE_EDITOR_STATUS_REJECTED_WITH_CORRECTION:
		draftStatus = config.ARTICLE_REVIEWER_STATUS_BACK_FOR_CORRECTION
		draftStep = config.DRAFT_STEP_AUTHOR
	case config.ARTICLE_EDITOR_STATUS_APPROVED:
		draftStatus = config.ARTICLE_STATUS_CONFIRMED
		draftStep = config.DRAFT_STEP_AUTHOR
	case config.ARTICLE_EDITOR_STATUS_APPROVED_WITH_CORRECTION:
		draftStatus = config.ARTICLE_REVIEWER_STATUS_BACK_FOR_CORRECTION
		draftStep = config.DRAFT_STEP_AUTHOR
	default:
		draftStatus = config.ARTICLE_STATUS_PENDING
	}

	_, err = h.services.ArticleService().UpdateArticle(
		c.Request.Context(),
		&pb.UpdateArticleReq{
			Status: draftStatus,
			Id:     articleId,
			Step:   draftStep,
		})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// UpdateArticleCheck godoc
// @ID update_article_check
// @Router /journal/{journal-id}/draft/{draft-id}/check [PUT]
// @Summary update Article check
// @Description update Article check
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal Id"
// @Param draft-id path string true "draft Id"
// @Param check body models.UpdateArticleCheckReq true "UpdateArticleCheckReq"
// @Success 201 {object} http.Response{data=pb.CreateArticleCheckerRes} "CreateArticleReviewerRes"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateArticleCheck(c *gin.Context) {
	var (
		fileComments []*pb.FileComment
		check        models.UpdateArticleCheckReq
	)

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		err := errors.New("journal-id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	articleId := c.Param("draft-id")
	if !util.IsValidUUID(articleId) {
		err := errors.New("draft-id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err := c.ShouldBindJSON(&check)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	for _, val := range check.FileComments {
		fileComments = append(fileComments, &pb.FileComment{
			Id:      val.Id,
			Type:    val.Type,
			FileId:  val.FileId,
			Comment: val.Comment,
		})
	}

	resp, err := h.services.CheckerService().UpdateArticleChecker(
		c.Request.Context(),
		&pb.UpdateArticleCheckerReq{
			Id:        check.Id,
			Status:    check.Status,
			ArticleId: articleId,
			Comments:  fileComments,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	draftStatus := ""
	draftStep := ""

	switch check.Status {
	case config.ARTICLE_EDITOR_STATUS_REJECTED:
		draftStatus = config.ARTICLE_STATUS_DENIED
		draftStep = config.DRAFT_STEP_AUTHOR
	case config.ARTICLE_EDITOR_STATUS_REJECTED_WITH_CORRECTION:
		draftStatus = config.ARTICLE_REVIEWER_STATUS_BACK_FOR_CORRECTION
		draftStep = config.DRAFT_STEP_AUTHOR
	case config.ARTICLE_EDITOR_STATUS_APPROVED:
		draftStatus = config.ARTICLE_STATUS_CONFIRMED
		draftStep = config.DRAFT_STEP_AUTHOR
	case config.ARTICLE_EDITOR_STATUS_APPROVED_WITH_CORRECTION:
		draftStatus = config.ARTICLE_REVIEWER_STATUS_BACK_FOR_CORRECTION
		draftStep = config.DRAFT_STEP_AUTHOR
	default:
		draftStatus = config.ARTICLE_STATUS_PENDING
	}

	_, err = h.services.ArticleService().UpdateArticle(
		c.Request.Context(),
		&pb.UpdateArticleReq{
			Status: draftStatus,
			Id:     articleId,
			Step:   draftStep,
		})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// DeleteArticleCheck godoc
// @ID delete_article_check
// @Router /journal/{journal-id}/draft/{draft-id}/check/{check-id} [DELETE]
// @Summary Delete Article reviewer
// @Description Get Article reviewer
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param draft-id path string true "draft-id"
// @Param check-id path string true "check-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteArticleCheck(c *gin.Context) {
	journalID := c.Param("journal-id")
	if !util.IsValidUUID(journalID) {
		h.handleResponse(c, http.InvalidArgument, "journal id is an invalid uuid")
		return
	}

	articleID := c.Param("draft-id")
	if !util.IsValidUUID(articleID) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	checkID := c.Param("check-id")
	if !util.IsValidUUID(checkID) {
		h.handleResponse(c, http.InvalidArgument, "check id is an invalid uuid")
		return
	}

	_, err := h.services.CheckerService().DeleteArticleChecker(
		c.Request.Context(),
		&pb.DeleteArticleCheckerReq{
			Id: checkID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}

// GetArticleCheckList godoc
// @ID get_article_check_list
// @Router /journal/{journal-id}/draft/{draft-id}/check [GET]
// @Summary Get article check list
// @Description  Get article check list
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param draft-id path string true "draft-id"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=pb.GetArticleCheckerListRes} "GetArticleReviewerListRes"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetArticleCheckList(c *gin.Context) {

	articleId := c.Param("draft-id")
	if !util.IsValidUUID(articleId) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

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

	resp, err := h.services.CheckerService().GetArticleCheckerList(
		c.Request.Context(),
		&pb.GetArticleCheckerListReq{
			Limit:     int32(limit),
			Offset:    int32(offset),
			Search:    c.DefaultQuery("search", ""),
			ArticleId: articleId,
			Type:      config.EDITOR,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetArticleCheckByID godoc
// @ID get_article_check_by_id
// @Router /journal/{journal-id}/draft/{draft-id}/check/{check-id} [GET]
// @Summary Get article check
// @Description  Get article check
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param draft-id path string true "draft-id"
// @Param check-id path string true "check-id"
// @Success 200 {object} http.Response{data=pb.GetArticleCheckerRes} "GetArticleReviewerRes"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetArticleCheckByID(c *gin.Context) {

	articleId := c.Param("draft-id")
	if !util.IsValidUUID(articleId) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	checkId := c.Param("check-id")
	if !util.IsValidUUID(checkId) {
		h.handleResponse(c, http.InvalidArgument, "check id is an invalid uuid")
		return
	}

	resp, err := h.services.CheckerService().GetArticleChecker(
		c.Request.Context(),
		&pb.GetArticleCheckerReq{
			Id: checkId,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
