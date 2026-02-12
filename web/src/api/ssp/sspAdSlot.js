import service from '@/utils/request'
// @Tags Ssp_ad_slot
// @Summary 创建媒体广告位
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Ssp_ad_slot true "创建媒体广告位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /adSlot/createSsp_ad_slot [post]
export const createSsp_ad_slot = (data) => {
  return service({
    url: '/adSlot/createSsp_ad_slot',
    method: 'post',
    data
  })
}

// @Tags Ssp_ad_slot
// @Summary 删除媒体广告位
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Ssp_ad_slot true "删除媒体广告位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adSlot/deleteSsp_ad_slot [delete]
export const deleteSsp_ad_slot = (params) => {
  return service({
    url: '/adSlot/deleteSsp_ad_slot',
    method: 'delete',
    params
  })
}

// @Tags Ssp_ad_slot
// @Summary 批量删除媒体广告位
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除媒体广告位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adSlot/deleteSsp_ad_slot [delete]
export const deleteSsp_ad_slotByIds = (params) => {
  return service({
    url: '/adSlot/deleteSsp_ad_slotByIds',
    method: 'delete',
    params
  })
}

// @Tags Ssp_ad_slot
// @Summary 更新媒体广告位
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Ssp_ad_slot true "更新媒体广告位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adSlot/updateSsp_ad_slot [put]
export const updateSsp_ad_slot = (data) => {
  return service({
    url: '/adSlot/updateSsp_ad_slot',
    method: 'put',
    data
  })
}

// @Tags Ssp_ad_slot
// @Summary 用id查询媒体广告位
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Ssp_ad_slot true "用id查询媒体广告位"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adSlot/findSsp_ad_slot [get]
export const findSsp_ad_slot = (params) => {
  return service({
    url: '/adSlot/findSsp_ad_slot',
    method: 'get',
    params
  })
}

// @Tags Ssp_ad_slot
// @Summary 分页获取媒体广告位列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取媒体广告位列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adSlot/getSsp_ad_slotList [get]
export const getSsp_ad_slotList = (params) => {
  return service({
    url: '/adSlot/getSsp_ad_slotList',
    method: 'get',
    params
  })
}

// @Tags Ssp_ad_slot
// @Summary 不需要鉴权的媒体广告位接口
// @Accept application/json
// @Produce application/json
// @Param data query sspReq.Ssp_ad_slotSearch true "分页获取媒体广告位列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /adSlot/getSsp_ad_slotPublic [get]
export const getSsp_ad_slotPublic = () => {
  return service({
    url: '/adSlot/getSsp_ad_slotPublic',
    method: 'get',
  })
}
