package models

type Favorite struct {
	IDFavorite uint `gorm:"column:id_favorite;primaryKey;autoIncrement" json:"id_favorite"`
	IDProduct  uint `gorm:"column:id_product" json:"id_product"`
	IDUser     uint `gorm:"column:id_user" json:"id_user"`

	User    User    `gorm:"foreignKey:IDUser;references:IDUser;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,-:all" json:"user"`
	Product Product `gorm:"foreignKey:IDProduct;references:IDProduct;constraint:OnUpdate:CASCADE,OnDelete:CASCADE,-:all" json:"product"`
}

func (Favorite) TableName() string {
	return "favorite"
}
