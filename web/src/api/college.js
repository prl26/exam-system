import service from '@/utils/request'

// @Tags College
// @Summary 创建College
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.College true "创建College"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /college/createCollege [post]
export const createCollege = (data) => {
  return service({
    url: '/college/createCollege',
    method: 'post',
    data
  })
}

// @Tags College
// @Summary 删除College
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.College true "删除College"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /college/deleteCollege [delete]
export const deleteCollege = (data) => {
  return service({
    url: '/college/deleteCollege',
    method: 'delete',
    data
  })
}

// @Tags College
// @Summary 删除College
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除College"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /college/deleteCollege [delete]
export const deleteCollegeByIds = (data) => {
  return service({
    url: '/college/deleteCollegeByIds',
    method: 'delete',
    data
  })
}

// @Tags College
// @Summary 更新College
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.College true "更新College"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /college/updateCollege [put]
export const updateCollege = (data) => {
  return service({
    url: '/college/updateCollege',
    method: 'put',
    data
  })
}

// @Tags College
// @Summary 用id查询College
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.College true "用id查询College"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /college/findCollege [get]
export const findCollege = (params) => {
  return service({
    url: '/college/findCollege',
    method: 'get',
    params
  })
}

// @Tags College
// @Summary 分页获取College列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取College列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /college/getCollegeList [get]
export const getCollegeList = (params) => {
  return service({
    url: '/college/getCollegeList',
    method: 'get',
    params
  })
}
