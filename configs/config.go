package configs

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Port      int
	Driver    string
	Name      string
	Address   string
	DB_Port   int
	Username  string
	Password  string
	S3_KEY    string
	S3_SECRET string
	S3_REGION string
	Midtrans  string
	LOC       string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()
	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = 8000
	defaultConfig.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Name = getEnv("NAME", "healthfit")
	defaultConfig.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.DB_Port = 3306
	defaultConfig.Username = getEnv("USERNAME", "root") /* "root" */
	defaultConfig.Password = getEnv("PASSWORD", "adol1122")
	defaultConfig.S3_KEY = getEnv("S3_KEY", "AKIAVOMUO3KKNSP4RXWR")
	defaultConfig.S3_SECRET = getEnv("S3_SECRET", "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")
	defaultConfig.S3_REGION = getEnv("S3_REGION", "ap-southeast-1")
	defaultConfig.Midtrans = getEnv("MIDTRANS_SERVER_KEY", "SB-Mid-server-6xAOB7aeOdPQL9EEbmV9hnqE")
	defaultConfig.LOC = getEnv("LOC", "Asia%2FJakarta")

	fmt.Println(defaultConfig)

	return &defaultConfig
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "Thinkpad X250" {
		fmt.Println(value)
		return value
	}

	return fallback
}
