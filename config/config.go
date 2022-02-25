package config

type ServerConfig struct {
	Port            int
	TokenSigningKey string
}

type Postgres struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
}

type config struct {
	Server   ServerConfig `json:"server"`
	Postgres Postgres     `json:"postgres"`
}
