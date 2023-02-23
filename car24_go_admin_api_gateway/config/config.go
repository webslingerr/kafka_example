package config

import (
	"os"

	"github.com/spf13/cast"
)

const (
	AdminClientTypeID        string = "5a3818a9-90f0-44e9-a053-3be0ba1e2c01"
	DefaultAdminRoleID       string = "a1ca1301-4da9-424d-a9e2-578ae6dcde01"
	TransferSubCategoryID           = "499e0e48-5890-4d00-9092-b5493f94448b"
	OrderPaymentCategoryID          = 1
	CostPaymentSubCategoryID        = ""

	RentCategoryID                = 1
	CarExpenditureCategoryID      = 2
	EmployeeExpenditureCategoryID = 3
	TransferCategoryID            = 4
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	KafkaUrl string

	CarServiceURL string
	LogLevel      string

	MinioEndpoint       string
	MinioAccessKeyID    string
	MinioSecretAccesKey string
	CDN                 string
	Bucket              string
	MinioExcelBucket    string
	MinioBucketName     string

	HttpPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", ""))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8001"))

	c.KafkaUrl = cast.ToString(getOrReturnDefault("KAFKA_URL", "localhost:9092"))

	c.CarServiceURL = cast.ToString(getOrReturnDefault("CAR_SERVICE_URL", "http://localhost:8004"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
