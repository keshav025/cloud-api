package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"

	"github.com/sirupsen/logrus"
)

type appConfig struct {
	MasterDatabase string `envconfig:"MASTER_DATABASE" default:"test"`
	DBHost         string `envconfig:"DB_HOST" default:"test"`
	DBUser         string `envconfig:"DBUSER" default:"test"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"test"`
	DBPort         string `envconfig:"DB_PORT" default:"5432"`
	DBName         string `envconfig:"DB_Name" default:"test-api"`
	IsLocalEnv     bool   `envconfig:"LOCAL_ENV" default:"false"`
	ServerPort     string `envconfig:"SERVER_PORT" default:"8080"`

	AzureSubscriptionId string `envconfig:"AZURE_SUBSCRIPTION_ID" default:"test"`
	AzureResourceGroup  string `envconfig:"AZURE_RESOURCE_GROUP" default:"test"`
	AzureClientId       string `envconfig:"AZURE_CLIENT_ID" default:"test"`
	AzureTenantId       string `envconfig:"AZURE_TENANT_ID" default:"test"`
	AzureClientSecret   string `envconfig:"AZURE_CLIENT_SECRET" default:"test"`
}

var appCfg *appConfig

func init() {
	appCfg = &appConfig{}
	err := InjectEnv(appCfg)
	if err != nil {
		logrus.Error("Error while loading app configurations")
	}
}

func GetConfig() *appConfig {
	return appCfg
}

func InjectEnvByPrefix(prefix string, spec interface{}) error {
	return envconfig.Process(prefix, spec)
}

func InjectEnv(spec interface{}) error {
	return InjectEnvByPrefix("", spec)
}

func GetEnv(name string) string {
	return os.Getenv(name)
}
