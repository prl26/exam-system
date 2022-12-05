package common

import (
	"github.com/prl26/exam-system/server/global"
	commonError "github.com/prl26/exam-system/server/model/common/error"
	ojBo "github.com/prl26/exam-system/server/model/oj/bo"
	questionBank "github.com/prl26/exam-system/server/model/questionBank/po"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 23:39

 * @Note:

 **/

type CommonService struct{}

func (s CommonService) FindProgram(id uint) (*ojBo.Program, error) {
	result := &ojBo.Program{}
	if r := global.GVA_DB.Model(&questionBank.Program{}).Find(&result, id); r.Error != nil {
		return nil, r.Error
	} else if r.RowsAffected == 0 {
		return nil, commonError.NotFoundError
	} else {
		return result, nil
	}
}

//
//func (c *CommonService) FindProgrammCase(programmId uint, languageId uint) (result []*questionBank.ProgrammCase, err error) {
//	r := global.GVA_DB.Where("programm_id = ? AND language_id = ?", programmId, languageId).Find(&result)
//	if r.Error != nil {
//		return nil, r.Error
//	}
//
//	return
//}
//
//func (c *CommonService) FindProgrammSupport(programmId uint) (result []*questionBank.ProgrammLanguageMerge, err error) {
//	var r = global.GVA_DB.Where("programm_id", programmId).Find(&result)
//	if r.Error != nil {
//		return nil, r.Error
//	}
//	return
//}
