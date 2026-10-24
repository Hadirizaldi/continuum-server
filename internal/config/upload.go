package config

import "time"


type UploadConfig struct {

	MaxFileSize int64
	ChunkSize   int64
	AllowedMimeTypes []string
	SessionTTL time.Duration
	UserIDHeader string

}


func DefaultUploadConfig() UploadConfig {
	return UploadConfig {
		MaxFileSize: 2 * 1024 * 1024 * 1024, // 2 GB
		ChunkSize: 5 * 1024 * 1024, // 5 MB
		AllowedMimeTypes: []string{
			"video/*", "image/*", "audio/*", "application/pdf",
			"application/msword",
			"application/vnd.openxmlformats-officedocument.*",
			"application/zip", "application/x-rar-compressed",
			"application/vnd.rar" },
		SessionTTL: 24 * time.Hour,
		UserIDHeader: "X-User-ID",
	}
}
