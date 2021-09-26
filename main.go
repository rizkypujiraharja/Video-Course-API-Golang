package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/config"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/controller"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/entity"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/middleware"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/repo"
	"github.com/rizkypujiraharja/Video-Course-API-Golang/service"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	orderedLessonRepo repo.OrderedLessonRepository = repo.NewOrderedLessonRepo(db)

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
	lessonService    service.LessonService       = service.NewLessonService(lessonRepo, orderedLessonRepo)
	lessonController controller.LessonController = controller.NewLessonController(lessonService, jwtService)

	subLessonRepo       repo.SubLessonRepository       = repo.NewSubLessonRepo(db)
	subLessonService    service.SubLessonService       = service.NewSubLessonService(subLessonRepo)
	subLessonController controller.SubLessonController = controller.NewSubLessonController(subLessonService, jwtService)

	videoRepo       repo.VideoRepository       = repo.NewVideoRepo(db)
	videoService    service.VideoService       = service.NewVideoService(videoRepo)
	videoController controller.VideoController = controller.NewVideoController(videoService, jwtService)

	orderRepo       repo.OrderRepository       = repo.NewOrderRepo(db)
	orderDetailRepo repo.OrderDetailRepository = repo.NewOrderDetailRepo(db)
	orderService    service.OrderService       = service.NewOrderService(orderRepo, orderDetailRepo, orderedLessonRepo, lessonRepo)
	orderController controller.OrderController = controller.NewOrderController(orderService, jwtService)
)

func main() {
	db.AutoMigrate(
		&entity.User{},
		&entity.Category{},
		&entity.Lesson{},
		&entity.SubLesson{},
		&entity.Video{},
		&entity.Order{},
		&entity.OrderDetail{},
		&entity.OrderedLesson{},
	)
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := server.Group("api/profile", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/", userController.Profile)
		userRoutes.PUT("/", userController.Update)
	}

	server.GET("/api/categories", categoryController.All)
	categoryRoutes := server.Group("api/categories", middleware.AuthorizeJWT(jwtService, "admin"))
	{
		categoryRoutes.GET("/", categoryController.All)
		categoryRoutes.POST("/", categoryController.CreateCategory)
		categoryRoutes.PUT("/:id", categoryController.UpdateCategory)
		categoryRoutes.DELETE("/:id", categoryController.DeleteCategory)
	}

	publicLessonRoutes := server.Group("api/lessons")
	{
		publicLessonRoutes.GET("/", lessonController.All)
		publicLessonRoutes.GET("/:id", lessonController.FindOneLessonByID)
	}

	server.GET("/api/my-lessons", middleware.AuthorizeJWT(jwtService, "user"), lessonController.MyLesson)
	lessonRoutes := server.Group("api/lessons", middleware.AuthorizeJWT(jwtService, "admin"))
	{
		lessonRoutes.POST("/", lessonController.CreateLesson)
		lessonRoutes.PUT("/:id", lessonController.UpdateLesson)
		lessonRoutes.DELETE("/:id", lessonController.DeleteLesson)
	}

	subLessonRoutes := server.Group("api/sub-lessons", middleware.AuthorizeJWT(jwtService, "admin"))
	{
		subLessonRoutes.POST("/", subLessonController.CreateSubLesson)
		subLessonRoutes.GET("/:id", subLessonController.FindOneSubLessonByID)
		subLessonRoutes.PUT("/:id", subLessonController.UpdateSubLesson)
		subLessonRoutes.DELETE("/:id", subLessonController.DeleteSubLesson)
	}

	videoRoutes := server.Group("api/videos", middleware.AuthorizeJWT(jwtService, "admin"))
	{
		videoRoutes.POST("/", videoController.CreateVideo)
		videoRoutes.GET("/:id", videoController.FindOneVideoByID)
		videoRoutes.PUT("/:id", videoController.UpdateVideo)
		videoRoutes.DELETE("/:id", videoController.DeleteVideo)
	}

	server.GET("/api/my-orders", middleware.AuthorizeJWT(jwtService, "user"), orderController.MyOrder)
	server.POST("/api/orders", middleware.AuthorizeJWT(jwtService, "user"), orderController.CreateOrder)
	server.GET("/api/orders/:id", middleware.AuthorizeJWT(jwtService, "user", "admin"), orderController.FindOneOrderByID)
	adminOrderRoutes := server.Group("api/orders", middleware.AuthorizeJWT(jwtService, "admin"))
	{
		adminOrderRoutes.GET("/", orderController.All)
		adminOrderRoutes.PUT("/paid/:id", orderController.UpdatePaidOrder)
		adminOrderRoutes.PUT("/unpaid/:id", orderController.UpdateUnpaidOrder)
	}

	server.Run()
}
