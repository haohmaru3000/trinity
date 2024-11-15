package voucher_model

import "time"

type Voucher struct {
	ID         int        `json:"id" gorm:"column:id;"`
	Code       string     `json:"code" gorm:"column:code;"`
	UserID     int        `json:"user_id" gorm:"column:user_id;"`
	CampaignID int        `json:"campaign_id" gorm:"column:campaign_id;"`
	ValidFrom  *time.Time `json:"valid_from" gorm:"column:valid_from;"`
	ValidTo    *time.Time `json:"valid_to" gorm:"column:valid_to;"`
	IsRedeemed bool       `json:"is_redeemed" gorm:"column:is_redeemed;"`
}
