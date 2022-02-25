package config

import "log"

func VerifyConfig() config {

	if parsedConfig.Server.Port == 0 {
		log.Fatal("No port specified")
	}
	if parsedConfig.Server.TokenSigningKey == "" {
		log.Fatal("No token signing key specified")
	}

	if parsedConfig.Postgres.User == "" {
		log.Fatal("Postgres user not specified")
	}
	if parsedConfig.Postgres.Password == "" {
		log.Fatal("Postgres password not specified")
	}
	if parsedConfig.Postgres.Host == "" {
		log.Fatal("Postgres host not specified")
	}
	if parsedConfig.Postgres.Port == 0 {
		log.Fatal("Postgres port not specified")
	}
	if parsedConfig.Postgres.DBName == "" {
		log.Fatal("Postgres dbName not specified")
	}

	return parsedConfig

}

var Config = VerifyConfig()
