package bootstrap

import (
	"echo-hello/internal/logger"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func NewPgDatabase(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v dbname=%v password=%v sslmode=%v",
		env.Db.Pg.Host,
		env.Db.Pg.User,
		env.Db.Pg.Database,
		env.Db.Pg.Password,
		env.Db.Pg.Ssl,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
		DryRun: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	pg, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = pg.Ping()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("[DB] Postgres Database has been initialize")
	return db
}
