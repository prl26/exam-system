import service from '@/utils/request'

// @Tags Knowledge
// @Summary 创建Knowledge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Knowledge true "创建Knowledge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /knowledge/createKnowledge [post]
export const createKnowledge = (data) => {
  return service({
    url: '/knowledge/createKnowledge',
    method: 'post',
    data
  })
}

// @Tags Knowledge
// @Summary 删除Knowledge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Knowledge true "删除Knowledge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /knowledge/deleteKnowledge [delete]
export const deleteKnowledge = (data) => {
  return service({
    url: '/knowledge/deleteKnowledge',
    method: 'delete',
    data
  })
}

// @Tags Knowledge
// @Summary 删除Knowledge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Knowledge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /knowledge/deleteKnowledge [delete]
export const deleteKnowledgeByIds = (data) => {
  return service({
    url: '/knowledge/deleteKnowledgeByIds',
    method: 'delete',
    data
  })
}

// @Tags Knowledge
// @Summary 更新Knowledge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Knowledge true "更新Knowledge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /knowledge/updateKnowledge [put]
export const updateKnowledge = (data) => {
  return service({
    url: '/knowledge/updateKnowledge',
    method: 'put',
    data
  })
}

// @Tags Knowledge
// @Summary 用id查询Knowledge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Knowledge true "用id查询Knowledge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /knowledge/findKnowledge [get]
export const findKnowledge = (params) => {
  return service({
    url: '/knowledge/findKnowledge',
    method: 'get',
    params
  })
}

// @Tags Knowledge
// @Summary 分页获取Knowledge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Knowledge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /knowledge/getKnowledgeList [get]
export const getKnowledgeList = (params) => {
  return service({
    url: '/knowledge/getKnowledgeList',
    method: 'get',
    params
  })
}
