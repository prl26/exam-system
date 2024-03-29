package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/prl26/exam-system/server/global"
	"github.com/prl26/exam-system/server/model/basicdata"
	"github.com/prl26/exam-system/server/model/common/response"
	"github.com/prl26/exam-system/server/model/system"
	systemReq "github.com/prl26/exam-system/server/model/system/request"
	systemRes "github.com/prl26/exam-system/server/model/system/response"
	"github.com/prl26/exam-system/server/utils"
	"go.uber.org/zap"
	"strconv"
)

// @Tags Base
// @Summary 学生登录
// @Produce  application/json
// @Param data body systemReq.Login true "学号, 密码, 验证码"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=string} "返回包括用户信息,token,过期时间"
// @Router /base/studentLogin [post]
func (b *BaseApi) StudentLogin(c *gin.Context) {
	var l systemReq.StudentLogin
	_ = c.ShouldBindJSON(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if store.Verify(l.CaptchaId, l.Captcha, true) {
	u := &basicdata.Student{
		GVA_MODEL: global.GVA_MODEL{ID: l.ID},
		Password:  l.Password,
	}
	if user, err := userService.StudentLogin(u); err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		//如需要冻结功能,可自加字段
		//if user.Enable != 1 {
		//	global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
		//	response.FailWithMessage("用户被禁止登录", c)
		//	return
		//}
		b.StudentTokenNext(c, *user)
	}
	//} else {
	//	response.FailWithMessage("验证码错误", c)
	//}
}

func (b *BaseApi) StudentTokenNext(c *gin.Context, user basicdata.Student) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateStudentClaims(systemReq.StudentBaseClaims{
		ID:          user.ID,
		Name:        user.Name,
		AuthorityId: 666,
	})
	token, err := j.CreateStudentToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithData(gin.H{
			"loginResponse": systemRes.StudentLoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
			},
			"状态": "登录成功",
		}, c)
		return
	}
	if jwtStr, err := jwtService.GetStudentRedisJWT(user.ID); err == redis.Nil {
		if err := jwtService.SetStudentRedisJWT(token, user.ID); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithData(gin.H{
			"loginResponse": systemRes.StudentLoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
			},
			"状态": "登录成功",
		}, c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetStudentRedisJWT(token, user.ID); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		if err != nil {
			response.FailWithMessage("查询班级失败", c)
			return
		}
		//response.OkWithDetailed(systemRes.StudentLoginResponse{
		//	User:      user,
		//	Token:     token,
		//	ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		//}, "登录成功", c)
		response.OkWithData(gin.H{
			"loginResponse": systemRes.StudentLoginResponse{
				User:      user,
				Token:     token,
				ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
			},
			"状态": "登录成功",
		}, c)
	}
}

//获取该学生所在的所有教学计划
func (b *BaseApi) GetTeachPlans(c *gin.Context) {
	termId := c.Query("termId")
	termID, _ := strconv.Atoi(termId)
	StudentId := utils.GetStudentId(c)
	Lessons, err := teachClassService.FindTeachClass(StudentId, termID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(Lessons, c)
	}
}
func (b *BaseApi) GetAllTeachPlans(c *gin.Context) {
	termId := c.Query("termId")
	termID, _ := strconv.Atoi(termId)
	StudentId := utils.GetStudentId(c)
	Lessons, err := teachClassService.FindAllTeachClass(StudentId, termID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(Lessons, c)
	}
}

//获取该学生所在的所有教学计划--靶场专区
func (b *BaseApi) GetTargetTeachPlans(c *gin.Context) {
	termId := c.Query("termId")
	termID, _ := strconv.Atoi(termId)
	StudentId := utils.GetStudentId(c)
	Lessons, err := teachClassService.FindTargetTeachClass(StudentId, termID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(Lessons, c)
	}
}
func (b *BaseApi) StudentChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	_ = c.ShouldBindJSON(&req)
	uid := utils.GetStudentId(c)
	u := &basicdata.Student{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: req.Password}
	if _, err := userService.ChangeStudetnPassword(u, req.NewPassword); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
