package dsp

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
    dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type DspSlotInfoApi struct {}



// CreateDspSlotInfo 创建预算位管理
// @Tags DspSlotInfo
// @Summary 创建预算位管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspSlotInfo true "创建预算位管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dspSlotInfo/createDspSlotInfo [post]
func (dspSlotInfoApi *DspSlotInfoApi) CreateDspSlotInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var dspSlotInfo dsp.DspSlotInfo
	err := c.ShouldBindJSON(&dspSlotInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dspSlotInfoService.CreateDspSlotInfo(ctx,&dspSlotInfo)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteDspSlotInfo 删除预算位管理
// @Tags DspSlotInfo
// @Summary 删除预算位管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspSlotInfo true "删除预算位管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dspSlotInfo/deleteDspSlotInfo [delete]
func (dspSlotInfoApi *DspSlotInfoApi) DeleteDspSlotInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := dspSlotInfoService.DeleteDspSlotInfo(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDspSlotInfoByIds 批量删除预算位管理
// @Tags DspSlotInfo
// @Summary 批量删除预算位管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dspSlotInfo/deleteDspSlotInfoByIds [delete]
func (dspSlotInfoApi *DspSlotInfoApi) DeleteDspSlotInfoByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := dspSlotInfoService.DeleteDspSlotInfoByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDspSlotInfo 更新预算位管理
// @Tags DspSlotInfo
// @Summary 更新预算位管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspSlotInfo true "更新预算位管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dspSlotInfo/updateDspSlotInfo [put]
func (dspSlotInfoApi *DspSlotInfoApi) UpdateDspSlotInfo(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var dspSlotInfo dsp.DspSlotInfo
	err := c.ShouldBindJSON(&dspSlotInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dspSlotInfoService.UpdateDspSlotInfo(ctx,dspSlotInfo)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDspSlotInfo 用id查询预算位管理
// @Tags DspSlotInfo
// @Summary 用id查询预算位管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询预算位管理"
// @Success 200 {object} response.Response{data=dsp.DspSlotInfo,msg=string} "查询成功"
// @Router /dspSlotInfo/findDspSlotInfo [get]
func (dspSlotInfoApi *DspSlotInfoApi) FindDspSlotInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	redspSlotInfo, err := dspSlotInfoService.GetDspSlotInfo(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(redspSlotInfo, c)
}
// GetDspSlotInfoList 分页获取预算位管理列表
// @Tags DspSlotInfo
// @Summary 分页获取预算位管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query dspReq.DspSlotInfoSearch true "分页获取预算位管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dspSlotInfo/getDspSlotInfoList [get]
func (dspSlotInfoApi *DspSlotInfoApi) GetDspSlotInfoList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo dspReq.DspSlotInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dspSlotInfoService.GetDspSlotInfoInfoList(ctx,pageInfo)
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

// GetDspSlotInfoPublic 不需要鉴权的预算位管理接口
// @Tags DspSlotInfo
// @Summary 不需要鉴权的预算位管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dspSlotInfo/getDspSlotInfoPublic [get]
func (dspSlotInfoApi *DspSlotInfoApi) GetDspSlotInfoPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    dspSlotInfoService.GetDspSlotInfoPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的预算位管理接口信息",
    }, "获取成功", c)
}
