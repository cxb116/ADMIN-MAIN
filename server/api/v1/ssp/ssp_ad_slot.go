package ssp

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ssp"
    sspReq "github.com/flipped-aurora/gin-vue-admin/server/model/ssp/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type Ssp_ad_slotApi struct {}



// CreateSsp_ad_slot 创建媒体广告位
// @Tags Ssp_ad_slot
// @Summary 创建媒体广告位
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ssp.Ssp_ad_slot true "创建媒体广告位"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /adSlot/createSsp_ad_slot [post]
func (adSlotApi *Ssp_ad_slotApi) CreateSsp_ad_slot(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var adSlot ssp.Ssp_ad_slot
	err := c.ShouldBindJSON(&adSlot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = adSlotService.CreateSsp_ad_slot(ctx,&adSlot)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteSsp_ad_slot 删除媒体广告位
// @Tags Ssp_ad_slot
// @Summary 删除媒体广告位
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ssp.Ssp_ad_slot true "删除媒体广告位"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /adSlot/deleteSsp_ad_slot [delete]
func (adSlotApi *Ssp_ad_slotApi) DeleteSsp_ad_slot(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := adSlotService.DeleteSsp_ad_slot(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSsp_ad_slotByIds 批量删除媒体广告位
// @Tags Ssp_ad_slot
// @Summary 批量删除媒体广告位
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /adSlot/deleteSsp_ad_slotByIds [delete]
func (adSlotApi *Ssp_ad_slotApi) DeleteSsp_ad_slotByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := adSlotService.DeleteSsp_ad_slotByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSsp_ad_slot 更新媒体广告位
// @Tags Ssp_ad_slot
// @Summary 更新媒体广告位
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ssp.Ssp_ad_slot true "更新媒体广告位"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /adSlot/updateSsp_ad_slot [put]
func (adSlotApi *Ssp_ad_slotApi) UpdateSsp_ad_slot(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var adSlot sspReq.Ssp_ad_slotUpdate
	err := c.ShouldBindJSON(&adSlot)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 转换为 Ssp_ad_slot 类型以兼容 Service 层
	var adSlotModel ssp.Ssp_ad_slot
	adSlotModel.GVA_MODEL = adSlot.GVA_MODEL
	adSlotModel.MediaId = adSlot.MediaId
	adSlotModel.AppId = adSlot.AppId
	adSlotModel.Name = adSlot.Name
	adSlotModel.NameAlise = adSlot.NameAlise
	adSlotModel.SceneId = adSlot.SceneId
	adSlotModel.SspPayType = adSlot.SspPayType
	adSlotModel.SspDealRatio = adSlot.SspDealRatio
	adSlotModel.InteractionType = adSlot.InteractionType
	adSlotModel.Height = adSlot.Height
	adSlotModel.Width = adSlot.Width
	adSlotModel.AdImage = adSlot.AdImage
	adSlotModel.Enable = adSlot.Enable
	adSlotModel.FlowConfig = adSlot.FlowConfig

	err = adSlotService.UpdateSsp_ad_slot(ctx,adSlotModel)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSsp_ad_slot 用id查询媒体广告位
// @Tags Ssp_ad_slot
// @Summary 用id查询媒体广告位
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询媒体广告位"
// @Success 200 {object} response.Response{data=ssp.Ssp_ad_slot,msg=string} "查询成功"
// @Router /adSlot/findSsp_ad_slot [get]
func (adSlotApi *Ssp_ad_slotApi) FindSsp_ad_slot(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	readSlot, err := adSlotService.GetSsp_ad_slot(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(readSlot, c)
}
// GetSsp_ad_slotList 分页获取媒体广告位列表
// @Tags Ssp_ad_slot
// @Summary 分页获取媒体广告位列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query sspReq.Ssp_ad_slotSearch true "分页获取媒体广告位列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /adSlot/getSsp_ad_slotList [get]
func (adSlotApi *Ssp_ad_slotApi) GetSsp_ad_slotList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo sspReq.Ssp_ad_slotSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := adSlotService.GetSsp_ad_slotInfoList(ctx,pageInfo)
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

// GetSsp_ad_slotPublic 不需要鉴权的媒体广告位接口
// @Tags Ssp_ad_slot
// @Summary 不需要鉴权的媒体广告位接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /adSlot/getSsp_ad_slotPublic [get]
func (adSlotApi *Ssp_ad_slotApi) GetSsp_ad_slotPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    adSlotService.GetSsp_ad_slotPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的媒体广告位接口信息",
    }, "获取成功", c)
}
// FindDSPInfo 根据 SSP 表的相关信息查询 DSP 信息
// @Tags Ssp_ad_slot
// @Summary 根据 SSP 表的相关信息查询 DSP 信息
// @Accept application/json
// @Produce application/json
// @Param id query string true "SSP 广告位 ID"
// @Success 200 {object} response.Response{data=[]ssp.DspLaunchInfo,msg=string} "成功"
// @Router /adSlot/findDSPInfo [GET]
func (adSlotApi *Ssp_ad_slotApi)FindDSPInfo(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 获取 id 参数
    id := c.Query("id")
    if id == "" {
        response.FailWithMessage("id 参数不能为空", c)
        return
    }

    list, err := adSlotService.FindDSPInfo(ctx, id)
    if err != nil {
        global.GVA_LOG.Error("查询 DSP 信息失败!", zap.Error(err))
        response.FailWithMessage("查询失败:" + err.Error(), c)
        return
    }
    response.OkWithData(list, c)
}


