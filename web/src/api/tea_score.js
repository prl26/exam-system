import service from '@/utils/request'

// @Tags Score
// @Summary 创建Score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Score true "创建Score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /score/createScore [post]
export const createScore = (data) => {
  return service({
    url: '/score/createScore',
    method: 'post',
    data
  })
}

// @Tags Score
// @Summary 删除Score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Score true "删除Score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /score/deleteScore [delete]
export const deleteScore = (data) => {
  return service({
    url: '/score/deleteScore',
    method: 'delete',
    data
  })
}

// @Tags Score
// @Summary 删除Score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /score/deleteScore [delete]
export const deleteScoreByIds = (data) => {
  return service({
    url: '/score/deleteScoreByIds',
    method: 'delete',
    data
  })
}

// @Tags Score
// @Summary 更新Score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Score true "更新Score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /score/updateScore [put]
export const updateScore = (data) => {
  return service({
    url: '/score/updateScore',
    method: 'put',
    data
  })
}

// @Tags Score
// @Summary 用id查询Score
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Score true "用id查询Score"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /score/findScore [get]
export const findScore = (params) => {
  return service({
    url: '/score/findScore',
    method: 'get',
    params
  })
}

// @Tags Score
// @Summary 分页获取Score列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Score列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /score/getScoreList [get]
export const getScoreList = (params) => {
  return service({
    url: '/score/getScoreList',
    method: 'get',
    params
  })
}
