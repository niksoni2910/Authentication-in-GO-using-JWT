package initializers

import "go_jwt/models"

func SyncDtabase() {
	DB.AutoMigrate(&models.User{})
}