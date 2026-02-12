package ssp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SspAppRouter struct {}

// InitSspAppRouter 初始化 媒体应用 路由信息
func (s *SspAppRouter) InitSspAppRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	AppRouter := Router.Group("App").Use(middleware.OperationRecord())
	AppRouterWithoutRecord := Router.Group("App")
	AppRouterWithoutAuth := PublicRouter.Group("App")
	{
		AppRouter.POST("createSspApp", AppApi.CreateSspApp)   // 新建媒体应用
		AppRouter.DELETE("deleteSspApp", AppApi.DeleteSspApp) // 删除媒体应用
		AppRouter.DELETE("deleteSspAppByIds", AppApi.DeleteSspAppByIds) // 批量删除媒体应用
		AppRouter.PUT("updateSspApp", AppApi.UpdateSspApp)    // 更新媒体应用
	}
	{
		AppRouterWithoutRecord.GET("findSspApp", AppApi.FindSspApp)        // 根据ID获取媒体应用
		AppRouterWithoutRecord.GET("getSspAppList", AppApi.GetSspAppList)  // 获取媒体应用列表
	}
	{
	    AppRouterWithoutAuth.GET("getSspAppPublic", AppApi.GetSspAppPublic)  // 媒体应用开放接口
	}
}
