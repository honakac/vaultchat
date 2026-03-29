package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type InboxMessage struct {
	Cuid         string `gorm:"primaryKey"`
	ReceiverAddr string `gorm:"index"`
	Payload      []byte
	CreatedAt    time.Time
}

type Database struct {
	db *gorm.DB
}

func (db *Database) Init(filepath string) {
	database, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.db = database

	db.db.AutoMigrate(&InboxMessage{})

	if err := database.Exec(`
		CREATE TRIGGER IF NOT EXISTS cleanup_old_rows
		AFTER INSERT ON inbox_messages
		BEGIN
				DELETE FROM inbox_messages 
				WHERE created_at < DATETIME('now', '-1 day');
		END;
	`).Error; err != nil {
		panic(err)
	}
}
