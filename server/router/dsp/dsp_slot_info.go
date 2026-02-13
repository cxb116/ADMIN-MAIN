package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DspSlotInfoRouter struct {}

// InitDspSlotInfoRouter 初始化 预算位管理 路由信息
func (s *DspSlotInfoRouter) InitDspSlotInfoRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	dspSlotInfoRouter := Router.Group("dspSlotInfo").Use(middleware.OperationRecord())
	dspSlotInfoRouterWithoutRecord := Router.Group("dspSlotInfo")
	dspSlotInfoRouterWithoutAuth := PublicRouter.Group("dspSlotInfo")
	{
		dspSlotInfoRouter.POST("createDspSlotInfo", dspSlotInfoApi.CreateDspSlotInfo)   // 新建预算位管理
		dspSlotInfoRouter.DELETE("deleteDspSlotInfo", dspSlotInfoApi.DeleteDspSlotInfo) // 删除预算位管理
		dspSlotInfoRouter.DELETE("deleteDspSlotInfoByIds", dspSlotInfoApi.DeleteDspSlotInfoByIds) // 批量删除预算位管理
		dspSlotInfoRouter.PUT("updateDspSlotInfo", dspSlotInfoApi.UpdateDspSlotInfo)    // 更新预算位管理
	}
	{
		dspSlotInfoRouterWithoutRecord.GET("findDspSlotInfo", dspSlotInfoApi.FindDspSlotInfo)        // 根据ID获取预算位管理
		dspSlotInfoRouterWithoutRecord.GET("getDspSlotInfoList", dspSlotInfoApi.GetDspSlotInfoList)  // 获取预算位管理列表
	}
	{
	    dspSlotInfoRouterWithoutAuth.GET("getDspSlotInfoPublic", dspSlotInfoApi.GetDspSlotInfoPublic)  // 预算位管理开放接口
	}
}
