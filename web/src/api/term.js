import service from '@/utils/request'

// @Tags Term
// @Summary 创建Term
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Term true "创建Term"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /term/createTerm [post]
export const createTerm = (data) => {
  return service({
    url: '/term/createTerm',
    method: 'post',
    data
  })
}

// @Tags Term
// @Summary 删除Term
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Term true "删除Term"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /term/deleteTerm [delete]
export const deleteTerm = (data) => {
  return service({
    url: '/term/deleteTerm',
    method: 'delete',
    data
  })
}

// @Tags Term
// @Summary 删除Term
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Term"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /term/deleteTerm [delete]
export const deleteTermByIds = (data) => {
  return service({
    url: '/term/deleteTermByIds',
    method: 'delete',
    data
  })
}

// @Tags Term
// @Summary 更新Term
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Term true "更新Term"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /term/updateTerm [put]
export const updateTerm = (data) => {
  return service({
    url: '/term/updateTerm',
    method: 'put',
    data
  })
}

// @Tags Term
// @Summary 用id查询Term
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Term true "用id查询Term"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /term/findTerm [get]
export const findTerm = (params) => {
  return service({
    url: '/term/findTerm',
    method: 'get',
    params
  })
}

// @Tags Term
// @Summary 分页获取Term列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Term列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /term/getTermList [get]
export const getTermList = (params) => {
  return service({
    url: '/term/getTermList',
    method: 'get',
    params
  })
}
