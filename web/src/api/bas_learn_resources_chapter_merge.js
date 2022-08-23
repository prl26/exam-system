import service from '@/utils/request'

// @Tags Learn_resources_chapter_merge
// @Summary 创建Learn_resources_chapter_merge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Learn_resources_chapter_merge true "创建Learn_resources_chapter_merge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /learnResourcesChapterMerge/createLearn_resources_chapter_merge [post]
export const createLearn_resources_chapter_merge = (data) => {
  return service({
    url: '/learnResourcesChapterMerge/createLearn_resources_chapter_merge',
    method: 'post',
    data
  })
}

// @Tags Learn_resources_chapter_merge
// @Summary 删除Learn_resources_chapter_merge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Learn_resources_chapter_merge true "删除Learn_resources_chapter_merge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /learnResourcesChapterMerge/deleteLearn_resources_chapter_merge [delete]
export const deleteLearn_resources_chapter_merge = (data) => {
  return service({
    url: '/learnResourcesChapterMerge/deleteLearn_resources_chapter_merge',
    method: 'delete',
    data
  })
}

// @Tags Learn_resources_chapter_merge
// @Summary 删除Learn_resources_chapter_merge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Learn_resources_chapter_merge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /learnResourcesChapterMerge/deleteLearn_resources_chapter_merge [delete]
export const deleteLearn_resources_chapter_mergeByIds = (data) => {
  return service({
    url: '/learnResourcesChapterMerge/deleteLearn_resources_chapter_mergeByIds',
    method: 'delete',
    data
  })
}

// @Tags Learn_resources_chapter_merge
// @Summary 更新Learn_resources_chapter_merge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Learn_resources_chapter_merge true "更新Learn_resources_chapter_merge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /learnResourcesChapterMerge/updateLearn_resources_chapter_merge [put]
export const updateLearn_resources_chapter_merge = (data) => {
  return service({
    url: '/learnResourcesChapterMerge/updateLearn_resources_chapter_merge',
    method: 'put',
    data
  })
}

// @Tags Learn_resources_chapter_merge
// @Summary 用id查询Learn_resources_chapter_merge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Learn_resources_chapter_merge true "用id查询Learn_resources_chapter_merge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /learnResourcesChapterMerge/findLearn_resources_chapter_merge [get]
export const findLearn_resources_chapter_merge = (params) => {
  return service({
    url: '/learnResourcesChapterMerge/findLearn_resources_chapter_merge',
    method: 'get',
    params
  })
}

// @Tags Learn_resources_chapter_merge
// @Summary 分页获取Learn_resources_chapter_merge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Learn_resources_chapter_merge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /learnResourcesChapterMerge/getLearn_resources_chapter_mergeList [get]
export const getLearn_resources_chapter_mergeList = (params) => {
  return service({
    url: '/learnResourcesChapterMerge/getLearn_resources_chapter_mergeList',
    method: 'get',
    params
  })
}
