import service from '@/utils/request'

// @Tags TeachClass
// @Summary 创建TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachClass true "创建TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClass/createTeachClass [post]
export const createTeachClass = (data) => {
  return service({
    url: '/teachClass/createTeachClass',
    method: 'post',
    data
  })
}

// @Tags TeachClass
// @Summary 删除TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachClass true "删除TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachClass/deleteTeachClass [delete]
export const deleteTeachClass = (data) => {
  return service({
    url: '/teachClass/deleteTeachClass',
    method: 'delete',
    data
  })
}

// @Tags TeachClass
// @Summary 删除TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teachClass/deleteTeachClass [delete]
export const deleteTeachClassByIds = (data) => {
  return service({
    url: '/teachClass/deleteTeachClassByIds',
    method: 'delete',
    data
  })
}

// @Tags TeachClass
// @Summary 更新TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeachClass true "更新TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teachClass/updateTeachClass [put]
export const updateTeachClass = (data) => {
  return service({
    url: '/teachClass/updateTeachClass',
    method: 'put',
    data
  })
}

// @Tags TeachClass
// @Summary 用id查询TeachClass
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeachClass true "用id查询TeachClass"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teachClass/findTeachClass [get]
export const findTeachClass = (params) => {
  return service({
    url: '/teachClass/findTeachClass',
    method: 'get',
    params
  })
}

// @Tags TeachClass
// @Summary 分页获取TeachClass列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeachClass列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teachClass/getTeachClassList [get]
export const getTeachClassList = (params) => {
  return service({
    url: '/teachClass/getTeachClassList',
    method: 'get',
    params
  })
}
