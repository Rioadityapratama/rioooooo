package models

type Keranjang struct {
	IDKeranjang uint `gorm:"column:id_keranjang;primaryKey;autoIncrement" json:"id_keranjang"`
	IDProduct   uint `gorm:"column:id_product" json:"id_product"`
	IDUser      uint `gorm:"column:id_user" json:"id_user"`
	Jumlah      int  `gorm:"column:jumlah" json:"jumlah"`

	Product Product `gorm:"foreignKey:IDProduct;references:IDProduct;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,-:all" json:"product"`
	User    User    `gorm:"foreignKey:IDUser;references:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,-:all" json:"user"`
}

func (Keranjang) TableName() string {
	return "keranjang"
}
