import service from '@/utils/request'

// @Tags QuestionBankKnowledgeMerge
// @Summary 创建QuestionBankKnowledgeMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankKnowledgeMerge true "创建QuestionBankKnowledgeMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankKnowledgeMerge/createQuestionBankKnowledgeMerge [post]
export const createQuestionBankKnowledgeMerge = (data) => {
  return service({
    url: '/questionBankKnowledgeMerge/createQuestionBankKnowledgeMerge',
    method: 'post',
    data
  })
}

// @Tags QuestionBankKnowledgeMerge
// @Summary 删除QuestionBankKnowledgeMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankKnowledgeMerge true "删除QuestionBankKnowledgeMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBankKnowledgeMerge/deleteQuestionBankKnowledgeMerge [delete]
export const deleteQuestionBankKnowledgeMerge = (data) => {
  return service({
    url: '/questionBankKnowledgeMerge/deleteQuestionBankKnowledgeMerge',
    method: 'delete',
    data
  })
}

// @Tags QuestionBankKnowledgeMerge
// @Summary 删除QuestionBankKnowledgeMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankKnowledgeMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBankKnowledgeMerge/deleteQuestionBankKnowledgeMerge [delete]
export const deleteQuestionBankKnowledgeMergeByIds = (data) => {
  return service({
    url: '/questionBankKnowledgeMerge/deleteQuestionBankKnowledgeMergeByIds',
    method: 'delete',
    data
  })
}

// @Tags QuestionBankKnowledgeMerge
// @Summary 更新QuestionBankKnowledgeMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankKnowledgeMerge true "更新QuestionBankKnowledgeMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBankKnowledgeMerge/updateQuestionBankKnowledgeMerge [put]
export const updateQuestionBankKnowledgeMerge = (data) => {
  return service({
    url: '/questionBankKnowledgeMerge/updateQuestionBankKnowledgeMerge',
    method: 'put',
    data
  })
}

// @Tags QuestionBankKnowledgeMerge
// @Summary 用id查询QuestionBankKnowledgeMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.QuestionBankKnowledgeMerge true "用id查询QuestionBankKnowledgeMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBankKnowledgeMerge/findQuestionBankKnowledgeMerge [get]
export const findQuestionBankKnowledgeMerge = (params) => {
  return service({
    url: '/questionBankKnowledgeMerge/findQuestionBankKnowledgeMerge',
    method: 'get',
    params
  })
}

// @Tags QuestionBankKnowledgeMerge
// @Summary 分页获取QuestionBankKnowledgeMerge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取QuestionBankKnowledgeMerge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBankKnowledgeMerge/getQuestionBankKnowledgeMergeList [get]
export const getQuestionBankKnowledgeMergeList = (params) => {
  return service({
    url: '/questionBankKnowledgeMerge/getQuestionBankKnowledgeMergeList',
    method: 'get',
    params
  })
}
