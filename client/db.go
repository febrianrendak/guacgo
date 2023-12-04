package client

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GuacamoleDB *gorm.DB

type GuacamoleEntity struct {
	EntityID int    `json:"entity_id" gorm:"column:entity_id"`
	Name     string `json:"name" gorm:"column:name"`
	Type     string `json:"type" gorm:"column:type"`
}

func (g *GuacamoleEntity) TableName() string {
	return "guacamole_entity"
}

func InitDBConn(dbName, username, password, host string, port int, paramsString string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", username, password, host, port, dbName)
	dialect := postgres.Open(dsn)

	GuacamoleDB, err = gorm.Open(dialect, &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})

	return GuacamoleDB, err
}

func UpdateUsername(oldUsername, newUsername string) error {
	mapOfData := map[string]any{
		"name": newUsername,
	}

	return GuacamoleDB.
		Model(&GuacamoleEntity{}).
		Where("name = ?", oldUsername).
		Where("type = ?", "USER").
		Updates(mapOfData).
		Error
}
