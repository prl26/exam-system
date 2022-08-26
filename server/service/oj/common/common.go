package common

import (
	"exam-system/global"
	"exam-system/model/questionBank"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 23:39

 * @Note:

 **/

type CommonService struct{}

func (c *CommonService) FindProgrammCase(programmId uint, languageId uint) (result []*questionBank.ProgrammCase, err error) {
	r := global.GVA_DB.Where("programm_id = ? AND language_id = ?", programmId, languageId).Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}

	return
}

func (c *CommonService) FindProgrammSupport(programmId uint) (result []*questionBank.ProgrammLanguageMerge, err error) {
	var r = global.GVA_DB.Where("programm_id", programmId).Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return
}
