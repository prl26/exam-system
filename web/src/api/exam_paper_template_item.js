import service from '@/utils/request'

// @Tags PaperTemplateItem
// @Summary 创建PaperTemplateItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperTemplateItem true "创建PaperTemplateItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /paperTemplateItem/createPaperTemplateItem [post]
export const createPaperTemplateItem = (data) => {
  return service({
    url: '/paperTemplateItem/createPaperTemplateItem',
    method: 'post',
    data
  })
}

// @Tags PaperTemplateItem
// @Summary 删除PaperTemplateItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperTemplateItem true "删除PaperTemplateItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /paperTemplateItem/deletePaperTemplateItem [delete]
export const deletePaperTemplateItem = (data) => {
  return service({
    url: '/paperTemplateItem/deletePaperTemplateItem',
    method: 'delete',
    data
  })
}

// @Tags PaperTemplateItem
// @Summary 删除PaperTemplateItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PaperTemplateItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /paperTemplateItem/deletePaperTemplateItem [delete]
export const deletePaperTemplateItemByIds = (data) => {
  return service({
    url: '/paperTemplateItem/deletePaperTemplateItemByIds',
    method: 'delete',
    data
  })
}

// @Tags PaperTemplateItem
// @Summary 更新PaperTemplateItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperTemplateItem true "更新PaperTemplateItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /paperTemplateItem/updatePaperTemplateItem [put]
export const updatePaperTemplateItem = (data) => {
  return service({
    url: '/paperTemplateItem/updatePaperTemplateItem',
    method: 'put',
    data
  })
}

// @Tags PaperTemplateItem
// @Summary 用id查询PaperTemplateItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PaperTemplateItem true "用id查询PaperTemplateItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /paperTemplateItem/findPaperTemplateItem [get]
export const findPaperTemplateItem = (params) => {
  return service({
    url: '/paperTemplateItem/findPaperTemplateItem',
    method: 'get',
    params
  })
}

// @Tags PaperTemplateItem
// @Summary 分页获取PaperTemplateItem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PaperTemplateItem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /paperTemplateItem/getPaperTemplateItemList [get]
export const getPaperTemplateItemList = (params) => {
  return service({
    url: '/paperTemplateItem/getPaperTemplateItemList',
    method: 'get',
    params
  })
}
