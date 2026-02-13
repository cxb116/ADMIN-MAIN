package dsp

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
    dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type DspCompanyApi struct {}



// CreateDspCompany 创建公司管理
// @Tags DspCompany
// @Summary 创建公司管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspCompany true "创建公司管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dspCompany/createDspCompany [post]
func (dspCompanyApi *DspCompanyApi) CreateDspCompany(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var dspCompany dsp.DspCompany
	err := c.ShouldBindJSON(&dspCompany)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dspCompanyService.CreateDspCompany(ctx,&dspCompany)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteDspCompany 删除公司管理
// @Tags DspCompany
// @Summary 删除公司管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspCompany true "删除公司管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dspCompany/deleteDspCompany [delete]
func (dspCompanyApi *DspCompanyApi) DeleteDspCompany(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := dspCompanyService.DeleteDspCompany(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDspCompanyByIds 批量删除公司管理
// @Tags DspCompany
// @Summary 批量删除公司管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dspCompany/deleteDspCompanyByIds [delete]
func (dspCompanyApi *DspCompanyApi) DeleteDspCompanyByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := dspCompanyService.DeleteDspCompanyByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDspCompany 更新公司管理
// @Tags DspCompany
// @Summary 更新公司管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspCompany true "更新公司管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dspCompany/updateDspCompany [put]
func (dspCompanyApi *DspCompanyApi) UpdateDspCompany(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var dspCompany dsp.DspCompany
	err := c.ShouldBindJSON(&dspCompany)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dspCompanyService.UpdateDspCompany(ctx,dspCompany)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDspCompany 用id查询公司管理
// @Tags DspCompany
// @Summary 用id查询公司管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询公司管理"
// @Success 200 {object} response.Response{data=dsp.DspCompany,msg=string} "查询成功"
// @Router /dspCompany/findDspCompany [get]
func (dspCompanyApi *DspCompanyApi) FindDspCompany(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	redspCompany, err := dspCompanyService.GetDspCompany(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(redspCompany, c)
}
// GetDspCompanyList 分页获取公司管理列表
// @Tags DspCompany
// @Summary 分页获取公司管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query dspReq.DspCompanySearch true "分页获取公司管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dspCompany/getDspCompanyList [get]
func (dspCompanyApi *DspCompanyApi) GetDspCompanyList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo dspReq.DspCompanySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dspCompanyService.GetDspCompanyInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetDspCompanyPublic 不需要鉴权的公司管理接口
// @Tags DspCompany
// @Summary 不需要鉴权的公司管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dspCompany/getDspCompanyPublic [get]
func (dspCompanyApi *DspCompanyApi) GetDspCompanyPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    dspCompanyService.GetDspCompanyPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的公司管理接口信息",
    }, "获取成功", c)
}
