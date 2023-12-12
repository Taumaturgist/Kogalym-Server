package router

import (
	"Kogalym/backend/app/middleware"
	"Kogalym/backend/business"
	"Kogalym/backend/business/students"
	"Kogalym/backend/config"
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func Init(init *config.Initialization, embeddedFiles embed.FS) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(SetupSession())
	router.Use(middleware.CsrfCheckMiddleware())
	//router.Use(SetupCsrfTokens())

	createRoutesForStaticFiles(router, embeddedFiles)

	//api := router.Group("/api")
	//{
	//	user := api.Group("/user")
	//	user.GET("", init.UserCtrl.GetAllUserData)
	//	user.POST("", init.UserCtrl.AddUserData)
	//	user.GET("/:userID", init.UserCtrl.GetUserById)
	//	user.PUT("/:userID", init.UserCtrl.UpdateUserData)
	//	user.DELETE("/:userID", init.UserCtrl.DeleteUser)
	//}

	webRoutes(router, init)

	return router
}

func webRoutes(router *gin.Engine, init *config.Initialization) {
	web := router.Group("")

	// Generate CSRF token
	//web.Use(middleware.CsrfCheckMiddleware())

	// Auth
	{
		web.GET("/login", init.AuthCtrl.LoginPage)
		web.POST("/login", init.AuthCtrl.WebLogin)
	}

	authenticatedWeb := web.Group("")
	authenticatedWeb.Use(middleware.WebAuthMiddleware())
	{
		authenticatedWeb.GET("", business.Home)
		authenticatedWeb.POST("/logout", init.AuthCtrl.WebLogout)

		authenticatedWeb.GET("/groups", init.GroupCtrl.IndexGroups)
		authenticatedWeb.GET("/students", students.IndexStudents)
	}

	//api := authenticatedWeb.Group("/api")
	//{
	//api.GET("/groups", groups.GetGroups)
	//api.POST("/groups", groups.CreateGroup)
	//api.PUT("/groups/:id", groups.UpdateGroup)
	//
	//api.GET("/students", students.GetStudents)
	//api.POST("/students", students.CreateStudent)
	//api.PUT("/students/:id", students.UpdateStudent)
	//}
}

func createRoutesForStaticFiles(router *gin.Engine, embeddedFiles embed.FS) {
	templates := template.Must(template.New("").ParseFS(embeddedFiles, "templates/*.html"))
	router.SetHTMLTemplate(templates)
	router.StaticFS("/public", http.Dir("templates/public"))
}
