package campaignbusiness

import (
	"context"
	"time"

	campaignmodel "github.com/haohmaru3000/trinity/module/campaign/model"
)

type RegisterUserStore interface {
	Register(ctx context.Context, campaignId, userId int) error
}

type FindCampaignRepo interface {
	FindCampaign(ctx context.Context, campaignId int) (*campaignmodel.Campaign, error)
}

type CountRegisteredUsersRepo interface {
	CountRegisteredUsers(ctx context.Context, campaignId int) (int, error)
}

type registerUserBiz struct {
	store          RegisterUserStore
	findCampRepo   FindCampaignRepo
	countUsersRepo CountRegisteredUsersRepo
}

func NewRegisterUserBiz(store RegisterUserStore) *registerUserBiz {
	return &registerUserBiz{store: store}
}

func (biz *registerUserBiz) RegisterUser(ctx context.Context, campaignId, userId int) (bool, error) {
	userCount, err := biz.countUsersRepo.CountRegisteredUsers(ctx, campaignId)
	if err != nil {
		return false, err
	}

	campaign, err := biz.findCampRepo.FindCampaign(ctx, campaignId)
	if err != nil {
		return false, err
	}

	if userCount >= campaign.MaxUsers || !campaign.Active || time.Now().After(*campaign.Expiry) {
		return false, nil
	}

	err = biz.store.Register(ctx, campaignId, userId)
	if err != nil {
		return false, err
	}

	return true, nil
}
