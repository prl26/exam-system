import service from '@/utils/request'

// @Tags QuestionBankSupplyBlank
// @Summary 创建QuestionBankSupplyBlank
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankSupplyBlank true "创建QuestionBankSupplyBlank"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank_supply_blank/createQuestionBankSupplyBlank [post]
export const createQuestionBankSupplyBlank = (data) => {
  return service({
    url: '/questionBank_supply_blank/createQuestionBankSupplyBlank',
    method: 'post',
    data
  })
}

// @Tags QuestionBankSupplyBlank
// @Summary 删除QuestionBankSupplyBlank
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankSupplyBlank true "删除QuestionBankSupplyBlank"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBank_supply_blank/deleteQuestionBankSupplyBlank [delete]
export const deleteQuestionBankSupplyBlank = (data) => {
  return service({
    url: '/questionBank_supply_blank/deleteQuestionBankSupplyBlank',
    method: 'delete',
    data
  })
}

// @Tags QuestionBankSupplyBlank
// @Summary 删除QuestionBankSupplyBlank
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除QuestionBankSupplyBlank"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /questionBank_supply_blank/deleteQuestionBankSupplyBlank [delete]
export const deleteQuestionBankSupplyBlankByIds = (data) => {
  return service({
    url: '/questionBank_supply_blank/deleteQuestionBankSupplyBlankByIds',
    method: 'delete',
    data
  })
}

// @Tags QuestionBankSupplyBlank
// @Summary 更新QuestionBankSupplyBlank
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QuestionBankSupplyBlank true "更新QuestionBankSupplyBlank"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /questionBank_supply_blank/updateQuestionBankSupplyBlank [put]
export const updateQuestionBankSupplyBlank = (data) => {
  return service({
    url: '/questionBank_supply_blank/updateQuestionBankSupplyBlank',
    method: 'put',
    data
  })
}

// @Tags QuestionBankSupplyBlank
// @Summary 用id查询QuestionBankSupplyBlank
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.QuestionBankSupplyBlank true "用id查询QuestionBankSupplyBlank"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /questionBank_supply_blank/findQuestionBankSupplyBlank [get]
export const findQuestionBankSupplyBlank = (params) => {
  return service({
    url: '/questionBank_supply_blank/findQuestionBankSupplyBlank',
    method: 'get',
    params
  })
}

// @Tags QuestionBankSupplyBlank
// @Summary 分页获取QuestionBankSupplyBlank列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取QuestionBankSupplyBlank列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /questionBank_supply_blank/getQuestionBankSupplyBlankList [get]
export const getQuestionBankSupplyBlankList = (params) => {
  return service({
    url: '/questionBank_supply_blank/getQuestionBankSupplyBlankList',
    method: 'get',
    params
  })
}
