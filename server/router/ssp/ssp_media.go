package ssp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SspMediaRouter struct{}

// InitSspMediaRouter 初始化 媒体 路由信息
func (s *SspMediaRouter) InitSspMediaRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	mediaRouter := Router.Group("media").Use(middleware.OperationRecord())
	mediaRouterWithoutRecord := Router.Group("media")
	mediaRouterWithoutAuth := PublicRouter.Group("media")
	{
		mediaRouter.POST("createSspMedia", mediaApi.CreateSspMedia)             // 新建媒体
		mediaRouter.DELETE("deleteSspMedia", mediaApi.DeleteSspMedia)           // 删除媒体
		mediaRouter.DELETE("deleteSspMediaByIds", mediaApi.DeleteSspMediaByIds) // 批量删除媒体
		mediaRouter.PUT("updateSspMedia", mediaApi.UpdateSspMedia)              // 更新媒体

	}
	{
		mediaRouterWithoutRecord.GET("findSspMedia", mediaApi.FindSspMedia)       // 根据ID获取媒体
		mediaRouterWithoutRecord.GET("getSspMediaList", mediaApi.GetSspMediaList) // 获取媒体列表
		mediaRouterWithoutRecord.GET("getAllSspMediaList", mediaApi.GetAllSspMediaList)
	}
	{
		mediaRouterWithoutAuth.GET("getSspMediaPublic", mediaApi.GetSspMediaPublic) // 媒体开放接口
	}
}
