package campaignrepo

import (
	"context"

	"github.com/haohmaru3000/trinity/common"
	campaignmodel "github.com/haohmaru3000/trinity/module/campaign/model"
)

type FindCampaignStore interface {
	FindCampaign(context context.Context, condition map[string]interface{}, moreKeys ...string) (*campaignmodel.Campaign, error)
}

type findCampaignRepo struct {
	store FindCampaignStore
}

func NewFindCampaignRepo(store FindCampaignStore) *findCampaignRepo {
	return &findCampaignRepo{store: store}
}

func (repo *findCampaignRepo) FindCampaign(ctx context.Context, campaignId int) (*campaignmodel.Campaign, error) {
	result, err := repo.store.FindCampaign(ctx, map[string]interface{}{"id": campaignId})

	if err != nil {
		return nil, common.ErrCannotGetEntity(campaignmodel.EntityName, err)
	}

	return result, nil
}
