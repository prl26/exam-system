import service from '@/utils/request'

// @Tags ExamPlan
// @Summary 创建ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamPlan true "创建ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPlan/createExamPlan [post]
export const createExamPlan = (data) => {
  return service({
    url: '/examPlan/createExamPlan',
    method: 'post',
    data
  })
}

// @Tags ExamPlan
// @Summary 删除ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamPlan true "删除ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examPlan/deleteExamPlan [delete]
export const deleteExamPlan = (data) => {
  return service({
    url: '/examPlan/deleteExamPlan',
    method: 'delete',
    data
  })
}

// @Tags ExamPlan
// @Summary 删除ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /examPlan/deleteExamPlan [delete]
export const deleteExamPlanByIds = (data) => {
  return service({
    url: '/examPlan/deleteExamPlanByIds',
    method: 'delete',
    data
  })
}

// @Tags ExamPlan
// @Summary 更新ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExamPlan true "更新ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /examPlan/updateExamPlan [put]
export const updateExamPlan = (data) => {
  return service({
    url: '/examPlan/updateExamPlan',
    method: 'put',
    data
  })
}

// @Tags ExamPlan
// @Summary 用id查询ExamPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ExamPlan true "用id查询ExamPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /examPlan/findExamPlan [get]
export const findExamPlan = (params) => {
  return service({
    url: '/examPlan/findExamPlan',
    method: 'get',
    params
  })
}

// @Tags ExamPlan
// @Summary 分页获取ExamPlan列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ExamPlan列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /examPlan/getExamPlanList [get]
export const getExamPlanList = (params) => {
  return service({
    url: '/examPlan/getExamPlanList',
    method: 'get',
    params
  })
}
