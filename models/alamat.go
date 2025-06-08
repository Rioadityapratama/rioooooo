package models

type Alamat struct {
	IDAlamat uint   `gorm:"column:id_alamat;primaryKey;autoIncrement" json:"id_alamat"`
	IDUser   uint   `gorm:"column:id_user" json:"id_user"`
	Alamat   string `gorm:"column:alamat" json:"alamat"`

	User User `gorm:"foreignKey:IDUser;references:IDUser"` // Relasi ke user
}

func (Alamat) TableName() string {
	return "alamat"
}
