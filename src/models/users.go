package models

	type Users struct {
		IDUser   uint   `gorm:"primaryKey;autoIncrement" json:"id_user"`
		NamaUser string `gorm:"not null" json:"nama_user"`
		Email    string `gorm:"unique;not null" json:"email"`
	}
