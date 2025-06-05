package models

type History struct {
	IDHistory  uint `gorm:"column:id_history;primaryKey;autoIncrement" json:"id_history"`
	IDUser     uint `gorm:"column:id_user" json:"id_user"`
	IDCheckout uint `gorm:"column:id_checkout" json:"id_checkout"`

	User     User     `gorm:"foreignKey:IDUser;references:IDUser" json:"user"`
	Checkout Checkout `gorm:"foreignKey:IDCheckout;references:IDCheckout" json:"checkout"`
}

func (History) TableName() string {
	return "history"
}