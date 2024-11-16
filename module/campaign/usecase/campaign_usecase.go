package campaignusecase

import (
	"time"

	campaignrepo "github.com/haohmaru3000/trinity/module/campaign/repository"
)

type CampaignService struct {
	repo campaignrepo.CampaignRepository
}

func NewCampaignService(repo campaignrepo.CampaignRepository) *CampaignService {
	return &CampaignService{repo: repo}
}

func (s *CampaignService) IsCampaignActive(campaignID int) (bool, error) {
	campaign, err := s.repo.GetCampaignByID(campaignID)
	if err != nil {
		return false, err
	}
	if !campaign.Active || time.Now().After(*campaign.Expiry) {
		return false, nil
	}
	return true, nil
}

func (s *CampaignService) RegisterUser(campaignID, userID int) (bool, error) {
	userCount, err := s.repo.CountRegisteredUsers(campaignID)
	if err != nil {
		return false, err
	}
	campaign, err := s.repo.GetCampaignByID(campaignID)
	if err != nil {
		return false, err
	}
	if userCount >= campaign.MaxUsers || !campaign.Active || time.Now().After(*campaign.Expiry) {
		return false, nil
	}
	err = s.repo.RegisterUserToCampaign(campaignID, userID)
	if err != nil {
		return false, err
	}
	return true, nil
}
