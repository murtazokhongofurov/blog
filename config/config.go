package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HttpPort     string
	DbUser       string
	DbPassword   string
	DbPort       string
	DbHost       string
	DbName       string
	LogLevel     string
	AuthFilePath string
	CsvFilePath  string
	SigningKey   string
	MailUsername string
	MailPassword string
	SmtpHost     string
}

func Load(path string) Config {
	c := Config{}
	if err := godotenv.Load(path + ".env"); err != nil {
		log.Fatalln("Error read .env: ", err)
	}

	c.HttpPort = cast.ToString(getDefaultKey("HTTP_PORT", "3333"))
	c.DbHost = cast.ToString(getDefaultKey("DB_HOST", "host"))
	c.DbPort = cast.ToString(getDefaultKey("DB_PORT", "5432"))
	c.DbUser = cast.ToString(getDefaultKey("DB_USER", "username"))
	c.DbPassword = cast.ToString(getDefaultKey("DB_PASS", "password"))
	c.DbName = cast.ToString(getDefaultKey("DB_NAME", "db_name"))
	c.AuthFilePath = cast.ToString(getDefaultKey("AUTH_FILE_PATH", "file_path"))
	c.CsvFilePath = cast.ToString(getDefaultKey("CSV_FILE_PATH", "file_path"))
	c.SigningKey = cast.ToString(getDefaultKey("SIGNING_KEY", "signing_key"))
	c.MailUsername = cast.ToString(getDefaultKey("MAIL_USERNAME", "username@gmail.com"))
	c.MailPassword = cast.ToString(getDefaultKey("MAIL_PASSWORD", "password"))
	c.SmtpHost = cast.ToString(getDefaultKey("SMTP_HOST", "host.smtp"))
	c.LogLevel = cast.ToString(getDefaultKey("LOG_LEVEL", "debug"))

	return c

}

func getDefaultKey(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
