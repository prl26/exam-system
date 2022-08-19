import service from '@/utils/request'

// @Tags Professional
// @Summary 创建Professional
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Professional true "创建Professional"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /professional/createProfessional [post]
export const createProfessional = (data) => {
  return service({
    url: '/professional/createProfessional',
    method: 'post',
    data
  })
}

// @Tags Professional
// @Summary 删除Professional
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Professional true "删除Professional"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /professional/deleteProfessional [delete]
export const deleteProfessional = (data) => {
  return service({
    url: '/professional/deleteProfessional',
    method: 'delete',
    data
  })
}

// @Tags Professional
// @Summary 删除Professional
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Professional"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /professional/deleteProfessional [delete]
export const deleteProfessionalByIds = (data) => {
  return service({
    url: '/professional/deleteProfessionalByIds',
    method: 'delete',
    data
  })
}

// @Tags Professional
// @Summary 更新Professional
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Professional true "更新Professional"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /professional/updateProfessional [put]
export const updateProfessional = (data) => {
  return service({
    url: '/professional/updateProfessional',
    method: 'put',
    data
  })
}

// @Tags Professional
// @Summary 用id查询Professional
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Professional true "用id查询Professional"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /professional/findProfessional [get]
export const findProfessional = (params) => {
  return service({
    url: '/professional/findProfessional',
    method: 'get',
    params
  })
}

// @Tags Professional
// @Summary 分页获取Professional列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Professional列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /professional/getProfessionalList [get]
export const getProfessionalList = (params) => {
  return service({
    url: '/professional/getProfessionalList',
    method: 'get',
    params
  })
}
