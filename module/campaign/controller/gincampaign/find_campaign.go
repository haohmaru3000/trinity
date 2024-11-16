package gincampaign

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/haohmaru3000/trinity/common"
	"github.com/haohmaru3000/trinity/components/appctx"
	"github.com/haohmaru3000/trinity/module/campaign/business"
	"github.com/haohmaru3000/trinity/module/campaign/repository"
	"github.com/haohmaru3000/trinity/module/campaign/storage"
)

func FindCampaign(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := campaignstorage.NewSQLStore(db)

		repo := campaignrepo.NewFindCampaignRepo(store)
		biz := campaignbusiness.NewFindCampaignBiz(repo)

		result, err := biz.FindCampaign(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
