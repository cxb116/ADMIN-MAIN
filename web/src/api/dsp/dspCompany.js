import service from '@/utils/request'
// @Tags DspCompany
// @Summary 创建公司管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspCompany true "创建公司管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dspCompany/createDspCompany [post]
export const createDspCompany = (data) => {
  return service({
    url: '/dspCompany/createDspCompany',
    method: 'post',
    data
  })
}

// @Tags DspCompany
// @Summary 删除公司管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspCompany true "删除公司管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dspCompany/deleteDspCompany [delete]
export const deleteDspCompany = (params) => {
  return service({
    url: '/dspCompany/deleteDspCompany',
    method: 'delete',
    params
  })
}

// @Tags DspCompany
// @Summary 批量删除公司管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公司管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dspCompany/deleteDspCompany [delete]
export const deleteDspCompanyByIds = (params) => {
  return service({
    url: '/dspCompany/deleteDspCompanyByIds',
    method: 'delete',
    params
  })
}

// @Tags DspCompany
// @Summary 更新公司管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspCompany true "更新公司管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dspCompany/updateDspCompany [put]
export const updateDspCompany = (data) => {
  return service({
    url: '/dspCompany/updateDspCompany',
    method: 'put',
    data
  })
}

// @Tags DspCompany
// @Summary 用id查询公司管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DspCompany true "用id查询公司管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dspCompany/findDspCompany [get]
export const findDspCompany = (params) => {
  return service({
    url: '/dspCompany/findDspCompany',
    method: 'get',
    params
  })
}

// @Tags DspCompany
// @Summary 分页获取公司管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公司管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dspCompany/getDspCompanyList [get]
export const getDspCompanyList = (params) => {
  return service({
    url: '/dspCompany/getDspCompanyList',
    method: 'get',
    params
  })
}

// @Tags DspCompany
// @Summary 不需要鉴权的公司管理接口
// @Accept application/json
// @Produce application/json
// @Param data query dspReq.DspCompanySearch true "分页获取公司管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dspCompany/getDspCompanyPublic [get]
export const getDspCompanyPublic = () => {
  return service({
    url: '/dspCompany/getDspCompanyPublic',
    method: 'get',
  })
}
