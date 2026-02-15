import service from '@/utils/request'
// @Tags DspAdScene
// @Summary 创建广告类型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspAdScene true "创建广告类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /adScene/createDspAdScene [post]
export const createDspAdScene = (data) => {
  return service({
    url: '/adScene/createDspAdScene',
    method: 'post',
    data
  })
}

// @Tags DspAdScene
// @Summary 删除广告类型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspAdScene true "删除广告类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adScene/deleteDspAdScene [delete]
export const deleteDspAdScene = (params) => {
  return service({
    url: '/adScene/deleteDspAdScene',
    method: 'delete',
    params
  })
}

// @Tags DspAdScene
// @Summary 批量删除广告类型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除广告类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adScene/deleteDspAdScene [delete]
export const deleteDspAdSceneByIds = (params) => {
  return service({
    url: '/adScene/deleteDspAdSceneByIds',
    method: 'delete',
    params
  })
}

// @Tags DspAdScene
// @Summary 更新广告类型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.DspAdScene true "更新广告类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adScene/updateDspAdScene [put]
export const updateDspAdScene = (data) => {
  return service({
    url: '/adScene/updateDspAdScene',
    method: 'put',
    data
  })
}

// @Tags DspAdScene
// @Summary 用id查询广告类型
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.DspAdScene true "用id查询广告类型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adScene/findDspAdScene [get]
export const findDspAdScene = (params) => {
  return service({
    url: '/adScene/findDspAdScene',
    method: 'get',
    params
  })
}

// @Tags DspAdScene
// @Summary 分页获取广告类型列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取广告类型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adScene/getDspAdSceneList [get]
export const getDspAdSceneList = (params) => {
  return service({
    url: '/adScene/getDspAdSceneList',
    method: 'get',
    params
  })
}

// @Tags DspAdScene
// @Summary 不需要鉴权的广告类型接口
// @Accept application/json
// @Produce application/json
// @Param data query dspReq.DspAdSceneSearch true "分页获取广告类型列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /adScene/getDspAdScenePublic [get]
export const getDspAdScenePublic = () => {
  return service({
    url: '/adScene/getDspAdScenePublic',
    method: 'get',
  })
}
// GetDictionaryTreeListByType 参考数据字典协议，查询业务类
// @Tags DspAdScene
// @Summary 参考数据字典协议，查询业务类
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /adScene/getDictionaryTreeListByType [GET]
export const getDictionaryTreeListByType = () => {
  return service({
    url: '/adScene/getDictionaryTreeListByType',
    method: 'GET'
  })
}
