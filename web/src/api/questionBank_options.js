import service from '@/utils/request'

// @Tags QuestionBankOptions
// @Summary 创建QuestionBankOptions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankOptions true "创建QuestionBankOptions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank_options/createQuestionBankOptions [post]
export const createQuestionBankOptions = (data) => {
  return service({
    url: '/questionBank_options/createQuestionBankOptions',
    method: 'post',
    data
  })
}

// @Tags QuestionBankOptions
// @Summary 删除QuestionBankOptions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankOptions true "删除QuestionBankOptions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBank_options/deleteQuestionBankOptions [delete]
export const deleteQuestionBankOptions = (data) => {
  return service({
    url: '/questionBank_options/deleteQuestionBankOptions',
    method: 'delete',
    data
  })
}

// @Tags QuestionBankOptions
// @Summary 删除QuestionBankOptions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankOptions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBank_options/deleteQuestionBankOptions [delete]
export const deleteQuestionBankOptionsByIds = (data) => {
  return service({
    url: '/questionBank_options/deleteQuestionBankOptionsByIds',
    method: 'delete',
    data
  })
}

// @Tags QuestionBankOptions
// @Summary 更新QuestionBankOptions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankOptions true "更新QuestionBankOptions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBank_options/updateQuestionBankOptions [put]
export const updateQuestionBankOptions = (data) => {
  return service({
    url: '/questionBank_options/updateQuestionBankOptions',
    method: 'put',
    data
  })
}

// @Tags QuestionBankOptions
// @Summary 用id查询QuestionBankOptions
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.QuestionBankOptions true "用id查询QuestionBankOptions"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBank_options/findQuestionBankOptions [get]
export const findQuestionBankOptions = (params) => {
  return service({
    url: '/questionBank_options/findQuestionBankOptions',
    method: 'get',
    params
  })
}

// @Tags QuestionBankOptions
// @Summary 分页获取QuestionBankOptions列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取QuestionBankOptions列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank_options/getQuestionBankOptionsList [get]
export const getQuestionBankOptionsList = (params) => {
  return service({
    url: '/questionBank_options/getQuestionBankOptionsList',
    method: 'get',
    params
  })
}
