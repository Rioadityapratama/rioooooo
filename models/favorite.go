package models

type Favorite struct {
	IDFavorite uint `gorm:"column:id_favorite;primaryKey;autoIncrement" json:"id_favorite"`
	IDProduct  uint `gorm:"column:id_product" json:"id_product"`
	IDUser     uint `gorm:"column:id_user" json:"id_user"`

	User    User    `gorm:"foreignKey:IDUser;references:IDUser"`
	Product Product `gorm:"foreignKey:IDProduct;references:IDProduct"`
}

func (Favorite) TableName() string {
	return "favorite"
}
