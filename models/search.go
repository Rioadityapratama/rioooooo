package models

type Search struct {
	IDSearch  uint    `gorm:"column:id_search;primaryKey;autoIncrement" json:"id_search"`
	IDProduct uint    `gorm:"column:id_product" json:"id_product"`
	Product   Product `gorm:"foreignKey:IDProduct;references:IDProduct" json:"product"`
}

func (Search) TableName() string {
	return "search"
}
