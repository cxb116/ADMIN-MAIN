package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DspAdSceneRouter struct{}

func (s *DspAdSceneRouter) InitDspAdSceneRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	adSceneRouter := Router.Group("adScene").Use(middleware.OperationRecord())
	adSceneRouterWithoutRecord := Router.Group("adScene")
	adSceneRouterWithoutAuth := PublicRouter.Group("adScene")
	{
		adSceneRouter.POST("createDspAdScene", adSceneApi.CreateDspAdScene)
		adSceneRouter.DELETE("deleteDspAdScene", adSceneApi.DeleteDspAdScene)
		adSceneRouter.DELETE("deleteDspAdSceneByIds", adSceneApi.DeleteDspAdSceneByIds)
		adSceneRouter.PUT("updateDspAdScene", adSceneApi.UpdateDspAdScene)
		adSceneRouter.GET("getDictionaryTreeListByType", adSceneApi.GetDictionaryTreeListByType)
	}
	{
		adSceneRouterWithoutRecord.GET("findDspAdScene", adSceneApi.FindDspAdScene)
		adSceneRouterWithoutRecord.GET("getDspAdSceneList", adSceneApi.GetDspAdSceneList)
	}
	{
		adSceneRouterWithoutAuth.GET("getDspAdScenePublic", adSceneApi.GetDspAdScenePublic)
	}
}
