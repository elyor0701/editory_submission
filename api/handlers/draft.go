package handlers

import (
	"editory_submission/api/http"
	"editory_submission/api/models"
	"editory_submission/config"
	"editory_submission/genproto/auth_service"
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

	if articlePB.Status == "DRAFT" {
		articlePB.Step = "AUTHOR"
	} else if articlePB.Status == "NEW" {
		articlePB.Step = "EDITOR"
	}

	resp, err := h.services.ArticleService().CreateArticle(
		c.Request.Context(),
		&articlePB,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	for _, val := range article.Files {
		_, err := h.services.ArticleService().AddFiles(
			c.Request.Context(),
			&submission_service.AddFilesReq{
				Url:       val.Url,
				Type:      val.Type,
				ArticleId: resp.Id,
			},
		)

		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
	}

	for _, val := range article.Coauthors {
		userPb, err := h.services.UserService().GetUser(
			c.Request.Context(),
			&auth_service.GetUserReq{
				Email: val.Email,
			},
		)
		if err != nil {
			if util.IsErrNoRows(err) {
				userPb, err = h.services.UserService().CreateUser(
					c.Request.Context(),
					&auth_service.User{
						FirstName:    val.FirstName,
						LastName:     val.LastName,
						Email:        val.Email,
						UniversityId: val.UniversityId,
						CountryId:    val.CountryId,
						Password:     config.DEFAULT_PASSWORD,
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
			RoleType: config.AUTHOR,
			UserId:   userPb.GetId(),
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

		_, err = h.services.ArticleService().AddCoAuthor(
			c.Request.Context(),
			&submission_service.AddCoAuthorReq{
				ArticleId: resp.Id,
				UserId:    userPb.Id,
			},
		)

		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
	}

	h.handleResponse(c, http.Created, resp)
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
// @Param draft body models.UpdateUserDraftReq true "UpdateDraftRequestBody"
// @Success 200 {object} http.Response{data=submission_service.UpdateArticleRes} "Draft data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUserDraft(c *gin.Context) {
	var (
		articlePB submission_service.UpdateArticleReq
		article   models.UpdateUserDraftReq
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

	if articlePB.Status == "DRAFT" {
		articlePB.Step = "AUTHOR"
	} else if articlePB.Status == "NEW" {
		articlePB.Step = "EDITOR"
	}

	resp, err := h.services.ArticleService().UpdateArticle(
		c.Request.Context(),
		&articlePB,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	for _, val := range article.Files {
		_, err := h.services.ArticleService().AddFiles(
			c.Request.Context(),
			&submission_service.AddFilesReq{
				Url:       val.Url,
				Type:      val.Type,
				ArticleId: article.Id,
			},
		)

		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
	}

	h.handleResponse(c, http.OK, resp)
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

	files, err := h.services.ArticleService().GetFiles(
		c.Request.Context(),
		&submission_service.GetFilesReq{
			ArticleId: articleID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	resp.Files = files.GetFiles()

	coauthor, err := h.services.ArticleService().GetCoAuthors(
		c.Request.Context(),
		&submission_service.GetCoAuthorsReq{
			DraftId: articleID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	resp.Coauthors = coauthor.GetCoauthors()

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
// @Success 200 {object} http.Response{data=submission_service.GetArticleRes} "DraftBody"
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

	files, err := h.services.ArticleService().GetFiles(
		c.Request.Context(),
		&submission_service.GetFilesReq{
			ArticleId: articleID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	resp.Files = files.GetFiles()

	coauthor, err := h.services.ArticleService().GetCoAuthors(
		c.Request.Context(),
		&submission_service.GetCoAuthorsReq{
			DraftId: articleID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	resp.Coauthors = coauthor.GetCoauthors()

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
// @Param draft body models.UpdateJournalDraftReq true "UpdateDraftRequestBody"
// @Success 200 {object} http.Response{data=submission_service.UpdateArticleRes} "Draft data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateJournalDraft(c *gin.Context) {
	var (
		article      models.UpdateJournalDraftReq
		fileComments []*submission_service.FileComment
	)

	journalId := c.Param("journal-id")
	if !util.IsValidUUID(journalId) {
		h.handleResponse(c, http.BadRequest, errors.New("journal id is not valid"))
		return
	}

	err := c.ShouldBindJSON(&article)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	checker, err := h.services.CheckerService().GetArticleCheckerList(
		c.Request.Context(),
		&submission_service.GetArticleCheckerListReq{
			CheckerId: article.EditorId,
			Type:      "EDITOR",
			ArticleId: article.Id,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	for _, val := range article.FileComments {
		fileComments = append(fileComments, &submission_service.FileComment{
			Id:      val.Id,
			Type:    val.Type,
			FileId:  val.FileId,
			Comment: val.Comment,
		})
	}

	if len(checker.GetArticleCheckers()) > 0 {
		_, err = h.services.CheckerService().UpdateArticleChecker(
			c.Request.Context(),
			&submission_service.UpdateArticleCheckerReq{
				Id:       checker.ArticleCheckers[0].Id,
				Comment:  article.Comment,
				Status:   article.CheckerStatus,
				Comments: fileComments,
			},
		)
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
	} else {
		_, err = h.services.CheckerService().CreateArticleChecker(
			c.Request.Context(),
			&submission_service.CreateArticleCheckerReq{
				CheckerId: article.EditorId,
				ArticleId: article.Id,
				Comment:   article.Comment,
				Type:      "EDITOR",
				Status:    article.CheckerStatus,
				Comments:  fileComments,
			},
		)
		if err != nil {
			h.handleResponse(c, http.GRPCError, err.Error())
			return
		}
	}

	resp, err := h.services.ArticleService().UpdateArticle(
		c.Request.Context(),
		&submission_service.UpdateArticleReq{
			Status: article.Status,
			Id:     article.Id,
		},
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

	files, err := h.services.ArticleService().GetFiles(
		c.Request.Context(),
		&submission_service.GetFilesReq{
			ArticleId: articleID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	resp.Files = files.GetFiles()

	h.handleResponse(c, http.OK, resp)
}

// AddDraftFile godoc
// @ID add_draft_file
// @Router /user/{user-id}/draft/{draft-id}/file [POST]
// @Summary Add Draft File
// @Description Add Draft File
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft-id path string true "draft-id"
// @Param file body models.AddFileReq true "AddFileReq"
// @Success 201 {object} http.Response{data=submission_service.AddFilesRes} "Draft data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddDraftFile(c *gin.Context) {
	var (
		file   models.AddFileReq
		filePB submission_service.AddFilesReq
	)

	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		err := errors.New("user id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	draftId := c.Param("draft-id")
	if !util.IsValidUUID(draftId) {
		err := errors.New("draft id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err := c.ShouldBindJSON(&file)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	fileJSON, err := json.Marshal(file)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = json.Unmarshal(fileJSON, &filePB)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	filePB.ArticleId = draftId

	resp, err := h.services.ArticleService().AddFiles(
		c.Request.Context(),
		&filePB,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetDraftFiles godoc
// @ID get_draft_file
// @Router /user/{user-id}/draft/{draft-id}/file [GET]
// @Summary Get Draft File
// @Description Get Draft File
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft-id path string true "draft-id"
// @Success 201 {object} http.Response{data=submission_service.GetFilesRes} "Draft data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetDraftFiles(c *gin.Context) {

	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		err := errors.New("user id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	draftId := c.Param("draft-id")
	if !util.IsValidUUID(draftId) {
		err := errors.New("draft id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ArticleService().GetFiles(
		c.Request.Context(),
		&submission_service.GetFilesReq{
			ArticleId: draftId,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// DeleteDraftFiles godoc
// @ID delete_draft_file
// @Router /user/{user-id}/draft/{draft-id}/file/{file-id} [DELETE]
// @Summary Delete Draft File
// @Description Delete Draft File
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft-id path string true "draft-id"
// @Param file-id path string true "file-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteDraftFiles(c *gin.Context) {
	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		err := errors.New("user id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	draftId := c.Param("draft-id")
	if !util.IsValidUUID(draftId) {
		err := errors.New("draft id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	fileId := c.Param("file-id")
	if !util.IsValidUUID(fileId) {
		err := errors.New("file id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err := h.services.ArticleService().DeleteFiles(
		c.Request.Context(),
		&submission_service.DeleteFilesReq{
			Ids: fileId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}

// AddDraftCoAuthors godoc
// @ID add_draft_coauthor
// @Router /user/{user-id}/draft/{draft-id}/coauthor [POST]
// @Summary Add Draft Coauthor
// @Description Add Draft Coauthor
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft-id path string true "draft-id"
// @Param author body models.AddCoAuthorReq true "AddCoAuthorReq"
// @Success 201 {object} http.Response{data=submission_service.AddCoAuthorRes} "Draft data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddDraftCoAuthors(c *gin.Context) {
	var (
		author   models.AddCoAuthorReq
		authorPb submission_service.AddCoAuthorReq
	)

	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		err := errors.New("user id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	draftId := c.Param("draft-id")
	if !util.IsValidUUID(draftId) {
		err := errors.New("draft id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err := c.ShouldBindJSON(&author)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	authorJSON, err := json.Marshal(author)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = json.Unmarshal(authorJSON, &authorPb)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	authorPb.ArticleId = draftId

	userPb, err := h.services.UserService().GetUser(
		c.Request.Context(),
		&auth_service.GetUserReq{
			Email: author.Email,
		},
	)
	if err != nil {
		if util.IsErrNoRows(err) {
			userPb, err = h.services.UserService().CreateUser(
				c.Request.Context(),
				&auth_service.User{
					FirstName:    author.FirstName,
					LastName:     author.LastName,
					Email:        author.Email,
					UniversityId: author.UniversityId,
					CountryId:    author.CountryId,
					Password:     config.DEFAULT_PASSWORD,
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
		RoleType: config.AUTHOR,
		UserId:   userPb.GetId(),
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

	authorPb.UserId = userPb.Id

	resp, err := h.services.ArticleService().AddCoAuthor(
		c.Request.Context(),
		&authorPb,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetDraftCoAuthors godoc
// @ID get_draft_coauthor
// @Router /user/{user-id}/draft/{draft-id}/coauthor [GET]
// @Summary Get Draft coauthor
// @Description Get Draft coauthor
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft-id path string true "draft-id"
// @Success 201 {object} http.Response{data=submission_service.GetCoAuthorsRes} "Draft data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetDraftCoAuthors(c *gin.Context) {

	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		err := errors.New("user id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	draftId := c.Param("draft-id")
	if !util.IsValidUUID(draftId) {
		err := errors.New("draft id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ArticleService().GetCoAuthors(
		c.Request.Context(),
		&submission_service.GetCoAuthorsReq{
			DraftId: draftId,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// DeleteDraftCoAuthors godoc
// @ID delete_draft_coauthor
// @Router /user/{user-id}/draft/{draft-id}/coauthor/{coauthor-id} [DELETE]
// @Summary Delete Draft coauthor
// @Description Delete Draft coauthor
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user Id"
// @Param draft-id path string true "draft-id"
// @Param coauthor-id path string true "coauthor-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteDraftCoAuthors(c *gin.Context) {
	userId := c.Param("user-id")
	if !util.IsValidUUID(userId) {
		err := errors.New("user id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	draftId := c.Param("draft-id")
	if !util.IsValidUUID(draftId) {
		err := errors.New("draft id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	coauthorId := c.Param("coauthor-id")
	if !util.IsValidUUID(coauthorId) {
		err := errors.New("coauthor id is not valid")
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	_, err := h.services.ArticleService().DeleteCoAuthor(
		c.Request.Context(),
		&submission_service.DeleteCoAuthorReq{
			Ids: coauthorId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, "")
}
