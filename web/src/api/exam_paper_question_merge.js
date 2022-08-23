import service from '@/utils/request'

// @Tags PaperQuestionMerge
// @Summary 创建PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperQuestionMerge true "创建PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /paperQuestionMerge/createPaperQuestionMerge [post]
export const createPaperQuestionMerge = (data) => {
  return service({
    url: '/paperQuestionMerge/createPaperQuestionMerge',
    method: 'post',
    data
  })
}

// @Tags PaperQuestionMerge
// @Summary 删除PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperQuestionMerge true "删除PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /paperQuestionMerge/deletePaperQuestionMerge [delete]
export const deletePaperQuestionMerge = (data) => {
  return service({
    url: '/paperQuestionMerge/deletePaperQuestionMerge',
    method: 'delete',
    data
  })
}

// @Tags PaperQuestionMerge
// @Summary 删除PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /paperQuestionMerge/deletePaperQuestionMerge [delete]
export const deletePaperQuestionMergeByIds = (data) => {
  return service({
    url: '/paperQuestionMerge/deletePaperQuestionMergeByIds',
    method: 'delete',
    data
  })
}

// @Tags PaperQuestionMerge
// @Summary 更新PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperQuestionMerge true "更新PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /paperQuestionMerge/updatePaperQuestionMerge [put]
export const updatePaperQuestionMerge = (data) => {
  return service({
    url: '/paperQuestionMerge/updatePaperQuestionMerge',
    method: 'put',
    data
  })
}

// @Tags PaperQuestionMerge
// @Summary 用id查询PaperQuestionMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PaperQuestionMerge true "用id查询PaperQuestionMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /paperQuestionMerge/findPaperQuestionMerge [get]
export const findPaperQuestionMerge = (params) => {
  return service({
    url: '/paperQuestionMerge/findPaperQuestionMerge',
    method: 'get',
    params
  })
}

// @Tags PaperQuestionMerge
// @Summary 分页获取PaperQuestionMerge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PaperQuestionMerge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /paperQuestionMerge/getPaperQuestionMergeList [get]
export const getPaperQuestionMergeList = (params) => {
  return service({
    url: '/paperQuestionMerge/getPaperQuestionMergeList',
    method: 'get',
    params
  })
}
