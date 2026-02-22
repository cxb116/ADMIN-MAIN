package ssp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Ssp_ad_slotRouter struct{}

func (s *Ssp_ad_slotRouter) InitSsp_ad_slotRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	adSlotRouter := Router.Group("adSlot").Use(middleware.OperationRecord())
	adSlotRouterWithoutRecord := Router.Group("adSlot")
	adSlotRouterWithoutAuth := PublicRouter.Group("adSlot")
	{
		adSlotRouter.POST("createSsp_ad_slot", adSlotApi.CreateSsp_ad_slot)
		adSlotRouter.DELETE("deleteSsp_ad_slot", adSlotApi.DeleteSsp_ad_slot)
		adSlotRouter.DELETE("deleteSsp_ad_slotByIds", adSlotApi.DeleteSsp_ad_slotByIds)
		adSlotRouter.PUT("updateSsp_ad_slot", adSlotApi.UpdateSsp_ad_slot)
		adSlotRouter.GET("findDSPInfo", adSlotApi.FindDSPInfo)
	}
	{
		adSlotRouterWithoutRecord.GET("findSsp_ad_slot", adSlotApi.FindSsp_ad_slot)
		adSlotRouterWithoutRecord.GET("getSsp_ad_slotList", adSlotApi.GetSsp_ad_slotList)
	}
	{
		adSlotRouterWithoutAuth.GET("getSsp_ad_slotPublic", adSlotApi.GetSsp_ad_slotPublic)
	}
}
