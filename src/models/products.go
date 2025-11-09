package models

type Products struct {
	IDProduct   uint   `gorm:"primaryKey;autoIncrement" json:"id_product"`
	NamaProduct string `gorm:"unique;not null" json:"nama_product"`
	Harga       uint   `gorm:"not null" json:"harga"`
	Stok        uint   `gorm:"not null" json:"stok"`
}
