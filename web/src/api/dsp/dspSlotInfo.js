import service from '@/utils/request'
// @Tags DspSlotInfo
// @Summary 创建预算位管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspSlotInfo true "创建预算位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dspSlotInfo/createDspSlotInfo [post]
export const createDspSlotInfo = (data) => {
  return service({
    url: '/dspSlotInfo/createDspSlotInfo',
    method: 'post',
    data
  })
}

// @Tags DspSlotInfo
// @Summary 删除预算位管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspSlotInfo true "删除预算位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dspSlotInfo/deleteDspSlotInfo [delete]
export const deleteDspSlotInfo = (params) => {
  return service({
    url: '/dspSlotInfo/deleteDspSlotInfo',
    method: 'delete',
    params
  })
}

// @Tags DspSlotInfo
// @Summary 批量删除预算位管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除预算位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dspSlotInfo/deleteDspSlotInfo [delete]
export const deleteDspSlotInfoByIds = (params) => {
  return service({
    url: '/dspSlotInfo/deleteDspSlotInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags DspSlotInfo
// @Summary 更新预算位管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspSlotInfo true "更新预算位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dspSlotInfo/updateDspSlotInfo [put]
export const updateDspSlotInfo = (data) => {
  return service({
    url: '/dspSlotInfo/updateDspSlotInfo',
    method: 'put',
    data
  })
}

// @Tags DspSlotInfo
// @Summary 用id查询预算位管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DspSlotInfo true "用id查询预算位管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dspSlotInfo/findDspSlotInfo [get]
export const findDspSlotInfo = (params) => {
  return service({
    url: '/dspSlotInfo/findDspSlotInfo',
    method: 'get',
    params
  })
}

// @Tags DspSlotInfo
// @Summary 分页获取预算位管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取预算位管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dspSlotInfo/getDspSlotInfoList [get]
export const getDspSlotInfoList = (params) => {
  return service({
    url: '/dspSlotInfo/getDspSlotInfoList',
    method: 'get',
    params
  })
}

// @Tags DspSlotInfo
// @Summary 不需要鉴权的预算位管理接口
// @Accept application/json
// @Produce application/json
// @Param data query dspReq.DspSlotInfoSearch true "分页获取预算位管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dspSlotInfo/getDspSlotInfoPublic [get]
export const getDspSlotInfoPublic = () => {
  return service({
    url: '/dspSlotInfo/getDspSlotInfoPublic',
    method: 'get',
  })
}
