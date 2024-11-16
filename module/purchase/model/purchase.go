package purchase_model

import "time"

type Purchase struct {
	ID        int        `json:"id" gorm:"column:id;"`
	UserId    int        `json:"user_id" gorm:"column:user_id;"`
	Plan      string     `json:"plan" gorm:"column:plan;"`
	Price     float64    `json:"price" gorm:"column:price;"`
	Discount  float64    `json:"discount" gorm:"column:discount;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
}

func (Purchase) TableName() string {
	return "subscriptions"
}
