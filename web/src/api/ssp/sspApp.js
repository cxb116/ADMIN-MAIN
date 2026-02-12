import service from '@/utils/request'
// @Tags SspApp
// @Summary 创建媒体应用
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SspApp true "创建媒体应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /App/createSspApp [post]
export const createSspApp = (data) => {
  return service({
    url: '/App/createSspApp',
    method: 'post',
    data
  })
}

// @Tags SspApp
// @Summary 删除媒体应用
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SspApp true "删除媒体应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /App/deleteSspApp [delete]
export const deleteSspApp = (params) => {
  return service({
    url: '/App/deleteSspApp',
    method: 'delete',
    params
  })
}

// @Tags SspApp
// @Summary 批量删除媒体应用
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除媒体应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /App/deleteSspApp [delete]
export const deleteSspAppByIds = (params) => {
  return service({
    url: '/App/deleteSspAppByIds',
    method: 'delete',
    params
  })
}

// @Tags SspApp
// @Summary 更新媒体应用
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SspApp true "更新媒体应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /App/updateSspApp [put]
export const updateSspApp = (data) => {
  return service({
    url: '/App/updateSspApp',
    method: 'put',
    data
  })
}

// @Tags SspApp
// @Summary 用id查询媒体应用
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SspApp true "用id查询媒体应用"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /App/findSspApp [get]
export const findSspApp = (params) => {
  return service({
    url: '/App/findSspApp',
    method: 'get',
    params
  })
}

// @Tags SspApp
// @Summary 分页获取媒体应用列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取媒体应用列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /App/getSspAppList [get]
export const getSspAppList = (params) => {
  return service({
    url: '/App/getSspAppList',
    method: 'get',
    params
  })
}

// @Tags SspApp
// @Summary 不需要鉴权的媒体应用接口
// @Accept application/json
// @Produce application/json
// @Param data query sspReq.SspAppSearch true "分页获取媒体应用列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /App/getSspAppPublic [get]
export const getSspAppPublic = () => {
  return service({
    url: '/App/getSspAppPublic',
    method: 'get',
  })
}
