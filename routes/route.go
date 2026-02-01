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

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "This is Kasir API",
		})
	})

	// HEALTH CHECK
	e.GET("/health-check", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "API is running",
		})
	})

	// PRODUCT ROUTES
	ProductController := controllers.NewProductController(db)
	e.GET("/products", ProductController.GetAllProducts)
	e.GET("/products/:id", ProductController.GetProductByID)
}
