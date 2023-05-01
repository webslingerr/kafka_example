package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.udevs.io/car24/car24_go_car_service/api/docs"
	v1 "gitlab.udevs.io/car24/car24_go_car_service/api/v1"
	"gitlab.udevs.io/car24/car24_go_car_service/config"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
	"gitlab.udevs.io/car24/car24_go_car_service/storage"
)

type RouterOptions struct {
	Log     logger.Logger
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cfg := cors.DefaultConfig()

	cfg.AllowHeaders = append(cfg.AllowHeaders, "*")
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true

	router.Use(cors.New(cfg))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:     opt.Log,
		Cfg:     opt.Cfg,
		Storage: opt.Storage,
	})

	apiV1 := router.Group("/v1")

	apiV1.Use()
	{
		// Car --->
		apiV1.GET("/car/:id", handlerV1.GetCar)
		apiV1.GET("/car", handlerV1.GetAllCars)
		// <---

		// Brand --->
		apiV1.GET("/brand/:id", handlerV1.GetBrand)
		apiV1.GET("/brand", handlerV1.GetAllBrands)
		// <---

		// Brand --->
		apiV1.GET("/mark/:id", handlerV1.GetMark)
		apiV1.GET("/mark", handlerV1.GetAllMarks)
		// <---
	}

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
