import service from '@/utils/request'

// @Tags TeachClassStudent
// @Summary 创建TeachClassStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachClassStudent true "创建TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClassStudent/createTeachClassStudent [post]
export const createTeachClassStudent = (data) => {
  return service({
    url: '/teachClassStudent/createTeachClassStudent',
    method: 'post',
    data
  })
}

// @Tags TeachClassStudent
// @Summary 删除TeachClassStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachClassStudent true "删除TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachClassStudent/deleteTeachClassStudent [delete]
export const deleteTeachClassStudent = (data) => {
  return service({
    url: '/teachClassStudent/deleteTeachClassStudent',
    method: 'delete',
    data
  })
}

// @Tags TeachClassStudent
// @Summary 删除TeachClassStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachClassStudent/deleteTeachClassStudent [delete]
export const deleteTeachClassStudentByIds = (data) => {
  return service({
    url: '/teachClassStudent/deleteTeachClassStudentByIds',
    method: 'delete',
    data
  })
}

// @Tags TeachClassStudent
// @Summary 更新TeachClassStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachClassStudent true "更新TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teachClassStudent/updateTeachClassStudent [put]
export const updateTeachClassStudent = (data) => {
  return service({
    url: '/teachClassStudent/updateTeachClassStudent',
    method: 'put',
    data
  })
}

// @Tags TeachClassStudent
// @Summary 用id查询TeachClassStudent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeachClassStudent true "用id查询TeachClassStudent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teachClassStudent/findTeachClassStudent [get]
export const findTeachClassStudent = (params) => {
  return service({
    url: '/teachClassStudent/findTeachClassStudent',
    method: 'get',
    params
  })
}

// @Tags TeachClassStudent
// @Summary 分页获取TeachClassStudent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeachClassStudent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClassStudent/getTeachClassStudentList [get]
export const getTeachClassStudentList = (params) => {
  return service({
    url: '/teachClassStudent/getTeachClassStudentList',
    method: 'get',
    params
  })
}
