import service from '@/utils/request'

// @Tags ExamPaper
// @Summary 创建ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamPaper true "创建ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPaper/createExamPaper [post]
export const createExamPaper = (data) => {
  return service({
    url: '/examPaper/createExamPaper',
    method: 'post',
    data
  })
}

// @Tags ExamPaper
// @Summary 删除ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamPaper true "删除ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examPaper/deleteExamPaper [delete]
export const deleteExamPaper = (data) => {
  return service({
    url: '/examPaper/deleteExamPaper',
    method: 'delete',
    data
  })
}

// @Tags ExamPaper
// @Summary 删除ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examPaper/deleteExamPaper [delete]
export const deleteExamPaperByIds = (data) => {
  return service({
    url: '/examPaper/deleteExamPaperByIds',
    method: 'delete',
    data
  })
}

// @Tags ExamPaper
// @Summary 更新ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamPaper true "更新ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /examPaper/updateExamPaper [put]
export const updateExamPaper = (data) => {
  return service({
    url: '/examPaper/updateExamPaper',
    method: 'put',
    data
  })
}

// @Tags ExamPaper
// @Summary 用id查询ExamPaper
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExamPaper true "用id查询ExamPaper"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examPaper/findExamPaper [get]
export const findExamPaper = (params) => {
  return service({
    url: '/examPaper/findExamPaper',
    method: 'get',
    params
  })
}

// @Tags ExamPaper
// @Summary 分页获取ExamPaper列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ExamPaper列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPaper/getExamPaperList [get]
export const getExamPaperList = (params) => {
  return service({
    url: '/examPaper/getExamPaperList',
    method: 'get',
    params
  })
}
