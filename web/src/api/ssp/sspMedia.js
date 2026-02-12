import service from '@/utils/request'
// @Tags SspMedia
// @Summary 创建媒体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SspMedia true "创建媒体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /media/createSspMedia [post]
export const createSspMedia = (data) => {
  return service({
    url: '/media/createSspMedia',
    method: 'post',
    data
  })
}

// @Tags SspMedia
// @Summary 删除媒体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SspMedia true "删除媒体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /media/deleteSspMedia [delete]
export const deleteSspMedia = (params) => {
  return service({
    url: '/media/deleteSspMedia',
    method: 'delete',
    params
  })
}

// @Tags SspMedia
// @Summary 批量删除媒体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除媒体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /media/deleteSspMedia [delete]
export const deleteSspMediaByIds = (params) => {
  return service({
    url: '/media/deleteSspMediaByIds',
    method: 'delete',
    params
  })
}

// @Tags SspMedia
// @Summary 更新媒体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.SspMedia true "更新媒体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /media/updateSspMedia [put]
export const updateSspMedia = (data) => {
  return service({
    url: '/media/updateSspMedia',
    method: 'put',
    data
  })
}

// @Tags SspMedia
// @Summary 用id查询媒体
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.SspMedia true "用id查询媒体"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /media/findSspMedia [get]
export const findSspMedia = (params) => {
  return service({
    url: '/media/findSspMedia',
    method: 'get',
    params
  })
}

// @Tags SspMedia
// @Summary 分页获取媒体列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取媒体列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /media/getSspMediaList [get]
export const getSspMediaList = (params) => {
  return service({
    url: '/media/getSspMediaList',
    method: 'get',
    params
  })
}


export const getAllSspMediaList = (params) => {
  return service({
    url: '/media/getAllSspMediaList',
    method: 'get',
    params
  })
}


// @Tags SspMedia
// @Summary 不需要鉴权的媒体接口
// @Accept application/json
// @Produce application/json
// @Param data query sspReq.SspMediaSearch true "分页获取媒体列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /media/getSspMediaPublic [get]
export const getSspMediaPublic = () => {
  return service({
    url: '/media/getSspMediaPublic',
    method: 'get',
  })
}
