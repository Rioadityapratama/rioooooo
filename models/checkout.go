package models

type Checkout struct {
	IDCheckout       uint   `gorm:"column:id_checkout;primaryKey;autoIncrement" json:"id_checkout"`
	IDUser           uint   `gorm:"column:id_user" json:"id_user"`
	IDProduct        uint   `gorm:"column:id_product" json:"id_product"`
	IDKeranjang      *uint  `gorm:"column:id_keranjang" json:"id_keranjang"`
	Alamat           string `gorm:"column:alamat" json:"alamat"`
	MetodePengiriman string `gorm:"column:metode_pengiriman" json:"metode_pengiriman"`
	Pembayaran       string `gorm:"column:pembayaran" json:"pembayaran"`
	Jumlah           int    `gorm:"column:jumlah" json:"jumlah"`

	User      User      `gorm:"foreignKey:IDUser;references:IDUser"`
	Product   Product   `gorm:"foreignKey:IDProduct;references:IDProduct"`
	Keranjang Keranjang `gorm:"foreignKey:IDKeranjang;references:IDKeranjang"`
}

func (Checkout) TableName() string {
	return "checkout"
}
