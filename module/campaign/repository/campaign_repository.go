package campaignrepo

import (
	"github.com/haohmaru3000/trinity/module/campaign/model"
)

type CampaignRepository interface {
	GetCampaignByID(id int) (*campaignmodel.Campaign, error)
	RegisterUserToCampaign(campaignID, userID int) error
	CountRegisteredUsers(campaignID int) (int, error)
}
