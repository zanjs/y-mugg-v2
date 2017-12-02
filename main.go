package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/zanjs/y-mugg-v2/app/controllers"
	mid "github.com/zanjs/y-mugg-v2/app/middleware"
	"github.com/zanjs/y-mugg-v2/app/monitor"
	"github.com/zanjs/y-mugg-v2/config"
)

var (
	appConfig = config.Config.App
	jwtConfig = config.Config.JWT
)

func main() {

	e := echo.New()

	e.HTTPErrorHandler = monitor.CustomHTTPErrorHandler
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(mid.ServerHeader)

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Routes
	e.GET("/", controllers.GetHome)

	e.POST("/user/add", controllers.CreateUser)

	e.GET("/records/jobs", controllers.AllProductWareroom)

	v0 := e.Group("/v0")

	v0.GET("/", controllers.CreateTable)

	v1 := e.Group("/v1")
	v1.POST("/login", controllers.PostLogin)

	v1.Use(middleware.JWT([]byte(jwtConfig.Secret)))

	// Users
	v1.GET("/users", controllers.ArticlesController{}.GetAll)
	v1.POST("/users", controllers.CreateUser)
	v1.GET("/users/:id", controllers.ShowUser)
	v1.PUT("/users/:id", controllers.UpdateUser)
	v1.DELETE("/users/:id", controllers.DeleteUser)

	// Articles
	v1.GET("/articles", controllers.ArticlesController{}.GetAll)
	v1.POST("/articles", controllers.ArticlesController{}.Create)
	v1.GET("/articles/:id", controllers.ArticlesController{}.Get)
	v1.PUT("/articles/:id", controllers.ArticlesController{}.Update)
	v1.DELETE("/articles/:id", controllers.ArticlesController{}.Delete)

	// Products
	v1.GET("/products", controllers.AllProducts)
	v1.POST("/products", controllers.CreateProduct)
	v1.GET("/products/:id", controllers.ShowProduct)
	v1.PUT("/products/:id", controllers.UpdateProduct)
	v1.DELETE("/products/:id", controllers.DeleteProduct)

	// Wareroom
	v1.GET("/warerooms", controllers.AllWarerooms)
	v1.POST("/warerooms", controllers.CreateWareroom)
	v1.GET("/warerooms/:id", controllers.ShowWareroom)
	v1.PUT("/warerooms/:id", controllers.UpdateWareroom)
	v1.DELETE("/warerooms/:id", controllers.DeleteWareroom)

	// qm 库存销量更新
	v1.GET("/records", controllers.AllRecordsPage)
	v1.GET("/records/all", controllers.AllRecords)
	v1.GET("/records/q", controllers.GetRecordWhereTime)
	v1.GET("/records/q2", controllers.AllProductWareroomRecordsTime)
	v1.GET("/records/excel", controllers.AllProductWareroomRecords)
	v1.PUT("/records/:id", controllers.UpdateRecord)
	v1.DELETE("/records/:id", controllers.DeleteRecord)
	// Server
	if err := e.Start(fmt.Sprintf("%s:%s", appConfig.HttpAddr, appConfig.HttpPort)); err != nil {
		e.Logger.Fatal(err.Error())
	}

}
