package campaignbusiness

import (
	"context"
	"time"
)

type isCampaignActiveBiz struct {
	store FindCampaignStore
}

func NewIsCampaignActiveBiz(store FindCampaignStore) *isCampaignActiveBiz {
	return &isCampaignActiveBiz{store: store}
}

func (biz *isCampaignActiveBiz) IsCampaignActive(ctx context.Context, campaignId int) (bool, error) {
	campaign, err := biz.store.FindCampaign(ctx, campaignId)
	if err != nil {
		return false, err
	}

	if !campaign.Active || time.Now().After(*campaign.Expiry) {
		return false, nil
	}

	return true, nil
}
