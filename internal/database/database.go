package database

import (
	"context"
	"log"
	"sync"

	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	lock     = &sync.Mutex{}
)

// Singleton design pattern for creating and providing a single database connection instance
func ConnectDatabase(ctx context.Context, dialector gorm.Dialector) *gorm.DB {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		log.Println("Creating database instance...")
		var err error
		instance, err = gorm.Open(dialector, &gorm.Config{})
		if err != nil {
			log.Panic("Failed to connect to database!")
		}
	} else {
		log.Println("Connected to database!")
	}

	return instance
}
