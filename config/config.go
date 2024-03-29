package config

import (
	"github.com/spf13/cast"
	"os"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	HTTPPort   string
	HTTPScheme string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	PostgresMaxConnections int32

	DefaultOffset string
	DefaultLimit  string

	SecretKey string

	PasscodePool   string
	PasscodeLength int

	AuthServiceHost string
	AuthGRPCPort    string

	MigrationPath string

	EmailUsername string
	EmailPassword string

	MinioEndpoint        string
	MinioAccessKeyID     string
	MinioSecretAccessKey string
	MinioProtocol        bool
}

// Load ...
func Load() Config {
	//if err := godotenv.Load("/app/.env"); err != nil {
	//	fmt.Println("No .env file found [/app/.env]")
	//} else if err := godotenv.Load(".env"); err != nil {
	//	fmt.Println("No .env file found [.env]")
	//}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "editory_submission"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":9107"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "164.92.169.168"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 30032))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "es_backend"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "mob5jahb0Zeeriem"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "es_backend"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "Here$houldBe$ome$ecretKey"))

	config.PasscodePool = cast.ToString(getOrReturnDefaultValue("PASSCODE_POOL", "0123456789"))
	config.PasscodeLength = cast.ToInt(getOrReturnDefaultValue("PASSCODE_LENGTH", "6"))

	config.AuthServiceHost = cast.ToString(getOrReturnDefaultValue("AUTH_SERVICE_HOST", "localhost"))
	config.AuthGRPCPort = cast.ToString(getOrReturnDefaultValue("AUTH_GRPC_PORT", ":8998"))

	config.MigrationPath = cast.ToString(getOrReturnDefaultValue("MIGRATION_PATH", "/home/euler/Documents/projects/editory_submission/migrations/postgres"))

	config.EmailUsername = cast.ToString(getOrReturnDefaultValue("EMAIL_USERNAME", "editorysubmission@gmail.com"))
	config.EmailPassword = cast.ToString(getOrReturnDefaultValue("EMAIL_PASSWORD", "occo pbku zktc oqqt"))

	config.MinioAccessKeyID = cast.ToString(getOrReturnDefaultValue("MINIO_ACCESS_KEY", "fczbKQdzXNSjxCDu7aEatAnKjqpxHXp7km7HGveQyKCSZFPK"))
	config.MinioSecretAccessKey = cast.ToString(getOrReturnDefaultValue("MINIO_SECRET_KEY", "kQffPzZRcEz8UNyzcV9WMEGFb2fhUAKXMxCJbCXJhKrdGLWY"))
	config.MinioEndpoint = cast.ToString(getOrReturnDefaultValue("MINIO_ENDPOINT", "test.cdn.editorypress.uz"))
	config.MinioProtocol = cast.ToBool(getOrReturnDefaultValue("MINIO_PROTOCOL", true))

	return config
}

//func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
//	val, exists := os.LookupEnv(key)
//
//	if exists {
//		return val
//	}
//
//	return defaultValue
//}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val := os.Getenv(key)

	if val != "" {
		return val
	}

	return defaultValue
}
