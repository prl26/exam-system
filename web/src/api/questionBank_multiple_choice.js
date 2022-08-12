import service from '@/utils/request'

// @Tags QuestionBankMultipleChoice
// @Summary 创建QuestionBankMultipleChoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankMultipleChoice true "创建QuestionBankMultipleChoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank_multiple_choice/createQuestionBankMultipleChoice [post]
export const createQuestionBankMultipleChoice = (data) => {
  return service({
    url: '/questionBank_multiple_choice/createQuestionBankMultipleChoice',
    method: 'post',
    data
  })
}

// @Tags QuestionBankMultipleChoice
// @Summary 删除QuestionBankMultipleChoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankMultipleChoice true "删除QuestionBankMultipleChoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBank_multiple_choice/deleteQuestionBankMultipleChoice [delete]
export const deleteQuestionBankMultipleChoice = (data) => {
  return service({
    url: '/questionBank_multiple_choice/deleteQuestionBankMultipleChoice',
    method: 'delete',
    data
  })
}

// @Tags QuestionBankMultipleChoice
// @Summary 删除QuestionBankMultipleChoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankMultipleChoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBank_multiple_choice/deleteQuestionBankMultipleChoice [delete]
export const deleteQuestionBankMultipleChoiceByIds = (data) => {
  return service({
    url: '/questionBank_multiple_choice/deleteQuestionBankMultipleChoiceByIds',
    method: 'delete',
    data
  })
}

// @Tags QuestionBankMultipleChoice
// @Summary 更新QuestionBankMultipleChoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankMultipleChoice true "更新QuestionBankMultipleChoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBank_multiple_choice/updateQuestionBankMultipleChoice [put]
export const updateQuestionBankMultipleChoice = (data) => {
  return service({
    url: '/questionBank_multiple_choice/updateQuestionBankMultipleChoice',
    method: 'put',
    data
  })
}

// @Tags QuestionBankMultipleChoice
// @Summary 用id查询QuestionBankMultipleChoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.QuestionBankMultipleChoice true "用id查询QuestionBankMultipleChoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBank_multiple_choice/findQuestionBankMultipleChoice [get]
export const findQuestionBankMultipleChoice = (params) => {
  return service({
    url: '/questionBank_multiple_choice/findQuestionBankMultipleChoice',
    method: 'get',
    params
  })
}

// @Tags QuestionBankMultipleChoice
// @Summary 分页获取QuestionBankMultipleChoice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取QuestionBankMultipleChoice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank_multiple_choice/getQuestionBankMultipleChoiceList [get]
export const getQuestionBankMultipleChoiceList = (params) => {
  return service({
    url: '/questionBank_multiple_choice/getQuestionBankMultipleChoiceList',
    method: 'get',
    params
  })
}
