package voucherstorage

import (
	"context"

	"github.com/haohmaru3000/trinity/common"
	vouchermodel "github.com/haohmaru3000/trinity/module/voucher/model"
)

// Count the number of Users which were registered for a Campaign
func (s *sqlStore) CountRegisteredUsers(ctx context.Context, campaignId int) (int, error) {
	result := make(map[int]int)

	type sqlData struct {
		CampaignId      int `gorm:"column:campaign_id;"`
		RegisteredCount int `gorm:"column:count;"`
	}

	var voucherList []sqlData

	if err := s.db.Table(vouchermodel.Voucher{}.TableName()).
		Select("campaign_id, count(campaign_id) as count").
		Where("campaign_id = (?)", campaignId).
		Group("campaign_id").Find(&voucherList).Error; err != nil {
		return 0, common.ErrDB(err)
	}

	for _, item := range voucherList {
		result[item.CampaignId] = item.RegisteredCount
	}

	return result[0], nil
}
