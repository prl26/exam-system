package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}
type GetByStudentId struct {
	StudentId uint `json:"studentId" form:"studentId"`
}
type GetByTeachClassId struct {
	TeachClassId uint `json:"teachClassId" form:"teachClassId"`
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []uint `json:"ids" form:"ids"`
}
type PrePlanReq struct {
	PlanId uint     `json:"planId"`
	Ids    []string `json:"ids" form:"ids"`
}
type IdReq struct {
	Id int `json:"id" form:"id"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
