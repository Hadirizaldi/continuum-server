package config

import "github.com/joho/godotenv"


type Config struct {

	App AppConfig
	Database DatabaseConfig
	Storage StorageConfig
	Upload UploadConfig

}

func Load() (*Config, error) {

	godotenv.Load()

	cfg := &Config {
		App: DefaultAppConfig(),
		Database: DefaultDatabaseConfig(),
		Storage: DefaultStorageConfig(),
		Upload: DefaultUploadConfig(),
	}

	// App
	cfg.App.Env = getEnv("APP_ENV", cfg.App.Env)
	cfg.App.Port = getEnvInt("APP_PORT", cfg.App.Port)
	cfg.App.CORSAllowedOrigins = getEnvStringSlice("CORS_ALLOWED_ORIGINS", cfg.App.CORSAllowedOrigins)
	cfg.App.CORSAllowedHeaders = getEnvStringSlice("CORS_ALLOWED_HEADERS", cfg.App.CORSAllowedHeaders)

	// Upload
	cfg.Upload.MaxFileSize = getEnvInt64("UPLOAD_MAX_FILE_SIZE", cfg.Upload.MaxFileSize)
	cfg.Upload.ChunkSize = getEnvInt64("UPLOAD_CHUNK_SIZE", cfg.Upload.ChunkSize)
	cfg.Upload.AllowedMimeTypes = getEnvStringSlice("UPLOAD_ALLOWED_MIME_TYPES", cfg.Upload.AllowedMimeTypes)
	cfg.Upload.SessionTTL = getEnvDuration("UPLOAD_SESSION_TTL", cfg.Upload.SessionTTL)
	cfg.Upload.UserIDHeader = getEnv("UPLOAD_USER_ID_HEADER", cfg.Upload.UserIDHeader)

	// Storage
	cfg.Storage.Type = getEnv("STORAGE_TYPE", cfg.Storage.Type)
	cfg.Storage.FinalPathPattern = getEnv("STORAGE_FINAL_PATH_PATTERN", cfg.Storage.FinalPathPattern)
	cfg.Storage.PublicURLPrefix = getEnv("STORAGE_PUBLIC_URL_PREFIX", cfg.Storage.PublicURLPrefix)
	cfg.Storage.Local.BasePath = getEnv("STORAGE_LOCAL_BASE_PATH", cfg.Storage.Local.BasePath)

	// Database
	cfg.Database.Type = getEnv("DB_TYPE", cfg.Database.Type)
	cfg.Database.Host = getEnv("DB_HOST", cfg.Database.Host)
	cfg.Database.Port = getEnvInt("DB_PORT", cfg.Database.Port)
	cfg.Database.User = getEnv("DB_USER", cfg.Database.User)
	cfg.Database.Password = getEnv("DB_PASSWORD", cfg.Database.Password)
	cfg.Database.Name = getEnv("DB_NAME", cfg.Database.Name)
	cfg.Database.SSLMode = getEnv("DB_SSL_MODE", cfg.Database.SSLMode)
	cfg.Database.MaxConnections = getEnvInt("DB_MAX_CONNECTIONS", cfg.Database.MaxConnections)

	return cfg, nil
}