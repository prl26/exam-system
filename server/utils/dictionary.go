package utils

import "github.com/prl26/exam-system/server/global"

func GetDictionaryName(Type string, value int) (name string) {
	global.GVA_DB.Raw("select label from sys_dictionary_details,sys_dictionaries "+
		"WHERE sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id and "+
		"sys_dictionaries.type = ? and sys_dictionary_details.`value` =?", Type, value).Scan(&name)
	return
}
