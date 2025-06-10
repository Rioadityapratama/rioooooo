package models

type Penjual struct {
	IDPenjual uint   `gorm:"column:id_penjual;primaryKey" json:"id_penjual"`
	Nama      string `gorm:"column:nama" json:"nama"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	Telepon   string `gorm:"column:telepon" json:"telepon"`
	Profil    string `gorm:"column:profil" json:"profil"`
}

func (Penjual) TableName() string {
	return "penjual"
}
