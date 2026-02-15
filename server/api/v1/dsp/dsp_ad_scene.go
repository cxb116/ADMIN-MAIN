package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
	dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DspAdSceneApi struct{}

// CreateDspAdScene 创建广告类型
// @Tags DspAdScene
// @Summary 创建广告类型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspAdScene true "创建广告类型"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /adScene/createDspAdScene [post]
func (adSceneApi *DspAdSceneApi) CreateDspAdScene(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var adScene dsp.DspAdScene
	err := c.ShouldBindJSON(&adScene)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = adSceneService.CreateDspAdScene(ctx, &adScene)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteDspAdScene 删除广告类型
// @Tags DspAdScene
// @Summary 删除广告类型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspAdScene true "删除广告类型"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /adScene/deleteDspAdScene [delete]
func (adSceneApi *DspAdSceneApi) DeleteDspAdScene(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := adSceneService.DeleteDspAdScene(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDspAdSceneByIds 批量删除广告类型
// @Tags DspAdScene
// @Summary 批量删除广告类型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /adScene/deleteDspAdSceneByIds [delete]
func (adSceneApi *DspAdSceneApi) DeleteDspAdSceneByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := adSceneService.DeleteDspAdSceneByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDspAdScene 更新广告类型
// @Tags DspAdScene
// @Summary 更新广告类型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspAdScene true "更新广告类型"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /adScene/updateDspAdScene [put]
func (adSceneApi *DspAdSceneApi) UpdateDspAdScene(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var adScene dsp.DspAdScene
	err := c.ShouldBindJSON(&adScene)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = adSceneService.UpdateDspAdScene(ctx, adScene)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDspAdScene 用id查询广告类型
// @Tags DspAdScene
// @Summary 用id查询广告类型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询广告类型"
// @Success 200 {object} response.Response{data=dsp.DspAdScene,msg=string} "查询成功"
// @Router /adScene/findDspAdScene [get]
func (adSceneApi *DspAdSceneApi) FindDspAdScene(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	readScene, err := adSceneService.GetDspAdScene(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(readScene, c)
}

// GetDspAdSceneList 分页获取广告类型列表
// @Tags DspAdScene
// @Summary 分页获取广告类型列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query dspReq.DspAdSceneSearch true "分页获取广告类型列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /adScene/getDspAdSceneList [get]
func (adSceneApi *DspAdSceneApi) GetDspAdSceneList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo dspReq.DspAdSceneSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := adSceneService.GetDspAdSceneInfoList(ctx, pageInfo)
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

// GetDspAdScenePublic 不需要鉴权的广告类型接口
// @Tags DspAdScene
// @Summary 不需要鉴权的广告类型接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /adScene/getDspAdScenePublic [get]
func (adSceneApi *DspAdSceneApi) GetDspAdScenePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	adSceneService.GetDspAdScenePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的广告类型接口信息",
	}, "获取成功", c)
}

// GetDictionaryTreeListByType 参考数据字典协议，查询广告场景
// @Tags DspAdScene
// @Summary 参考数据字典协议，查询广告场景
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=DictionaryResponse,msg=string} "成功"
// @Router /adScene/getDictionaryTreeListByType [GET]
func (adSceneApi *DspAdSceneApi) GetDictionaryTreeListByType(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 查询广告场景数据
	list, err := adSceneService.GetDictionaryTreeListByType(ctx)
	if err != nil {
		global.GVA_LOG.Error("获取广告场景列表失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	// 转换为数据字典格式
	var dictionaryList []DictionaryItem
	for _, item := range list {
		if item.Name != nil {
			dictionaryList = append(dictionaryList, DictionaryItem{
				Label: *item.Name,
				Value: int64(item.ID),
			})
		}
	}

	response.OkWithDetailed(DictionaryResponse{
		List: dictionaryList,
	}, "获取成功", c)
}
