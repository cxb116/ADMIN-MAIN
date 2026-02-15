package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp"
	dspReq "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DspProductApi struct{}



// CreateDspProduct 创建预算产品
// @Tags DspProduct
// @Summary 创建预算产品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspProduct true "创建预算产品"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /dProduct/createDspProduct [post]
func (dProductApi *DspProductApi) CreateDspProduct(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var dProduct dsp.DspProduct
	err := c.ShouldBindJSON(&dProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dProductService.CreateDspProduct(ctx,&dProduct)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteDspProduct 删除预算产品
// @Tags DspProduct
// @Summary 删除预算产品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspProduct true "删除预算产品"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /dProduct/deleteDspProduct [delete]
func (dProductApi *DspProductApi) DeleteDspProduct(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := dProductService.DeleteDspProduct(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDspProductByIds 批量删除预算产品
// @Tags DspProduct
// @Summary 批量删除预算产品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /dProduct/deleteDspProductByIds [delete]
func (dProductApi *DspProductApi) DeleteDspProductByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := dProductService.DeleteDspProductByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDspProduct 更新预算产品
// @Tags DspProduct
// @Summary 更新预算产品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body dsp.DspProduct true "更新预算产品"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /dProduct/updateDspProduct [put]
func (dProductApi *DspProductApi) UpdateDspProduct(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var dProduct dsp.DspProduct
	err := c.ShouldBindJSON(&dProduct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dProductService.UpdateDspProduct(ctx,dProduct)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDspProduct 用id查询预算产品
// @Tags DspProduct
// @Summary 用id查询预算产品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询预算产品"
// @Success 200 {object} response.Response{data=dsp.DspProduct,msg=string} "查询成功"
// @Router /dProduct/findDspProduct [get]
func (dProductApi *DspProductApi) FindDspProduct(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	redProduct, err := dProductService.GetDspProduct(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(redProduct, c)
}
// GetDspProductList 分页获取预算产品列表
// @Tags DspProduct
// @Summary 分页获取预算产品列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query dspReq.DspProductSearch true "分页获取预算产品列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /dProduct/getDspProductList [get]
func (dProductApi *DspProductApi) GetDspProductList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo dspReq.DspProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := dProductService.GetDspProductInfoList(ctx,pageInfo)
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

// GetDspProductPublic 不需要鉴权的预算产品接口
// @Tags DspProduct
// @Summary 不需要鉴权的预算产品接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dProduct/getDspProductPublic [get]
func (dProductApi *DspProductApi) GetDspProductPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    dProductService.GetDspProductPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的预算产品接口信息",
    }, "获取成功", c)
}
// GetDictionaryTreeListByType 参考数据字典协议，查询产品
// @Tags DspProduct
// @Summary 参考数据字典协议，查询产品
// @Description 当没有传递 dsp_company_id 时，查询所有公司的产品；当传递了 dsp_company_id 时，只查询指定公司的产品
// @Accept application/json
// @Produce application/json
// @Param dsp_company_id query string false "公司ID（可选）"
// @Success 200 {object} response.Response{data=DictionaryResponse,msg=string} "获取成功"
// @Router /dProduct/getDictionaryTreeListByType [get]
func (dProductApi *DspProductApi) GetDictionaryTreeListByType(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 获取可选的公司ID参数
	dspCompanyId := c.Query("dsp_company_id")

	// 查询产品数据
	list, err := dProductService.GetDictionaryTreeListByType(ctx, &dspCompanyId)
	if err != nil {
		global.GVA_LOG.Error("获取产品列表失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	// 转换为数据字典格式
	var dictionaryList []DictionaryItem
	for _, item := range list {
		dictionaryList = append(dictionaryList, DictionaryItem{
			Label: *item.Name,
			Value: int64(item.ID),
		})
	}

	response.OkWithDetailed(DictionaryResponse{
		List: dictionaryList,
	}, "获取成功", c)
}


