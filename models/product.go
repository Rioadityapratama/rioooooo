package models

type Product struct {
	IDProduct  uint   `gorm:"column:id_product;primaryKey;autoIncrement" json:"id_product"`
	IDPenjual  uint   `gorm:"column:id_penjual" json:"id_penjual"`
	NamaProduk string `gorm:"column :namaproduk" json:"namaproduk"`
	Kategori   string `gorm:"column:kategori" json:"kategori"`
	Size       string `gorm:"column:size" json:"size"`
	Deskripsi  string `gorm:"column:deskripsi" json:"deskripsi"`
	Brand      string `gorm:"column:brand" json:"brand"`
	Price      int    `gorm:"column:price" json:"price"`
	Image      string `gorm:"column:image" json:"image"`
	Warna      string `gorm:"column:warna" json:"warna"`

	Penjual Penjual `gorm:"foreignKey:IDPenjual;references:IDPenjual" json:"penjual"`
}

func (Product) TableName() string {
	return "product"
}