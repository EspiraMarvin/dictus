package initializers

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// type PostgreSQL struct {
// 	ORM *gorm.DB
// }

var DB *gorm.DB

// func ConnectToDB(*gorm.DB, error) {
func ConnectToDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Nairobi",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	var err error

	// connect to db
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("can't open connection to the local database: %v", err)
	} else {
		// Logger: logger.Default.LogMode(logger.Info),

		// Logger: logger.Default.LogMode(logger.Info),
		// log.Info("Connected to db successfully!", db)
		log.Info("Connected to db successfully!")

		// log.Infof("Connected to db successful!: %s", db)

	}

	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	// DriverName: "cloudsqlpostgres",
	// 	// DSN:        dsn,
	// 	PrepareStmt: true,
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// 	Logger: logger.Default.LogMode(logger.Info),
	// }))

	// return db
}
