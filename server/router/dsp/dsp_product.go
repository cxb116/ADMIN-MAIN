package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DspProductRouter struct{}

func (s *DspProductRouter) InitDspProductRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	dProductRouter := Router.Group("dProduct").Use(middleware.OperationRecord())
	dProductRouterWithoutRecord := Router.Group("dProduct")
	dProductRouterWithoutAuth := PublicRouter.Group("dProduct")
	{
		dProductRouter.POST("createDspProduct", dProductApi.CreateDspProduct)
		dProductRouter.DELETE("deleteDspProduct", dProductApi.DeleteDspProduct)
		dProductRouter.DELETE("deleteDspProductByIds", dProductApi.DeleteDspProductByIds)
		dProductRouter.PUT("updateDspProduct", dProductApi.UpdateDspProduct)
		dProductRouter.GET("getDictionaryTreeListByType", dProductApi.GetDictionaryTreeListByType)
	}
	{
		dProductRouterWithoutRecord.GET("findDspProduct", dProductApi.FindDspProduct)
		dProductRouterWithoutRecord.GET("getDspProductList", dProductApi.GetDspProductList)
	}
	{
		dProductRouterWithoutAuth.GET("getDspProductPublic", dProductApi.GetDspProductPublic)
	}
}
