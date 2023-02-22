package questionBank

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/questionBank/enum/problemType"
	"github.com/xuri/excelize/v2"

	questionBankPo "github.com/prl26/exam-system/server/model/questionBank/po"
	questionBankReq "github.com/prl26/exam-system/server/model/questionBank/vo/request"
	questionBankResp "github.com/prl26/exam-system/server/model/questionBank/vo/response"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
)

type TargetApi struct {
}

var TargetService = service.ServiceGroupApp.QuestionBankServiceGroup.TargetService
var ChapterService = service.ServiceGroupApp.BasicdataApiGroup.ChapterService
var KnowledgeService = service.ServiceGroupApp.LessondataServiceGroup.KnowledgeService

// Create 创建靶场题
func (api *TargetApi) Create(c *gin.Context) {
	var req questionBankReq.TargetCreate
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ProblemType": {utils.NotEmpty()},
		"CanPractice": {utils.NotEmpty()},
		"CanExam":     {utils.NotEmpty()},
		"Title":       {utils.NotEmpty()},
		"Describe":    {utils.NotEmpty()},
		"ByteCode":    {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	if err := TargetService.Create(&req.Target); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("创建失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("创建成功", c)
	}
}

// Delete 删除靶场题
func (api *TargetApi) Delete(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := TargetService.Delete(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("批量删除失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新靶场题
func (api *TargetApi) Update(c *gin.Context) {
	var req questionBankPo.Target
	_ = c.ShouldBindJSON(&req)
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TargetService.Update(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("更新失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithMessage("更新成功", c)
	}
}

// FindList  分页查找靶场题
func (api *TargetApi) FindList(c *gin.Context) {
	var pageInfo questionBankReq.TargetSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := TargetService.FindTargetList(pageInfo.TargetSearchCriteria, pageInfo.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
	} else {
		questionBankResp.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// FindDetail  获取靶场题详细
func (api *TargetApi) FindDetail(c *gin.Context) {
	var req questionBankReq.DetailFind
	_ = c.ShouldBindQuery(&req)
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}

	if Target, err := TargetService.FindDetail(req.Id, true); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		questionBankResp.ErrorHandle(c, fmt.Errorf("获取失败:%s", err.Error()))
		return
	} else {
		questionBankResp.OkWithDetailed(Target, "获取成功", c)
	}
}

func (api *TargetApi) Import(c *gin.Context) {
	var req questionBankReq.TargetExcel
	_ = c.ShouldBind(&req)
	verify := utils.Rules{
		"File":     {utils.NotEmpty()},
		"LessonId": {utils.NotEmpty()},
	}
	if err := utils.Verify(req, verify); err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	// 此处可以增加对file的限制
	file, err := req.File.Open()
	if err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	reader, err := excelize.OpenReader(file)
	if err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	rows, err := reader.GetRows("Sheet1")
	n := len(rows) - 1
	if err != nil || n <= 0 {
		questionBankResp.CheckHandle(c, err)
		return
	}
	lessonId := req.LessonId
	chapterTable := make(map[string]uint)
	knowledgeTable := make(map[uint]map[string]uint)
	ProblemTypeTable := map[string]uint{"困难": 3, "中等": 2, "简单": 1}
	yes := 1
	no := 0
	whetherTable := map[string]*int{"是": &yes, "否": &no}
	targetList := make([]*questionBankPo.Target, 0, n)
	for _, row := range rows[1:] {
		chapterName := row[6]
		knowledgeName := row[7]
		var chapterId, knowledgeId uint
		if id, ok := chapterTable[chapterName]; !ok {
			id, err = ChapterService.AccessOrCreateByName(chapterName, int(lessonId))
			if err != nil {
				questionBankResp.CheckHandle(c, err)
				return
			}
			chapterId = id
			chapterTable[chapterName] = id
			knowledgeTable[id] = make(map[string]uint)
		} else {
			chapterId = id
		}
		if id, ok := knowledgeTable[chapterId][knowledgeName]; !ok {
			id, err = KnowledgeService.AccessOrCreateByName(knowledgeName, chapterId)
			if err != nil {
				questionBankResp.CheckHandle(c, err)
				return
			}
			knowledgeTable[chapterId][knowledgeName] = id
			knowledgeId = id
		} else {
			knowledgeId = id
		}
		target := questionBankPo.Target{
			TargetModel: questionBankPo.TargetModel{
				BasicModel: questionBankPo.BasicModel{
					SimpleModel: questionBankPo.SimpleModel{
						SerNo:       row[0],
						ProblemType: problemType.ProblemType(ProblemTypeTable[row[3]]),
						PracticeExamSupport: questionBankPo.PracticeExamSupport{
							IsCheck:     whetherTable[row[10]],
							CanPractice: whetherTable[row[8]],
							CanExam:     whetherTable[row[9]],
						},
						Title: row[1],
					},
					Describe: row[2],
				},
				Code:     row[4],
				ByteCode: row[5],
			},
			CourseSupport: questionBankPo.CourseSupport{
				LessonId:    lessonId,
				ChapterId:   chapterId,
				KnowledgeId: knowledgeId,
			},
		}
		targetList = append(targetList, &target)
	}
	number, err := TargetService.CreateList(targetList)
	if err != nil {
		questionBankResp.CheckHandle(c, err)
		return
	}
	questionBankResp.OkWithMessage(fmt.Sprintf("成功创建了%d个题目", number), c)
}
