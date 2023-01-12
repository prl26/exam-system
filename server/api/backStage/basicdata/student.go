package basicdata

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	basicdataReq "github.com/prl26/exam-system/server/model/basicdata/request"
	"github.com/prl26/exam-system/server/model/common/request"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/teachplan"
	"github.com/prl26/exam-system/server/service"
	"github.com/prl26/exam-system/server/utils"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"log"
	"strconv"
)

type StudentApi struct {
}

var studentService = service.ServiceGroupApp.BasicdataApiGroup.StudentService

// CreateStudent 创建Student
// @Tags Student
// @Summary 创建Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Student true "创建Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/createStudent [post]
func (studentApi *StudentApi) CreateStudent(c *gin.Context) {
	var student basicdata.Student
	_ = c.ShouldBindJSON(&student)
	verify := utils.Rules{
		"ID":   {utils.NotEmpty()},
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(student, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	student.Password = utils.BcryptHash(strconv.Itoa(int(student.ID)))
	if err := studentService.CreateStudent(student); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStudent 删除Student
// @Tags Student
// @Summary 删除Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Student true "删除Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /student/deleteStudent [delete]
func (studentApi *StudentApi) DeleteStudent(c *gin.Context) {
	var student basicdata.Student
	_ = c.ShouldBindJSON(&student)
	if err := studentService.DeleteStudent(student); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStudentByIds 批量删除Student
// @Tags Student
// @Summary 批量删除Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /student/deleteStudentByIds [delete]
func (studentApi *StudentApi) DeleteStudentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := studentService.DeleteStudentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStudent 更新Student
// @Tags Student
// @Summary 更新Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body basicdata.Student true "更新Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /student/updateStudent [put]
func (studentApi *StudentApi) UpdateStudent(c *gin.Context) {
	var student basicdata.Student
	_ = c.ShouldBindJSON(&student)
	if err := studentService.UpdateStudent(student); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStudent 用id查询Student
// @Tags Student
// @Summary 用id查询Student
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdata.Student true "用id查询Student"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /student/findStudent [get]
func (studentApi *StudentApi) FindStudent(c *gin.Context) {
	var student basicdata.Student
	_ = c.ShouldBindQuery(&student)
	if restudent, err := studentService.GetStudent(student.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restudent": restudent}, c)
	}
}

// GetStudentList 分页获取Student列表
// @Tags Student
// @Summary 分页获取Student列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.StudentSearch true "分页获取Student列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/getStudentList [get]
func (studentApi *StudentApi) GetStudentList(c *gin.Context) {
	var pageInfo basicdataReq.StudentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := studentService.GetStudentInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// AddStudentsByExcel 表格导入学生
// @Tags Student
// @Summary 表格导入学生
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query basicdataReq.StudentExcel true "表格导入学生"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /student/excel [post]
func (studentApi *StudentApi) AddStudentsByExcel(c *gin.Context) {
	var studentExcel basicdataReq.StudentExcel
	_ = c.ShouldBind(&studentExcel)
	verify := utils.Rules{
		"File": {utils.NotEmpty()},
		//"CollegeId": {utils.NotEmpty()},
		//"ProfessionalId": {utils.NotEmpty()}, 专业和学院去掉
		"TermId":   {utils.NotEmpty()},
		"CourseId": {utils.NotEmpty()},
		"ClassId":  {utils.NotEmpty()},
	}
	if err := utils.Verify(studentExcel, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 此处可以增加对file的限制
	file, err := studentExcel.File.Open()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reader, err := excelize.OpenReader(file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rows, err := reader.GetRows("Sheet1")
	n := len(rows) - 1
	if err != nil || n <= 0 {
		response.FailWithMessage(err.Error(), c)
		return
	}

	teachClassID := int(studentExcel.ClassId)
	termID := int(studentExcel.TermId)
	courseId := int(studentExcel.CourseId)
	var scoreService = service.ServiceGroupApp.TeachplanServiceGroup.ScoreService
	newScoreStudents := make([]*teachplan.Score, 0, n)
	NewStudents := make([]*basicdata.Student, 0, n)
	//ExitStudents := make([]*basicdata.Student, 0, n)
	rows = rows[1:]
	for i, row := range rows {
		length := len(rows[i])
		if length < 6 {
			response.FailWithMessage(fmt.Sprintf("表格格式有误!第%d行只有%d个数据", i+1, length), c)
			return
		}
		id, err := strconv.Atoi(row[0])
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("表格格式有误!第%d行学号为非数字(%s)", i+1, row[0]), c)
			continue
		}

		student := &basicdata.Student{}
		scoreStudent := &teachplan.Score{}
		student.ID = uint(id)
		scoreStudent.StudentId = &id
		scoreStudent.TeachClassId = &teachClassID
		scoreStudent.TermId = &termID
		scoreStudent.CourseId = &courseId
		student.IdCard = row[1]
		student.Name = row[2]
		student.Sex = row[3]
		student.ProfessionalName = row[4]
		student.CollegeName = row[5]
		student.Password = utils.BcryptHash(row[0]) //密码默认为学号
		//if length <= 5 || row[4] == "" {
		//	student.Password = utils.BcryptHash(row[0])
		//} else {
		//	student.Password = utils.BcryptHash(row[4])
		//}

		user := studentService.QueryStudentById(student.ID)
		if user.ID != 0 {
			// 学生已存在，直接关联
			_ = multiTableService.AssociationStudent(student, studentExcel.ClassId)
		} else {
			// 学生不存在
			NewStudents = append(NewStudents, student)
		}

		score := scoreService.QueryScoreByStudent(scoreStudent)
		if score.StudentId == nil && score.TeachClassId == nil {
			// 不存在这个 数据 则创建索引
			newScoreStudents = append(newScoreStudents, scoreStudent)
		}

	}

	// 学生不存在 先创建学生
	err = studentService.CreateStudents(NewStudents)
	if err != nil {
		log.Printf(err.Error())
	}
	err = multiTableService.AssociationStudents(NewStudents, studentExcel.ClassId)

	// 不存在的创建 tea-score 数据 创建索引
	_ = scoreService.CreateScores(newScoreStudents)

	if err != nil {
		response.FailWithMessage("学生导入失败", c)
		return
	} else {
		response.OkWithMessage("学生导入成功", c)
		return
	}
}

// ResetStudentsPassword 批量重置学生密码
// @Tags Student
// @Summary 批量重置学生密码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量重置学生密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"重置密码成功"}"
// @Router /student/deleteStudentByIds [post]
func (studentApi *StudentApi) ResetStudentPassword(c *gin.Context) {
	var id request.IdReq
	_ = c.ShouldBindJSON(&id)
	if err := studentService.ResetStudentsPassword(id); err != nil {
		global.GVA_LOG.Error("重置密码失败!", zap.Error(err))
		response.FailWithMessage("重置密码失败", c)
	} else {
		response.OkWithMessage("重置密码成功", c)
	}
}
