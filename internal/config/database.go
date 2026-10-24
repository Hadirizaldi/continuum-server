package config


type DatabaseConfig struct {
	
	Type string
	Host string
	Port int
	User string
	Password string
	Name string
	SSLMode string
	MaxConnections int

}


func DefaultDatabaseConfig() DatabaseConfig {
	return DatabaseConfig {
		Type : "postgres",
		Host : "localhost",
		Port : 5432,
		User : "postgres",
		Password : "",
		Name : "uploads",
		SSLMode : "disable",
		MaxConnections : 25,
	}
}