import service from '@/utils/request'
// @Tags DspProduct
// @Summary 创建预算产品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspProduct true "创建预算产品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dProduct/createDspProduct [post]
export const createDspProduct = (data) => {
  return service({
    url: '/dProduct/createDspProduct',
    method: 'post',
    data
  })
}

// @Tags DspProduct
// @Summary 删除预算产品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspProduct true "删除预算产品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dProduct/deleteDspProduct [delete]
export const deleteDspProduct = (params) => {
  return service({
    url: '/dProduct/deleteDspProduct',
    method: 'delete',
    params
  })
}

// @Tags DspProduct
// @Summary 批量删除预算产品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除预算产品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dProduct/deleteDspProduct [delete]
export const deleteDspProductByIds = (params) => {
  return service({
    url: '/dProduct/deleteDspProductByIds',
    method: 'delete',
    params
  })
}

// @Tags DspProduct
// @Summary 更新预算产品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspProduct true "更新预算产品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dProduct/updateDspProduct [put]
export const updateDspProduct = (data) => {
  return service({
    url: '/dProduct/updateDspProduct',
    method: 'put',
    data
  })
}

// @Tags DspProduct
// @Summary 用id查询预算产品
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DspProduct true "用id查询预算产品"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dProduct/findDspProduct [get]
export const findDspProduct = (params) => {
  return service({
    url: '/dProduct/findDspProduct',
    method: 'get',
    params
  })
}

// @Tags DspProduct
// @Summary 分页获取预算产品列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取预算产品列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dProduct/getDspProductList [get]
export const getDspProductList = (params) => {
  return service({
    url: '/dProduct/getDspProductList',
    method: 'get',
    params
  })
}

// @Tags DspProduct
// @Summary 不需要鉴权的预算产品接口
// @Accept application/json
// @Produce application/json
// @Param data query dspReq.DspProductSearch true "分页获取预算产品列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dProduct/getDspProductPublic [get]
export const getDspProductPublic = () => {
  return service({
    url: '/dProduct/getDspProductPublic',
    method: 'get',
  })
}
