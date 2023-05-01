package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gitlab.udevs.io/car24/car24_go_admin_api_gateway/api/docs" //for swagger
	v1 "gitlab.udevs.io/car24/car24_go_admin_api_gateway/api/handlers/v1"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/config"

	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/pkg/event"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/pkg/logger"
)

// Config ...
type Config struct {
	Logger logger.Logger
	Cfg    config.Config
	Kafka  *event.Kafka
}

// New
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	config.AllowHeaders = append(config.AllowHeaders, "*")

	r.Use(cors.New(config))

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger: cnf.Logger,
		Cfg:    cnf.Cfg,
		Kafka:  cnf.Kafka,
	})

	// Car endpoints
	r.POST("/v1/car", handlerV1.CreateCar)
	r.GET("/v1/car/:id", handlerV1.GetCar)
	r.GET("/v1/car", handlerV1.GetAllCars)
	r.PUT("/v1/car/:id", handlerV1.UpdateCar)
	r.DELETE("/v1/car/:id", handlerV1.DeleteCar)

	r.POST("/v1/mark", handlerV1.CreateMark)
	r.GET("/v1/mark/:id", handlerV1.GetMark)
	r.GET("/v1/mark", handlerV1.GetAllMarks)
	r.PUT("/v1/mark/:id", handlerV1.UpdateMark)
	r.DELETE("/v1/mark/:id", handlerV1.DeleteMark)

	r.POST("/v1/brand", handlerV1.CreateBrand)
	r.GET("/v1/brand/:id", handlerV1.GetBrand)
	r.GET("/v1/brand", handlerV1.GetAllBrands)
	r.PUT("/v1/brand/:id", handlerV1.UpdateBrand)
	r.DELETE("/v1/brand/:id", handlerV1.DeleteBrand)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
