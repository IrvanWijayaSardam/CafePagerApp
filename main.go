package main

import (
	"github.com/IrvanWijayaSardam/mynotesapp/config"
	"github.com/IrvanWijayaSardam/mynotesapp/controller"
	"github.com/IrvanWijayaSardam/mynotesapp/middleware"
	"github.com/IrvanWijayaSardam/mynotesapp/repository"
	"github.com/IrvanWijayaSardam/mynotesapp/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = config.SetupDatabaseConnection()
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	notesRepository repository.NotesRepository = repository.NewNotesRepository(db)
	pagerRepository repository.PagerRepository = repository.NewPagerRepository(db)
	jwtService      service.JWTService         = service.NewJWTService()
	userService     service.UserService        = service.NewUserService(userRepository)
	authService     service.AuthService        = service.NewAuthService(userRepository)
	noteService     service.NoteService        = service.NewNoteService(notesRepository)
	pagerService    service.PagerService       = service.NewPagerService(pagerRepository)
	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController  = controller.NewUserController(userService, jwtService)
	noteController  controller.NoteController  = controller.NewNotesController(noteService, jwtService)
	pagerController controller.PagerController = controller.NewPagerController(pagerService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	noteRoutes := r.Group("api/notes", middleware.AuthorizeJWT(jwtService))
	{
		noteRoutes.GET("/", noteController.All)
		noteRoutes.POST("/", noteController.Insert)
		noteRoutes.GET("/:id", noteController.FindById)
		noteRoutes.PUT("/:id", noteController.Update)
		noteRoutes.DELETE("/:id", noteController.Delete)
	}

	pagerRoutes := r.Group("api/pager", middleware.AuthorizeJWT(jwtService))
	{
		pagerRoutes.GET("/", pagerController.All)
		pagerRoutes.POST("/", pagerController.Insert)
		pagerRoutes.GET("/:id", pagerController.FindById)
		pagerRoutes.PUT("/:id", pagerController.Update)
		pagerRoutes.DELETE("/:id", pagerController.Delete)
	}

	pagerStatusRoutes := r.Group("/api/status")
	{
		pagerStatusRoutes.GET("/:id", pagerController.FindStatusById)
	}

	r.Run()
}
