import service from '@/utils/request'

// @Tags QuestionBankJudge
// @Summary 创建QuestionBankJudge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankJudge true "创建QuestionBankJudge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank_judge/createQuestionBankJudge [post]
export const createQuestionBankJudge = (data) => {
  return service({
    url: '/questionBank_judge/createQuestionBankJudge',
    method: 'post',
    data
  })
}

// @Tags QuestionBankJudge
// @Summary 删除QuestionBankJudge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankJudge true "删除QuestionBankJudge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBank_judge/deleteQuestionBankJudge [delete]
export const deleteQuestionBankJudge = (data) => {
  return service({
    url: '/questionBank_judge/deleteQuestionBankJudge',
    method: 'delete',
    data
  })
}

// @Tags QuestionBankJudge
// @Summary 删除QuestionBankJudge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankJudge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBank_judge/deleteQuestionBankJudge [delete]
export const deleteQuestionBankJudgeByIds = (data) => {
  return service({
    url: '/questionBank_judge/deleteQuestionBankJudgeByIds',
    method: 'delete',
    data
  })
}

// @Tags QuestionBankJudge
// @Summary 更新QuestionBankJudge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankJudge true "更新QuestionBankJudge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBank_judge/updateQuestionBankJudge [put]
export const updateQuestionBankJudge = (data) => {
  return service({
    url: '/questionBank_judge/updateQuestionBankJudge',
    method: 'put',
    data
  })
}

// @Tags QuestionBankJudge
// @Summary 用id查询QuestionBankJudge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.QuestionBankJudge true "用id查询QuestionBankJudge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBank_judge/findQuestionBankJudge [get]
export const findQuestionBankJudge = (params) => {
  return service({
    url: '/questionBank_judge/findQuestionBankJudge',
    method: 'get',
    params
  })
}

// @Tags QuestionBankJudge
// @Summary 分页获取QuestionBankJudge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取QuestionBankJudge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank_judge/getQuestionBankJudgeList [get]
export const getQuestionBankJudgeList = (params) => {
  return service({
    url: '/questionBank_judge/getQuestionBankJudgeList',
    method: 'get',
    params
  })
}
