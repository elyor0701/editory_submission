package handlers

import (
	"editory_submission/api/http"
	"editory_submission/api/models"
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
	pb "editory_submission/genproto/submission_service"
	"editory_submission/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
)

// CreateArticleReviewer godoc
// @ID create_article_reviewer
// @Router /journal/{journal-id}/draft/{draft-id}/reviewer [POST]
// @Summary Create Article Reviewer
// @Description Create Article Reviewer
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal Id"
// @Param draft-id path string true "draft Id"
// @Param article body models.CreateArticleReviewerReq true "CreateArticleRequestBody"
// @Success 201 {object} http.Response{data=pb.CreateArticleCheckerRes} "CreateArticleReviewerRes"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateArticleReviewer(c *gin.Context) {
	var reviewer models.CreateArticleReviewerReq

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

	err := c.ShouldBindJSON(&reviewer)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	userPb, err := h.services.UserService().GetUser(
		c.Request.Context(),
		&auth_service.GetUserReq{
			Email: reviewer.Email,
		},
	)
	if err != nil {
		if util.IsErrNoRows(err) {
			userPb, err = h.services.UserService().CreateUser(
				c.Request.Context(),
				&auth_service.User{
					FirstName: reviewer.FirstName,
					Email:     reviewer.Email,
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
		RoleType:  config.REVIEWER,
		JournalId: journalId,
		UserId:    userPb.GetId(),
	}

	_, err = h.services.RoleService().CreateRole(
		c.Request.Context(),
		&role,
	)
	if err != nil {
		if !util.IsErrDuplicateKey(err) {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
	}

	resp, err := h.services.CheckerService().CreateArticleChecker(
		c.Request.Context(),
		&pb.CreateArticleCheckerReq{
			CheckerId: userPb.Id,
			ArticleId: articleId,
			Status:    "NEW",
			Type:      config.REVIEWER,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	_, err = h.services.ArticleService().UpdateArticle(
		c.Request.Context(),
		&pb.UpdateArticleReq{
			Id:             articleId,
			Step:           "REVIEWER",
			ReviewerStatus: "PENDING",
			EditorStatus:   "PENDING",
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	_, err = h.services.ArticleService().UpdateArticle(
		c.Request.Context(),
		&pb.UpdateArticleReq{
			Id:   articleId,
			Step: config.DRAFT_STEP_REVIEWER,
		})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// DeleteArticleReviewer godoc
// @ID delete_article_reviewer
// @Router /journal/{journal-id}/draft/{draft-id}/reviewer/{reviewer-id} [DELETE]
// @Summary Delete Article reviewer
// @Description Get Article reviewer
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param draft-id path string true "draft-id"
// @Param reviewer-id path string true "reviewer-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteArticleReviewer(c *gin.Context) {
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

	reviewerID := c.Param("reviewer-id")
	if !util.IsValidUUID(reviewerID) {
		h.handleResponse(c, http.InvalidArgument, "reviewer id is an invalid uuid")
		return
	}

	_, err := h.services.CheckerService().DeleteArticleChecker(
		c.Request.Context(),
		&pb.DeleteArticleCheckerReq{
			Id: reviewerID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}

// GetArticleReviewList godoc
// @ID get_article_review_list
// @Router /journal/{journal-id}/draft/{draft-id}/review [GET]
// @Summary Get article review List
// @Description  Get article review List
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
func (h *Handler) GetArticleReviewList(c *gin.Context) {

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
			Type:      config.REVIEWER,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetArticleReviewByID godoc
// @ID get_article_review_by_id
// @Router /journal/{journal-id}/draft/{draft-id}/review/{review-id} [GET]
// @Summary Get article review
// @Description  Get article review
// @Tags Journal
// @Accept json
// @Produce json
// @Param journal-id path string true "journal-id"
// @Param draft-id path string true "draft-id"
// @Param review-id path string true "review-id"
// @Success 200 {object} http.Response{data=pb.GetArticleCheckerRes} "GetArticleReviewerRes"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetArticleReviewByID(c *gin.Context) {

	articleId := c.Param("draft-id")
	if !util.IsValidUUID(articleId) {
		h.handleResponse(c, http.InvalidArgument, "article id is an invalid uuid")
		return
	}

	reviewId := c.Param("review-id")
	if !util.IsValidUUID(reviewId) {
		h.handleResponse(c, http.InvalidArgument, "review id is an invalid uuid")
		return
	}

	resp, err := h.services.CheckerService().GetArticleChecker(
		c.Request.Context(),
		&pb.GetArticleCheckerReq{
			Id: reviewId,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetUserReviewList godoc
// @ID get_user_review_list
// @Router /user/{user-id}/review [GET]
// @Summary Get user review List
// @Description  Get user review List
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user-id"
// @Param status query string false "status"
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=pb.GetArticleCheckerListRes} "GetArticleReviewerListRes"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserReviewList(c *gin.Context) {

	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
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
			CheckerId: userId,
			Type:      config.REVIEWER,
			Status:    c.Query("status"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetUserReviewByID godoc
// @ID get_user_review_by_id
// @Router /user/{user-id}/review/{review-id} [GET]
// @Summary Get user review
// @Description  Get user review
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user-id"
// @Param review-id path string true "review-id"
// @Success 200 {object} http.Response{data=pb.GetArticleCheckerRes} "GetArticleReviewerRes"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserReviewByID(c *gin.Context) {

	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	reviewId := c.Param("review-id")
	if !util.IsValidUUID(reviewId) {
		h.handleResponse(c, http.InvalidArgument, "review id is an invalid uuid")
		return
	}

	resp, err := h.services.CheckerService().GetArticleChecker(
		c.Request.Context(),
		&pb.GetArticleCheckerReq{
			Id: reviewId,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	files, err := h.services.ArticleService().GetFiles(
		c.Request.Context(),
		&pb.GetFilesReq{
			ArticleId: resp.ArticleId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	resp.ArticleIdData.Files = files.GetFiles()

	h.handleResponse(c, http.OK, resp)
}

// UpdateUserReview godoc
// @ID update_user_review
// @Router /user/{user-id}/review [PUT]
// @Summary Update User Review
// @Description Update User Review
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user-id"
// @Param review body models.UpdateUserReviewReq true "UpdateUserReviewReq"
// @Success 200 {object} http.Response{data=pb.UpdateArticleCheckerRes} "UpdateArticleReviewerRes"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUserReview(c *gin.Context) {
	var (
		review       models.UpdateUserReviewReq
		fileComments []*pb.FileComment
	)

	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&review)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	for _, val := range review.FileComments {
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
			Id:       review.Id,
			Status:   review.Status,
			Comment:  review.Comment,
			Comments: fileComments,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
