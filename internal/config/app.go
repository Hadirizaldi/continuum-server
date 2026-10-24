package config


type AppConfig struct {

	Env string
	Port int
	CORSAllowedOrigins []string
    CORSAllowedHeaders []string
	
}

func DefaultAppConfig() AppConfig {
	return AppConfig {
		Env: "development",
		Port: 8080,
		CORSAllowedOrigins: []string{"*"},
		CORSAllowedHeaders: []string{},
	}
}