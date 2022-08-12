// 自动生成模板ArticleResources
package lessondata

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// ArticleResources 结构体
type ArticleResources struct {
      global.GVA_MODEL
      Data  string `json:"data" form:"data" gorm:"column:data;comment:博客文章数据;size:999;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:文章的标题;size:32;"`
      Author  string `json:"author" form:"author" gorm:"column:author;comment:作者;size:255;"`
      IsReference  *int `json:"isReference" form:"isReference" gorm:"column:is_reference;comment:是否为转载文章;size:8;"`
      ReferenceUrl  string `json:"referenceUrl" form:"referenceUrl" gorm:"column:reference_url;comment:转载地址;size:255;"`
}


// TableName ArticleResources 表名
func (ArticleResources) TableName() string {
  return "article_resources"
}

