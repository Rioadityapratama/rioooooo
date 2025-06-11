package models

type Product struct {
	IDProduct uint   `gorm:"column:id_product;primaryKey;autoIncrement" json:"id_product"`
	IDPenjual uint   `gorm:"column:id_penjual" json:"id_penjual"`
	Name      string `gorm:"column :name" json:"name"`
	Kategori  string `gorm:"column:kategori" json:"kategori"`
	Size      string `gorm:"column:size" json:"size"`
	Deskripsi string `gorm:"column:deskripsi" json:"deskripsi"`
	Brand     string `gorm:"column:brand" json:"brand"`
	Price     int    `gorm:"column:price" json:"price"`
	Image     string `gorm:"column:image" json:"image"`
	Warna     string `gorm:"column:warna" json:"warna"`
	Stok      int    `gorm:"column:stok" json:"stok"`

	Penjual    Penjual `gorm:"foreignKey:IDPenjual;references:IDPenjual" json:"penjual"`
	IDFavorite *uint   `json:"id_favorite" gorm:"-"`
}

func (Product) TableName() string {
	return "product"
}
