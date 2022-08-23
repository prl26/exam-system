import service from '@/utils/request'

// @Tags ExamPaperTemplate
// @Summary 创建ExamPaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamPaperTemplate true "创建ExamPaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPaperTemplate/createExamPaperTemplate [post]
export const createExamPaperTemplate = (data) => {
  return service({
    url: '/examPaperTemplate/createExamPaperTemplate',
    method: 'post',
    data
  })
}

// @Tags ExamPaperTemplate
// @Summary 删除ExamPaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamPaperTemplate true "删除ExamPaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examPaperTemplate/deleteExamPaperTemplate [delete]
export const deleteExamPaperTemplate = (data) => {
  return service({
    url: '/examPaperTemplate/deleteExamPaperTemplate',
    method: 'delete',
    data
  })
}

// @Tags ExamPaperTemplate
// @Summary 删除ExamPaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExamPaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examPaperTemplate/deleteExamPaperTemplate [delete]
export const deleteExamPaperTemplateByIds = (data) => {
  return service({
    url: '/examPaperTemplate/deleteExamPaperTemplateByIds',
    method: 'delete',
    data
  })
}

// @Tags ExamPaperTemplate
// @Summary 更新ExamPaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamPaperTemplate true "更新ExamPaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /examPaperTemplate/updateExamPaperTemplate [put]
export const updateExamPaperTemplate = (data) => {
  return service({
    url: '/examPaperTemplate/updateExamPaperTemplate',
    method: 'put',
    data
  })
}

// @Tags ExamPaperTemplate
// @Summary 用id查询ExamPaperTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExamPaperTemplate true "用id查询ExamPaperTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examPaperTemplate/findExamPaperTemplate [get]
export const findExamPaperTemplate = (params) => {
  return service({
    url: '/examPaperTemplate/findExamPaperTemplate',
    method: 'get',
    params
  })
}

// @Tags ExamPaperTemplate
// @Summary 分页获取ExamPaperTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ExamPaperTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPaperTemplate/getExamPaperTemplateList [get]
export const getExamPaperTemplateList = (params) => {
  return service({
    url: '/examPaperTemplate/getExamPaperTemplateList',
    method: 'get',
    params
  })
}
