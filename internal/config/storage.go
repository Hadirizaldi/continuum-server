package config


type StorageConfig struct {
	
	Type string
	FinalPathPattern string
	PublicURLPrefix string
	Local LocalStorageConfig
	S3 S3StorageConfig

}

type LocalStorageConfig struct {
	BasePath string
}

type S3StorageConfig struct {
	Bucket string
	Region string
	Prefix string
}


func DefaultStorageConfig() StorageConfig {
	return StorageConfig {
		Type: "local",
		FinalPathPattern: "{category}/{session_id}/{filename}",
		PublicURLPrefix: "/files",
		Local: LocalStorageConfig {
			BasePath: "/var/storage",
		},
		
		// TODO: next version
		S3: S3StorageConfig {},
	}
}

