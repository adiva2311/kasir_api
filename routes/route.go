package routes

import (
	"kasir_api/config"
	"kasir_api/controllers"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
)

func APIRoutes(e *echo.Echo) {
	// INIT Database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed Connect to Database", err)
	}

	g := e.Group("/api")

	g.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "This is Kasir API",
		})
	})

	// HEALTH CHECK
	g.GET("/health-check", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "API is running",
		})
	})

	// PRODUCT ROUTES
	ProductController := controllers.NewProductController(db)
	g.GET("/products", ProductController.GetAllProducts)
	g.POST("/products", ProductController.CreateProduct)
	g.GET("/products/:id", ProductController.GetProductByID)
	g.PUT("/products/:id", ProductController.UpdateProduct)
	g.DELETE("/products/:id", ProductController.DeleteProduct)

	// CATEGORY ROUTES
	CategoryController := controllers.NewCategoryController(db)
	g.GET("/categories", CategoryController.GetAllCategories)
	g.POST("/categories", CategoryController.CreateCategory)
	g.PUT("/categories/:id", CategoryController.UpdateCategory)
	g.GET("/categories/:id", CategoryController.GetCategoryByID)
	g.DELETE("/categories/:id", CategoryController.DeleteCategory)
}
