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
		r.GET("/ping", h.Ping)
		r.GET("/config", h.GetConfig)
		r.POST("/upload", h.Upload)

		r.POST("/login", h.Login)
		r.DELETE("/logout", h.Logout)
		r.PUT("/refresh", h.RefreshToken)
		r.POST("/has-access", h.HasAccess)
	}

	// auth
	r.POST("/user", h.CreateUser)
	r.GET("/user", h.GetUserList)
	r.GET("/user/:user-id", h.GetUserByID)
	r.PUT("/user", h.UpdateUser)
	r.DELETE("/user/:user-id", h.DeleteUser)
	//r.PUT("/user/reset-password", h.ResetPassword)
	//r.POST("/user/send-message", h.SendMessageToUserEmail)
	r.POST("/send-verification-message", h.SendVerificationMessage)
	r.PUT("/verification", h.EmailVerification)

	{
		// journal
		journal := r.Group("/journal")
		journal.POST("", h.CreateJournal)
		journal.GET("", h.GetJournalList)
		journal.GET("/:journal-id", h.GetJournalByID)
		journal.PUT("", h.UpdateJournal)
		journal.DELETE("/:journal-id", h.DeleteJournal)

		journal.POST("/:journal-id/article", h.CreateArticle)
		journal.GET("/:journal-id/article", h.GetArticleList)
		journal.GET("/:journal-id/article/:article-id", h.GetArticleByID)
		journal.PUT("/:journal-id/article", h.UpdateArticle)
		journal.DELETE("/:journal-id/article/:article-id", h.DeleteArticle)

		journal.POST("/:journal-id/edition", h.CreateEdition)
		journal.GET("/:journal-id/edition", h.GetEditionList)
		journal.GET("/:journal-id/edition/:edition-id", h.GetEditionByID)
		journal.PUT("/:journal-id/edition", h.UpdateEdition)
		journal.DELETE("/:journal-id/edition/:edition-id", h.DeleteEdition)

		journal.GET("/user")  //@TODO
		journal.POST("/user") //@TODO
	}
	{
		// admin
		admin := r.Group("/admin")
		admin.POST("/journal")
		admin.GET("/journal")
		admin.GET("/journal/:journal-id")
		admin.PUT("/journal")
		admin.DELETE("/journal/:journal-id")

		admin.POST("/article")
		admin.GET("/article")
		admin.GET("/article/:article-id")
		admin.PUT("/article")
		admin.DELETE("/article/:article-id")

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

		admin.POST("/user")
		admin.PUT("/user")
		admin.GET("/user")
		admin.GET("/user/:user-id")
		admin.DELETE("/user/:user-id")
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
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
