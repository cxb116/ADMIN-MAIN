package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DspAdSceneRouter struct {}

// InitDspAdSceneRouter 初始化 广告类型 路由信息
func (s *DspAdSceneRouter) InitDspAdSceneRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	adSceneRouter := Router.Group("adScene").Use(middleware.OperationRecord())
	adSceneRouterWithoutRecord := Router.Group("adScene")
	adSceneRouterWithoutAuth := PublicRouter.Group("adScene")
	{
		adSceneRouter.POST("createDspAdScene", adSceneApi.CreateDspAdScene)   // 新建广告类型
		adSceneRouter.DELETE("deleteDspAdScene", adSceneApi.DeleteDspAdScene) // 删除广告类型
		adSceneRouter.DELETE("deleteDspAdSceneByIds", adSceneApi.DeleteDspAdSceneByIds) // 批量删除广告类型
		adSceneRouter.PUT("updateDspAdScene", adSceneApi.UpdateDspAdScene)    // 更新广告类型
	}
	{
		adSceneRouterWithoutRecord.GET("findDspAdScene", adSceneApi.FindDspAdScene)        // 根据ID获取广告类型
		adSceneRouterWithoutRecord.GET("getDspAdSceneList", adSceneApi.GetDspAdSceneList)  // 获取广告类型列表
	}
	{
	    adSceneRouterWithoutAuth.GET("getDspAdScenePublic", adSceneApi.GetDspAdScenePublic)  // 广告类型开放接口
	}
}
