package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DspCompanyRouter struct{}

func (s *DspCompanyRouter) InitDspCompanyRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	dspCompanyRouter := Router.Group("dspCompany").Use(middleware.OperationRecord())
	dspCompanyRouterWithoutRecord := Router.Group("dspCompany")
	dspCompanyRouterWithoutAuth := PublicRouter.Group("dspCompany")
	{
		dspCompanyRouter.POST("createDspCompany", dspCompanyApi.CreateDspCompany)
		dspCompanyRouter.DELETE("deleteDspCompany", dspCompanyApi.DeleteDspCompany)
		dspCompanyRouter.DELETE("deleteDspCompanyByIds", dspCompanyApi.DeleteDspCompanyByIds)
		dspCompanyRouter.PUT("updateDspCompany", dspCompanyApi.UpdateDspCompany)
	}
	{
		dspCompanyRouterWithoutRecord.GET("findDspCompany", dspCompanyApi.FindDspCompany)
		dspCompanyRouterWithoutRecord.GET("getDspCompanyList", dspCompanyApi.GetDspCompanyList)
		dspCompanyRouterWithoutRecord.GET("getDictionaryTreeListByType", dspCompanyApi.GetDictionaryTreeListByType)
	}
	{
		dspCompanyRouterWithoutAuth.GET("getDspCompanyPublic", dspCompanyApi.GetDspCompanyPublic)
	}
}
