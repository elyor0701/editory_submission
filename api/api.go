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

	r.GET("/ping", h.Ping)
	r.GET("/config", h.GetConfig)
	r.POST("/upload", h.Upload)

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

	r.POST("/login", h.Login)
	r.DELETE("/logout", h.Logout)
	r.PUT("/refresh", h.RefreshToken)
	r.POST("/has-access", h.HasAccess)

	// content
	r.POST("/journal", h.CreateJournal)
	r.GET("/journal", h.GetJournalList)
	r.GET("/journal/:journal-id", h.GetJournalByID)
	r.PUT("/journal", h.UpdateJournal)
	r.DELETE("/journal/:journal-id", h.DeleteJournal)

	r.POST("/article", h.CreateArticle)
	r.GET("/article", h.GetArticleList)
	r.GET("/article/:article-id", h.GetArticleByID)
	r.PUT("/article", h.UpdateArticle)
	r.DELETE("/article/:article-id", h.DeleteArticle)

	r.GET("/country", h.GetCountryList)
	r.GET("/city", h.GetCityList)

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
