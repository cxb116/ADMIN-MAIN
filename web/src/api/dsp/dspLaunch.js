import service from '@/utils/request'
// @Tags DspLaunch
// @Summary 创建预算投放表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspLaunch true "创建预算投放表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dspLaunch/createDspLaunch [post]
export const createDspLaunch = (data) => {
  return service({
    url: '/dspLaunch/createDspLaunch',
    method: 'post',
    data
  })
}

// @Tags DspLaunch
// @Summary 删除预算投放表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspLaunch true "删除预算投放表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dspLaunch/deleteDspLaunch [delete]
export const deleteDspLaunch = (params) => {
  return service({
    url: '/dspLaunch/deleteDspLaunch',
    method: 'delete',
    params
  })
}

// @Tags DspLaunch
// @Summary 批量删除预算投放表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除预算投放表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dspLaunch/deleteDspLaunch [delete]
export const deleteDspLaunchByIds = (params) => {
  return service({
    url: '/dspLaunch/deleteDspLaunchByIds',
    method: 'delete',
    params
  })
}

// @Tags DspLaunch
// @Summary 更新预算投放表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspLaunch true "更新预算投放表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dspLaunch/updateDspLaunch [put]
export const updateDspLaunch = (data) => {
  return service({
    url: '/dspLaunch/updateDspLaunch',
    method: 'put',
    data
  })
}

// @Tags DspLaunch
// @Summary 用id查询预算投放表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DspLaunch true "用id查询预算投放表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dspLaunch/findDspLaunch [get]
export const findDspLaunch = (params) => {
  return service({
    url: '/dspLaunch/findDspLaunch',
    method: 'get',
    params
  })
}

// @Tags DspLaunch
// @Summary 分页获取预算投放表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取预算投放表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dspLaunch/getDspLaunchList [get]
export const getDspLaunchList = (params) => {
  return service({
    url: '/dspLaunch/getDspLaunchList',
    method: 'get',
    params
  })
}

// @Tags DspLaunch
// @Summary 不需要鉴权的预算投放表接口
// @Accept application/json
// @Produce application/json
// @Param data query dspReq.DspLaunchSearch true "分页获取预算投放表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /dspLaunch/getDspLaunchPublic [get]
export const getDspLaunchPublic = () => {
  return service({
    url: '/dspLaunch/getDspLaunchPublic',
    method: 'get',
  })
}

// 批量保存预算投放配置
export const batchSaveDspLaunch = (data) => {
  return service({
    url: '/dspLaunch/batchSave',
    method: 'post',
    data
  })
}

// 根据预算位ID获取配置
export const getDspLaunchByDspSlotId = (params) => {
  return service({
    url: '/dspLaunch/getByDspSlotId',
    method: 'get',
    params
  })
}
