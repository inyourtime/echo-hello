package bootstrap

import (
	"echo-hello/internal/logger"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Config    Config
	DcWebhook DiscordWebhook
	Db        Db
	S3        S3
	Jwt       Jwt
	Google    Google
}

func NewEnv() *Env {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}

	env := Env{
		Config: Config{
			Env:        os.Getenv("APP_ENV"),
			ServerPort: os.Getenv("PORT"),
			ApiPrefix:  os.Getenv("API_PREFIX"),
		},
		DcWebhook: DiscordWebhook{
			ID:    os.Getenv("DISCORD_ID"),
			Token: os.Getenv("DISCORD_TOKEN"),
		},
		Db: Db{
			Pg: Postgres{
				Host:     os.Getenv("PGHOST"),
				Database: os.Getenv("PGDATABASE"),
				User:     os.Getenv("PGUSER"),
				Password: os.Getenv("PGPASSWORD"),
				Ssl:      os.Getenv("PGSSL"),
			},
		},
		S3: S3{
			AccountID:       os.Getenv("OBJECTSTORAGE_ACCOUNTID"),
			AccessKeyID:     os.Getenv("OBJECTSTORAGE_ACCESSKEYID"),
			AccessKeySecret: os.Getenv("OBJECTSTORAGE_SECRETACCESSKEY"),
			Bucket:          os.Getenv("OBJECTSTORAGE_BUCKET"),
		},
		Jwt: Jwt{
			AccessSecret:  os.Getenv("JWT_ACCESS_SECRET"),
			RefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
		},
		Google: Google{
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			RedirectUrl:  os.Getenv("GOOGLE_REDIRECT_URL"),
		},
	}

	if env.Config.Env == "development" {
		logger.Info("[ENV] The App is running in development env")
	}

	return &env
}

type Config struct {
	Env        string
	ServerPort string
	ApiPrefix  string
}

type DiscordWebhook struct {
	ID    string
	Token string
}

type Db struct {
	Pg Postgres
}

type Postgres struct {
	Host     string
	Database string
	User     string
	Password string
	Ssl      string
}

type S3 struct {
	AccountID       string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
}

type Jwt struct {
	AccessSecret  string
	RefreshSecret string
}

type Google struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
}
