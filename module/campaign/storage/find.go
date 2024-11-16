package campaignstorage

import (
	"context"
	"gorm.io/gorm"

	"github.com/haohmaru3000/trinity/common"
	campaignmodel "github.com/haohmaru3000/trinity/module/campaign/model"
)

func (s *sqlStore) FindCampaign(context context.Context, condition map[string]interface{}, moreKeys ...string) (*campaignmodel.Campaign, error) {
	var data campaignmodel.Campaign

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
