package request

type ConvertDraft struct {
	PlanId       int    `json:"planId" form:"planId"`
	Name         string `json:"name" form:"name"`
	DraftPaperId uint   `json:"draftPaperId" form:"draftPaperId"`
}
type IdsReq struct {
	Ids []uint `json:"ids" form:"ids"`
}
