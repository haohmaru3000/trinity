package campaignmodel

import "time"

const EntityName = "Campaign"

type Campaign struct {
	Id                 int        `json:"id" gorm:"column:id;"`
	Name               string     `json:"name" gorm:"column:name;"`
	DiscountPercentage int        `json:"discount_percentage" gorm:"column:discount_percentage;"`
	MaxUsers           int        `json:"max_users" gorm:"column:max_users;"`
	Expiry             *time.Time `json:"expiry" gorm:"column:expiry;"`
	Active             bool       `json:"active" gorm:"column:active;"`
}
