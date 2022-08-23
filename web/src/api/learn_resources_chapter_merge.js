import service from '@/utils/request'

// @Tags LearnResourcesChapterMerge
// @Summary 创建LearnResourcesChapterMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LearnResourcesChapterMerge true "创建LearnResourcesChapterMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /learnResourcesChapterMerge/createLearnResourcesChapterMerge [post]
export const createLearnResourcesChapterMerge = (data) => {
  return service({
    url: '/learnResourcesChapterMerge/createLearnResourcesChapterMerge',
    method: 'post',
    data
  })
}

// @Tags LearnResourcesChapterMerge
// @Summary 删除LearnResourcesChapterMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LearnResourcesChapterMerge true "删除LearnResourcesChapterMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /learnResourcesChapterMerge/deleteLearnResourcesChapterMerge [delete]
export const deleteLearnResourcesChapterMerge = (data) => {
  return service({
    url: '/learnResourcesChapterMerge/deleteLearnResourcesChapterMerge',
    method: 'delete',
    data
  })
}

// @Tags LearnResourcesChapterMerge
// @Summary 删除LearnResourcesChapterMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LearnResourcesChapterMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /learnResourcesChapterMerge/deleteLearnResourcesChapterMerge [delete]
export const deleteLearnResourcesChapterMergeByIds = (data) => {
  return service({
    url: '/learnResourcesChapterMerge/deleteLearnResourcesChapterMergeByIds',
    method: 'delete',
    data
  })
}

// @Tags LearnResourcesChapterMerge
// @Summary 更新LearnResourcesChapterMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LearnResourcesChapterMerge true "更新LearnResourcesChapterMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /learnResourcesChapterMerge/updateLearnResourcesChapterMerge [put]
export const updateLearnResourcesChapterMerge = (data) => {
  return service({
    url: '/learnResourcesChapterMerge/updateLearnResourcesChapterMerge',
    method: 'put',
    data
  })
}

// @Tags LearnResourcesChapterMerge
// @Summary 用id查询LearnResourcesChapterMerge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LearnResourcesChapterMerge true "用id查询LearnResourcesChapterMerge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /learnResourcesChapterMerge/findLearnResourcesChapterMerge [get]
export const findLearnResourcesChapterMerge = (params) => {
  return service({
    url: '/learnResourcesChapterMerge/findLearnResourcesChapterMerge',
    method: 'get',
    params
  })
}

// @Tags LearnResourcesChapterMerge
// @Summary 分页获取LearnResourcesChapterMerge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LearnResourcesChapterMerge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /learnResourcesChapterMerge/getLearnResourcesChapterMergeList [get]
export const getLearnResourcesChapterMergeList = (params) => {
  return service({
    url: '/learnResourcesChapterMerge/getLearnResourcesChapterMergeList',
    method: 'get',
    params
  })
}
