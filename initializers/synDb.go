package initializers

import "go-auth/m/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
