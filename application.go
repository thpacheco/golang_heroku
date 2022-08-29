package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thpacheco/golang_heroku/config"
	v1 "github.com/thpacheco/golang_heroku/handler/v1"
	"github.com/thpacheco/golang_heroku/middleware"
	"github.com/thpacheco/golang_heroku/repo"
	"github.com/thpacheco/golang_heroku/service"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                = config.SetupDatabaseConnection()
	userRepo        repo.UserRepository     = repo.NewUserRepo(db)
	productRepo     repo.ProductRepository  = repo.NewProductRepo(db)
	custumerRepo    repo.CustumerRepository = repo.NewCustumerRepo(db)
	authService     service.AuthService     = service.NewAuthService(userRepo)
	jwtService      service.JWTService      = service.NewJWTService()
	userService     service.UserService     = service.NewUserService(userRepo)
	productService  service.ProductService  = service.NewProductService(productRepo)
	custumerService service.CustumerService = service.NewCustumerService(custumerRepo)
	authHandler     v1.AuthHandler          = v1.NewAuthHandler(authService, jwtService, userService)
	userHandler     v1.UserHandler          = v1.NewUserHandler(userService, jwtService)
	productHandler  v1.ProductHandler       = v1.NewProductHandler(productService, jwtService)
	custumerHandler v1.CustumerHandler      = v1.NewCustumerHandler(custumerService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()
	server.Use(CORSMiddleware())

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	productRoutes := server.Group("api/product", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.GET("/", productHandler.All)
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.FindOneProductByID)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}
	custumerRoutes := server.Group("api/custumer", middleware.AuthorizeJWT(jwtService))
	{
		custumerRoutes.GET("/", custumerHandler.All)
		custumerRoutes.POST("/", custumerHandler.Createcustumer)
		custumerRoutes.GET("/:id", custumerHandler.FindOnecustumerByID)
		custumerRoutes.PUT("/:id", custumerHandler.Updatecustumer)
		custumerRoutes.DELETE("/:id", custumerHandler.Deletecustumer)
		custumerRoutes.GET("/count", custumerHandler.CountAllCustumer)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
