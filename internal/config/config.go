package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/techforge-lat/errortrace/v2"
)

const (
	serverPortEnvKey     = "BASTION_SERVER_PORT"
	appEnvKey            = "BASTION_ENV"
	allowedDomainsEnvKey = "BASTION_ALLOWED_DOMAINS"
	allowedMethodsEnvKey = "BASTION_ALLOWED_METHODS"

	databaseDriverEnvKey   = "BASTION_DATABASE_DRIVER"
	databaseUserEnvKey     = "BASTION_DATABASE_USER"
	databasePasswordEnvKey = "BASTION_DATABASE_PASSWORD"
	databaseHostEnvKey     = "BASTION_DATABASE_HOST"
	databasePortEnvKey     = "BASTION_DATABASE_PORT"
	databaseNameEnvKey     = "BASTION_DATABASE_NAME"
	databaseSSLModeEnvKey  = "BASTION_DATABASE_SSLMODE"
)

var defaultConfig = map[string]any{
	serverPortEnvKey:      8080, //nolint:gomnd
	appEnvKey:             "local",
	allowedDomainsEnvKey:  "*",
	allowedMethodsEnvKey:  "POST,PUT,PATCH,DELETE,GET",
	databaseDriverEnvKey:  "postgres",
	databaseUserEnvKey:    "crm_user",
	databaseHostEnvKey:    "127.0.0.1",
	databasePortEnvKey:    5432, //nolint:gomnd
	databaseNameEnvKey:    "crm",
	databaseSSLModeEnvKey: "disable",
}

func Load() Root {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("cannot load env file, %v", errortrace.Wrap(err))
	}

	serverPort, err := strconv.Atoi(os.Getenv(serverPortEnvKey))
	if err != nil {
		serverPort, _ = defaultConfig[serverPortEnvKey].(int)
	}

	databasePort, err := strconv.Atoi(os.Getenv(databasePortEnvKey))
	if err != nil {
		databasePort, _ = defaultConfig[databasePortEnvKey].(int)
	}

	localConfig := Root{
		ServerPort:     uint(serverPort),
		Env:            readStrEnvOrDefault(appEnvKey),
		AllowedDomains: readStrEnvOrDefault(allowedDomainsEnvKey),
		AllowedMethods: readStrEnvOrDefault(allowedMethodsEnvKey),
		Database: Database{
			Driver:   readStrEnvOrDefault(databaseDriverEnvKey),
			Host:     readStrEnvOrDefault(databaseHostEnvKey),
			Port:     uint(databasePort),
			User:     readStrEnvOrDefault(databaseUserEnvKey),
			Password: os.Getenv(databasePasswordEnvKey),
			Name:     readStrEnvOrDefault(databaseNameEnvKey),
			SSLMode:  readStrEnvOrDefault(databaseSSLModeEnvKey),
		},
	}

	return localConfig
}

func readStrEnvOrDefault(env string) string {
	value := os.Getenv(env)
	if value != "" {
		return value
	}

	defaultVal, ok := defaultConfig[env]
	if !ok {
		log.Fatal(errortrace.Wrap(fmt.Errorf("not found default value for env %q", env)))
	}

	defaultValStr, ok := defaultVal.(string)
	if !ok {
		log.Fatal(errortrace.Wrap(fmt.Errorf("default value for env %q is not `string`", env)))
	}

	return defaultValStr
}
