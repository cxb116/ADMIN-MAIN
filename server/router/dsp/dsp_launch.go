package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DspLaunchRouter struct {}

// InitDspLaunchRouter 初始化 预算投放表 路由信息
func (s *DspLaunchRouter) InitDspLaunchRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	dspLaunchRouter := Router.Group("dspLaunch").Use(middleware.OperationRecord())
	dspLaunchRouterWithoutRecord := Router.Group("dspLaunch")
	dspLaunchRouterWithoutAuth := PublicRouter.Group("dspLaunch")
	{
		dspLaunchRouter.POST("createDspLaunch", dspLaunchApi.CreateDspLaunch)   // 新建预算投放表
		dspLaunchRouter.DELETE("deleteDspLaunch", dspLaunchApi.DeleteDspLaunch) // 删除预算投放表
		dspLaunchRouter.DELETE("deleteDspLaunchByIds", dspLaunchApi.DeleteDspLaunchByIds) // 批量删除预算投放表
		dspLaunchRouter.PUT("updateDspLaunch", dspLaunchApi.UpdateDspLaunch)    // 更新预算投放表
		dspLaunchRouter.POST("batchSave", dspLaunchApi.BatchSaveDspLaunch)      // 批量保存预算投放配置
	}
	{
		dspLaunchRouterWithoutRecord.GET("findDspLaunch", dspLaunchApi.FindDspLaunch)        // 根据ID获取预算投放表
		dspLaunchRouterWithoutRecord.GET("getDspLaunchList", dspLaunchApi.GetDspLaunchList)  // 获取预算投放表列表
		dspLaunchRouterWithoutRecord.GET("getByDspSlotId", dspLaunchApi.GetDspLaunchByDspSlotId)  // 根据预算位ID获取配置
	}
	{
	    dspLaunchRouterWithoutAuth.GET("getDspLaunchPublic", dspLaunchApi.GetDspLaunchPublic)  // 预算投放表开放接口
	}
}
