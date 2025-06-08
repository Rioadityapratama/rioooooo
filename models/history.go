package models

type History struct {
	IDHistory  uint `gorm:"column:id_history;primaryKey;autoIncrement" json:"id_history"`
	IDCheckout uint `gorm:"column:id_checkout" json:"id_checkout"`

	Checkout Checkout `gorm:"foreignKey:IDCheckout;references:IDCheckout;constraint:OnUpdate:CASCADE,OnDelete:CASCADE," json:"checkout"`
}

func (History) TableName() string {
	return "history"
}
