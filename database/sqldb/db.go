package sqldb

import (
	"github.com/OAB/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	// PoolMigrationName is the name of the migration used by packr to pack the migration file
	PoolMigrationName = "OAB-pool-db"
)

type DB struct {
	gorm         *gorm.DB
	ChallengesDB database.ChallengesDB
}

var (
	Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
)

func NewChallengesDB(dsn string) (*DB, error) {
	gorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: Logger,
	})
	if err != nil {
		return nil, err
	}
	return &DB{
		gorm:         gorm,
		ChallengesDB: database.NewChallengesProofDB(gorm),
	}, nil
}
