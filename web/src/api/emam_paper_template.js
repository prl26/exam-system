import service from '@/utils/request'

// @Tags PaperTemplate
// @Summary 创建PaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperTemplate true "创建PaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Papertemplate/createPaperTemplate [post]
export const createPaperTemplate = (data) => {
  return service({
    url: '/Papertemplate/createPaperTemplate',
    method: 'post',
    data
  })
}

// @Tags PaperTemplate
// @Summary 删除PaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperTemplate true "删除PaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Papertemplate/deletePaperTemplate [delete]
export const deletePaperTemplate = (data) => {
  return service({
    url: '/Papertemplate/deletePaperTemplate',
    method: 'delete',
    data
  })
}

// @Tags PaperTemplate
// @Summary 删除PaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Papertemplate/deletePaperTemplate [delete]
export const deletePaperTemplateByIds = (data) => {
  return service({
    url: '/Papertemplate/deletePaperTemplateByIds',
    method: 'delete',
    data
  })
}

// @Tags PaperTemplate
// @Summary 更新PaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PaperTemplate true "更新PaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Papertemplate/updatePaperTemplate [put]
export const updatePaperTemplate = (data) => {
  return service({
    url: '/Papertemplate/updatePaperTemplate',
    method: 'put',
    data
  })
}

// @Tags PaperTemplate
// @Summary 用id查询PaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PaperTemplate true "用id查询PaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Papertemplate/findPaperTemplate [get]
export const findPaperTemplate = (params) => {
  return service({
    url: '/Papertemplate/findPaperTemplate',
    method: 'get',
    params
  })
}

// @Tags PaperTemplate
// @Summary 分页获取PaperTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取PaperTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Papertemplate/getPaperTemplateList [get]
export const getPaperTemplateList = (params) => {
  return service({
    url: '/Papertemplate/getPaperTemplateList',
    method: 'get',
    params
  })
}
