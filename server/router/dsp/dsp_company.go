package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DspCompanyRouter struct {}

// InitDspCompanyRouter 初始化 公司管理 路由信息
func (s *DspCompanyRouter) InitDspCompanyRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	dspCompanyRouter := Router.Group("dspCompany").Use(middleware.OperationRecord())
	dspCompanyRouterWithoutRecord := Router.Group("dspCompany")
	dspCompanyRouterWithoutAuth := PublicRouter.Group("dspCompany")
	{
		dspCompanyRouter.POST("createDspCompany", dspCompanyApi.CreateDspCompany)   // 新建公司管理
		dspCompanyRouter.DELETE("deleteDspCompany", dspCompanyApi.DeleteDspCompany) // 删除公司管理
		dspCompanyRouter.DELETE("deleteDspCompanyByIds", dspCompanyApi.DeleteDspCompanyByIds) // 批量删除公司管理
		dspCompanyRouter.PUT("updateDspCompany", dspCompanyApi.UpdateDspCompany)    // 更新公司管理
	}
	{
		dspCompanyRouterWithoutRecord.GET("findDspCompany", dspCompanyApi.FindDspCompany)        // 根据ID获取公司管理
		dspCompanyRouterWithoutRecord.GET("getDspCompanyList", dspCompanyApi.GetDspCompanyList)  // 获取公司管理列表
	}
	{
	    dspCompanyRouterWithoutAuth.GET("getDspCompanyPublic", dspCompanyApi.GetDspCompanyPublic)  // 公司管理开放接口
	}
}
