package common

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/questionBank"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/8/24 23:39

 * @Note:

 **/

type Service struct{}

//
// FindProgrammCase
//  @Description:  通过程序题的ID和语言ID获取该语言的所有用例
//  @receiver c
//  @param programmId
//  @param language
//  @return result
//  @return err
//
func (c *Service) FindProgrammCase(programmId uint, languageId uint) (result []*questionBank.ProgrammCase, err error) {
	r := global.GVA_DB.Where("programm_id = ? AND language_id = ?", programmId, languageId).Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}

	return
}

//
// FindProgrammSupport
//  @Description: 通过programmId 获取 程序题所能支持的编程语言 以及编程语言的默认代码 参考答案等信息
//  @receiver c*Service
//  @param programmId
//  @return result
//  @return err
//
func (c *Service) FindProgrammSupport(programmId uint) (result []*questionBank.ProgrammLanguageMerge, err error) {
	var r = global.GVA_DB.Where("programm_id", programmId).Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return
}
