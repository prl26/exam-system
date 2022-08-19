import service from '@/utils/request'

// @Tags TeachAttendance
// @Summary 创建TeachAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachAttendance true "创建TeachAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachAttendance/createTeachAttendance [post]
export const createTeachAttendance = (data) => {
  return service({
    url: '/teachAttendance/createTeachAttendance',
    method: 'post',
    data
  })
}

// @Tags TeachAttendance
// @Summary 删除TeachAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachAttendance true "删除TeachAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachAttendance/deleteTeachAttendance [delete]
export const deleteTeachAttendance = (data) => {
  return service({
    url: '/teachAttendance/deleteTeachAttendance',
    method: 'delete',
    data
  })
}

// @Tags TeachAttendance
// @Summary 删除TeachAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeachAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachAttendance/deleteTeachAttendance [delete]
export const deleteTeachAttendanceByIds = (data) => {
  return service({
    url: '/teachAttendance/deleteTeachAttendanceByIds',
    method: 'delete',
    data
  })
}

// @Tags TeachAttendance
// @Summary 更新TeachAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachAttendance true "更新TeachAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teachAttendance/updateTeachAttendance [put]
export const updateTeachAttendance = (data) => {
  return service({
    url: '/teachAttendance/updateTeachAttendance',
    method: 'put',
    data
  })
}

// @Tags TeachAttendance
// @Summary 用id查询TeachAttendance
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeachAttendance true "用id查询TeachAttendance"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teachAttendance/findTeachAttendance [get]
export const findTeachAttendance = (params) => {
  return service({
    url: '/teachAttendance/findTeachAttendance',
    method: 'get',
    params
  })
}

// @Tags TeachAttendance
// @Summary 分页获取TeachAttendance列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeachAttendance列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachAttendance/getTeachAttendanceList [get]
export const getTeachAttendanceList = (params) => {
  return service({
    url: '/teachAttendance/getTeachAttendanceList',
    method: 'get',
    params
  })
}
