package campaignbusiness

import (
	"context"

	campaignmodel "github.com/haohmaru3000/trinity/module/campaign/model"
)

type FindCampaignStore interface {
	FindCampaign(ctx context.Context, campaignId int) (*campaignmodel.Campaign, error)
}

type findCampaignBiz struct {
	store FindCampaignStore
}

func NewFindCampaignBiz(store FindCampaignStore) *findCampaignBiz {
	return &findCampaignBiz{store: store}
}

func (biz *findCampaignBiz) FindCampaign(ctx context.Context, campaignId int) (*campaignmodel.Campaign, error) {
	campaign, err := biz.store.FindCampaign(ctx, campaignId)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
