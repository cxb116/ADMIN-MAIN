package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DspProductRouter struct {}

// InitDspProductRouter 初始化 预算产品 路由信息
func (s *DspProductRouter) InitDspProductRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	dProductRouter := Router.Group("dProduct").Use(middleware.OperationRecord())
	dProductRouterWithoutRecord := Router.Group("dProduct")
	dProductRouterWithoutAuth := PublicRouter.Group("dProduct")
	{
		dProductRouter.POST("createDspProduct", dProductApi.CreateDspProduct)   // 新建预算产品
		dProductRouter.DELETE("deleteDspProduct", dProductApi.DeleteDspProduct) // 删除预算产品
		dProductRouter.DELETE("deleteDspProductByIds", dProductApi.DeleteDspProductByIds) // 批量删除预算产品
		dProductRouter.PUT("updateDspProduct", dProductApi.UpdateDspProduct)    // 更新预算产品
	}
	{
		dProductRouterWithoutRecord.GET("findDspProduct", dProductApi.FindDspProduct)        // 根据ID获取预算产品
		dProductRouterWithoutRecord.GET("getDspProductList", dProductApi.GetDspProductList)  // 获取预算产品列表
	}
	{
	    dProductRouterWithoutAuth.GET("getDspProductPublic", dProductApi.GetDspProductPublic)  // 预算产品开放接口
	}
}
