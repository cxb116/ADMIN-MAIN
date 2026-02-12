package ssp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type Ssp_ad_slotRouter struct {}

// InitSsp_ad_slotRouter 初始化 媒体广告位 路由信息
func (s *Ssp_ad_slotRouter) InitSsp_ad_slotRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	adSlotRouter := Router.Group("adSlot").Use(middleware.OperationRecord())
	adSlotRouterWithoutRecord := Router.Group("adSlot")
	adSlotRouterWithoutAuth := PublicRouter.Group("adSlot")
	{
		adSlotRouter.POST("createSsp_ad_slot", adSlotApi.CreateSsp_ad_slot)   // 新建媒体广告位
		adSlotRouter.DELETE("deleteSsp_ad_slot", adSlotApi.DeleteSsp_ad_slot) // 删除媒体广告位
		adSlotRouter.DELETE("deleteSsp_ad_slotByIds", adSlotApi.DeleteSsp_ad_slotByIds) // 批量删除媒体广告位
		adSlotRouter.PUT("updateSsp_ad_slot", adSlotApi.UpdateSsp_ad_slot)    // 更新媒体广告位
	}
	{
		adSlotRouterWithoutRecord.GET("findSsp_ad_slot", adSlotApi.FindSsp_ad_slot)        // 根据ID获取媒体广告位
		adSlotRouterWithoutRecord.GET("getSsp_ad_slotList", adSlotApi.GetSsp_ad_slotList)  // 获取媒体广告位列表
	}
	{
	    adSlotRouterWithoutAuth.GET("getSsp_ad_slotPublic", adSlotApi.GetSsp_ad_slotPublic)  // 媒体广告位开放接口
	}
}
