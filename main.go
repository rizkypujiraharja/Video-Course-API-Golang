package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/config"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/controller"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/middleware"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/service"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepo)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService, userService)

	userRepo       repo.UserRepository       = repo.NewUserRepo(db)
	authService    service.AuthService       = service.NewAuthService(userRepo)
	userController controller.UserController = controller.NewUserController(userService, jwtService)

	categoryRepo       repo.CategoryRepository       = repo.NewCategoryRepo(db)
	categoryService    service.CategoryService       = service.NewCategoryService(categoryRepo)
	categoryController controller.CategoryController = controller.NewCategoryController(categoryService, jwtService)

	lessonRepo       repo.LessonRepository       = repo.NewLessonRepo(db)
	lessonService    service.LessonService       = service.NewLessonService(lessonRepo)
	lessonController controller.LessonController = controller.NewLessonController(lessonService, jwtService)

	subLessonRepo       repo.SubLessonRepository       = repo.NewSubLessonRepo(db)
	subLessonService    service.SubLessonService       = service.NewSubLessonService(subLessonRepo)
	subLessonController controller.SubLessonController = controller.NewSubLessonController(subLessonService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	categoryRoutes := server.Group("api/category", middleware.AuthorizeJWT(jwtService))
	{
		categoryRoutes.GET("/", categoryController.All)
		categoryRoutes.POST("/", categoryController.CreateCategory)
		categoryRoutes.PUT("/:id", categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}

	lessonRoutes := server.Group("api/lesson", middleware.AuthorizeJWT(jwtService))
	{
		lessonRoutes.GET("/", lessonController.All)
		lessonRoutes.POST("/", lessonController.CreateLesson)
		lessonRoutes.GET("/:id", lessonController.FindOneLessonByID)
		lessonRoutes.PUT("/:id", lessonController.UpdateLesson)
		lessonRoutes.DELETE("/:id", lessonController.DeleteLesson)
	}

	subLessonRoutes := server.Group("api/sub-lesson", middleware.AuthorizeJWT(jwtService))
	{
		subLessonRoutes.POST("/", subLessonController.CreateSubLesson)
		subLessonRoutes.GET("/:id", subLessonController.FindOneSubLessonByID)
		subLessonRoutes.PUT("/:id", subLessonController.UpdateSubLesson)
		subLessonRoutes.DELETE("/:id", subLessonController.DeleteSubLesson)
	}

	server.Run()
}
