package ssp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ssp"
	sspReq "github.com/flipped-aurora/gin-vue-admin/server/model/ssp/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SspMediaApi struct{}

// CreateSspMedia 创建媒体
// @Tags SspMedia
// @Summary 创建媒体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ssp.SspMedia true "创建媒体"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /media/createSspMedia [post]
func (mediaApi *SspMediaApi) CreateSspMedia(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var media ssp.SspMedia
	err := c.ShouldBindJSON(&media)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mediaService.CreateSspMedia(ctx, &media)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSspMedia 删除媒体
// @Tags SspMedia
// @Summary 删除媒体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ssp.SspMedia true "删除媒体"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /media/deleteSspMedia [delete]
func (mediaApi *SspMediaApi) DeleteSspMedia(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := mediaService.DeleteSspMedia(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSspMediaByIds 批量删除媒体
// @Tags SspMedia
// @Summary 批量删除媒体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /media/deleteSspMediaByIds [delete]
func (mediaApi *SspMediaApi) DeleteSspMediaByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := mediaService.DeleteSspMediaByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSspMedia 更新媒体
// @Tags SspMedia
// @Summary 更新媒体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body ssp.SspMedia true "更新媒体"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /media/updateSspMedia [put]
func (mediaApi *SspMediaApi) UpdateSspMedia(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var media ssp.SspMedia
	err := c.ShouldBindJSON(&media)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = mediaService.UpdateSspMedia(ctx, media)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSspMedia 用id查询媒体
// @Tags SspMedia
// @Summary 用id查询媒体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询媒体"
// @Success 200 {object} response.Response{data=ssp.SspMedia,msg=string} "查询成功"
// @Router /media/findSspMedia [get]
func (mediaApi *SspMediaApi) FindSspMedia(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	remedia, err := mediaService.GetSspMedia(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(remedia, c)
}

// GetSspMediaList 分页获取媒体列表
// @Tags SspMedia
// @Summary 分页获取媒体列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query sspReq.SspMediaSearch true "分页获取媒体列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /media/getSspMediaList [get]
func (mediaApi *SspMediaApi) GetSspMediaList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo sspReq.SspMediaSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := mediaService.GetSspMediaInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (mediaApi *SspMediaApi) GetAllSspMediaList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo sspReq.SspMediaSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, _, err := mediaService.GetSspMediaInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

// GetSspMediaPublic 不需要鉴权的媒体接口
// @Tags SspMedia
// @Summary 不需要鉴权的媒体接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /media/getSspMediaPublic [get]
func (mediaApi *SspMediaApi) GetSspMediaPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	mediaService.GetSspMediaPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的媒体接口信息",
	}, "获取成功", c)
}
