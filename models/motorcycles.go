package models

type Motorcycles struct {
	ID          uint   `gorm:"primary_key;autoIncrement" json:"id"`
	ModelName   string `json:"model_name"`
	MachineType string `json:"machine_type"`
	Year        int    `json:"year"`
	Color       string `json:"color"`
}

// func MigrateMotorcycles(db *gorm.DB) error {
// 	err := db.AutoMigrate(&Motorcycles{})
// 	return err
// }
