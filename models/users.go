package models

type Users struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password []byte `json:"-"`
}

// func MigrateUsers(db *gorm.DB) error {
// 	err := db.AutoMigrate(&Users{})
// 	return err
// }
