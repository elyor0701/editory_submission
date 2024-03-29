package api

import (
	"editory_submission/api/docs"
	"editory_submission/api/handlers"
	"editory_submission/config"
	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetUpRouter godoc
// @description This is a api gateway
// @termsOfService https://udevs.io
func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = cfg.Version
	// docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}

	r.Use(customCORSMiddleware())

	{
		// general
		r.GET("/country", h.GetCountryList)
		r.GET("/university", h.GetUniversityList)
		r.GET("/city", h.GetCityList)
		r.GET("/subject", h.GetSubjectList)
		r.GET("/keyword", h.GetKeywordList)
		r.GET("/ping", h.Ping)
		r.GET("/config", h.GetConfig)
		r.POST("/upload", h.Upload)

		r.POST("/register/email", h.RegistrationEmail)
		r.POST("/register/resend", h.ResendVerificationMessage)
		r.POST("/register/verify", h.EmailVerification)
		r.POST("/register/details", h.RegisterDetail)
		r.POST("/login", h.Login)
		r.DELETE("/logout", h.Logout)
		r.PUT("/refresh", h.RefreshToken)
		r.POST("/has-access", h.HasAccess)

		r.PUT("/profile", h.UpdateProfile)
		r.GET("/profile/:profile-id", h.GetProfileByID)

		r.GET("/article")
		r.GET("/article/:article-id")
		r.GET("/general/journal", h.GetGeneralJournalList)
		r.GET("/general/journal/:journal-id", h.GetGeneralJournalByID)
		r.GET("/general/journal/:journal-id/author", h.GetGeneralJournalAuthorList)
		r.GET("/general/journal/:journal-id/author/:author-id", h.GetGeneralJournalAuthorByID)
	}

	// auth
	//r.POST("/user", h.CreateUser)
	//r.GET("/user", h.GetUserList)
	//r.GET("/user/:user-id", h.GetUserByID)
	//r.PUT("/user", h.UpdateUser)
	//r.DELETE("/user/:user-id", h.DeleteUser)
	//r.PUT("/user/reset-password", h.ResetPassword)
	//r.POST("/user/send-message", h.SendMessageToUserEmail)
	//r.POST("/send-verification-message", h.SendVerificationMessage)
	//r.PUT("/verification", h.EmailVerification)

	r.POST("/user/:user-id/draft", h.CreateUserDraft)
	r.GET("/user/:user-id/draft", h.GetUserDraftList)
	r.GET("/user/:user-id/draft/:draft-id", h.GetUserDraftByID)
	r.PUT("/user/:user-id/draft", h.UpdateUserDraft)
	r.DELETE("/user/:user-id/draft/:draft-id", h.DeleteUserDraft)

	r.POST("/user/:user-id/draft/:draft-id/file", h.AddDraftFile)
	r.DELETE("/user/:user-id/draft/:draft-id/file/:file-id", h.DeleteDraftFiles)

	r.POST("/user/:user-id/draft/:draft-id/coauthor", h.AddDraftCoAuthors)
	r.DELETE("/user/:user-id/draft/:draft-id/coauthor/:coauthor-id", h.DeleteDraftCoAuthors)

	r.GET("/user/:user-id/draft/:draft-id/check", h.GetUserDraftCheckList)
	r.GET("/user/:user-id/draft/:draft-id/check/:check-id", h.GetUserDraftCheckByID)

	r.GET("/user/:user-id/review", h.GetUserReviewList)
	r.GET("/user/:user-id/review/:review-id", h.GetUserReviewByID)
	r.PUT("/user/:user-id/review", h.UpdateUserReview)

	{
		// journal
		journal := r.Group("/journal")
		journal.GET("/:journal-id", h.GetJournalByID)
		journal.PUT("", h.UpdateJournal)
		journal.DELETE("/:journal-id", h.DeleteJournal) // @TODO

		// journal author
		journal.POST("/:journal-id/author", h.CreateJournalAuthor)
		journal.GET("/:journal-id/author", h.GetJournalAuthorList)
		journal.GET("/:journal-id/author/:author-id", h.GetJournalArticleByID)
		journal.PUT("/:journal-id/author", h.UpdateJournalAuthor)
		journal.DELETE("/:journal-id/author/:author-id", h.DeleteJournalAuthor)

		//journal.POST("/:journal-id/draft", h.CreateJournalArticle)
		journal.GET("/:journal-id/draft", h.GetJournalDraftList)
		journal.GET("/:journal-id/draft/:draft-id", h.GetJournalDraftByID)
		journal.PUT("/:journal-id/draft", h.UpdateJournalDraft)
		//journal.DELETE("/:journal-id/draft/:draft-id", h.DeleteJournalArticle)

		journal.POST("/:journal-id/draft/:draft-id/check", h.CreateArticleCheck)
		journal.PUT("/:journal-id/draft/:draft-id/check", h.UpdateArticleCheck)
		journal.GET("/:journal-id/draft/:draft-id/check", h.GetArticleCheckList)
		journal.GET("/:journal-id/draft/:draft-id/check/:check-id", h.GetArticleCheckByID)
		journal.DELETE("/:journal-id/draft/:draft-id/check/:check-id", h.DeleteArticleCheck)

		journal.POST("/:journal-id/draft/:draft-id/reviewer", h.CreateArticleReviewer)
		journal.DELETE("/:journal-id/draft/:draft-id/reviewer/:reviewer-id", h.DeleteArticleReviewer)
		journal.GET("/:journal-id/draft/:draft-id/review", h.GetArticleReviewList)
		journal.GET("/:journal-id/draft/:draft-id/review/:review-id", h.GetArticleReviewByID)

		journal.POST("/:journal-id/article", h.CreateJournalArticle)
		journal.GET("/:journal-id/article", h.GetJournalArticleList)
		journal.GET("/:journal-id/article/:article-id", h.GetJournalArticleByID)
		journal.PUT("/:journal-id/article", h.UpdateJournalArticle)
		journal.DELETE("/:journal-id/article/:article-id", h.DeleteJournalArticle)

		journal.POST("/:journal-id/edition", h.CreateEdition)
		journal.GET("/:journal-id/edition", h.GetEditionList)
		journal.GET("/:journal-id/edition/:edition-id", h.GetEditionByID)
		journal.PUT("/:journal-id/edition", h.UpdateEdition)
		journal.DELETE("/:journal-id/edition/:edition-id", h.DeleteEdition)

		journal.GET("/:journal-id/user", h.GetJournalUserList)
		journal.POST("/:journal-id/user", h.CreateJournalUser)
		journal.GET("/:journal-id/user/:user-id", h.GetJournalUserByID)
		journal.PUT("/:journal-id/user", h.UpdateJournalUser)
		journal.DELETE("/:journal-id/user/:user-id", h.DeleteJournalUser)

		journal.GET("/author")
	}
	{
		// admin
		admin := r.Group("/admin")
		admin.POST("/journal", h.CreateAdminJournal)
		admin.GET("/journal", h.GetAdminJournalList)
		admin.GET("/journal/:journal-id", h.GetAdminJournalByID)
		admin.PUT("/journal", h.UpdateAdminJournal)
		admin.DELETE("/journal/:journal-id", h.DeleteAdminJournal)

		admin.GET("/draft", h.GetAdminDraftList)
		admin.GET("/draft/:draft-id", h.GetAdminDraftByID)

		admin.GET("/article")
		admin.GET("/article/:article-id")

		admin.POST("/university", h.CreateAdminUniversity)
		admin.GET("/university", h.GetAdminUniversityList)
		admin.GET("/university/:university-id", h.GetAdminUniversityByID)
		admin.PUT("/university", h.UpdateAdminUniversity)
		admin.DELETE("/university/:university-id", h.DeleteAdminUniversity)

		admin.POST("/subject", h.CreateAdminSubject)
		admin.GET("/subject", h.GetAdminSubjectList)
		admin.GET("/subject/:subject-id", h.GetAdminSubjectByID)
		admin.PUT("/subject", h.UpdateAdminSubject)
		admin.DELETE("/subject/:subject-id", h.DeleteAdminSubject)

		admin.POST("/editor", h.CreateEditor)
		admin.PUT("/editor", h.UpdateEditor)
		admin.GET("/editor", h.GetEditorList)
		admin.GET("/editor/:editor-id", h.GetEditorByID)
		admin.DELETE("/editor/:editor-id", h.DeleteEditor)

		admin.GET("/author", h.GetAdminAuthorList)
		admin.GET("/author/:author-id", h.GetAdminAuthorByID)
		admin.POST("/author", h.CreateAdminAuthor)
		admin.PUT("/author", h.UpdateAdminAuthor)

		admin.POST("/keyword", h.CreateAdminKeyword)
		admin.GET("/keyword", h.GetAdminKeywordList)
		admin.GET("/keyword/:keyword-id", h.GetAdminKeywordByID)
		admin.PUT("/keyword", h.UpdateAdminKeyword)
		admin.DELETE("/keyword/:keyword-id", h.DeleteAdminKeyword)

		admin.POST("/email/template", h.CreateAdminEmailTmp)
		admin.GET("/email/template", h.GetAdminEmailTmpList)
		admin.GET("/email/template/:template-id", h.GetAdminEmailTmpByID)
		admin.PUT("/email/template", h.UpdateAdminEmailTmp)
		admin.DELETE("/email/template/:template-id", h.DeleteAdminEmailTmp)
	}

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Role")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
