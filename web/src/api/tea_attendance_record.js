import service from '@/utils/request'

// @Tags TeachAttendanceRecord
// @Summary 创建TeachAttendanceRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachAttendanceRecord true "创建TeachAttendanceRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachAttendanceRecord/createTeachAttendanceRecord [post]
export const createTeachAttendanceRecord = (data) => {
  return service({
    url: '/teachAttendanceRecord/createTeachAttendanceRecord',
    method: 'post',
    data
  })
}

// @Tags TeachAttendanceRecord
// @Summary 删除TeachAttendanceRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachAttendanceRecord true "删除TeachAttendanceRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachAttendanceRecord/deleteTeachAttendanceRecord [delete]
export const deleteTeachAttendanceRecord = (data) => {
  return service({
    url: '/teachAttendanceRecord/deleteTeachAttendanceRecord',
    method: 'delete',
    data
  })
}

// @Tags TeachAttendanceRecord
// @Summary 删除TeachAttendanceRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeachAttendanceRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachAttendanceRecord/deleteTeachAttendanceRecord [delete]
export const deleteTeachAttendanceRecordByIds = (data) => {
  return service({
    url: '/teachAttendanceRecord/deleteTeachAttendanceRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags TeachAttendanceRecord
// @Summary 更新TeachAttendanceRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachAttendanceRecord true "更新TeachAttendanceRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teachAttendanceRecord/updateTeachAttendanceRecord [put]
export const updateTeachAttendanceRecord = (data) => {
  return service({
    url: '/teachAttendanceRecord/updateTeachAttendanceRecord',
    method: 'put',
    data
  })
}

// @Tags TeachAttendanceRecord
// @Summary 用id查询TeachAttendanceRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeachAttendanceRecord true "用id查询TeachAttendanceRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teachAttendanceRecord/findTeachAttendanceRecord [get]
export const findTeachAttendanceRecord = (params) => {
  return service({
    url: '/teachAttendanceRecord/findTeachAttendanceRecord',
    method: 'get',
    params
  })
}

// @Tags TeachAttendanceRecord
// @Summary 分页获取TeachAttendanceRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeachAttendanceRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachAttendanceRecord/getTeachAttendanceRecordList [get]
export const getTeachAttendanceRecordList = (params) => {
  return service({
    url: '/teachAttendanceRecord/getTeachAttendanceRecordList',
    method: 'get',
    params
  })
}
