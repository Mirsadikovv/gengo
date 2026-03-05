package pg

import (
	"fmt"
	"log"
	"sync"

	"github.com/Mirsadikovv/gengo/shared_service/sharedutil"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instances = make(map[string]*gorm.DB)
	instance  *gorm.DB
	mu        sync.Mutex
)

func Primary(gormConfig *GormConfig, cfg ...ConnectionConfig) *gorm.DB {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = sharedutil.MustValue(newDB(gormConfig, cfg...))
	}

	return instance
}

func DBFactory(name string, gormConfig *GormConfig, cfg ...ConnectionConfig) (*gorm.DB, error) {
	if name == "" {
		name = "default"
	}

	mu.Lock()
	defer mu.Unlock()

	if instance, exists := instances[name]; exists {
		return instance, nil
	}

	db, err := newDB(gormConfig, cfg...)
	{
		if err != nil {
			return nil, err
		}
	}

	instances[name] = db
	return db, nil
}

func Must(name string, gormConfig *GormConfig, cfg ...ConnectionConfig) *gorm.DB {
	db, err := DBFactory(name, gormConfig, cfg...)
	{
		if err != nil {
			log.Fatalf("Failed to initialize database connection [%s]: %v", name, err)
		}
	}

	return db
}

func newDB(gormConfig *GormConfig, cfg ...ConnectionConfig) (*gorm.DB, error) {
	config := defaultConfig
	{
		if len(cfg) > 0 {
			config = cfg[0]
		}
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	{
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
	}

	return db, nil
}
